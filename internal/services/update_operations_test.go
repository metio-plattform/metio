package services

import (
	"context"
	"testing"
	"time"

	"github.com/metio-plattform/metio/internal/db"
	"github.com/metio-plattform/metio/internal/dbtypes"
	"github.com/metio-plattform/metio/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBackupCoordinator_TriggerWorldSave(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, nil)
	mockDB.On("UpdateStatus", mock.Anything, "test-instance", mock.MatchedBy(func(s db.Status) bool {
		return s.PendingCommand == "save" && s.PendingCommandResult == ""
	})).Return(nil)

	err := bc.TriggerWorldSave(context.Background(), "test-instance")
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestBackupCoordinator_TriggerWorldSave_GetStatusError(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, assert.AnError)

	err := bc.TriggerWorldSave(context.Background(), "test-instance")
	assert.Error(t, err)
}

func TestBackupCoordinator_TriggerWorldSave_NoStatusDocument(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, db.ErrNotFound)

	err := bc.TriggerWorldSave(context.Background(), "test-instance")
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestBackupCoordinator_TriggerWorldSave_UpdateStatusError(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, nil)
	mockDB.On("UpdateStatus", mock.Anything, "test-instance", mock.AnythingOfType("dbtypes.Status")).Return(assert.AnError)

	err := bc.TriggerWorldSave(context.Background(), "test-instance")
	assert.Error(t, err)
}

func TestBackupCoordinator_WaitForCommandAck_Completed(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{PendingCommandResult: "completed"}, nil).Once()

	result, err := bc.WaitForCommandAck(context.Background(), "test-instance", 5*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, "completed", result)
}

func TestBackupCoordinator_WaitForCommandAck_ErrorResult(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{PendingCommandResult: "failed"}, nil).Once()

	result, err := bc.WaitForCommandAck(context.Background(), "test-instance", 5*time.Second)
	assert.Error(t, err)
	assert.Equal(t, "failed", result)
}

func TestBackupCoordinator_WaitForCommandAck_GetStatusError(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, assert.AnError).Once()

	_, err := bc.WaitForCommandAck(context.Background(), "test-instance", 5*time.Second)
	assert.Error(t, err)
}

func TestBackupCoordinator_WaitForCommandAck_NoStatusDocument(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, db.ErrNotFound).Once()

	result, err := bc.WaitForCommandAck(context.Background(), "test-instance", 5*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, "", result)
	mockDB.AssertExpectations(t)
}

func TestBackupCoordinator_WaitForCommandAck_Timeout(t *testing.T) {
	mockDB := new(testutil.MockDB)
	bc := NewBackupCoordinator(mockDB)

	mockDB.On("GetStatus", mock.Anything, "test-instance").Return(db.Status{}, nil)

	_, err := bc.WaitForCommandAck(context.Background(), "test-instance", 1*time.Second)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestWaitForServerHealthy_Success(t *testing.T) {
	mockDB := new(testutil.MockDB)

	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{ServerState: "RUNNING"}, nil).Once()

	err := WaitForServerHealthy(context.Background(), mockDB, "test", 5*time.Second)
	assert.NoError(t, err)
}

func TestWaitForServerHealthy_GetStatusError(t *testing.T) {
	mockDB := new(testutil.MockDB)

	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{}, assert.AnError).Once()

	err := WaitForServerHealthy(context.Background(), mockDB, "test", 1*time.Second)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestWaitForServerHealthy_Timeout(t *testing.T) {
	mockDB := new(testutil.MockDB)

	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{ServerState: "STARTING"}, nil)

	err := WaitForServerHealthy(context.Background(), mockDB, "test", 1*time.Second)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

// shortenModpackPollInterval speeds up WaitForModpackInstall tests.
func shortenModpackPollInterval(t *testing.T) {
	t.Helper()
	original := modpackInstallPollInterval
	modpackInstallPollInterval = time.Millisecond
	t.Cleanup(func() { modpackInstallPollInterval = original })
}

func TestWaitForModpackInstall_Success(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	// A pack on Minecraft 1.20.1: the server answers the RCON player list but
	// has no `version` command, so the player count is the readiness signal.
	notBefore := time.Now().Add(-time.Minute)
	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "RUNNING",
		Version:     "Unknown",
		Players:     dbtypes.Players{Current: 0, Max: 20},
		Timestamp:   time.Now(),
	}, nil)

	err := WaitForModpackInstall(context.Background(), mockDB, "test", notBefore, time.Second)
	assert.NoError(t, err)
}

func TestWaitForModpackInstall_SuccessViaKnownVersion(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	notBefore := time.Now().Add(-time.Minute)
	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "RUNNING",
		Version:     "1.21.11",
		Timestamp:   time.Now(),
	}, nil)

	err := WaitForModpackInstall(context.Background(), mockDB, "test", notBefore, time.Second)
	assert.NoError(t, err)
}

func TestWaitForModpackInstall_StaleStatusIgnored(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	// The status of the VM that the recreate replaced: running and reachable,
	// but pushed before the wait started.
	notBefore := time.Now()
	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "RUNNING",
		Version:     "1.21.11",
		Players:     dbtypes.Players{Current: 0, Max: 20},
		Timestamp:   notBefore.Add(-time.Minute),
	}, nil)

	err := WaitForModpackInstall(context.Background(), mockDB, "test", notBefore, 50*time.Millisecond)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestWaitForModpackInstall_UnknownVersionWaits(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	// The container is already running while the itzg image still downloads and
	// installs the modpack: no RCON answer yet, neither as a player count nor as
	// a version. The timestamp only has to be after notBefore, so a future one
	// is safe here.
	notBefore := time.Now().Add(-time.Minute)
	fresh := time.Now().Add(time.Hour)
	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "RUNNING",
		Version:     "Unknown",
		Timestamp:   fresh,
	}, nil).Once()
	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "RUNNING",
		Version:     "Unknown",
		Players:     dbtypes.Players{Current: 0, Max: 20},
		Timestamp:   fresh,
	}, nil)

	err := WaitForModpackInstall(context.Background(), mockDB, "test", notBefore, time.Second)
	assert.NoError(t, err)
}

func TestWaitForModpackInstall_NotRunningWaits(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	// A reachable server that is not running (for example a shut down server
	// that kept its last RCON answer) is not an installed pack.
	notBefore := time.Now().Add(-time.Minute)
	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "STOPPED",
		Version:     "1.21.11",
		Players:     dbtypes.Players{Current: 0, Max: 20},
		Timestamp:   time.Now(),
	}, nil)

	err := WaitForModpackInstall(context.Background(), mockDB, "test", notBefore, 50*time.Millisecond)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestWaitForModpackInstall_GetStatusError(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{}, assert.AnError)

	err := WaitForModpackInstall(context.Background(), mockDB, "test", time.Now().Add(-time.Minute), 50*time.Millisecond)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestWaitForModpackInstall_ContextCancelled(t *testing.T) {
	shortenModpackPollInterval(t)
	mockDB := new(testutil.MockDB)

	mockDB.On("GetStatus", mock.Anything, "test").Return(db.Status{
		ServerState: "RUNNING",
		Version:     "Unknown",
		Timestamp:   time.Now(),
	}, nil)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := WaitForModpackInstall(ctx, mockDB, "test", time.Now().Add(-time.Minute), time.Minute)
	assert.ErrorIs(t, err, context.Canceled)
}
