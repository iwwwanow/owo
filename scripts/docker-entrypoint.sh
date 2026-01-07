#!/bin/sh
set -e

mkdir -p /root/.ssh
chmod 700 /root/.ssh

# TODO: erro if not exist
if [ -f /root/.ssh/authorized_keys ]; then
    chown root:root /root/.ssh/authorized_keys
    chmod 600 /root/.ssh/authorized_keys
fi

chmod 600 /etc/ssh/ssh_host_*_key 2>/dev/null || true
chmod 644 /etc/ssh/ssh_host_*_key.pub 2>/dev/null || true

echo "Starting SSH daemon..."
/usr/sbin/sshd -D &

echo "Starting application..."
exec /app/server
