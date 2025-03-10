#!/bin/bash

# TroyOps - GitOps-driven Kubernetes deployment tool
# This script provides a convenient way to use TroyOps

set -e

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    echo "Installation instructions: https://golang.org/doc/install"
    exit 1
fi

# Build TroyOps if the binary doesn't exist
if [ ! -f "./bin/troyops" ]; then
    echo "Building TroyOps..."
    mkdir -p bin
    go build -o bin/troyops cmd/troyops.go
fi

# Run TroyOps with the provided arguments
./bin/troyops "$@"
