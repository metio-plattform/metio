import { AlertTriangle } from 'lucide-react';
import { cn } from '../../lib/utils';
import type { StatusResponse } from '../../types/server';

interface ServerHealthWarningProps {
  status?: StatusResponse | null;
  className?: string;
}

export function ServerHealthWarning({
  status,
  className,
}: ServerHealthWarningProps) {
  if (status?.health !== 'DEGRADED') {
    return null;
  }

  return (
    <div
      role="alert"
      className={cn(
        'rounded-lg border border-destructive/30 bg-destructive/10 px-4 py-3',
        className
      )}
    >
      <div className="flex gap-3">
        <AlertTriangle className="mt-0.5 h-5 w-5 shrink-0 text-destructive" />
        <div>
          <p className="text-sm font-semibold text-destructive">
            Minecraft is not starting
          </p>
          <p className="mt-1 text-sm text-destructive/90">
            {status.healthMessage ??
              'The container is restarting repeatedly and has not answered RCON.'}
          </p>
        </div>
      </div>
    </div>
  );
}
