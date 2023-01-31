#!/bin/zsh
unset IFS
set -eufo pipefail

# trap 'echo cleaning' EXIT

TRAPEXIT() {
    echo 'cleaning'
}

echo "Hello"
