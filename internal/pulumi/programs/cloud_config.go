package programs

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed server_cloud_config.yml
var cloudConfigTemplate string

const (
	// sparkPreStartHostDir lives on the Metio data disk, next to data/ rather
	// than inside it so the hook never appears in world backups. It cannot live
	// on the boot disk: Container-Optimized OS mounts the root filesystem
	// read-only, and cloud-init writes write_files entries before its mounts
	// module attaches the data disk, so the hook is created from runcmd.
	sparkPreStartHostDir = "/mnt/disks/minecraft/pre-start"
	// sparkPreStartContainerDir is where the hook directory is mounted into
	// the Minecraft container.
	sparkPreStartContainerDir = "/hooks"
	sparkPreStartScript       = "spark-profiler.sh"
)

type TemplateConfig struct {
	Region               string
	GCPProject           string
	Environment          string
	InstanceName         string
	BackupBucket         string
	ServerID             string
	BackupRetentionDays  int
	ResticPassword       string
	MachineAgentImage    string
	BackupImage          string
	BackupInterval       string
	PruneResticRetention string
	BackupServiceEnable  string
	MinecraftVersion     string
	RCONPassword         string
	ControllerURL        string
	AgentToken           string

	// Modpack makes the server pack-driven (ADR-0006). When set, the rendered
	// Minecraft unit passes MODRINTH_MODPACK instead of VERSION, because the
	// pack manifest controls the loader, Minecraft version and mod set.
	Modpack *ModpackConfig

	// RestoreSnapshotID is the restic snapshot to restore before Minecraft
	// starts on first boot. When set, a one-off restore step is added to the
	// cloud-config runcmd section.
	RestoreSnapshotID string

	// RestoreSourcePrefix is the restic repository prefix inside the backup
	// bucket that contains the snapshot to restore (e.g.
	// "servers/{old-server-id}/restic"). It differs from the new server's own
	// prefix so the restore reads from the source backup's repository.
	RestoreSourcePrefix string
}

func imageRegistryHost(image string, fallbackRegion string) string {
	host, _, _ := strings.Cut(image, "/")
	if strings.Contains(host, ".") {
		return host
	}
	return fallbackRegion + "-docker.pkg.dev"
}

// normalizeBackupPrefix strips a trailing slash from a restic repository
// prefix. Stored prefixes end in "/" (the backup hook stamps
// "servers/<id>/restic/"), so normalizing keeps the IAM condition and the
// restore repository URL from rendering a doubled slash. The wire format is
// unchanged.
func normalizeBackupPrefix(prefix string) string {
	return strings.TrimSuffix(prefix, "/")
}

func RenderCloudConfig(config *TemplateConfig) (string, error) {
	imageHost := imageRegistryHost(config.MachineAgentImage, config.Region)

	// Normalize defaults so the rendered unit file always has valid values,
	// even when the caller only sets a subset of the template fields.
	if config.BackupRetentionDays == 0 {
		config.BackupRetentionDays = 90
	}
	if config.BackupImage == "" {
		config.BackupImage = "ghcr.io/itzg/mc-backup:latest"
	}
	if config.BackupInterval == "" {
		config.BackupInterval = "1h"
	}
	if config.PruneResticRetention == "" {
		config.PruneResticRetention = fmt.Sprintf("--keep-within %dd", config.BackupRetentionDays)
	}

	// Build the optional first-boot restore: a one-shot systemd unit that
	// restores the snapshot into the data directory before any services start.
	// Running as the metio user (as the service units do) gives a writable HOME
	// so docker-credential-gcr can write its config; ordering is guaranteed by
	// Minecraft's After=/Requires=.
	//
	// The unit is written to the disk on first boot. It runs exactly once: on an
	// already-restored machine the run-once sentinel keyed to the snapshot ID
	// makes ConditionPathExists skip it, and runcmd/cloud-init re-run on COS on
	// every boot (users-groups is once, scripts-user/runcmd are "always"). If
	// the restore fails the unit is skipped-but-satisfied, so Minecraft never
	// starts against an empty world and the machine agent still reports.
	afterRestore := ""
	requiresRestore := ""
	restoreWriteFiles := ""
	if config.RestoreSnapshotID != "" && config.RestoreSourcePrefix != "" {
		// The restore dependency lines are injected inline into the
		// minecraft.service unit, at the same block indentation, so an empty
		// (non-restore) render leaves no dangling blank line inside the YAML
		// literal block.
		afterRestore = "\n      After=metio-restore.service"
		requiresRestore = "\n      Requires=metio-restore.service"
		sourcePrefix := normalizeBackupPrefix(config.RestoreSourcePrefix)
		// Values interpolated into a systemd unit are escaped: '%' becomes '%%'
		// so %n-style specifiers are not mangled in the rendered content.
		restoreWriteFiles = fmt.Sprintf(`  - path: /etc/systemd/system/metio-restore.service
    permissions: '0644'
    content: |
      [Unit]
      Description=Metio One-Shot World Restore
      After=docker.service
      Requires=docker.service
      Before=minecraft.service
      ConditionPathExists=!/mnt/disks/minecraft/.metio-restore-%s.done

      [Service]
      Type=oneshot
      RemainAfterExit=yes
      User=metio
      WorkingDirectory=/home/metio
      ExecStartPre=/usr/bin/docker-credential-gcr configure-docker --registries %s
      ExecStartPre=/bin/bash -c 'for i in 1 2 3; do docker pull %s && exit 0; sleep 10; done; exit 1'
      ExecStart=/usr/bin/docker run --rm --name metio-restore --network host \
        -e RESTIC_REPOSITORY=gs:%s:/%s \
        -e RESTIC_PASSWORD=%s \
        -v /mnt/disks/minecraft/data:/data \
        --entrypoint /usr/bin/restic %s restore %s:/data --target /data
      ExecStartPost=+/bin/touch /mnt/disks/minecraft/.metio-restore-%s.done
`, config.RestoreSnapshotID, imageHost, config.BackupImage,
			config.BackupBucket, sourcePrefix, config.ResticPassword,
			config.BackupImage, config.RestoreSnapshotID, config.RestoreSnapshotID)
	}
	// Pack servers: spark 1.10.x bundles async-profiler, whose native code
	// segfaults on the Temurin 25 image the platform pins (metio#631), so the
	// server crash-loops before it ever answers RCON. The 1.10 line reads its
	// settings only from config/spark/config.json (environment variables and
	// system properties arrived after 1.10.53), so the file has to be written.
	//
	// The write happens in the image's PRE_START_SCRIPT hook rather than from
	// the host: container setup installs or re-syncs the pack (including its
	// overrides) after the container starts, so a file written before docker
	// run can be replaced by pack content. Hook scripts execute directly, so
	// the file must be executable, and the whole directory is mounted so a
	// missing file cannot turn into a docker-created directory that fails the
	// hook and blocks the server.
	//
	// The hook is created from runcmd, not write_files: cloud-init's
	// write-files module runs in the init stage, before its mounts module has
	// attached the data disk, so a write_files entry would target the
	// read-only root filesystem on Container-Optimized OS and abort the whole
	// module. Deferred write_files entries are not an option either, because
	// COS does not enable cc_write_files_deferred. runcmd executes after the
	// mount and before the services start.
	sparkRuncmd := ""
	if config.Modpack != nil {
		sparkRuncmd = fmt.Sprintf(`
  - |
    install -d -m 0755 %s
    cat > %s/%s <<'SPARK_EOF'
    #!/bin/sh
    # Keep spark's async profiler off the JVM: its native code segfaults on
    # Temurin 25 (metio#631). /spark profiler --force-java-sampler is the
    # safe manual profiling path.
    set -eu
    cfg=/data/config/spark/config.json
    mkdir -p /data/config/spark
    if [ -f "$cfg" ] && grep -qE '"backgroundProfiler"[[:space:]]*:[[:space:]]*false|"backgroundProfilerEngine"[[:space:]]*:[[:space:]]*"java"' "$cfg"; then
      exit 0
    fi
    printf '{ "backgroundProfiler": false }\n' > "$cfg"
    SPARK_EOF
    chmod 0755 %s/%s
`, sparkPreStartHostDir, sparkPreStartHostDir, sparkPreStartScript, sparkPreStartHostDir, sparkPreStartScript)
	}
	startCommand := fmt.Sprintf("  - systemctl start minecraft %smetio-machine-agent", config.BackupServiceEnable)

	// The Minecraft unit identifies what to run either by vanilla version or by
	// Modrinth pack (ADR-0006). A pack-driven server omits VERSION entirely and
	// passes the pack, plus its pinned version when one was chosen; the image
	// then resolves loader, Minecraft version and mods from the pack manifest.
	//
	// TYPE must be set to MODRINTH: the image only treats MODRINTH_MODPACK as a
	// modpack when the server type (or MODPACK_PLATFORM) says so, and otherwise
	// keeps its VANILLA default and boots unmodded.
	//
	// The multi-line form keeps the systemd line-continuation backslashes
	// consistent with the surrounding ExecStart block.
	minecraftArgs := fmt.Sprintf("-e VERSION=%s \\", config.MinecraftVersion)
	if config.Modpack != nil {
		minecraftArgs = fmt.Sprintf("-e TYPE=MODRINTH \\\n        -e MODRINTH_MODPACK=%s \\", config.Modpack.ProjectID)
		if config.Modpack.VersionID != "" {
			minecraftArgs += fmt.Sprintf("\n        -e MODRINTH_VERSION=%s \\", config.Modpack.VersionID)
		}
		minecraftArgs += fmt.Sprintf("\n        -e PRE_START_SCRIPT=%s/%s \\", sparkPreStartContainerDir, sparkPreStartScript)
		minecraftArgs += fmt.Sprintf("\n        -v %s:%s:ro \\", sparkPreStartHostDir, sparkPreStartContainerDir)
	}

	replacements := map[string]string{
		"${restoreWriteFiles}":    restoreWriteFiles,
		"${sparkRuncmd}":          sparkRuncmd,
		"${afterRestore}":         afterRestore,
		"${requiresRestore}":      requiresRestore,
		"${startCommand}":         startCommand,
		"${minecraftArgs}":        minecraftArgs,
		"${region}":               config.Region,
		"${gcpProject}":           config.GCPProject,
		"${environment}":          config.Environment,
		"${instanceName}":         config.InstanceName,
		"${backupBucket}":         config.BackupBucket,
		"${serverId}":             config.ServerID,
		"${backupRetentionDays}":  strconv.Itoa(config.BackupRetentionDays),
		"${resticPassword}":       config.ResticPassword,
		"${machineAgentImage}":    config.MachineAgentImage,
		"${backupImage}":          config.BackupImage,
		"${backupInterval}":       config.BackupInterval,
		"${pruneResticRetention}": config.PruneResticRetention,
		"${backupServiceEnable}":  config.BackupServiceEnable,
		"${rconPassword}":         config.RCONPassword,
		"${controllerUrl}":        config.ControllerURL,
		"${agentToken}":           config.AgentToken,
		"${imageRegistryHost}":    imageHost,
	}

	result := cloudConfigTemplate
	for k, v := range replacements {
		result = strings.ReplaceAll(result, k, v)
	}

	return result, nil
}
