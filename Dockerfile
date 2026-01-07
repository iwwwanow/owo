FROM golang:1.24.6 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /server ./cmd/main.go

FROM alpine:3.22.1
WORKDIR /app

RUN apk add --no-cache \
	git \
	openssh \
	&& echo "root:root" | chpasswd \
	&& mkdir -p /var/run/sshd \
	&& mkdir -p /root/.ssh \
	&& chmod 700 /root/.ssh

COPY --from=builder /server /app/server
COPY --from=builder /app/web/static /var/www/owo/static
COPY --from=builder /app/web/templates /var/www/owo/templates

# Копируем entrypoint
COPY --from=builder /app/scripts/docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# TODO: Опционально: дефолтный authorized_keys
# COPY docker/authorized_keys.default /authorized_keys.default

COPY configs/sshd_config /etc/ssh/sshd_config
# COPY authorized_keys /root/.ssh/authorized_keys

# RUN chmod 600 /root/.ssh/authorized_keys
RUN chmod 600 /etc/ssh/sshd_config

EXPOSE 22 8080

ENV PORT=8080 \
	TZ=Europe/Moscow \
	PUBLIC_DIR=/var/www/owwo/shared

# TODO: for what?
# VOLUME ["/var/www/owo/shared"]

# CMD ["sh", "-c", "/usr/sbin/sshd -D & /app/server"]
ENTRYPOINT ["/entrypoint.sh"]
