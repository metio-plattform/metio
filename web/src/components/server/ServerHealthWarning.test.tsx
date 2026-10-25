import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import { ServerHealthWarning } from './ServerHealthWarning';
import type { StatusResponse } from '../../types/server';

const baseStatus: StatusResponse = {
  serverState: 'RUNNING',
  players: { current: 0, max: 0 },
  uptime: '',
  version: 'Unknown',
  instanceIP: '',
};

describe('ServerHealthWarning', () => {
  it('renders nothing for a healthy status', () => {
    const { container } = render(<ServerHealthWarning status={baseStatus} />);
    expect(container).toBeEmptyDOMElement();
  });

  it('renders the crash-loop message', () => {
    render(
      <ServerHealthWarning
        status={{
          ...baseStatus,
          health: 'DEGRADED',
          healthMessage: 'Minecraft restarted 4 times.',
        }}
      />
    );

    expect(screen.getByRole('alert')).toBeInTheDocument();
    expect(screen.getByText('Minecraft is not starting')).toBeInTheDocument();
    expect(
      screen.getByText('Minecraft restarted 4 times.')
    ).toBeInTheDocument();
  });

  it('falls back to a generic message', () => {
    render(
      <ServerHealthWarning status={{ ...baseStatus, health: 'DEGRADED' }} />
    );

    expect(
      screen.getByText(
        'The container is restarting repeatedly and has not answered RCON.'
      )
    ).toBeInTheDocument();
  });
});
