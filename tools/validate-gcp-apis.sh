#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

PROJECT_ID=""

# Try from metio.env
if [ -z "$PROJECT_ID" ] && [ -f "$REPO_ROOT/metio.env" ]; then
  PROJECT_ID="$(grep -E '^GCP_PROJECT=' "$REPO_ROOT/metio.env" | cut -d'=' -f2- | tr -d '"' | tr -d "'" | head -n1)"
fi

# Try from deploy/metio.auto.tfvars
if [ -z "$PROJECT_ID" ] && [ -f "$REPO_ROOT/deploy/metio.auto.tfvars" ]; then
  PROJECT_ID="$(grep -m1 '^[[:space:]]*project_id[[:space:]]*=' "$REPO_ROOT/deploy/metio.auto.tfvars" | sed 's/.*=\s*"\([^"]*\)".*/\1/' | sed "s/.*=\s*'\([^']*\)'.*/\1/")"
fi

if [ -z "$PROJECT_ID" ] || [ "$PROJECT_ID" = "TBD" ]; then
  echo "ERROR: Could not determine GCP project ID."
  echo "Set it via 'gcloud config set project <PROJECT_ID>', define GCP_PROJECT in metio.env, or project_id in deploy/metio.auto.tfvars"
  exit 1
fi

echo "Validating required GCP APIs for project: $PROJECT_ID"
REQUIRED_APIS=(
  "cloudtasks.googleapis.com"
  "secretmanager.googleapis.com"
  "run.googleapis.com"
  "cloudscheduler.googleapis.com"
  "compute.googleapis.com"
  "iam.googleapis.com"
  "cloudresourcemanager.googleapis.com"
)

MISSING_APIS=()
for api in "${REQUIRED_APIS[@]}"; do
  if gcloud services list --enabled --project="$PROJECT_ID" --filter="name:$api" --format="value(name)" 2>/dev/null | grep -q "$api"; then
    echo "  ✓ $api enabled"
  else
    echo "  ✗ $api NOT enabled"
    MISSING_APIS+=("$api")
  fi
done

if [ ${#MISSING_APIS[@]} -gt 0 ]; then
  echo ""
  echo "ERROR: Missing required GCP APIs!"
  echo ""
  echo "Enable them with:"
  for api in "${MISSING_APIS[@]}"; do
    echo "  gcloud services enable $api --project=$PROJECT_ID"
  done
  echo ""
  exit 1
fi

echo ""
echo "All required GCP APIs are enabled."
exit 0
