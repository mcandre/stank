#!/bin/bash
unset IFS
set -eufEo pipefail

trap 'echo cleaning' EXIT

echo "Hello"
