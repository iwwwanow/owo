FROM golang:1.24.6 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /server ./cmd/server/main.go

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
COPY --from=builder /app/static /app/static
COPY --from=builder /app/templates /app/templates

COPY sshd_config /etc/ssh/sshd_config
COPY authorized_keys /root/.ssh/authorized_keys

RUN chmod 600 /root/.ssh/authorized_keys \
	&& chmod 600 /etc/ssh/sshd_config

EXPOSE 22 8080

ENV PORT=8080 \
	TZ=Europe/Moscow \
	PUBLIC_DIR=/var/www/owwo/shared

VOLUME ["/var/www/owwo/shared"]

CMD ["sh", "-c", "/usr/sbin/sshd -D & /app/server"]
