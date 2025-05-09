#!/usr/bin/env sh

# Get OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    Darwin) OS="darwin" ;;
    Linux) OS="linux" ;;
    Windows) OS="windows" ;;
esac

case $ARCH in
    aarch64) ARCH="arm64" ;;
esac

LATEST_VERSION=$(curl -s https://api.github.com/repos/alexandre-girault/helm-rufin/releases/latest |tr ',' '\n'|grep '"name":' |cut -d'"' -f4 |head -1)
ARCHIVE="rufin_${LATEST_VERSION}_${OS}_${ARCH}"
URL="https://github.com/alexandre-girault/helm-rufin/releases/download/${LATEST_VERSION}/rufin-${OS}-${LATEST_VERSION}-${ARCH}"

mkdir -p "$HELM_PLUGIN_DIR/bin"
chmod 755 "$HELM_PLUGIN_DIR/bin"
chmod +x "$HELM_PLUGIN_DIR/bin/run.sh"


echo "Downloading $URL"
curl -Lo $HELM_PLUGIN_DIR/bin/rufin "${URL}"
chmod +x "$HELM_PLUGIN_DIR/bin/rufin"
