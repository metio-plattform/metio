import { useQuery } from '@tanstack/react-query';
import type { APIError, ModrinthPack } from '../types/server';

async function fetchModpack(projectId: string): Promise<ModrinthPack> {
  const response = await fetch(`/api/modpacks/${projectId}`);

  if (response.status === 401) {
    window.location.href = '/auth/login';
    throw new Error('Session expired');
  }

  if (!response.ok) {
    const error: APIError = await response.json();
    throw new Error(error.error || 'Failed to load modpack');
  }

  return response.json();
}

/**
 * Resolves a single Modpack by project id, used to name the pack a server
 * already has installed. Disabled until a projectId is provided.
 */
export function useModrinthPack(projectId: string | null) {
  return useQuery({
    queryKey: ['modpack', projectId],
    queryFn: () => fetchModpack(projectId as string),
    enabled: Boolean(projectId),
    staleTime: 15 * 60 * 1000,
    gcTime: 30 * 60 * 1000,
  });
}
