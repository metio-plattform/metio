package servers

import (
	"testing"
	"time"

	"github.com/metio-plattform/metio/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestStatusHealth(t *testing.T) {
	now := time.Date(2026, 10, 24, 12, 0, 0, 0, time.UTC)
	recent := now.Add(-time.Minute)
	old := now.Add(-time.Hour)

	tests := []struct {
		name        string
		status      db.Status
		wantHealth  string
		wantMessage bool
	}{
		{
			name:   "healthy",
			status: db.Status{Version: "1.20.1", Players: db.Players{Max: 20}},
		},
		{
			name: "crash looping",
			status: db.Status{
				RestartCount:  2,
				LastRestartAt: &recent,
				Version:       "Unknown",
			},
			wantHealth:  "DEGRADED",
			wantMessage: true,
		},
		{
			name: "old restarts are not a current crash loop",
			status: db.Status{
				RestartCount:  2,
				LastRestartAt: &old,
				Version:       "Unknown",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			health, message := statusHealth(tt.status, now)
			assert.Equal(t, tt.wantHealth, health)
			if tt.wantMessage {
				assert.NotEmpty(t, message)
			} else {
				assert.Empty(t, message)
			}
		})
	}
}
