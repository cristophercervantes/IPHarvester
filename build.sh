#!/bin/bash
# build.sh — Build IPHarvester and drop it where it belongs

set -e

BINARY="ipharvester"
VERSION="v0.0.2"
TARGET_DIR="/usr/local/bin"
BUILD_DIR=$(mktemp -d)

echo "Building $BINARY $VERSION with Go $(go version | awk '{print $3,$4}')"

# Build static binary (no CGO bullshit)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=$VERSION" -o "$BUILD_DIR/$BINARY" .

# Strip & compress (optional, makes it tiny)
strip "$BUILD_DIR/$BINARY" 2>/dev/null || true
upx -9 "$BUILD_DIR/$BINARY" 2>/dev/null || echo "[!] upx not installed, skipping compression"

# Install to system (root required)
echo "Installing $BINARY to $TARGET_DIR"
sudo install -m 755 "$BUILD_DIR/$BINARY" "$TARGET_DIR/"

# Cleanup
rm -rf "$BUILD_DIR"

echo
echo "IPHarvester $VERSION installed globally!"
echo "Run: $BINARY version"
echo
$BINARY version
