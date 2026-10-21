import { useState } from 'react';
import { Package, Search } from 'lucide-react';
import { useModrinthSearch } from '../../hooks/useModrinthSearch';
import { useModrinthPackVersions } from '../../hooks/useModrinthPackVersions';
import { Badge } from '../ui/Badge';
import { Button } from '../ui/Button';
import { Input } from '../ui/Input';
import { Skeleton } from '../ui/Skeleton';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '../ui/Select';
import { cn } from '../../lib/utils';
import { SEARCH_MIN_LENGTH, type ModpackSelection } from '../../lib/modpack';
import type { ModrinthPack } from '../../types/server';

/** Formats a raw download count into a compact label (e.g. 12.3M). */
function formatDownloads(count: number): string {
  if (count >= 1_000_000) {
    return `${(count / 1_000_000).toFixed(1).replace(/\.0$/, '')}M`;
  }
  if (count >= 1_000) {
    return `${(count / 1_000).toFixed(1).replace(/\.0$/, '')}K`;
  }
  return `${count}`;
}

/** Pack icon with a fallback to a generic package glyph. */
export function PackIcon({ iconUrl }: { iconUrl: string }) {
  const [failed, setFailed] = useState(false);

  if (!iconUrl || failed) {
    return (
      <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-md bg-muted">
        <Package className="h-5 w-5 text-muted-foreground" />
      </div>
    );
  }

  return (
    <img
      src={iconUrl}
      alt=""
      className="h-10 w-10 shrink-0 rounded-md object-cover"
      onError={() => setFailed(true)}
    />
  );
}

export interface ModpackPickerProps {
  /** Current selection, or null when the server is (to be) vanilla. */
  value: ModpackSelection | null;
  onChange: (selection: ModpackSelection | null) => void;
  /** Disables the search and version controls while a change is submitting. */
  disabled?: boolean;
  className?: string;
}

/**
 * Modrinth modpack search and version picker. Shared between the setup wizard
 * and the update modal so pack selection behaves identically in both places.
 * It is uncontrolled with respect to the search box and controlled with respect
 * to the selection.
 */
export function ModpackPicker({
  value,
  onChange,
  disabled = false,
  className,
}: ModpackPickerProps) {
  const [query, setQuery] = useState('');
  const search = useModrinthSearch(query);
  const versions = useModrinthPackVersions(value?.projectId ?? null);

  const trimmedQuery = query.trim();
  const packs = search.data ?? [];
  const versionOptions = versions.data ?? [];
  const isSearching =
    search.isLoading ||
    (trimmedQuery !== search.debouncedQuery &&
      trimmedQuery.length >= SEARCH_MIN_LENGTH);

  const selectPack = (pack: ModrinthPack) => {
    onChange({
      platform: 'modrinth',
      projectId: pack.id,
      name: pack.name,
      iconUrl: pack.iconUrl,
      loader: pack.loader || undefined,
    });
  };

  const selectVersion = (versionId: string) => {
    if (!value) return;

    if (versionId === 'latest') {
      onChange({ ...value, versionId: undefined, versionName: undefined });
      return;
    }

    const version = versionOptions.find((option) => option.id === versionId);
    onChange({
      ...value,
      versionId,
      versionName: version?.name ?? version?.versionNumber,
    });
  };

  return (
    <div className={cn('space-y-6', className)}>
      <div className="space-y-2">
        <label htmlFor="modpack-search" className="text-sm font-medium">
          Search Modpacks
        </label>
        <div className="relative">
          <Search className="pointer-events-none absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            id="modpack-search"
            value={query}
            onChange={(event) => setQuery(event.target.value)}
            placeholder="Search Modrinth modpacks…"
            className="pl-9"
            autoComplete="off"
            disabled={disabled}
          />
        </div>
      </div>

      {value ? (
        <div className="space-y-4 rounded-lg border border-green-600 bg-green-600/10 p-4">
          <div className="flex items-start justify-between gap-4">
            <div className="flex min-w-0 items-center gap-3">
              <PackIcon iconUrl={value.iconUrl ?? ''} />
              <div className="min-w-0">
                <div className="truncate font-medium text-foreground">
                  {value.name}
                </div>
                <div className="text-sm text-muted-foreground">
                  {value.loader ? `${value.loader} · ` : ''}Modrinth
                </div>
              </div>
            </div>
            <Button
              type="button"
              variant="ghost"
              size="sm"
              onClick={() => onChange(null)}
              disabled={disabled}
            >
              Clear pack
            </Button>
          </div>

          <div className="space-y-2">
            <label htmlFor="modpack-version" className="text-sm font-medium">
              Pack Version
            </label>
            <Select
              value={value.versionId ?? 'latest'}
              onValueChange={selectVersion}
              disabled={disabled || versions.isLoading}
            >
              <SelectTrigger id="modpack-version" className="w-full sm:w-72">
                <SelectValue
                  placeholder={
                    versions.isLoading ? 'Loading versions…' : 'Latest'
                  }
                />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="latest">Latest</SelectItem>
                {versionOptions.map((version) => (
                  <SelectItem key={version.id} value={version.id}>
                    {version.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {versions.isError ? (
              <p className="text-sm text-destructive">
                Failed to load pack versions.
              </p>
            ) : null}
          </div>
        </div>
      ) : null}

      {trimmedQuery.length < SEARCH_MIN_LENGTH ? (
        <p className="text-sm text-muted-foreground">
          Type at least {SEARCH_MIN_LENGTH} characters to search the Modrinth
          catalogue.
        </p>
      ) : isSearching ? (
        <div className="grid grid-cols-1 gap-3 sm:grid-cols-2">
          {[0, 1, 2, 3].map((index) => (
            <Skeleton key={index} className="h-20 rounded-lg" />
          ))}
        </div>
      ) : search.isError ? (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4 text-sm text-destructive">
          <p>Failed to search modpacks.</p>
          <Button
            type="button"
            variant="outline"
            size="sm"
            className="mt-3"
            onClick={() => search.refetch()}
          >
            Try again
          </Button>
        </div>
      ) : packs.length === 0 ? (
        <p className="text-sm text-muted-foreground">
          No modpacks found. Try a different search.
        </p>
      ) : (
        <div
          className="grid grid-cols-1 gap-3 sm:grid-cols-2"
          role="group"
          aria-label="Modpack results"
        >
          {packs.map((pack) => {
            const isSelected = value?.projectId === pack.id;
            return (
              <button
                key={pack.id}
                type="button"
                aria-pressed={isSelected}
                onClick={() => selectPack(pack)}
                disabled={disabled}
                className={cn(
                  'flex items-start gap-3 rounded-lg border p-3 text-left transition-all disabled:cursor-not-allowed disabled:opacity-60',
                  isSelected
                    ? 'border-green-600 bg-green-600/10 ring-1 ring-green-600'
                    : 'border-border bg-muted hover:border-ring'
                )}
              >
                <PackIcon iconUrl={pack.iconUrl} />
                <div className="min-w-0 flex-1">
                  <div className="truncate font-medium text-foreground">
                    {pack.name}
                  </div>
                  <div className="truncate text-sm text-muted-foreground">
                    by {pack.author}
                  </div>
                  <div className="mt-1 flex flex-wrap items-center gap-2">
                    {pack.loader ? (
                      <Badge variant="secondary">{pack.loader}</Badge>
                    ) : null}
                    <span className="text-xs text-muted-foreground">
                      {formatDownloads(pack.downloads)} downloads
                    </span>
                  </div>
                </div>
              </button>
            );
          })}
        </div>
      )}
    </div>
  );
}
