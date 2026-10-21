import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { UpdateModal } from './UpdateModal';
import type { ServerConfig } from '../../types/server';
import {
  useBackupSettings,
  useUpdateBackupSettings,
} from '../../hooks/useBackupSettings';
import { useServerOptions } from '../../hooks/useServerOptions';
import { useModrinthPack } from '../../hooks/useModrinthPack';
import { useModrinthPackVersions } from '../../hooks/useModrinthPackVersions';
import { useModrinthSearch } from '../../hooks/useModrinthSearch';

vi.mock('react-router-dom', () => ({
  useNavigate: () => vi.fn(),
}));

vi.mock('../../hooks/useServerOptions', () => ({
  useServerOptions: vi.fn(),
}));

vi.mock('../../hooks/useBackupSettings', () => ({
  useBackupSettings: vi.fn(),
  useUpdateBackupSettings: vi.fn(),
}));

vi.mock('../../hooks/useModrinthPack', () => ({
  useModrinthPack: vi.fn(),
}));

vi.mock('../../hooks/useModrinthPackVersions', () => ({
  useModrinthPackVersions: vi.fn(),
}));

// The picker searches through this hook; tests provide results directly.
vi.mock('../../hooks/useModrinthSearch', () => ({
  useModrinthSearch: vi.fn(),
}));

const config: ServerConfig = {
  name: 'test-server',
  region: 'us-central1',
  zone: 'us-central1-a',
  machineType: 'e2-small',
  minecraftVersion: '1.21.11',
  diskSizeGB: 50,
} as ServerConfig;

const machineTypes = [
  { id: 'e2-small', vcpus: 2, memoryGB: 2 },
  { id: 'e2-medium', vcpus: 2, memoryGB: 4 },
];

function renderModal(
  configOverrides: Partial<ServerConfig> = {},
  handlerOverrides: Partial<Parameters<typeof UpdateModal>[0]> = {}
) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false, gcTime: Infinity } },
  });

  return render(
    <QueryClientProvider client={queryClient}>
      <UpdateModal
        open
        serverId="srv1"
        serverName="test-server"
        config={{ ...config, ...configOverrides }}
        currentInfraVersion={1}
        outdated={false}
        onClose={vi.fn()}
        onUpdate={vi.fn()}
        isPending={false}
        {...handlerOverrides}
      />
    </QueryClientProvider>
  );
}

async function chooseOption(
  user: ReturnType<typeof userEvent.setup>,
  label: string,
  option: string
) {
  await user.click(screen.getByRole('combobox', { name: label }));
  await user.click(screen.getByRole('option', { name: option }));
}

beforeEach(() => {
  vi.clearAllMocks();
  vi.mocked(useServerOptions).mockReturnValue({
    data: {
      minecraftVersions: ['26.2', '1.21.11', '1.21.10'],
      machineTypes,
    },
    isLoading: false,
  } as never);
  vi.mocked(useBackupSettings).mockReturnValue({
    data: { enabled: true, backupIntervalHours: 6, keep: 3, keepUnit: 'daily' },
    isLoading: false,
  } as never);
  vi.mocked(useUpdateBackupSettings).mockReturnValue({
    mutate: vi.fn(),
    isPending: false,
  } as never);
  vi.mocked(useModrinthPack).mockReturnValue({
    data: undefined,
    isLoading: false,
    isError: false,
  } as never);
  vi.mocked(useModrinthPackVersions).mockReturnValue({
    data: [],
    isLoading: false,
    isError: false,
  } as never);
  vi.mocked(useModrinthSearch).mockReturnValue({
    data: [],
    isLoading: false,
    isError: false,
    refetch: vi.fn(),
    debouncedQuery: '',
  } as never);
});

describe('UpdateModal dialog behavior', () => {
  it('renders an accessible dialog with a labelled title', () => {
    renderModal();

    const dialog = screen.getByRole('dialog', { name: 'Server Settings' });
    expect(dialog).toHaveAttribute('aria-labelledby');
    expect(dialog).toHaveAttribute('aria-describedby');
  });

  it('closes through the close button, Cancel, and Escape', async () => {
    const user = userEvent.setup();
    const onClose = vi.fn();
    renderModal({}, { onClose });

    await user.click(screen.getByRole('button', { name: 'Close' }));
    expect(onClose).toHaveBeenCalledTimes(1);

    onClose.mockClear();
    renderModal({}, { onClose });
    await user.click(screen.getByRole('button', { name: 'Cancel' }));
    expect(onClose).toHaveBeenCalledTimes(1);

    onClose.mockClear();
    renderModal({}, { onClose });
    await user.keyboard('{Escape}');
    expect(onClose).toHaveBeenCalledTimes(1);
  });

  it('keeps focus inside the dialog when tabbing', async () => {
    const user = userEvent.setup();
    renderModal();

    await waitFor(() => {
      expect(screen.getByRole('dialog')).toContainElement(
        document.activeElement
      );
    });
    await user.tab({ shift: true });
    expect(screen.getByRole('dialog')).toContainElement(document.activeElement);
  });
});

describe('UpdateModal select fields', () => {
  it('renders versions in API order', async () => {
    const user = userEvent.setup();
    renderModal();

    await user.click(
      screen.getByRole('combobox', { name: 'Minecraft Version' })
    );
    expect(
      screen.getAllByRole('option').map((option) => option.textContent)
    ).toEqual(['26.2', '1.21.11', '1.21.10']);
  });

  it('includes the current version when the API no longer offers it', async () => {
    const user = userEvent.setup();
    renderModal({ minecraftVersion: '1.7.10' });

    await user.click(
      screen.getByRole('combobox', { name: 'Minecraft Version' })
    );
    expect(
      screen.getAllByRole('option').map((option) => option.textContent)
    ).toEqual(['1.7.10', '26.2', '1.21.11', '1.21.10']);
  });

  it('keeps the current machine type when the API no longer offers it', async () => {
    const user = userEvent.setup();
    renderModal({ machineType: 'e2-standard-2' });

    await user.click(screen.getByRole('combobox', { name: 'Machine Type' }));
    expect(
      screen.getAllByRole('option').map((option) => option.textContent)
    ).toEqual([
      'e2-standard-2',
      'e2-small (2 vCPU · 2 GB RAM)',
      'e2-medium (2 vCPU · 4 GB RAM)',
    ]);
  });

  it('disables both selects while options are loading', () => {
    vi.mocked(useServerOptions).mockReturnValue({
      data: undefined,
      isLoading: true,
    } as never);

    renderModal();

    expect(
      screen.getByRole('combobox', { name: 'Machine Type' })
    ).toBeDisabled();
    expect(
      screen.getByRole('combobox', { name: 'Minecraft Version' })
    ).toBeDisabled();
  });
});

describe('UpdateModal tabs and submit behavior', () => {
  it('shows Settings by default and switches to Backup', async () => {
    const user = userEvent.setup();
    renderModal();

    expect(screen.getByRole('tab', { name: 'Settings' })).toHaveAttribute(
      'aria-selected',
      'true'
    );
    expect(screen.getByDisplayValue('test-server')).toBeInTheDocument();

    await user.click(screen.getByRole('tab', { name: 'Backup' }));
    expect(screen.getByRole('tab', { name: 'Backup' })).toHaveAttribute(
      'aria-selected',
      'true'
    );
    expect(screen.getByLabelText('Backup interval (hours)')).toHaveValue(6);
  });

  it('submits only changed fields', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    renderModal({}, { onUpdate });

    await user.clear(screen.getByLabelText('Server Name'));
    await user.type(screen.getByLabelText('Server Name'), 'renamed-server');
    await user.click(screen.getByRole('button', { name: 'Update Server' }));

    expect(onUpdate).toHaveBeenCalledWith({ name: 'renamed-server' });
  });

  it('submits machine type and Minecraft version changes', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    renderModal({}, { onUpdate });

    await chooseOption(user, 'Machine Type', 'e2-medium (2 vCPU · 4 GB RAM)');
    await chooseOption(user, 'Minecraft Version', '26.2');
    await user.click(screen.getByRole('button', { name: 'Update Server' }));

    expect(onUpdate).toHaveBeenCalledWith({
      machineType: 'e2-medium',
      minecraftVersion: '26.2',
    });
  });

  it('submits disk size changes', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    renderModal({}, { onUpdate });

    const diskInput = screen.getByLabelText('Disk Size (GB)');
    await user.clear(diskInput);
    await user.type(diskInput, '100');
    await user.click(screen.getByRole('button', { name: 'Update Server' }));

    expect(onUpdate).toHaveBeenCalledWith({ diskSizeGB: 100 });
  });

  it('disables Update when unchanged and permits an outdated empty update', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    const { unmount } = renderModal({}, { onUpdate });
    expect(
      screen.getByRole('button', { name: 'Update Server' })
    ).toBeDisabled();

    unmount();
    renderModal({}, { outdated: true, onUpdate });
    await user.click(screen.getByRole('button', { name: 'Update Server' }));
    expect(onUpdate).toHaveBeenCalledWith({});
  });
});

const installedPack = { platform: 'modrinth', projectId: 'sodium-pack' };

const sodiumPack = {
  id: 'sodium-pack',
  slug: 'sodium',
  name: 'Sodium Optimized',
  author: 'jellysquid',
  iconUrl: 'https://cdn.modrinth.com/sodium.png',
  downloads: 1_500_000,
  gameVersions: ['1.21.11'],
  loader: 'fabric',
};

const sodiumVersions = [
  {
    id: 'ver-2',
    name: 'Sodium 2.0',
    versionNumber: '2.0.0',
    gameVersions: ['1.21.11'],
    loaders: ['fabric'],
    latest: true,
  },
  {
    id: 'ver-1',
    name: 'Sodium 1.0',
    versionNumber: '1.0.0',
    gameVersions: ['1.21.10'],
    loaders: ['fabric'],
    latest: false,
  },
];

function mockPackDetails({
  pack = sodiumPack,
  versions = sodiumVersions,
}: {
  pack?: typeof sodiumPack;
  versions?: typeof sodiumVersions;
} = {}) {
  vi.mocked(useModrinthPack).mockReturnValue({
    data: pack,
    isLoading: false,
    isError: false,
  } as never);
  vi.mocked(useModrinthPackVersions).mockReturnValue({
    data: versions,
    isLoading: false,
    isError: false,
  } as never);
}

function mockPackSearch(
  data: (typeof sodiumPack)[] = [sodiumPack],
  debouncedQuery = 'sodium'
) {
  vi.mocked(useModrinthSearch).mockReturnValue({
    data,
    isLoading: false,
    isError: false,
    refetch: vi.fn(),
    debouncedQuery,
  } as never);
}

async function openModpackTab(user: ReturnType<typeof userEvent.setup>) {
  await user.click(screen.getByRole('tab', { name: 'Modpack' }));
}

describe('UpdateModal modpack tab', () => {
  it('is offered alongside settings and backup', async () => {
    const user = userEvent.setup();
    renderModal();

    await openModpackTab(user);

    expect(screen.getByRole('tab', { name: 'Modpack' })).toHaveAttribute(
      'aria-selected',
      'true'
    );
    expect(screen.getByText('Vanilla Minecraft')).toBeInTheDocument();
  });

  it('shows the installed pack and its version', async () => {
    const user = userEvent.setup();
    mockPackDetails();
    renderModal({ modpack: installedPack, minecraftVersion: '' });

    await openModpackTab(user);

    expect(screen.getByText('Sodium Optimized')).toBeInTheDocument();
    expect(screen.getByText(/Installed version: Latest/)).toBeInTheDocument();
  });

  it('resolves the installed version when one is pinned', async () => {
    const user = userEvent.setup();
    mockPackDetails();
    renderModal({
      modpack: { ...installedPack, versionId: 'ver-2' },
      minecraftVersion: '',
    });

    await openModpackTab(user);

    expect(
      screen.getByText(/Installed version: Sodium 2\.0/)
    ).toBeInTheDocument();
  });

  it('falls back to the project id when the catalogue cannot resolve the pack', async () => {
    const user = userEvent.setup();
    vi.mocked(useModrinthPack).mockReturnValue({
      data: undefined,
      isLoading: false,
      isError: true,
    } as never);
    renderModal({ modpack: installedPack, minecraftVersion: '' });

    await openModpackTab(user);

    expect(screen.getByText('sodium-pack')).toBeInTheDocument();
    expect(
      screen.getByText(/Could not load the pack's details/)
    ).toBeInTheDocument();
  });

  it('hides the Minecraft version field on pack-driven servers', () => {
    renderModal({ modpack: installedPack, minecraftVersion: '' });

    expect(
      screen.queryByRole('combobox', { name: 'Minecraft Version' })
    ).not.toBeInTheDocument();
    expect(screen.getByText('Controlled by modpack')).toBeInTheDocument();
  });

  it('confirms before adding a pack and submits it on confirm', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    mockPackSearch();
    renderModal({}, { onUpdate });

    await openModpackTab(user);
    await user.type(screen.getByLabelText('Search Modpacks'), 'sodium');
    await user.click(
      await screen.findByRole('button', { name: /Sodium Optimized/ })
    );
    await user.click(screen.getByRole('button', { name: 'Update Server' }));

    const dialog = await screen.findByRole('alertdialog');
    expect(
      screen.getByText('Rebuild server to change the modpack')
    ).toBeInTheDocument();
    expect(dialog).toHaveTextContent('Vanilla → Sodium Optimized (Latest)');
    expect(dialog).toHaveTextContent(/persistent data disk and are kept/);
    expect(onUpdate).not.toHaveBeenCalled();

    await user.click(screen.getByRole('button', { name: 'Rebuild server' }));

    expect(onUpdate).toHaveBeenCalledWith({
      modpack: { platform: 'modrinth', projectId: 'sodium-pack' },
    });
  });

  it('cancels the rebuild without submitting', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    mockPackSearch();
    renderModal({}, { onUpdate });

    await openModpackTab(user);
    await user.type(screen.getByLabelText('Search Modpacks'), 'sodium');
    await user.click(
      await screen.findByRole('button', { name: /Sodium Optimized/ })
    );
    await user.click(screen.getByRole('button', { name: 'Update Server' }));
    await user.click(await screen.findByRole('button', { name: 'Cancel' }));

    expect(onUpdate).not.toHaveBeenCalled();
    expect(screen.queryByRole('alertdialog')).not.toBeInTheDocument();
  });

  it('confirms before changing the version of the installed pack', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    mockPackDetails();
    renderModal({ modpack: installedPack, minecraftVersion: '' }, { onUpdate });

    await openModpackTab(user);
    await chooseOption(user, 'Pack Version', 'Sodium 1.0');
    await user.click(screen.getByRole('button', { name: 'Update Server' }));
    await user.click(
      await screen.findByRole('button', { name: 'Rebuild server' })
    );

    expect(onUpdate).toHaveBeenCalledWith({
      modpack: {
        platform: 'modrinth',
        projectId: 'sodium-pack',
        versionId: 'ver-1',
      },
    });
  });

  it('confirms before changing to a different pack', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    mockPackDetails();
    mockPackSearch([
      sodiumPack,
      {
        id: 'other-pack',
        slug: 'other',
        name: 'Other Pack',
        author: 'someone',
        iconUrl: '',
        downloads: 10,
        gameVersions: ['1.21.11'],
        loader: 'fabric',
      },
    ]);
    renderModal({ modpack: installedPack, minecraftVersion: '' }, { onUpdate });

    await openModpackTab(user);
    await user.click(screen.getByRole('button', { name: 'Change pack' }));
    await user.type(screen.getByLabelText('Search Modpacks'), 'sodium');
    await user.click(await screen.findByRole('button', { name: /Other Pack/ }));
    await user.click(screen.getByRole('button', { name: 'Update Server' }));

    const dialog = await screen.findByRole('alertdialog');
    expect(dialog).toHaveTextContent(
      'Sodium Optimized (Latest) → Other Pack (Latest)'
    );
    expect(onUpdate).not.toHaveBeenCalled();

    await user.click(screen.getByRole('button', { name: 'Rebuild server' }));
    expect(onUpdate).toHaveBeenCalledWith({
      modpack: { platform: 'modrinth', projectId: 'other-pack' },
    });
  });

  it('requires a Minecraft version before removing the pack and submits both', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    mockPackDetails();
    renderModal({ modpack: installedPack, minecraftVersion: '' }, { onUpdate });

    await openModpackTab(user);
    await user.click(screen.getByRole('button', { name: 'Remove pack' }));

    expect(
      screen.getByText(/turns this into a vanilla server/i)
    ).toBeInTheDocument();
    expect(
      screen.getByRole('combobox', { name: 'Minecraft Version' })
    ).toBeInTheDocument();

    await user.click(screen.getByRole('button', { name: 'Update Server' }));
    const dialog = await screen.findByRole('alertdialog');
    expect(dialog).toHaveTextContent(
      'Sodium Optimized (Latest) → Vanilla (26.2)'
    );
    await user.click(screen.getByRole('button', { name: 'Rebuild server' }));

    expect(onUpdate).toHaveBeenCalledWith({
      modpack: null,
      minecraftVersion: '26.2',
    });
  });

  it('abandons a staged removal when the pack is kept', async () => {
    const user = userEvent.setup();
    mockPackDetails();
    renderModal({ modpack: installedPack, minecraftVersion: '' });

    await openModpackTab(user);
    await user.click(screen.getByRole('button', { name: 'Remove pack' }));
    await user.click(screen.getByRole('button', { name: 'Keep current pack' }));

    expect(screen.getByText(/Installed version: Latest/)).toBeInTheDocument();
    expect(
      screen.queryByText(/turns this into a vanilla server/i)
    ).not.toBeInTheDocument();
    expect(
      screen.getByRole('button', { name: 'Update Server' })
    ).toBeDisabled();
  });

  it('disables pack changes while the server is not running', async () => {
    const user = userEvent.setup();
    mockPackDetails();
    renderModal(
      { modpack: installedPack, minecraftVersion: '' },
      { isRunning: false }
    );

    await openModpackTab(user);

    expect(screen.getByRole('button', { name: 'Change pack' })).toBeDisabled();
    expect(screen.getByRole('button', { name: 'Remove pack' })).toBeDisabled();
    expect(screen.getByText(/requires it to be running/)).toBeInTheDocument();
  });

  it('does not ask for confirmation when the pack is unchanged', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();
    mockPackDetails();
    renderModal({ modpack: installedPack, minecraftVersion: '' }, { onUpdate });

    await openModpackTab(user);
    await user.click(screen.getByRole('button', { name: 'Update Server' }));

    expect(screen.queryByRole('alertdialog')).not.toBeInTheDocument();
    expect(onUpdate).not.toHaveBeenCalled();
  });
});
