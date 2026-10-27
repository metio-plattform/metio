// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/metio-plattform/metio/internal/db"
	metiov1alpha1 "github.com/metio-plattform/metio/internal/operator/api/v1alpha1"
	"github.com/metio-plattform/metio/internal/operator/controllerapi"
)

const (
	syncFieldOwner = "metio-sync-controller"
	serverNameAnno = "metio.dev/server-name"
)

// SyncReconciler reconciles controller API state into MinecraftServer CRs.
type SyncReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ControllerClient controllerapi.Client
	Namespace        string
	Interval         time.Duration

	lastStatus map[string]string
}

// NewSyncReconciler creates a sync reconciler with defaults.
func NewSyncReconciler(c client.Client, scheme *runtime.Scheme, cc controllerapi.Client, namespace string, interval time.Duration) *SyncReconciler {
	if namespace == "" {
		namespace = "metio"
	}
	if interval <= 0 {
		interval = 30 * time.Second
	}
	return &SyncReconciler{
		Client:           c,
		Scheme:           scheme,
		ControllerClient: cc,
		Namespace:        namespace,
		Interval:         interval,
		lastStatus:       map[string]string{},
	}
}

// +kubebuilder:rbac:groups=minecraft.metio.dev,resources=minecraftservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=minecraft.metio.dev,resources=minecraftservers/status,verbs=get;update;patch

// Reconcile is a no-op at the per-object level; syncing is driven by Sync.
func (r *SyncReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// SetupWithManager registers the sync controller with the manager.
func (r *SyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.lastStatus == nil {
		r.lastStatus = map[string]string{}
	}
	return mgr.Add(r)
}

// Start runs the periodic sync loop.
func (r *SyncReconciler) Start(ctx context.Context) error {
	if r.Interval <= 0 {
		r.Interval = 30 * time.Second
	}
	ticker := time.NewTicker(r.Interval)
	defer ticker.Stop()
	// Run once immediately.
	if err := r.Sync(ctx); err != nil {
		log.FromContext(ctx).Error(err, "sync failed")
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := r.Sync(ctx); err != nil {
				log.FromContext(ctx).Error(err, "sync failed")
			}
		}
	}
}

// Need implements manager.Runnable? manager.Add requires Start(context.Context) error. Good.

// Sync performs one sync pass.
func (r *SyncReconciler) Sync(ctx context.Context) error {
	if r.ControllerClient == nil {
		return nil
	}
	servers, err := r.ControllerClient.ListServers(ctx)
	if err != nil {
		return fmt.Errorf("list servers: %w", err)
	}

	current := map[string]struct{}{}
	for _, s := range servers {
		if s.Config.Runtime != db.RuntimeKubernetes {
			continue
		}
		current[s.ID] = struct{}{}
		ms := r.buildMinecraftServer(s)
		data, err := json.Marshal(ms)
		if err != nil {
			return fmt.Errorf("marshal %s: %w", s.ID, err)
		}
		patch := client.RawPatch(types.ApplyPatchType, data)
		if err := r.Client.Patch(ctx, ms, patch, client.FieldOwner(syncFieldOwner), client.ForceOwnership); err != nil {
			return fmt.Errorf("apply %s: %w", s.ID, err)
		}
		// Read back status and report if changed.
		got := &metiov1alpha1.MinecraftServer{}
		if err := r.Client.Get(ctx, types.NamespacedName{Name: s.ID, Namespace: r.Namespace}, got); err != nil {
			continue
		}
		status := crStatusToDB(got.Status)
		b, _ := json.Marshal(status)
		key := string(b)
		if r.lastStatus[s.ID] == key {
			continue
		}
		if err := r.ControllerClient.UpdateServerStatus(ctx, s.ID, status); err != nil {
			log.FromContext(ctx).Error(err, "report status", "server", s.ID)
			continue
		}
		r.lastStatus[s.ID] = key
	}

	// Delete CRs for removed servers.
	var list metiov1alpha1.MinecraftServerList
	if err := r.Client.List(ctx, &list, client.InNamespace(r.Namespace)); err != nil {
		return fmt.Errorf("list crs: %w", err)
	}
	for i := range list.Items {
		cr := &list.Items[i]
		if _, ok := current[cr.Name]; ok {
			continue
		}
		if err := r.Client.Delete(ctx, cr); err != nil && !errors.IsNotFound(err) {
			return fmt.Errorf("delete %s: %w", cr.Name, err)
		}
		delete(r.lastStatus, cr.Name)
	}
	return nil
}

func (r *SyncReconciler) buildMinecraftServer(s controllerapi.ServerResponse) *metiov1alpha1.MinecraftServer {
	ms := &metiov1alpha1.MinecraftServer{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "minecraft.metio.dev/v1alpha1",
			Kind:       "MinecraftServer",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        s.ID,
			Namespace:   r.Namespace,
			Annotations: map[string]string{serverNameAnno: s.Config.Name},
		},
		Spec: metiov1alpha1.MinecraftServerSpec{
			ServerID:         s.ID,
			MinecraftVersion: s.Config.MinecraftVersion,
			Modpack:          toV1Modpack(s.Config.Modpack),
			Backup:           toV1Backup(s.Config.Backup),
			ShutdownSchedule: toV1Shutdown(s.Config.ShutdownSchedule),
		},
	}
	return ms
}

func toV1Modpack(m *db.ModpackConfig) *metiov1alpha1.ModpackConfig {
	if m == nil {
		return nil
	}
	return &metiov1alpha1.ModpackConfig{
		Platform:  m.Platform,
		ProjectID: m.ProjectID,
		VersionID: m.VersionID,
	}
}

func toV1Backup(b *db.BackupConfig) *metiov1alpha1.BackupConfig {
	if b == nil {
		return nil
	}
	return &metiov1alpha1.BackupConfig{
		Enabled:             b.Enabled,
		BackupIntervalHours: b.BackupIntervalHours,
		Keep:                b.Keep,
		KeepUnit:            b.KeepUnit,
	}
}

func toV1Shutdown(s *db.ShutdownSchedule) *metiov1alpha1.ShutdownSchedule {
	if s == nil {
		return nil
	}
	return &metiov1alpha1.ShutdownSchedule{
		Enabled: s.Enabled,
		Time:    s.Time,
	}
}

func crStatusToDB(st metiov1alpha1.MinecraftServerStatus) db.Status {
	status := db.Status{
		Timestamp:   time.Now(),
		ServerState: db.ServerState(st.State),
		Version:     st.Version,
		Players: db.Players{
			Current: st.Players.Current,
			Max:     st.Players.Max,
		},
	}
	if status.ServerState == "" {
		status.ServerState = db.ServerStateStarting
	}
	return status
}
