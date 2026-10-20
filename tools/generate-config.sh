#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

if [ ! -f "$REPO_ROOT/metio.env" ]; then
  echo "Error: $REPO_ROOT/metio.env not found."
  echo "Copy metio.env.example to metio.env and fill in your values."
  exit 1
fi

# Parse metio.env safely (no sourcing to avoid shell interpretation)
get_var() {
  local key="$1"
  local default="$2"
  local val
  val="$(grep -E "^${key}=" "$REPO_ROOT/metio.env" | tail -n1 | cut -d'=' -f2-)"
  # Remove surrounding quotes if present
  val="$(echo "$val" | sed -e 's/^"//' -e 's/"$//' -e "s/^'//" -e "s/'$//")"
  if [ -z "$val" ]; then
    echo "$default"
  else
    echo "$val"
  fi
}

GCP_PROJECT="$(get_var GCP_PROJECT TBD)"
GCP_REGION="$(get_var GCP_REGION TBD)"
GCP_ZONE="$(get_var GCP_ZONE TBD)"
ENVIRONMENT="$(get_var ENVIRONMENT development)"
ADMIN_USERS="$(get_var ADMIN_USERS "$(get_var ALLOWED_USERS TBD)")"
ALLOWED_USERS="$(get_var ALLOWED_USERS "$ADMIN_USERS")"
GOOGLE_CLIENT_ID="$(get_var GOOGLE_CLIENT_ID TBD)"
GOOGLE_CLIENT_SECRET="$(get_var GOOGLE_CLIENT_SECRET TBD)"
SESSION_KEY="$(get_var SESSION_KEY demo1234)"
BASE_URL="$(get_var BASE_URL http://localhost:3000)"
INSTANCE_NAME="$(get_var INSTANCE_NAME TBD)"
POSTGRES_MODE="$(get_var POSTGRES_MODE cloudsql)"
POSTGRES_CONNECTION_STRING_SECRET_ID="$(get_var POSTGRES_CONNECTION_STRING_SECRET_ID "")"
BACKUP_DELETED_SERVER_RETENTION_DAYS="$(get_var BACKUP_DELETED_SERVER_RETENTION_DAYS 30)"
BACKUP_CLEANUP_SCHEDULE="$(get_var BACKUP_CLEANUP_SCHEDULE "0 0 * * *")"
OPENCODE_SERVER_PASSWORD="$(get_var OPENCODE_SERVER_PASSWORD change-me)"
GITHUB_ORG="$(get_var GITHUB_ORG "")"
GITHUB_REPO="$(get_var GITHUB_REPO "")"
GCP_ARTIFACT_REGISTRY_PROJECT="$(get_var GCP_ARTIFACT_REGISTRY_PROJECT "")"
GCP_ARTIFACT_REGISTRY_LOCATION="$(get_var GCP_ARTIFACT_REGISTRY_LOCATION "")"
GCP_ARTIFACT_REGISTRY_REPO="$(get_var GCP_ARTIFACT_REGISTRY_REPO "")"
DISTRO_REGISTRY="$(get_var DISTRO_REGISTRY "")"

# Derive git info if not explicitly set
if [ -z "$GITHUB_ORG" ] || [ -z "$GITHUB_REPO" ]; then
  REMOTE_URL="$(git -C "$REPO_ROOT" remote get-url origin 2>/dev/null || echo "")"
  if [ -n "$REMOTE_URL" ]; then
    if echo "$REMOTE_URL" | grep -q "github.com"; then
      GITHUB_ORG_DERIVED="$(echo "$REMOTE_URL" | sed -E 's|.*github.com[:/]+([^/]+)/.*|\1|')"
      GITHUB_REPO_DERIVED="$(echo "$REMOTE_URL" | sed -E 's|.*github.com[:/]+[^/]+/([^./]+).*|\1|')"
      [ -z "$GITHUB_ORG" ] && GITHUB_ORG="$GITHUB_ORG_DERIVED"
      [ -z "$GITHUB_REPO" ] && GITHUB_REPO="$GITHUB_REPO_DERIVED"
    fi
  fi
fi

GITHUB_ORG="${GITHUB_ORG:-metio-plattform}"
GITHUB_REPO="${GITHUB_REPO:-metio}"
GCP_ARTIFACT_REGISTRY_PROJECT="${GCP_ARTIFACT_REGISTRY_PROJECT:-minecraftbyl}"
GCP_ARTIFACT_REGISTRY_LOCATION="${GCP_ARTIFACT_REGISTRY_LOCATION:-europe-west3}"
GCP_ARTIFACT_REGISTRY_REPO="${GCP_ARTIFACT_REGISTRY_REPO:-metio}"
DISTRO_REGISTRY="${DISTRO_REGISTRY:-europe-docker.pkg.dev/metio-distribution/metio}"

mkdir -p "$REPO_ROOT/build" "$REPO_ROOT/.devcontainer" "$REPO_ROOT/deploy"

# Generate devcontainer.env
cat > "$REPO_ROOT/.devcontainer/devcontainer.env" << EOF
GCP_PROJECT=${GCP_PROJECT}
GCP_ZONE=${GCP_ZONE}
INSTANCE_NAME=${INSTANCE_NAME}
ALLOWED_USERS=${ALLOWED_USERS}
GOOGLE_CLIENT_ID=${GOOGLE_CLIENT_ID}
GOOGLE_CLIENT_SECRET=${GOOGLE_CLIENT_SECRET}
SESSION_KEY=${SESSION_KEY}
BASE_URL=${BASE_URL}

# Dapr (local Dapr sidecar + Datastore emulator, started by 'make develop')
# DAPR_STATE_STORE_NAME=statestore

OPENCODE_SERVER_PASSWORD=${OPENCODE_SERVER_PASSWORD}
EOF

echo "Generated: .devcontainer/devcontainer.env"

# Generate metio.auto.tfvars
cat > "$REPO_ROOT/deploy/metio.auto.tfvars" << EOF
# Generated from metio.env - do not edit manually
project_id                           = "${GCP_PROJECT}"
region                               = "${GCP_REGION}"
zone                                 = "${GCP_ZONE}"
admin_users                          = "${ADMIN_USERS}"
environment                          = "${ENVIRONMENT}"
postgres_mode                        = "${POSTGRES_MODE}"
postgres_connection_string_secret_id = "${POSTGRES_CONNECTION_STRING_SECRET_ID}"
backup_deleted_server_retention_days = ${BACKUP_DELETED_SERVER_RETENTION_DAYS}
backup_cleanup_schedule              = "${BACKUP_CLEANUP_SCHEDULE}"
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
