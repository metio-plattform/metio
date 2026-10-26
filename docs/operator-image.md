# Metio operator image

The `metio-operator` image is built lazily via the `metio-operator-image` Makefile target and pushed to Google Artifact Registry, like the other component images.

## k3s pull path

k3s on the standalone COS runtime pulls the operator image from Artifact Registry. Configure one of:

1. **Image pull secret** (Kubernetes): create a secret of type `kubernetes.io/dockerconfigjson` from a service account key with `roles/artifactregistry.reader` and reference it as `imagePullSecrets` on the operator pod.
2. **containerd config** (k3s node): add the Artifact Registry host to `/etc/rancher/k3s/registries.yaml` (or the containerd `config.toml` `plugins."io.containerd.grpc.v1.cri".registry.mirrors`) with the matching `configs.<host>.auth` credentials, then restart k3s.

The image runs as non-root (UID 65532).