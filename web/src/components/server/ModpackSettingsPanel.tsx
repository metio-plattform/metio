import { useState } from 'react';
import { AlertTriangle, Package, RefreshCw } from 'lucide-react';
import { useModrinthPack } from '../../hooks/useModrinthPack';
import { useModrinthPackVersions } from '../../hooks/useModrinthPackVersions';
import { Badge } from '../ui/Badge';
import { Button } from '../ui/Button';
import { Label } from '../ui/Label';
import { Skeleton } from '../ui/Skeleton';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '../ui/Select';
import { ModpackPicker, PackIcon } from './ModpackPicker';
import type { ModpackSelection } from '../../lib/modpack';
import type { ModpackConfig, ModrinthVersion } from '../../types/server';

export interface ModpackSettingsPanelProps {
  /** The pack the server currently has installed, if any. */
  installed: ModpackConfig | undefined;
  /** The staged selection: null means the server will be vanilla. */
  selection: ModpackSelection | null;
  onSelectionChange: (selection: ModpackSelection | null) => void;
  /** Minecraft version to use when a pack is removed (vanilla is required). */
  minecraftVersion: string;
  onMinecraftVersionChange: (version: string) => void;
  minecraftVersions: string[];
  /** True while an update is submitting. */
  disabled?: boolean;
  isRunning: boolean;
}

/** Human-readable label for a pinned pack version. */
function versionLabel(
  versionId: string | undefined,
  versions: ModrinthVersion[] | undefined
): string {
  if (!versionId) return 'Latest';
  const match = versions?.find((version) => version.id === versionId);
  return match?.name ?? match?.versionNumber ?? versionId;
}

/**
 * Modpack tab of the server update modal. Shows the installed pack and lets the
 * user add, change the version of, change, or remove it. Every such change
 * recreates the server, so this panel only stages the selection — the modal
 * asks for explicit confirmation before submitting.
 */
export function ModpackSettingsPanel({
  installed,
  selection,
  onSelectionChange,
  minecraftVersion,
  onMinecraftVersionChange,
  minecraftVersions,
  disabled = false,
  isRunning,
}: ModpackSettingsPanelProps) {
  const [picking, setPicking] = useState(false);

  const pack = useModrinthPack(installed?.projectId ?? null);
  const versions = useModrinthPackVersions(installed?.projectId ?? null);

  const installedPack = pack.data;
  const installedName = installedPack?.name ?? installed?.projectId ?? '';
  const versionsList = versions.data;
  const isRemoving = Boolean(installed) && selection === null;
  const isChangingPack =
    Boolean(installed) &&
    selection !== null &&
    selection.projectId !== installed?.projectId;
  const showPicker = picking || !installed || isChangingPack;

  const controlsDisabled = disabled || !isRunning;

  const startRemoving = () => {
    onSelectionChange(null);
    // A vanilla server must name a version, so preselect the newest one rather
    // than leaving the removal invalid.
    if (!minecraftVersion && minecraftVersions.length > 0) {
      onMinecraftVersionChange(minecraftVersions[0]);
    }
  };

  const installedSelection = (): ModpackSelection | null => {
    if (!installed) return null;
    return {
      platform: 'modrinth',
      projectId: installed.projectId,
      versionId: installed.versionId,
      name: installedName,
      iconUrl: installedPack?.iconUrl,
      loader: installedPack?.loader,
    };
  };

  /** Cancels any staged change and returns to the installed pack. */
  const restoreInstalled = () => {
    onSelectionChange(installedSelection());
    // Pack-driven servers keep an empty Minecraft version (ADR-0006); clearing
    // it undoes a removal that was staged and then abandoned.
    onMinecraftVersionChange('');
    setPicking(false);
  };

  /** Opens the picker with the installed pack preselected. */
  const startPicking = () => {
    onSelectionChange(installedSelection());
    onMinecraftVersionChange('');
    setPicking(true);
  };

  const selectVersion = (value: string) => {
    if (!selection) return;
    const match =
      value === 'latest'
        ? undefined
        : versionsList?.find((version) => version.id === value);
    onSelectionChange({
      ...selection,
      versionId: value === 'latest' ? undefined : value,
      versionName: match?.name ?? match?.versionNumber,
    });
  };

  return (
    <div className="space-y-4">
      {!isRunning && (
        <div className="flex items-start gap-3 rounded-lg border border-yellow-500/30 bg-yellow-500/10 p-3">
          <AlertTriangle className="mt-0.5 h-4 w-4 shrink-0 text-yellow-400" />
          <p className="text-xs text-muted-foreground">
            Changing the modpack rebuilds the server, which requires it to be
            running. Start the server before changing its modpack.
          </p>
        </div>
      )}

      {installed ? (
        <div className="space-y-3 rounded-lg border border-border bg-muted/40 p-4">
          <div className="flex items-start justify-between gap-4">
            <div className="flex min-w-0 items-center gap-3">
              {pack.isLoading ? (
                <Skeleton className="h-10 w-10 rounded-md" />
              ) : (
                <PackIcon iconUrl={installedPack?.iconUrl ?? ''} />
              )}
              <div className="min-w-0">
                <div className="flex items-center gap-2">
                  <span className="truncate font-medium text-foreground">
                    {installedName}
                  </span>
                  <Badge variant="secondary">Modrinth</Badge>
                </div>
                <div className="text-sm text-muted-foreground">
                  {versions.isLoading ? (
                    <Skeleton className="mt-1 h-4 w-32" />
                  ) : (
                    <>
                      Installed version:{' '}
                      {versionLabel(installed.versionId, versionsList)}
                    </>
                  )}
                </div>
              </div>
            </div>
          </div>

          {pack.isError ? (
            <p className="text-xs text-muted-foreground">
              Could not load the pack's details from Modrinth. The installed
              pack is shown by its project id.
            </p>
          ) : null}

          {selection && !isRemoving && !picking ? (
            <div className="space-y-2">
              <Label htmlFor="installed-pack-version">Pack Version</Label>
              <Select
                value={selection.versionId ?? 'latest'}
                onValueChange={selectVersion}
                disabled={controlsDisabled || versions.isLoading}
              >
                <SelectTrigger
                  id="installed-pack-version"
                  className="w-full sm:w-72"
                >
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="latest">Latest</SelectItem>
                  {(versionsList ?? []).map((version) => (
                    <SelectItem key={version.id} value={version.id}>
                      {version.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          ) : null}

          <div className="flex flex-wrap gap-2">
            {picking ? (
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={() => {
                  restoreInstalled();
                }}
                disabled={controlsDisabled}
              >
                Keep current pack
              </Button>
            ) : (
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={startPicking}
                disabled={controlsDisabled}
              >
                <RefreshCw className="h-4 w-4" />
                Change pack
              </Button>
            )}
            {!isRemoving ? (
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={startRemoving}
                disabled={controlsDisabled}
              >
                Remove pack
              </Button>
            ) : null}
          </div>
        </div>
      ) : (
        <div className="flex items-center gap-3 rounded-lg border border-border bg-muted/40 p-4">
          <Package className="h-5 w-5 text-muted-foreground" />
          <div>
            <p className="font-medium text-foreground">Vanilla Minecraft</p>
            <p className="text-sm text-muted-foreground">
              This server runs no modpack. Pick one below to install it.
            </p>
          </div>
        </div>
      )}

      {isRemoving ? (
        <div className="space-y-4 rounded-lg border border-destructive/40 bg-destructive/10 p-4">
          <div className="flex items-start gap-3">
            <AlertTriangle className="mt-0.5 h-4 w-4 shrink-0 text-destructive" />
            <p className="text-sm text-foreground">
              Removing <strong>{installedName}</strong> turns this into a
              vanilla server. Choose the Minecraft version it should run.
            </p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="removal-minecraft-version">Minecraft Version</Label>
            <Select
              value={minecraftVersion}
              onValueChange={onMinecraftVersionChange}
              disabled={controlsDisabled}
            >
              <SelectTrigger
                id="removal-minecraft-version"
                className="w-full sm:w-72"
              >
                <SelectValue placeholder="Select a version" />
              </SelectTrigger>
              <SelectContent>
                {minecraftVersions.map((version) => (
                  <SelectItem key={version} value={version}>
                    {version}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <Button
            type="button"
            variant="outline"
            size="sm"
            onClick={restoreInstalled}
            disabled={controlsDisabled}
          >
            Keep current pack
          </Button>
        </div>
      ) : null}

      {showPicker ? (
        <ModpackPicker
          value={selection}
          onChange={(next) => {
            onSelectionChange(next);
            if (!next) setPicking(false);
            // A pack controls the Minecraft version, so a version kept for a
            // staged removal must not survive a new pick.
            if (next) onMinecraftVersionChange('');
          }}
          disabled={controlsDisabled}
        />
      ) : null}
    </div>
  );
}
