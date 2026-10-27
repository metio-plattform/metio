// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	metiov1alpha1 "github.com/metio-plattform/metio/internal/operator/api/v1alpha1"
)

var (
	cfg       *rest.Config
	k8sClient client.Client
	testEnv   *envtest.Environment
	ctx       context.Context
	cancel    context.CancelFunc
)

func TestControllers(t *testing.T) {
	if os.Getenv("KUBEBUILDER_ASSETS") == "" {
		t.Skip("KUBEBUILDER_ASSETS not set; skipping envtest suite")
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(context.Background())

	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "..", "..", "config", "crd", "bases")},
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = metiov1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	// Start a manager with both controllers registered.
	mgr, err := manager.New(cfg, manager.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())

	err = (&SyncReconciler{Client: mgr.GetClient(), Scheme: mgr.GetScheme()}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())
	err = (&ReconcileReconciler{Client: mgr.GetClient(), Scheme: mgr.GetScheme()}).SetupWithManager(mgr)
	Expect(err).NotTo(HaveOccurred())

	go func() {
		defer GinkgoRecover()
		Expect(mgr.Start(ctx)).To(Succeed())
	}()
})

var _ = Describe("MinecraftServer", func() {
	It("creates, updates, and deletes a CR without errors", func() {
		ms := &metiov1alpha1.MinecraftServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-server",
				Namespace: "default",
			},
			Spec: metiov1alpha1.MinecraftServerSpec{
				ServerID:         "srv-1",
				MinecraftVersion: "1.21.1",
			},
		}
		Expect(k8sClient.Create(ctx, ms)).To(Succeed())

		// Update
		Eventually(func() error {
			got := &metiov1alpha1.MinecraftServer{}
			if err := k8sClient.Get(ctx, client.ObjectKeyFromObject(ms), got); err != nil {
				return err
			}
			got.Spec.MinecraftVersion = "1.21.4"
			return k8sClient.Update(ctx, got)
		}, time.Second*5, time.Millisecond*250).Should(Succeed())

		// Delete
		Expect(k8sClient.Delete(ctx, ms)).To(Succeed())
		Eventually(func() bool {
			got := &metiov1alpha1.MinecraftServer{}
			err := k8sClient.Get(ctx, client.ObjectKeyFromObject(ms), got)
			return err != nil
		}, time.Second*5, time.Millisecond*250).Should(BeTrue())
	})

	It("rejects an object that sets both minecraftVersion and modpack", func() {
		ms := &metiov1alpha1.MinecraftServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bad-both",
				Namespace: "default",
			},
			Spec: metiov1alpha1.MinecraftServerSpec{
				ServerID:         "srv-2",
				MinecraftVersion: "1.21.1",
				Modpack: &metiov1alpha1.ModpackConfig{
					Platform:  "modrinth",
					ProjectID: "abc123",
				},
			},
		}
		err := k8sClient.Create(ctx, ms)
		Expect(err).To(HaveOccurred())
	})

	It("rejects an object that sets neither minecraftVersion nor modpack", func() {
		ms := &metiov1alpha1.MinecraftServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bad-neither",
				Namespace: "default",
			},
			Spec: metiov1alpha1.MinecraftServerSpec{
				ServerID: "srv-3",
			},
		}
		err := k8sClient.Create(ctx, ms)
		Expect(err).To(HaveOccurred())
	})

	It("rejects a shutdown schedule enabled without a time", func() {
		ms := &metiov1alpha1.MinecraftServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bad-shutdown",
				Namespace: "default",
			},
			Spec: metiov1alpha1.MinecraftServerSpec{
				ServerID:         "srv-4",
				MinecraftVersion: "1.21.1",
				ShutdownSchedule: &metiov1alpha1.ShutdownSchedule{
					Enabled: true,
				},
			},
		}
		err := k8sClient.Create(ctx, ms)
		Expect(err).To(HaveOccurred())
	})

	It("accepts a valid modpack server", func() {
		ms := &metiov1alpha1.MinecraftServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pack-server",
				Namespace: "default",
			},
			Spec: metiov1alpha1.MinecraftServerSpec{
				ServerID: "srv-5",
				Modpack: &metiov1alpha1.ModpackConfig{
					Platform:  "modrinth",
					ProjectID: "abc123",
				},
			},
		}
		Expect(k8sClient.Create(ctx, ms)).To(Succeed())
		Expect(k8sClient.Delete(ctx, ms)).To(Succeed())
	})
})

var _ = AfterSuite(func() {
	if cancel != nil {
		cancel()
	}
	if testEnv != nil {
		Eventually(func() error { return testEnv.Stop() }, time.Minute).Should(Succeed())
	}
})
