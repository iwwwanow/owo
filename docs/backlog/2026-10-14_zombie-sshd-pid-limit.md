# Зомби sshd выедают лимит pids — sftp в контейнере умирает

Вывод (2026-10-14, прод на ru-timeweb): в обоих контейнерах owo (1.12.0 и 1.7.0) по ~1095 зомби `sshd-auth`, `pids.current` упёрся в `pids.max` 1107 — новые процессы не создаются (`docker exec … sh` → `can't fork`), sftp не принимает подключения. Причина — `scripts/docker-entrypoint.sh`: `sshd -D &`, затем `exec /app/server`, и PID 1 становится Go-сервер, который не подбирает осиротевших детей (`wait`). Каждое подключение к sftp-порту — в основном сканеры из интернета — оставляет зомби, и примерно через пару месяцев аптайма лимит исчерпан.

Фикс в образе — нормальный init как PID 1: `tini` (`apk add tini`, `ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]`) или supervisor для двух процессов. До тех пор в compose — `init: true` (так сделано в `iwwwanow_infrastructure/stacks/ru-timeweb/owo/compose.yml`). Заодно: если один из двух процессов падает, контейнер этого не замечает — sshd в фоне никто не перезапускает.
