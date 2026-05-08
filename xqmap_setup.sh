#!/bin/bash
BINARY_NAME="xqmap"
SOURCE_FILE="main.go"
INSTALL_PATH="$HOME/.local/bin/$BINARY_NAME"

echo "Compiling $SOURCE_FILE..."
go build -o "$BINARY_NAME" "$SOURCE_FILE"

if [ $? -ne 0 ]; then
    echo "Compilation failed"
    exit 1
fi

mkdir -p "$HOME/.local/bin"
mv "$BINARY_NAME" "$INSTALL_PATH"
chmod +x "$INSTALL_PATH"

echo "🚀 xqmap installed to $INSTALL_PATH"