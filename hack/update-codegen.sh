#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${SCRIPT_ROOT}/hack

source "${CODEGEN_PKG}/kube_codegen.sh"

THIS_PKG="github.com/fluxcd/flagger"

# Print the absolute path of the input directory for debugging
INPUT_DIR=$(realpath "${SCRIPT_ROOT}/pkg/apis")
echo ">> Input directory: ${INPUT_DIR}"

# Ensure the input directory exists
if [ ! -d "${INPUT_DIR}" ]; then
    echo "Input directory ${INPUT_DIR} does not exist" >&2
    exit 1
fi

# Generate helpers
kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${INPUT_DIR}"

# Generate client code
kube::codegen::gen_client \
    --with-watch \
    --output-dir "$(realpath "${SCRIPT_ROOT}/pkg/client")" \
    --output-pkg "${THIS_PKG}/pkg/client" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${INPUT_DIR}"