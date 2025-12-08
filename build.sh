#!/usr/bin/env bash
set -e

BINARY="ipharvester"
VERSION="v1.0"
TARGET="/usr/local/bin/$BINARY"

echo "Building IPHarvester $VERSION"
echo "Go: $(go version | awk '{print $3,$4}')"

# Build with correct version injected
CGO_ENABLED=0 go build \
    -ldflags="-s -w -X github.com/cristophercervantes/IPHarvester/cmd.version=$VERSION" \
    -o "$BINARY"

# Optional UPX
if command -v upx &>/dev/null; then
    echo "Compressing with UPX..."
    upx --best --lzma "$BINARY" >/dev/null 2>&1
fi

echo "Installing to $TARGET"
sudo install -m 755 "$BINARY" "$TARGET"
rm -f "$BINARY"

echo
echo "ipharvester $VERSION installed."
echo
$TARGET version
