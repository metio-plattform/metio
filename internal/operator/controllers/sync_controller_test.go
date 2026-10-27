// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/metio-plattform/metio/internal/db"
	"github.com/metio-plattform/metio/internal/operator/api/v1alpha1"
	"github.com/metio-plattform/metio/internal/operator/controllerapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func newScheme(t *testing.T) *runtime.Scheme {
	s := runtime.NewScheme()
	if err := v1alpha1.AddToScheme(s); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestSyncCreatesCRForKubernetesServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/servers" {
			json.NewEncoder(w).Encode([]map[string]interface{}{
				{"id": "srv-1", "config": map[string]interface{}{"name": "alpha", "runtime": "kubernetes", "minecraftVersion": "1.21.1"}},
			})
			return
		}
		if r.URL.Path == "/api/servers/srv-1/status" {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	scheme := newScheme(t)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	cc := controllerapi.New(srv.URL, "key")
	r := NewSyncReconciler(fakeClient, scheme, cc, "metio", 0)

	if err := r.Sync(context.Background()); err != nil {
		t.Fatal(err)
	}

	got := &v1alpha1.MinecraftServer{}
	if err := fakeClient.Get(context.Background(), client.ObjectKey{Name: "srv-1", Namespace: "metio"}, got); err != nil {
		t.Fatalf("cr not created: %v", err)
	}
	if got.Annotations["metio.dev/server-name"] != "alpha" {
		t.Fatalf("annotation missing: %v", got.Annotations)
	}
	if got.Spec.MinecraftVersion != "1.21.1" {
		t.Fatalf("version not synced: %s", got.Spec.MinecraftVersion)
	}
}

func TestSyncIsIdempotent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/servers" {
			json.NewEncoder(w).Encode([]map[string]interface{}{
				{"id": "srv-4", "config": map[string]interface{}{"name": "delta", "runtime": "kubernetes", "minecraftVersion": "1.21.1"}},
			})
			return
		}
		if r.URL.Path == "/api/servers/srv-4/status" {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	scheme := newScheme(t)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	cc := controllerapi.New(srv.URL, "key")
	r := NewSyncReconciler(fakeClient, scheme, cc, "metio", 0)

	if err := r.Sync(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.Sync(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestSyncSkipsDockerServers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/servers" {
			json.NewEncoder(w).Encode([]map[string]interface{}{
				{"id": "srv-2", "config": map[string]interface{}{"name": "beta", "runtime": "docker", "minecraftVersion": "1.21.1"}},
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	scheme := newScheme(t)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	cc := controllerapi.New(srv.URL, "key")
	r := NewSyncReconciler(fakeClient, scheme, cc, "metio", 0)

	if err := r.Sync(context.Background()); err != nil {
		t.Fatal(err)
	}

	got := &v1alpha1.MinecraftServer{}
	err := fakeClient.Get(context.Background(), client.ObjectKey{Name: "srv-2", Namespace: "metio"}, got)
	if err == nil {
		t.Fatal("docker server should not create a CR")
	}
}

func TestSyncDeletesRemovedCR(t *testing.T) {
	scheme := newScheme(t)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	// Pre-create a CR that is no longer in the controller API.
	pre := &v1alpha1.MinecraftServer{ObjectMeta: metav1.ObjectMeta{Name: "old", Namespace: "metio"}, Spec: v1alpha1.MinecraftServerSpec{ServerID: "old", MinecraftVersion: "1.21.1"}}
	if err := fakeClient.Create(context.Background(), pre); err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/servers" {
			json.NewEncoder(w).Encode([]map[string]interface{}{})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	cc := controllerapi.New(srv.URL, "key")
	r := NewSyncReconciler(fakeClient, scheme, cc, "metio", 0)

	if err := r.Sync(context.Background()); err != nil {
		t.Fatal(err)
	}

	got := &v1alpha1.MinecraftServer{}
	err := fakeClient.Get(context.Background(), client.ObjectKey{Name: "old", Namespace: "metio"}, got)
	if err == nil {
		t.Fatal("removed server CR should be deleted")
	}
}

func TestSyncReportsStatus(t *testing.T) {
	statusCalled := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/servers" {
			json.NewEncoder(w).Encode([]map[string]interface{}{
				{"id": "srv-3", "config": map[string]interface{}{"name": "gamma", "runtime": "kubernetes", "minecraftVersion": "1.21.1"}},
			})
			return
		}
		if r.URL.Path == "/api/servers/srv-3/status" && r.Method == http.MethodPut {
			statusCalled = true
			var st db.Status
			json.NewDecoder(r.Body).Decode(&st)
			if st.ServerState != "running" {
				t.Errorf("unexpected state: %s", st.ServerState)
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	scheme := newScheme(t)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	// Pre-create CR with status.
	pre := &v1alpha1.MinecraftServer{
		ObjectMeta: metav1.ObjectMeta{Name: "srv-3", Namespace: "metio"},
		Spec:       v1alpha1.MinecraftServerSpec{ServerID: "srv-3", MinecraftVersion: "1.21.1"},
		Status:     v1alpha1.MinecraftServerStatus{State: "running", Version: "1.21.1"},
	}
	if err := fakeClient.Create(context.Background(), pre); err != nil {
		t.Fatal(err)
	}

	cc := controllerapi.New(srv.URL, "key")
	r := NewSyncReconciler(fakeClient, scheme, cc, "metio", 0)

	if err := r.Sync(context.Background()); err != nil {
		t.Fatal(err)
	}
	if !statusCalled {
		t.Fatal("status should have been reported")
	}
}