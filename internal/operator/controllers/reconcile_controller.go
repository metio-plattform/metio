// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	metiov1alpha1 "github.com/metio-plattform/metio/internal/operator/api/v1alpha1"
)

// ReconcileReconciler reconciles a MinecraftServer object (reconcile controller).
type ReconcileReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=minecraft.metio.dev,resources=minecraftservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=minecraft.metio.dev,resources=minecraftservers/status,verbs=get;update;patch

// Reconcile is a no-op stub for the scaffold.
func (r *ReconcileReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log.FromContext(ctx).Info("reconcile reconcile", "request", req.NamespacedName)
	return ctrl.Result{}, nil
}

// SetupWithManager registers the reconcile controller with the manager.
func (r *ReconcileReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&metiov1alpha1.MinecraftServer{}).
		Named("reconcile").
		Complete(r)
}