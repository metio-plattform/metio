package dbtypes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStatusIsCrashLooping(t *testing.T) {
	now := time.Date(2026, 10, 24, 12, 0, 0, 0, time.UTC)
	recent := now.Add(-time.Minute)
	old := now.Add(-CrashLoopWindow - time.Minute)

	tests := []struct {
		name   string
		status Status
		want   bool
	}{
		{
			name:   "no restart data",
			status: Status{Version: "Unknown"},
			want:   false,
		},
		{
			name: "single restart",
			status: Status{
				RestartCount:  1,
				LastRestartAt: &recent,
				Version:       "Unknown",
			},
			want: false,
		},
		{
			name: "two recent restarts without RCON",
			status: Status{
				RestartCount:  CrashLoopRestartThreshold,
				LastRestartAt: &recent,
				Version:       "Unknown",
			},
			want: true,
		},
		{
			name: "empty version is also not ready",
			status: Status{
				RestartCount:  CrashLoopRestartThreshold,
				LastRestartAt: &recent,
			},
			want: true,
		},
		{
			name: "answered RCON",
			status: Status{
				RestartCount:  CrashLoopRestartThreshold,
				LastRestartAt: &recent,
				Players:       Players{Max: 20},
				Version:       "Unknown",
			},
			want: false,
		},
		{
			name: "known version",
			status: Status{
				RestartCount:  CrashLoopRestartThreshold,
				LastRestartAt: &recent,
				Version:       "1.20.1",
			},
			want: false,
		},
		{
			name: "restart outside the window",
			status: Status{
				RestartCount:  CrashLoopRestartThreshold,
				LastRestartAt: &old,
				Version:       "Unknown",
			},
			want: false,
		},
		{
			name: "nil last restart",
			status: Status{
				RestartCount: CrashLoopRestartThreshold,
				Version:      "Unknown",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.status.IsCrashLooping(now))
		})
	}
}

func TestStatusCrashLoopMessage(t *testing.T) {
	now := time.Date(2026, 10, 24, 12, 0, 0, 0, time.UTC)
	recent := now.Add(-time.Minute)

	status := Status{
		RestartCount:  3,
		LastRestartAt: &recent,
		Version:       "Unknown",
	}

	assert.Contains(t, status.CrashLoopMessage(now), "restarted 3 times")
	assert.Contains(t, status.CrashLoopMessage(now), "has not answered RCON")
	assert.Empty(t, Status{}.CrashLoopMessage(now))
}
