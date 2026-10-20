#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Source unified config if it exists
if [ -f "$REPO_ROOT/metio.env" ]; then
  # shellcheck source=/dev/null
  set -a
  . "$REPO_ROOT/metio.env"
  set +a
else
  echo "Error: $REPO_ROOT/metio.env not found."
  echo "Copy metio.env.example to metio.env and fill in your values."
  exit 1
fi

# Derive git info if not explicitly set
if [ -z "${GITHUB_ORG:-}" ] || [ -z "${GITHUB_REPO:-}" ]; then
  REMOTE_URL="$(git -C "$REPO_ROOT" remote get-url origin 2>/dev/null || echo "")"
  if [ -n "$REMOTE_URL" ]; then
    # Handle https://github.com/org/repo.git or git@github.com:org/repo.git
    if echo "$REMOTE_URL" | grep -q "github.com"; then
      GITHUB_ORG_DERIVED="$(echo "$REMOTE_URL" | sed -E 's|.*github.com[:/]+([^/]+)/.*|\1|')"
      GITHUB_REPO_DERIVED="$(echo "$REMOTE_URL" | sed -E 's|.*github.com[:/]+[^/]+/([^./]+).*|\1|')"
      GITHUB_ORG="${GITHUB_ORG:-$GITHUB_ORG_DERIVED}"
      GITHUB_REPO="${GITHUB_REPO:-$GITHUB_REPO_DERIVED}"
    fi
  fi
fi

GITHUB_ORG="${GITHUB_ORG:-metio-plattform}"
GITHUB_REPO="${GITHUB_REPO:-metio}"
GCP_ARTIFACT_REGISTRY_PROJECT="${GCP_ARTIFACT_REGISTRY_PROJECT:-minecraftbyl}"
GCP_ARTIFACT_REGISTRY_LOCATION="${GCP_ARTIFACT_REGISTRY_LOCATION:-europe-west3}"
GCP_ARTIFACT_REGISTRY_REPO="${GCP_ARTIFACT_REGISTRY_REPO:-metio}"
DISTRO_REGISTRY="${DISTRO_REGISTRY:-europe-docker.pkg.dev/metio-distribution/metio}"

# Generate devcontainer.env
cat > "$REPO_ROOT/.devcontainer/devcontainer.env" << EOF
GCP_PROJECT=${GCP_PROJECT:-TBD}
GCP_ZONE=${GCP_ZONE:-TBD}
INSTANCE_NAME=${INSTANCE_NAME:-TBD}
ALLOWED_USERS=${ALLOWED_USERS:-${ADMIN_USERS:-TBD}}
GOOGLE_CLIENT_ID=${GOOGLE_CLIENT_ID:-TBD}
GOOGLE_CLIENT_SECRET=${GOOGLE_CLIENT_SECRET:-TBD}
SESSION_KEY=${SESSION_KEY:-demo1234}
BASE_URL=${BASE_URL:-http://localhost:3000}

# Dapr (local Dapr sidecar + Datastore emulator, started by 'make develop')
# DAPR_STATE_STORE_NAME=statestore

OPENCODE_SERVER_PASSWORD=${OPENCODE_SERVER_PASSWORD:-change-me}
EOF

echo "Generated: .devcontainer/devcontainer.env"

# Generate metio.auto.tfvars
cat > "$REPO_ROOT/deploy/metio.auto.tfvars" << EOF
# Generated from metio.env - do not edit manually
project_id                           = "${GCP_PROJECT:-TBD}"
region                               = "${GCP_REGION:-${GCP_ZONE%???}TBD}" # try to derive region from zone if similar
region                               = "${GCP_REGION:-TBD}"
zone                                 = "${GCP_ZONE:-TBD}"
admin_users                          = "${ADMIN_USERS:-TBD}"
environment                          = "${ENVIRONMENT:-development}"
postgres_mode                        = "${POSTGRES_MODE:-cloudsql}"
postgres_connection_string_secret_id = "${POSTGRES_CONNECTION_STRING_SECRET_ID:-}"
backup_deleted_server_retention_days = ${BACKUP_DELETED_SERVER_RETENTION_DAYS:-30}
backup_cleanup_schedule              = "${BACKUP_CLEANUP_SCHEDULE:-0 0 * * *}"
EOF

# Fix the region line duplication - write clean version
cat > "$REPO_ROOT/deploy/metio.auto.tfvars" << EOF
# Generated from metio.env - do not edit manually
project_id                           = "${GCP_PROJECT:-TBD}"
region                               = "${GCP_REGION:-TBD}"
zone                                 = "${GCP_ZONE:-TBD}"
admin_users                          = "${ADMIN_USERS:-TBD}"
environment                          = "${ENVIRONMENT:-development}"
postgres_mode                        = "${POSTGRES_MODE:-cloudsql}"
postgres_connection_string_secret_id = "${POSTGRES_CONNECTION_STRING_SECRET_ID:-}"
backup_deleted_server_retention_days = ${BACKUP_DELETED_SERVER_RETENTION_DAYS:-30}
backup_cleanup_schedule              = "${BACKUP_CLEANUP_SCHEDULE:-0 0 * * *}"
EOF

echo "Generated: deploy/metio.auto.tfvars"

# Write derived registry info for Makefile to consume
cat > "$REPO_ROOT/build/registry.env" << EOF
GITHUB_ORG=${GITHUB_ORG}
GITHUB_REPO=${GITHUB_REPO}
GCP_ARTIFACT_REGISTRY_PROJECT=${GCP_ARTIFACT_REGISTRY_PROJECT}
GCP_ARTIFACT_REGISTRY_LOCATION=${GCP_ARTIFACT_REGISTRY_LOCATION}
GCP_ARTIFACT_REGISTRY_REPO=${GCP_ARTIFACT_REGISTRY_REPO}
DISTRO_REGISTRY=${DISTRO_REGISTRY}
EOF
echo "Generated: build/registry.env"
