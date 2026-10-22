package programs

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestImageRegistryHost(t *testing.T) {
	tests := []struct {
		name           string
		image          string
		fallbackRegion string
		want           string
	}{
		{
			name:           "fully qualified AR image",
			image:          "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
			fallbackRegion: "europe-west2",
			want:           "europe-west3-docker.pkg.dev",
		},
		{
			name:           "multi-region AR host",
			image:          "europe-docker.pkg.dev/metio-distribution/metio/machine-agent:1.5.0",
			fallbackRegion: "europe-west2",
			want:           "europe-docker.pkg.dev",
		},
		{
			name:           "bare image with no host (guard fallback)",
			image:          "machine-agent:latest",
			fallbackRegion: "europe-west2",
			want:           "europe-west2-docker.pkg.dev",
		},
		{
			name:           "short bare image fallback",
			image:          "alpine",
			fallbackRegion: "europe-west3",
			want:           "europe-west3-docker.pkg.dev",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := imageRegistryHost(tt.image, tt.fallbackRegion)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRenderCloudConfig_ImageRegistryHost(t *testing.T) {
	cfg := &TemplateConfig{
		Region:            "europe-west2",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)
	assert.Contains(t, result, "--registries europe-west3-docker.pkg.dev")
}

func TestRenderCloudConfig_ImageRegistryHostFallback(t *testing.T) {
	cfg := &TemplateConfig{
		Region:            "europe-west2",
		MachineAgentImage: "machine-agent:latest",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)
	assert.Contains(t, result, "--registries europe-west2-docker.pkg.dev")
}

func TestRenderCloudConfig_CentralBackupSettings(t *testing.T) {
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "0dcbaca4-2a26-489c-b4a3-d2fad8bb6483",
		BackupRetentionDays: 90,
		ResticPassword:      "deployment-wide-restic-password",
		RCONPassword:        "rcon-password",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	repo := "gs:my-project-development-backups:/servers/0dcbaca4-2a26-489c-b4a3-d2fad8bb6483/restic"
	assert.Contains(t, result, "RESTIC_REPOSITORY="+repo)
	assert.Contains(t, result, "PRUNE_RESTIC_RETENTION=\"--keep-within 90d\"")
	assert.Contains(t, result, "RESTIC_PASSWORD=deployment-wide-restic-password")
	// The backup hook needs the server ID to stamp manifests for the
	// machine-agent's backup reporting (ADR-0004).
	assert.Contains(t, result, "METIO_SERVER_ID=0dcbaca4-2a26-489c-b4a3-d2fad8bb6483")
	assert.Contains(t, result, "RCON_PASSWORD="+cfg.RCONPassword)
	assert.NotContains(t, result, "RESTIC_REPOSITORY=gs:"+cfg.BackupBucket+":/ \\")
	assert.NotContains(t, result, "keep-within 3m")
}

func TestRenderCloudConfig_DiskLayoutAndManifestMount(t *testing.T) {
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "0dcbaca4-2a26-489c-b4a3-d2fad8bb6483",
		BackupRetentionDays: 90,
		ResticPassword:      "deployment-wide-restic-password",
		RCONPassword:        "rcon-password",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	// The data disk now mounts at /mnt/disks/minecraft with world data in a
	// data/ subdirectory (InfraVersion 4 layout).
	assert.Contains(t, result, "[minecraft_data, /mnt/disks/minecraft, ext4, \"defaults\", \"0\", \"2\"]")
	assert.NotContains(t, result, "[minecraft_data, /mnt/disks/minecraft/data, ext4")

	// The manifest queue is a parallel directory on the same disk, shared by
	// the backup and machine-agent containers as /manifests.
	assert.Contains(t, result, "-v /mnt/disks/minecraft/backup-manifests:/manifests")
	assert.Contains(t, result, "mkdir -p /mnt/disks/minecraft/data /mnt/disks/minecraft/backup-manifests")
	assert.Contains(t, result, "chmod 0777 /mnt/disks/minecraft/backup-manifests")

	// Both backup and agent containers still mount the world data directory.
	assert.Contains(t, result, "-v /mnt/disks/minecraft/data:/data")
}

func TestRenderCloudConfig_MetioUserNoHardcodedUID(t *testing.T) {
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "0dcbaca4-2a26-489c-b4a3-d2fad8bb6483",
		BackupRetentionDays: 90,
		ResticPassword:      "deployment-wide-restic-password",
		RCONPassword:        "rcon-password",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	// UID 1000 collides with COS's `chronos` stub account; the user must be
	// auto-allocated so cloud-init can create it.
	assert.NotContains(t, result, "uid: 1000")

	// Only the one users entry must render; nothing else may pin a uid.
	var parsed map[string]interface{}
	assert.NoError(t, yaml.Unmarshal([]byte(result), &parsed))
	users, ok := parsed["users"].([]interface{})
	assert.True(t, ok, "users must be a list")
	assert.Len(t, users, 1)
	user, ok := users[0].(map[string]interface{})
	assert.True(t, ok, "user entry must be a map")
	assert.NotContains(t, user, "uid", "the metio user must not pin a UID")
}

func TestRenderCloudConfig_MetioUserGuardBeforeStart(t *testing.T) {
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "0dcbaca4-2a26-489c-b4a3-d2fad8bb6483",
		BackupRetentionDays: 90,
		ResticPassword:      "deployment-wide-restic-password",
		RCONPassword:        "rcon-password",
		BackupServiceEnable: "minecraft-backup ",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	// The id/useradd guard must run before systemctl starts any User=metio
	// unit so a failed user creation cannot silently take down all services.
	guardIdx := strings.Index(result, "id -u metio >/dev/null 2>&1 || useradd")
	startIdx := strings.Index(result, "systemctl start minecraft")
	assert.True(t, guardIdx >= 0, "runcmd must contain the metio user guard")
	assert.True(t, startIdx >= 0, "runcmd must start services")
	assert.True(t, guardIdx < startIdx, "metio user guard must precede service start")
}

func TestRenderCloudConfig_BackupOverrides(t *testing.T) {
	cfg := &TemplateConfig{
		Region:               "europe-west3",
		MachineAgentImage:    "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupImage:          "europe-west3-docker.pkg.dev/minecraftbyl/metio/mc-backup:1.0.0",
		BackupInterval:       "6h",
		PruneResticRetention: "--keep-last 3 --keep-daily 10",
		BackupServiceEnable:  "minecraft-backup ",
		RCONPassword:         "rcon-password",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.Contains(t, result, "docker pull europe-west3-docker.pkg.dev/minecraftbyl/metio/mc-backup:1.0.0")
	assert.Contains(t, result, "BACKUP_INTERVAL=\"6h\"")
	assert.Contains(t, result, "PRUNE_RESTIC_RETENTION=\"--keep-last 3 --keep-daily 10\"")
	assert.Contains(t, result, "systemctl enable minecraft minecraft-backup metio-machine-agent")
	assert.Contains(t, result, "systemctl start minecraft minecraft-backup metio-machine-agent")
}

func TestRenderCloudConfig_BackupDisabled(t *testing.T) {
	cfg := &TemplateConfig{
		Region:               "europe-west3",
		MachineAgentImage:    "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupImage:          "ghcr.io/itzg/mc-backup:latest",
		BackupInterval:       "1h",
		PruneResticRetention: "--keep-within 90d",
		BackupServiceEnable:  "",
		RCONPassword:         "rcon-password",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.Contains(t, result, "systemctl enable minecraft metio-machine-agent")
	assert.Contains(t, result, "systemctl start minecraft metio-machine-agent")
	assert.NotContains(t, result, "minecraft-backup metio-machine-agent")
}

func TestRenderCloudConfig_RestoreSnapshotOmitted(t *testing.T) {
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "new-server-id",
		BackupRetentionDays: 90,
		ResticPassword:      "restic-pw",
		RCONPassword:        "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.NotContains(t, result, "metio-restore")
	assert.NotContains(t, result, "restic restore")
	// Without a restore configured, Minecraft depends only on docker (the
	// restore unit lines must not be injected).
	assert.Contains(t, result, "After=docker.service\n      Requires=docker.service\n\n      [Service]")
	assert.NotContains(t, result, "After=metio-restore.service")
	assert.NotContains(t, result, "Requires=metio-restore.service")
}

func TestRenderCloudConfig_RestoreSnapshotPresent(t *testing.T) {
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "new-server-id",
		BackupRetentionDays: 90,
		ResticPassword:      "restic-pw",
		RCONPassword:        "rcon-pw",
		BackupImage:         "ghcr.io/itzg/mc-backup:latest",
		BackupServiceEnable: "minecraft-backup ",
		RestoreSnapshotID:   "abc123-snapshot",
		RestoreSourcePrefix: "servers/old-server-id/restic",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	// The restore is shipped as a one-shot systemd unit running as the metio
	// user (writable HOME, docker access) and gated on a run-once sentinel
	// keyed to the snapshot.
	assert.Contains(t, result, "path: /etc/systemd/system/metio-restore.service")
	assert.Contains(t, result, "Description=Metio One-Shot World Restore")
	assert.Contains(t, result, "User=metio")
	assert.Contains(t, result, "ConditionPathExists=!/mnt/disks/minecraft/.metio-restore-abc123-snapshot.done")
	assert.Contains(t, result, "docker run --rm --name metio-restore")
	assert.Contains(t, result, "docker-credential-gcr configure-docker --registries europe-west3-docker.pkg.dev")
	assert.Contains(t, result, "RESTIC_REPOSITORY=gs:my-project-development-backups:/servers/old-server-id/restic")
	assert.Contains(t, result, "RESTIC_PASSWORD=restic-pw")
	assert.Contains(t, result, "--entrypoint /usr/bin/restic ghcr.io/itzg/mc-backup:latest restore abc123-snapshot:/data --target /data")
	assert.Contains(t, result, "-v /mnt/disks/minecraft/data:/data")
	assert.Contains(t, result, "ExecStartPost=+/bin/touch /mnt/disks/minecraft/.metio-restore-abc123-snapshot.done")

	// Minecraft must not start against an empty world: the restore unit runs
	// before it via After=/Requires=, and fails (leaving Minecraft stopped) if
	// the restore errors. On success the sentinel skips the unit on later boots.
	restoreIdx := strings.Index(result, "metio-restore")
	startIdx := strings.Index(result, "systemctl start minecraft")
	assert.True(t, restoreIdx < startIdx, "restore unit must appear before runcmd systemctl start")
	assert.Contains(t, result, "Before=minecraft.service")
	assert.Contains(t, result, "After=docker.service\n      After=metio-restore.service\n      Requires=docker.service\n      Requires=metio-restore.service")
	assert.Contains(t, result, "systemctl start minecraft minecraft-backup metio-machine-agent")

	// The old marker-based runcmd guard is gone.
	assert.NotContains(t, result, ".metio-restore-failed")
}

func TestRenderCloudConfig_RestoreSnapshotOnly(t *testing.T) {
	// When only RestoreSnapshotID is set but RestoreSourcePrefix is empty,
	// no restore step should be rendered (both must be present).
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "machine-agent:latest",
		RestoreSnapshotID: "abc123",
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.NotContains(t, result, "metio-restore")
}

func TestRenderCloudConfig_RestoreSourcePrefixTrailingSlash(t *testing.T) {
	// Stored repository prefixes end in "/" (the wire format is validated by
	// the backup report handler), so the restore URL and start guard must
	// not render a doubled slash.
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupBucket:        "my-project-development-backups",
		ServerID:            "new-server-id",
		BackupRetentionDays: 90,
		ResticPassword:      "restic-pw",
		RCONPassword:        "rcon-pw",
		BackupImage:         "ghcr.io/itzg/mc-backup:latest",
		RestoreSnapshotID:   "abc123-snapshot",
		RestoreSourcePrefix: "servers/old-server-id/restic/",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.Contains(t, result, "RESTIC_REPOSITORY=gs:my-project-development-backups:/servers/old-server-id/restic")
	assert.NotContains(t, result, "restic//")
}

func TestRenderCloudConfig_NoRestoreEmitsPlainStart(t *testing.T) {
	// Without a restore, the start command must be byte-identical to the
	// historical form so the normal boot path is untouched.
	cfg := &TemplateConfig{
		Region:              "europe-west3",
		MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		BackupServiceEnable: "minecraft-backup ",
		RCONPassword:        "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.Contains(t, result, "systemctl start minecraft minecraft-backup metio-machine-agent")
	assert.NotContains(t, result, "metio-restore-failed")
}

func TestRenderCloudConfig_Java25Image(t *testing.T) {
	// The 26.x releases require Java 25, so the Minecraft image must be the
	// stable-java25 variant. Java 25 also runs older releases (1.21.x needs
	// Java 17/21+), so a single tag covers every offered version.
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		MinecraftVersion:  "26.1",
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)
	assert.Contains(t, result, "itzg/minecraft-server:stable-java25")
	assert.NotContains(t, result, "itzg/minecraft-server:stable-java21")
}

func TestRenderCloudConfig_MaxMemoryPercentage(t *testing.T) {
	// The JVM heap is sized as a percentage of host RAM so user-data encodes
	// no absolute or machine-specific value: paying for a larger machine must
	// yield a larger server, and a machine-type change must not rewrite
	// user-data (which would promote the resize to a recreate).
	mkCfg := func() *TemplateConfig {
		return &TemplateConfig{
			Region:              "europe-west3",
			MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
			BackupBucket:        "my-project-development-backups",
			ServerID:            "server-a",
			BackupRetentionDays: 90,
			ResticPassword:      "restic-pw",
			RCONPassword:        "rcon-pw",
			MinecraftVersion:    "26.1",
		}
	}

	other, err := RenderCloudConfig(mkCfg())
	assert.NoError(t, err)
	assert.Contains(t, other, "MAX_MEMORY=75%")
	assert.NotContains(t, other, "MAX_MEMORY=3G")
	// No absolute size (MAX_MEMORY=<n>G/M/K) must survive into user-data.
	assert.NotRegexp(t, regexp.MustCompile(`MAX_MEMORY=\d+[gGmMkK]`), other)

	// The rendered MAX_MEMORY token must be identical across servers so
	// user-data stays byte-identical across machine types: reject any
	// machine-derived (absolute) value in favor of a fixed percentage.
	second := mkCfg()
	second.RCONPassword = "different-rcon-pw"
	second.ServerID = "server-b"
	second.MinecraftVersion = "1.21.1"
	secondRender, err := RenderCloudConfig(second)
	assert.NoError(t, err)
	assert.Contains(t, secondRender, "MAX_MEMORY=75%")
}

func TestRenderCloudConfig_ModpackLatest(t *testing.T) {
	// A pack-driven server (ADR-0006) replaces the VERSION env with
	// MODRINTH_MODPACK; the pack manifest controls the Minecraft version, so
	// no version reference may leak through.
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		MinecraftVersion:  "",
		Modpack:           &ModpackConfig{Platform: "modrinth", ProjectID: "abC123"},
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	// TYPE=MODRINTH is what makes the image install the pack: with the default
	// TYPE=VANILLA it would silently boot an unmodded server.
	assert.Contains(t, result, "-e TYPE=MODRINTH \\")
	assert.Contains(t, result, "-e MODRINTH_MODPACK=abC123 \\")
	assert.NotContains(t, result, "MODRINTH_VERSION")
	assert.NotContains(t, result, "-e VERSION=")
}

func TestRenderCloudConfig_ModpackPinned(t *testing.T) {
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		MinecraftVersion:  "",
		Modpack:           &ModpackConfig{Platform: "modrinth", ProjectID: "abC123", VersionID: "XyZ789"},
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.Contains(t, result, "-e TYPE=MODRINTH \\")
	assert.Contains(t, result, "-e MODRINTH_MODPACK=abC123 \\")
	assert.Contains(t, result, "-e MODRINTH_VERSION=XyZ789 \\")
	assert.NotContains(t, result, "-e VERSION=")
}

func TestRenderCloudConfig_ModpackDisablesSparkProfiler(t *testing.T) {
	// spark 1.10.x's async-profiler segfaults on the pinned Java 25 image
	// (metio#631). Pack servers get a PRE_START_SCRIPT hook that writes the
	// spark config after the pack install and before the JVM starts.
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		MinecraftVersion:  "",
		Modpack:           &ModpackConfig{Platform: "modrinth", ProjectID: "abC123"},
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	// The hook file is written with the boot disk and mounted read-only into
	// the container; spark only reads its settings from the config file.
	assert.Contains(t, result, "- path: /opt/metio/pre-start/spark-profiler.sh")
	assert.Contains(t, result, "permissions: '0755'")
	assert.Contains(t, result, "/data/config/spark/config.json")
	assert.Contains(t, result, `printf '{ "backgroundProfiler": false }\n' > "$cfg"`)
	assert.Contains(t, result, "-e PRE_START_SCRIPT=/hooks/spark-profiler.sh \\")
	assert.Contains(t, result, "-v /opt/metio/pre-start:/hooks:ro \\")

	// Parse the rendered document: the hook must land as an executable
	// write_files entry, otherwise the image refuses to start the server.
	var parsed struct {
		WriteFiles []struct {
			Path        string `yaml:"path"`
			Permissions string `yaml:"permissions"`
			Content     string `yaml:"content"`
		} `yaml:"write_files"`
	}
	assert.NoError(t, yaml.Unmarshal([]byte(result), &parsed))
	found := false
	for _, wf := range parsed.WriteFiles {
		if wf.Path != "/opt/metio/pre-start/spark-profiler.sh" {
			continue
		}
		found = true
		assert.Equal(t, "0755", wf.Permissions)
		assert.True(t, strings.HasPrefix(wf.Content, "#!/bin/sh\n"), "hook needs a shebang")
		assert.Contains(t, wf.Content, `"backgroundProfiler": false`)
	}
	assert.True(t, found, "pack servers must write the spark hook")
}

func TestRenderCloudConfig_VanillaHasNoSparkProfilerHook(t *testing.T) {
	// Vanilla servers never bundle spark, so their user-data stays unchanged.
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		MinecraftVersion:  "1.21.1",
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.NotContains(t, result, "spark-profiler")
	assert.NotContains(t, result, "PRE_START_SCRIPT")
	assert.NotContains(t, result, "/opt/metio/pre-start")
}

func TestRenderCloudConfig_VanillaStillRendersVersion(t *testing.T) {
	cfg := &TemplateConfig{
		Region:            "europe-west3",
		MachineAgentImage: "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
		MinecraftVersion:  "1.21.1",
		RCONPassword:      "rcon-pw",
	}
	result, err := RenderCloudConfig(cfg)
	assert.NoError(t, err)

	assert.Contains(t, result, "-e VERSION=1.21.1 \\")
	assert.NotContains(t, result, "MODRINTH_MODPACK")
	assert.NotContains(t, result, "TYPE=MODRINTH")
}

func TestRenderCloudConfig_YAMLValid(t *testing.T) {
	// The restore and guarded-start entries are plain YAML scalars; a stray
	// indicator would silently corrupt the whole user-data document. Parse the
	// rendered output to guarantee the VM would receive valid cloud-config.
	tests := []*TemplateConfig{
		{
			Region:              "europe-west3",
			MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
			BackupBucket:        "my-project-development-backups",
			ServerID:            "new-server-id",
			BackupRetentionDays: 90,
			ResticPassword:      "restic-pw",
			RCONPassword:        "rcon-pw",
			BackupImage:         "ghcr.io/itzg/mc-backup:latest",
			BackupServiceEnable: "minecraft-backup ",
			RestoreSnapshotID:   "abc123-snapshot",
			RestoreSourcePrefix: "servers/old-server-id/restic/",
		},
		{
			Region:              "europe-west3",
			MachineAgentImage:   "europe-west3-docker.pkg.dev/minecraftbyl/metio/machine-agent:tag",
			BackupBucket:        "my-project-development-backups",
			ServerID:            "new-server-id",
			BackupRetentionDays: 90,
			ResticPassword:      "restic-pw",
			RCONPassword:        "rcon-pw",
			BackupImage:         "ghcr.io/itzg/mc-backup:latest",
			RestoreSnapshotID:   "abc123-snapshot",
			RestoreSourcePrefix: "servers/old-server-id/restic",
		},
		{
			Region:              "europe-west3",
			MachineAgentImage:   "machine-agent:latest",
			BackupServiceEnable: "minecraft-backup ",
			RCONPassword:        "rcon-pw",
		},
		{
			// Pack servers add an extra write_files entry (the spark hook),
			// whose block scalar must not break the surrounding YAML.
			Region:            "europe-west3",
			MachineAgentImage: "machine-agent:latest",
			Modpack:           &ModpackConfig{Platform: "modrinth", ProjectID: "abC123"},
			RCONPassword:      "rcon-pw",
		},
	}
	for _, cfg := range tests {
		result, err := RenderCloudConfig(cfg)
		assert.NoError(t, err)

		var parsed map[string]interface{}
		assert.NoError(t, yaml.Unmarshal([]byte(result), &parsed))
		runcmd, ok := parsed["runcmd"].([]interface{})
		assert.True(t, ok, "runcmd must be a list")
		assert.NotEmpty(t, runcmd)
	}
}
