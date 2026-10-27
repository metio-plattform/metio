// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true
// ModpackConfig references a modpack a server boots with (ADR-0006). A server
// is either vanilla (Modpack nil) or pack-driven: when a pack is set, the
// Minecraft version is pack-controlled and stored empty.
type ModpackConfig struct {
	// Platform is the pack source, e.g. "modrinth".
	// +kubebuilder:validation:Enum=modrinth
	Platform string `json:"platform"`

	// ProjectID identifies the pack on the platform.
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9]+$`
	// +kubebuilder:validation:MinLength=1
	ProjectID string `json:"projectId"`

	// VersionID pins a specific version of the pack; empty means latest.
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9]+$`
	VersionID string `json:"versionId,omitempty"`
}

// +kubebuilder:object:generate=true
// BackupConfig holds a per-server override for the backup schedule and Restic
// retention policy. Zero/empty values fall back to the deployment defaults.
type BackupConfig struct {
	// Enabled controls whether backups run for this server.
	Enabled bool `json:"enabled"`

	// BackupIntervalHours is the backup interval in hours.
	// +kubebuilder:validation:Minimum=0
	BackupIntervalHours int `json:"backupIntervalHours,omitempty"`

	// Keep is the number of backups to retain.
	// +kubebuilder:validation:Minimum=0
	Keep int `json:"keep,omitempty"`

	// KeepUnit is the retention unit for Keep.
	// +kubebuilder:validation:Enum=hourly;daily;weekly;monthly;yearly
	KeepUnit string `json:"keepUnit,omitempty"`
}

// +kubebuilder:object:generate=true
// ShutdownSchedule describes a recurring daily shutdown time in UTC.
// +kubebuilder:validation:XValidation:rule="!self.enabled || self.time != ''",message="time is required when enabled"
type ShutdownSchedule struct {
	// Enabled controls whether the shutdown schedule is active.
	Enabled bool `json:"enabled"`

	// Time is the daily shutdown time in UTC, in HH:MM format.
	// +kubebuilder:validation:Pattern=`^([01]?[0-9]|2[0-3]):[0-5][0-9]$`
	Time string `json:"time"`
}

// +kubebuilder:object:generate=true
// WhitelistConfig controls the server whitelist.
type WhitelistConfig struct {
	// Enabled controls whether the whitelist is enforced.
	Enabled bool `json:"enabled"`

	// Players is the list of whitelisted players.
	Players []WhitelistEntry `json:"players,omitempty"`
}

// +kubebuilder:object:generate=true
// WhitelistEntry identifies a single whitelisted player.
type WhitelistEntry struct {
	// Username is the player's Minecraft username.
	// +kubebuilder:validation:MinLength=1
	Username string `json:"username"`

	// UUID is the player's Mojang UUID.
	UUID string `json:"uuid"`
}

// +kubebuilder:object:generate=true
// MinecraftServerSpec defines the desired state of a Minecraft server.
// +kubebuilder:validation:XValidation:rule="(has(self.modpack) && !has(self.minecraftVersion)) || (!has(self.modpack) && has(self.minecraftVersion) && self.minecraftVersion != '')",message="exactly one of minecraftVersion or modpack must be set"
type MinecraftServerSpec struct {
	// ServerID is the unique identifier of the server.
	// +kubebuilder:validation:MinLength=1
	ServerID string `json:"serverID"`

	// MinecraftVersion is the vanilla Minecraft version to run.
	MinecraftVersion string `json:"minecraftVersion,omitempty"`

	// Modpack references a modpack to boot with.
	Modpack *ModpackConfig `json:"modpack,omitempty"`

	// Backup holds per-server backup overrides.
	Backup *BackupConfig `json:"backup,omitempty"`

	// ShutdownSchedule describes a recurring daily shutdown time in UTC.
	ShutdownSchedule *ShutdownSchedule `json:"shutdownSchedule,omitempty"`

	// Whitelist controls the server whitelist.
	Whitelist *WhitelistConfig `json:"whitelist,omitempty"`
}

// MinecraftServerStatus defines the observed state of a Minecraft server.
type MinecraftServerStatus struct {
	// Ready indicates whether the server is ready.
	Ready bool `json:"ready,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=mcs,scope=Namespaced

// MinecraftServer is the Schema for the minecraftservers API.
type MinecraftServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinecraftServerSpec   `json:"spec"`
	Status MinecraftServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MinecraftServerList contains a list of MinecraftServer.
type MinecraftServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MinecraftServer `json:"items"`
}
