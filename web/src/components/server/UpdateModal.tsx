import { useState } from 'react';
import { Loader2 } from 'lucide-react';
import type { ServerConfig, UpdateServerRequest } from '../../types/server';
import { Button } from '../ui/Button';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '../ui/Dialog';
import { Input } from '../ui/Input';
import { Label } from '../ui/Label';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '../ui/Select';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '../ui/Tabs';
import { useServerOptions } from '../../hooks/useServerOptions';
import { BackupSettingsPanel } from './BackupSettingsPanel';
import { ModpackSettingsPanel } from './ModpackSettingsPanel';
import { ModpackChangeConfirmDialog } from './ModpackChangeConfirmDialog';
import type { ModpackSelection } from '../../lib/modpack';

export interface UpdateModalProps {
  open: boolean;
  serverId: string;
  serverName: string;
  config: ServerConfig;
  currentInfraVersion: number;
  outdated: boolean;
  onClose: () => void;
  onUpdate: (data: UpdateServerRequest) => void;
  isPending: boolean;
  /** Modpack changes recreate the VM, which the backend only does while running. */
  isRunning?: boolean;
}

type SettingsTab = 'settings' | 'modpack' | 'backup';

/** The staged selection for a server that already has a pack installed. */
function selectionFromConfig(config: ServerConfig): ModpackSelection | null {
  if (!config.modpack) return null;
  return {
    platform: 'modrinth',
    projectId: config.modpack.projectId,
    versionId: config.modpack.versionId,
    name: '',
  };
}

export function UpdateModal({
  open,
  serverId,
  serverName,
  config,
  currentInfraVersion,
  outdated,
  onClose,
  onUpdate,
  isPending,
  isRunning = true,
}: UpdateModalProps) {
  const [name, setName] = useState(config.name);
  const [machineType, setMachineType] = useState(config.machineType);
  const [diskSizeGB, setDiskSizeGB] = useState(config.diskSizeGB);
  const [minecraftVersion, setMinecraftVersion] = useState(
    config.minecraftVersion
  );
  const [packSelection, setPackSelection] = useState<ModpackSelection | null>(
    () => selectionFromConfig(config)
  );
  const [activeTab, setActiveTab] = useState<SettingsTab>('settings');
  const [confirmOpen, setConfirmOpen] = useState(false);
  const [pendingUpdate, setPendingUpdate] =
    useState<UpdateServerRequest | null>(null);
  const { data: options, isLoading: optionsLoading } = useServerOptions();

  // Always offer the current version so opening the modal cannot silently
  // switch the server when the API no longer lists that version. Pack-driven
  // servers store no version, so only real versions are offered.
  const availableVersions = options?.minecraftVersions ?? [];
  const versionOptions = availableVersions.includes(config.minecraftVersion)
    ? availableVersions
    : [config.minecraftVersion, ...availableVersions].filter(Boolean);

  // Keep the current machine type selectable when the dynamic list changes.
  const availableMachineTypes = options?.machineTypes ?? [];
  const machineTypeOptions = (() => {
    const withSpecs = availableMachineTypes.map((mt) => ({
      id: mt.id,
      label: `${mt.id} (${mt.vcpus} vCPU · ${mt.memoryGB} GB RAM)`,
    }));
    if (withSpecs.some((mt) => mt.id === config.machineType)) return withSpecs;
    return [
      { id: config.machineType, label: config.machineType },
      ...withSpecs,
    ];
  })();

  const installed = config.modpack;
  const packChanged = (() => {
    if (!installed && !packSelection) return false;
    if (!installed || !packSelection) return true;
    return (
      installed.projectId !== packSelection.projectId ||
      (installed.versionId ?? '') !== (packSelection.versionId ?? '')
    );
  })();

  // A pack controls the Minecraft version, so its field only counts as a change
  // while the server is (or becomes) vanilla.
  const versionChanged =
    !packSelection && minecraftVersion !== config.minecraftVersion;
  const removalMissingVersion = !packSelection && !minecraftVersion;

  const hasChanges =
    name !== config.name ||
    machineType !== config.machineType ||
    diskSizeGB !== config.diskSizeGB ||
    versionChanged ||
    packChanged;

  const buildUpdate = (): UpdateServerRequest => {
    const data: UpdateServerRequest = {};
    if (name !== config.name) data.name = name;
    if (machineType !== config.machineType) data.machineType = machineType;
    if (diskSizeGB !== config.diskSizeGB) data.diskSizeGB = diskSizeGB;

    if (packChanged) {
      // Adding or changing a pack makes the version pack-controlled, so it is
      // never sent alongside a pack. Removing one requires its replacement.
      if (packSelection) {
        data.modpack = {
          platform: packSelection.platform,
          projectId: packSelection.projectId,
          versionId: packSelection.versionId,
        };
      } else {
        data.modpack = null;
        data.minecraftVersion = minecraftVersion;
      }
    } else if (versionChanged) {
      data.minecraftVersion = minecraftVersion;
    }

    return data;
  };

  const handleSubmit = (e?: React.FormEvent) => {
    e?.preventDefault();
    if (removalMissingVersion) return;

    const data = buildUpdate();
    if (packChanged) {
      // Any pack change recreates the VM on the current runtime, so it is
      // confirmed explicitly before it is submitted.
      setPendingUpdate(data);
      setConfirmOpen(true);
      return;
    }

    onUpdate(data);
  };

  const confirmPackChange = () => {
    if (pendingUpdate) onUpdate(pendingUpdate);
    setPendingUpdate(null);
    setConfirmOpen(false);
  };

  const cancelPackChange = () => {
    setPendingUpdate(null);
    setConfirmOpen(false);
  };

  const footer = (
    <DialogFooter className="mx-0 mb-0 rounded-none border-0 bg-transparent p-0">
      <DialogClose asChild>
        <Button type="button" variant="outline" disabled={isPending}>
          Cancel
        </Button>
      </DialogClose>
      <Button
        type="button"
        onClick={() => handleSubmit()}
        disabled={
          (!hasChanges && !outdated) || isPending || removalMissingVersion
        }
      >
        {isPending && <Loader2 className="animate-spin" />}
        Update Server
      </Button>
    </DialogFooter>
  );

  return (
    <>
      <Dialog
        open={open}
        onOpenChange={(nextOpen) => {
          if (!nextOpen) onClose();
        }}
      >
        <DialogContent className="max-h-[85vh] max-w-2xl overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Server Settings</DialogTitle>
            <DialogDescription className="sr-only">
              Update settings for {serverName}.
            </DialogDescription>
          </DialogHeader>

          {outdated && (
            <div className="rounded-lg border border-yellow-500/30 bg-yellow-500/10 p-3 -mt-2 space-y-1">
              <p className="text-sm font-medium text-yellow-400">
                Infrastructure Update Available
              </p>
              <p className="text-xs text-yellow-300/70">
                Server version: v{config.infraVersion ?? '?'} → Controller
                version: v{currentInfraVersion}
              </p>
              <p className="text-xs text-muted-foreground">
                The server will continue running during the update. No
                configuration changes are required — submitting will trigger the
                upgrade.
              </p>
            </div>
          )}

          <Tabs
            value={activeTab}
            onValueChange={(value) => setActiveTab(value as SettingsTab)}
          >
            <TabsList variant="line">
              <TabsTrigger value="settings">Settings</TabsTrigger>
              <TabsTrigger value="modpack">Modpack</TabsTrigger>
              <TabsTrigger value="backup">Backup</TabsTrigger>
            </TabsList>
            <TabsContent value="settings" className="py-2">
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="server-name">Server Name</Label>
                  <Input
                    id="server-name"
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                  />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="machine-type">Machine Type</Label>
                  <Select value={machineType} onValueChange={setMachineType}>
                    <SelectTrigger
                      id="machine-type"
                      className="w-full"
                      disabled={optionsLoading}
                    >
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      {machineTypeOptions.map((mt) => (
                        <SelectItem key={mt.id} value={mt.id}>
                          {mt.label}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <p className="text-xs text-muted-foreground">
                    Changing the machine type will stop and restart the server.
                  </p>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="disk-size">Disk Size (GB)</Label>
                  <Input
                    id="disk-size"
                    type="number"
                    value={diskSizeGB}
                    onChange={(e) => setDiskSizeGB(Number(e.target.value))}
                    min={20}
                  />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="minecraft-version">Minecraft Version</Label>
                  {packSelection ? (
                    <div
                      id="minecraft-version"
                      className="flex h-9 w-full items-center rounded-md border border-input bg-muted px-3 text-sm text-muted-foreground"
                    >
                      Controlled by modpack
                    </div>
                  ) : (
                    <Select
                      value={minecraftVersion}
                      onValueChange={setMinecraftVersion}
                    >
                      <SelectTrigger
                        id="minecraft-version"
                        className="w-full"
                        disabled={optionsLoading}
                      >
                        <SelectValue placeholder="Select a version" />
                      </SelectTrigger>
                      <SelectContent>
                        {versionOptions.map((version) => (
                          <SelectItem key={version} value={version}>
                            {version}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  )}
                  <p className="text-xs text-muted-foreground">
                    {packSelection
                      ? 'The selected modpack chooses the loader and Minecraft version.'
                      : 'Changing the Minecraft version rebuilds the server.'}
                  </p>
                </div>

                <DialogFooter className="mx-0 mb-0 rounded-none border-0 bg-transparent p-0">
                  <DialogClose asChild>
                    <Button
                      type="button"
                      variant="outline"
                      disabled={isPending}
                    >
                      Cancel
                    </Button>
                  </DialogClose>
                  <Button
                    type="submit"
                    disabled={
                      (!hasChanges && !outdated) ||
                      isPending ||
                      removalMissingVersion
                    }
                  >
                    {isPending && <Loader2 className="animate-spin" />}
                    Update Server
                  </Button>
                </DialogFooter>
              </form>
            </TabsContent>
            <TabsContent value="modpack" className="py-2">
              <div className="space-y-4">
                <ModpackSettingsPanel
                  installed={installed}
                  selection={packSelection}
                  onSelectionChange={setPackSelection}
                  minecraftVersion={minecraftVersion}
                  onMinecraftVersionChange={setMinecraftVersion}
                  minecraftVersions={availableVersions}
                  disabled={isPending}
                  isRunning={isRunning}
                />
                {footer}
              </div>
            </TabsContent>
            <TabsContent value="backup" className="py-2">
              <BackupSettingsPanel
                serverId={serverId}
                serverName={serverName}
              />
            </TabsContent>
          </Tabs>
        </DialogContent>
      </Dialog>

      <ModpackChangeConfirmDialog
        open={confirmOpen}
        serverName={serverName}
        installed={installed}
        next={packSelection}
        nextMinecraftVersion={minecraftVersion}
        isPending={isPending}
        onCancel={cancelPackChange}
        onConfirm={confirmPackChange}
      />
    </>
  );
}
