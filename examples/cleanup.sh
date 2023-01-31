#!/bin/sh
unset IFS
set -euf

trap 'echo cleaning' EXIT

echo "Hello"
