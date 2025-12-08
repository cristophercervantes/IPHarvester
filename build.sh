#!/usr/bin/env bash
set -e

BINARY="ipharvester"
VERSION="v1.0"
TARGET="/usr/local/bin/$BINARY"

echo "Building IPHarvester $VERSION"
echo "Go: $(go version | awk '{print $3,$4}')"

# Build static binary + inject version
CGO_ENABLED=0 go build \
    -ldflags="-s -w -X github.com/cristophercervantes/IPHarvester/cmd.version=$VERSION" \
    -o "$BINARY"

# Optional: compress with UPX
if command -v upx &>/dev/null; then
    echo "Compressing with UPX..."
    upx --best --lzma "$BINARY" >/dev/null 2>&1
fi

echo "Installing → $TARGET"
sudo install -m 755 "$BINARY" "$TARGET"
rm -f "$BINARY"

echo
echo "ipharvester $VERSION installed globally"
echo
echo "Testing version:"
$TARGET -v          # ← correct way: use the -v flag
echo
echo "Done. Use: ipharvester -v   or   ipharvester --version"
