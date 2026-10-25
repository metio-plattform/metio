/**
 * A modpack selection as edited in the UI. Display fields (name, iconUrl,
 * versionName) are carried for summaries; only platform/projectId/versionId are
 * sent to the API. An empty versionId installs the pack's latest version.
 */
export interface ModpackSelection {
  platform: 'modrinth';
  projectId: string;
  name: string;
  iconUrl?: string;
  loader?: string;
  versionId?: string;
  versionName?: string;
}

/** Shortest query the modpack search accepts. */
export const SEARCH_MIN_LENGTH = 3;

/** Human-readable summary of a pack selection. */
export function modpackSummary(selection: ModpackSelection): string {
  const version = selection.versionName ?? selection.versionId;
  return version
    ? `${selection.name} · ${version}`
    : `${selection.name} · Latest`;
}
