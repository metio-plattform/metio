import { describe, it, expect, beforeEach, afterEach, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useModrinthPack } from './useModrinthPack';
import type { ModrinthPack } from '../types/server';

const mockFetch = vi.fn();
global.fetch = mockFetch;

function createQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: { retry: false, gcTime: Infinity },
    },
  });
}

function createWrapper(queryClient: QueryClient) {
  return function Wrapper({ children }: { children: React.ReactNode }) {
    return (
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    );
  };
}

const mockPack: ModrinthPack = {
  id: 'proj-1',
  slug: 'skyblock',
  name: 'SkyBlock Enhanced',
  author: '',
  iconUrl: 'https://cdn.modrinth.com/icon.png',
  downloads: 100,
  gameVersions: ['1.21.11'],
  loader: 'fabric',
};

describe('useModrinthPack', () => {
  let queryClient: QueryClient;

  beforeEach(() => {
    vi.clearAllMocks();
    queryClient = createQueryClient();
  });

  afterEach(() => {
    queryClient.clear();
  });

  it('fetches the pack when a projectId is provided', async () => {
    mockFetch.mockResolvedValue({
      ok: true,
      json: () => Promise.resolve(mockPack),
    });

    const { result } = renderHook(() => useModrinthPack('proj-1'), {
      wrapper: createWrapper(queryClient),
    });

    await waitFor(() => {
      expect(result.current.data).toEqual(mockPack);
    });
    expect(mockFetch).toHaveBeenCalledWith('/api/modpacks/proj-1');
  });

  it('stays disabled without a projectId', async () => {
    const { result } = renderHook(() => useModrinthPack(null), {
      wrapper: createWrapper(queryClient),
    });

    await waitFor(() => {
      expect(result.current.isFetching).toBe(false);
    });
    expect(mockFetch).not.toHaveBeenCalled();
    expect(result.current.data).toBeUndefined();
  });

  it('exposes an error when the request fails', async () => {
    mockFetch.mockResolvedValue({
      ok: false,
      status: 500,
      json: () => Promise.resolve({ error: 'boom' }),
    });

    const { result } = renderHook(() => useModrinthPack('proj-1'), {
      wrapper: createWrapper(queryClient),
    });

    await waitFor(() => {
      expect(result.current.isError).toBe(true);
    });
    expect(result.current.error?.message).toBe('boom');
  });
});
