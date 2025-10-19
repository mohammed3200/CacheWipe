#!/bin/bash

# Installer for Cache wipe

set -e

INSTALL_DIR="/usr/local/bin"
APP_NAME="cache-wipe"
RELEASE_URL="https://github.com/mohammed3200/cache-wipe/releases/latest"

echo "🚀 Installing Cache wipe..."

# Detect OS and Architecture
OS=$(uname -s)
ARCH=$(uname -m)

case $OS in
    Linux)
        if [ "$ARCH" = "x86_64" ]; then
            BINARY="cache-wipe-linux-amd64"
        elif [ "$ARCH" = "arm64" ]; then
            BINARY="cache-wipe-linux-arm64"
        fi
        ;;
    Darwin)
        if [ "$ARCH" = "x86_64" ]; then
            BINARY="cache-wipe-darwin-amd64"
        elif [ "$ARCH" = "arm64" ]; then
            BINARY="cache-wipe-darwin-arm64"
        fi
        ;;
    *)
        echo "❌ Unsupported OS: $OS"
        exit 1
        ;;
esac

if [ -z "$BINARY" ]; then
    echo "❌ Unsupported architecture: $ARCH"
    exit 1
fi

echo "📥 Downloading $BINARY..."

# Download the binary
if ! command -v curl &> /dev/null; then
    echo "❌ curl is required but not installed."
    exit 1
fi

curl -L "$RELEASE_URL/download/$BINARY" -o "/tmp/$BINARY"

echo "📦 Installing to $INSTALL_DIR..."
sudo mv "/tmp/$BINARY" "$INSTALL_DIR/$APP_NAME"
sudo chmod +x "$INSTALL_DIR/$APP_NAME"

echo "✅ Installation complete!"
echo "Run 'cache-wipe' to start the application"