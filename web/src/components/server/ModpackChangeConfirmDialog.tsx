import { AlertTriangle, Loader2, RefreshCw } from 'lucide-react';
import { useModrinthPack } from '../../hooks/useModrinthPack';
import { useModrinthPackVersions } from '../../hooks/useModrinthPackVersions';
import { Button } from '../ui/Button';
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '../ui/AlertDialog';
import type { ModpackConfig } from '../../types/server';
import type { ModpackSelection } from '../../lib/modpack';

export interface ModpackChangeConfirmDialogProps {
  open: boolean;
  serverName: string;
  /** The pack the server has installed right now, if any. */
  installed: ModpackConfig | undefined;
  /** The pack the server will have after the change, null for vanilla. */
  next: ModpackSelection | null;
  /** Minecraft version the server will run once the pack is removed. */
  nextMinecraftVersion?: string;
  isPending: boolean;
  onCancel: () => void;
  onConfirm: () => void;
}

/** Label for a pack, resolved from the catalogue when possible. */
function useInstalledLabel(
  installed: ModpackConfig | undefined,
  enabled: boolean
) {
  const projectId = enabled ? (installed?.projectId ?? null) : null;
  const pack = useModrinthPack(projectId);
  const versions = useModrinthPackVersions(projectId);

  if (!installed) return 'Vanilla';

  const name = pack.data?.name ?? installed.projectId;
  if (!installed.versionId) return `${name} (Latest)`;

  const version = versions.data?.find((v) => v.id === installed.versionId);
  return `${name} (${version?.name ?? version?.versionNumber ?? installed.versionId})`;
}

/**
 * Confirmation gate for modpack changes on an existing server. Any pack change
 * recreates the VM on the current runtime, so the rebuild is spelled out before
 * the update is submitted (see ADR-0006 and the note on #538).
 */
export function ModpackChangeConfirmDialog({
  open,
  serverName,
  installed,
  next,
  nextMinecraftVersion,
  isPending,
  onCancel,
  onConfirm,
}: ModpackChangeConfirmDialogProps) {
  const currentLabel = useInstalledLabel(installed, open);
  const nextLabel = next
    ? `${next.versionName || next.name || next.projectId}${next.versionId ? '' : ' (Latest)'}`
    : `Vanilla${nextMinecraftVersion ? ` (${nextMinecraftVersion})` : ''}`;

  return (
    <AlertDialog
      open={open}
      onOpenChange={(nextOpen) => {
        if (!nextOpen && !isPending) onCancel();
      }}
    >
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle className="flex items-center gap-2">
            <AlertTriangle className="h-5 w-5 text-destructive" />
            Rebuild server to change the modpack
          </AlertDialogTitle>
          <AlertDialogDescription>
            Changing the modpack recreates the machine behind{' '}
            <strong className="text-foreground">{serverName}</strong>. This
            usually takes a few minutes, and the server is unavailable while it
            is rebuilt and restarted.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <div className="space-y-3">
          <div className="rounded-lg border border-border bg-muted p-3 text-sm">
            <span className="text-muted-foreground">Modpack: </span>
            <span className="font-medium text-foreground">
              {currentLabel} → {nextLabel}
            </span>
          </div>

          <p className="text-sm text-muted-foreground">
            Your world and player data live on the persistent data disk and are
            kept. The server is saved before the rebuild starts.
          </p>
        </div>

        <AlertDialogFooter>
          <AlertDialogCancel asChild>
            <Button type="button" variant="outline" disabled={isPending}>
              Cancel
            </Button>
          </AlertDialogCancel>
          <Button
            type="button"
            variant="destructive"
            onClick={onConfirm}
            disabled={isPending}
          >
            {isPending ? <Loader2 className="animate-spin" /> : <RefreshCw />}
            Rebuild server
          </Button>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
