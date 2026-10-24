# Dev-режим owo + uiu: override-монтирование рабочей копии uiu

Чтобы одновременно править uiu и owo без коммитов в сабмодуль, нужен gitignored `docker-compose.override.yml` в owo. Он монтирует соседнюю рабочую копию `../iwwwanow_uiu` поверх `/var/www/owo/static/uiu`; docker compose подхватывает override автоматически. Правки CSS в uiu сразу видны в owo, указатель сабмодуля при этом не трогается. Когда uiu готов: тег в uiu (semantic-release) → `git -C web/static/uiu checkout vX.Y.Z` → коммит указателя в owo. Без override owo работает на закреплённой версии сабмодуля.

```yaml
# docker-compose.override.yml (gitignored)
services:
  air:
    volumes:
      - ../iwwwanow_uiu:/var/www/owo/static/uiu
```

Сервис `air` — dev-сервис из `docker-compose.yml`. Та же заметка есть в uiu.
