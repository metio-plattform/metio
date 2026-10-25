package dbtypes

import (
	"fmt"
	"time"
)

type ServerState string

const (
	ServerStateStopped  ServerState = "STOPPED"
	ServerStateStarting ServerState = "STARTING"
	ServerStateRunning  ServerState = "RUNNING"
	ServerStateStopping ServerState = "STOPPING"
)

func (s ServerState) String() string {
	return string(s)
}

func (s ServerState) IsRunning() bool {
	return s == ServerStateRunning
}

func (s ServerState) IsStopped() bool {
	return s == ServerStateStopped
}

func (s ServerState) IsTransitioning() bool {
	return s == ServerStateStarting || s == ServerStateStopping
}

type Players struct {
	Current int `json:"current"`
	Max     int `json:"max"`
}

type Status struct {
	Players              Players           `json:"players"`
	Timestamp            time.Time         `json:"timestamp"`
	Uptime               string            `json:"uptime"`
	ServerState          ServerState       `json:"server_state"`
	InstanceIP           string            `json:"instance_ip"`
	Version              string            `json:"version"`
	WhitelistEnabled     bool              `json:"whitelist_enabled"`
	ScheduledShutdown    *time.Time        `json:"scheduled_shutdown,omitempty"`
	PendingCommand       string            `json:"pending_command,omitempty"`
	PendingCommandArgs   map[string]string `json:"pending_command_args,omitempty"`
	PendingCommandResult string            `json:"pending_command_result,omitempty"`
	AgentVersion         string            `json:"agent_version,omitempty"`

	// The machine-agent cannot reach the host's systemd from its container, so
	// restart detection is based on the Docker container identity. The agent
	// compares StartedAt with the previous status document and persists the
	// observed count across its own restarts. BootID resets the counter when
	// the VM reboots, and explicit STARTING/STOPPING transitions reset it for
	// intentional lifecycle actions.
	ContainerState     string     `json:"container_state,omitempty"`
	ContainerStartedAt *time.Time `json:"container_started_at,omitempty"`
	RestartCount       int        `json:"restart_count,omitempty"`
	LastRestartAt      *time.Time `json:"last_restart_at,omitempty"`
	BootID             string     `json:"boot_id,omitempty"`
}

const (
	// CrashLoopRestartThreshold is the number of observed container restarts
	// that, combined with an unanswered RCON, indicates a crash loop rather
	// than a one-off restart.
	CrashLoopRestartThreshold = 2
	// CrashLoopWindow keeps old restarts from keeping a server marked as
	// degraded forever.
	CrashLoopWindow = 5 * time.Minute
)

// IsCrashLooping reports whether the status is consistent with a server whose
// container keeps restarting without ever becoming reachable over RCON.
func (s Status) IsCrashLooping(now time.Time) bool {
	if s.RestartCount < CrashLoopRestartThreshold || s.LastRestartAt == nil {
		return false
	}
	if now.Sub(*s.LastRestartAt) > CrashLoopWindow {
		return false
	}
	if s.Players.Max > 0 {
		return false
	}
	return s.Version == "" || s.Version == "Unknown"
}

// CrashLoopMessage returns the user-facing warning for a crash-looping server,
// or an empty string when the status is healthy enough.
func (s Status) CrashLoopMessage(now time.Time) string {
	if !s.IsCrashLooping(now) {
		return ""
	}
	return fmt.Sprintf(
		"Minecraft restarted %d times in the last %d minutes and has not answered RCON. The server is probably crash-looping; check the server log.",
		s.RestartCount,
		int(CrashLoopWindow.Minutes()),
	)
}

type WhitelistEntry struct {
	Username string    `json:"username"`
	UUID     string    `json:"uuid"`
	AddedAt  time.Time `json:"added_at"`
	AddedBy  string    `json:"added_by"`
}

type WhitelistConfig struct {
	Enabled bool `json:"enabled"`
}

type BackupStatus string

const (
	BackupStatusCompleted BackupStatus = "COMPLETED"
	BackupStatusFailed    BackupStatus = "FAILED"
)

func (b BackupStatus) String() string {
	return string(b)
}

// ModrinthPlatform is the only valid ModpackConfig platform today. It is
// retained as a named constant because the field exists to make a future
// second pack source an additive change.
const ModrinthPlatform = "modrinth"

// ModpackConfig references a modpack a server boots with (ADR-0006). A server
// is either vanilla (Modpack nil) or pack-driven: when a pack is set, the
// Minecraft version is pack-controlled and stored empty. Platform is retained
// even though only "modrinth" is valid today so a second pack source later is
// an additive change rather than a schema migration.
type ModpackConfig struct {
	// Platform is the pack source, e.g. "modrinth".
	Platform string `json:"platform"`
	// ProjectID identifies the pack on the platform.
	ProjectID string `json:"projectId"`
	// VersionID pins a specific version of the pack; empty means latest.
	VersionID string `json:"versionId,omitempty"`
}

type BackupSourceConfig struct {
	Region           string         `json:"region"`
	Zone             string         `json:"zone"`
	MachineType      string         `json:"machine_type"`
	DiskSizeGB       int            `json:"disk_size_gb"`
	MinecraftVersion string         `json:"minecraft_version"`
	Modpack          *ModpackConfig `json:"modpack,omitempty"`
}

type Backup struct {
	ID               string              `json:"id"`
	ServerID         string              `json:"server_id"`
	ServerName       string              `json:"server_name"`
	SnapshotID       string              `json:"snapshot_id"`
	RepositoryPrefix string              `json:"repository_prefix"`
	CreatedAt        time.Time           `json:"created_at"`
	DurationSeconds  int64               `json:"duration_seconds"`
	FileCount        int64               `json:"file_count"`
	RepositorySize   int64               `json:"repository_size"`
	MinecraftVersion string              `json:"minecraft_version"`
	Status           BackupStatus        `json:"status"`
	ServerDeletedAt  *time.Time          `json:"server_deleted_at,omitempty"`
	RetentionUntil   *time.Time          `json:"retention_until,omitempty"`
	SourceConfig     *BackupSourceConfig `json:"source_config,omitempty"`
}
