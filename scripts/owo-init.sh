#!/bin/bash
set -e

echo "=== OWO SSH Initialization ==="

OWO_DIR="./owo"
SSH_KEYS_DIR="$OWO_DIR/ssh_keys"

# Создать базовые директории
mkdir -p "$OWO_DIR/uploads"
mkdir -p "$SSH_KEYS_DIR"

# Определить публичный ключ
if [ -n "$1" ]; then
    # Использовать переданный ключ как аргумент
    PUBLIC_KEY="$1"
    echo "Using provided SSH public key"
# elif [ -n "$AUTHORIZED_KEY" ]; then
#     # Использовать ключ из переменной окружения
#     PUBLIC_KEY="$AUTHORIZED_KEY"
#     echo "Using AUTHORIZED_KEY from environment"
# elif [ -f ~/.ssh/id_ed25519.pub ]; then
#     # Использовать ключ по умолчанию
#     PUBLIC_KEY=$(cat ~/.ssh/id_ed25519.pub)
#     echo "Using default SSH key from ~/.ssh/id_ed25519.pub"
# elif [ -f ~/.ssh/id_rsa.pub ]; then
#     PUBLIC_KEY=$(cat ~/.ssh/id_rsa.pub)
#     echo "Using default SSH key from ~/.ssh/id_rsa.pub"
else
    echo "Error: No SSH public key found"
    echo "Usage: $0 [ssh-public-key]"
    echo "Or set AUTHORIZED_KEY environment variable"
    exit 1
fi

# Проверить формат ключа
if ! echo "$PUBLIC_KEY" | grep -q "^ssh-"; then
    echo "Error: Invalid SSH public key format"
    exit 1
fi

# Создать authorized_keys
echo "$PUBLIC_KEY" > "$OWO_DIR/authorized_keys"
chmod 600 "$OWO_DIR/authorized_keys"
echo "✓ Created $OWO_DIR/authorized_keys"

# Генерация только RSA SSH host key
if [ ! -f "$SSH_KEYS_DIR/ssh_host_rsa_key" ]; then
    echo "Generating SSH host RSA key..."
    ssh-keygen -t rsa -b 4096 -f "$SSH_KEYS_DIR/ssh_host_rsa_key" -N ""
    chmod 600 "$SSH_KEYS_DIR/ssh_host_rsa_key"
    chmod 644 "$SSH_KEYS_DIR/ssh_host_rsa_key.pub"
    echo "✓ Generated RSA host key"
else
    echo "✓ RSA host key already exists"
fi

echo ""
echo "=== Initialization Complete ==="
echo "Directory structure:"
echo "  $OWO_DIR/uploads/           - Uploaded files"
echo "  $OWO_DIR/authorized_keys    - SSH authorized keys"
echo "  $SSH_KEYS_DIR/              - SSH host keys (RSA only)"
echo ""
echo "To start: docker-compose up"
