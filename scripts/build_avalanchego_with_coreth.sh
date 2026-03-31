#!/bin/bash

# This script builds a new AvalancheGo binary with the Coreth dependency pointing to the local Coreth path
# Usage: ./build_avalanchego_with_coreth.sh with optional AVALANCHEGO_VERSION and METALGO_CLONE_PATH environment variables

set -euo pipefail

# Coreth root directory
CORETH_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Allow configuring the clone path to point to an existing clone
METALGO_CLONE_PATH="${METALGO_CLONE_PATH:-metalgo}"

# Load the version
source "$CORETH_PATH"/scripts/versions.sh

# Always return to the coreth path on exit
function cleanup {
  cd "${CORETH_PATH}"
}
trap cleanup EXIT

echo "checking out target MetalGo version ${METAL_VERSION}"
if [[ -d "${METALGO_CLONE_PATH}" ]]; then
  echo "updating existing clone"
  cd "${METALGO_CLONE_PATH}"
  git fetch
else
  echo "creating new clone"
  git clone https://github.com/MetalBlockchain/metalgo.git "${METALGO_CLONE_PATH}"
  cd "${METALGO_CLONE_PATH}"
fi
# Branch will be reset to $METAL_VERSION if it already exists
git checkout -B "test-${METAL_VERSION}" "${METAL_VERSION}"

echo "updating coreth dependency to point to ${CORETH_PATH}"
go mod edit -replace "github.com/MetalBlockchain/coreth=${CORETH_PATH}"
go mod tidy

echo "building metalgo"
./scripts/build.sh
