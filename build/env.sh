#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
eledir="$workspace/src/github.com/elementrem"
if [ ! -L "$eledir/go-elementrem" ]; then
    mkdir -p "$eledir"
    cd "$eledir"
    ln -s ../../../../../. go-elementrem
    cd "$root"
fi

# Set up the environment to use the workspace.
# Also add Godeps workspace so we build using canned dependencies.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$eledir/go-elementrem"
PWD="$eledir/go-elementrem"

# Launch the arguments with the configured environment.
exec "$@"
