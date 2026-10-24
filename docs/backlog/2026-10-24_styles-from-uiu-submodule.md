# Стили owo брать из uiu через сабмодуль

Идея: дизайн-систему ведёт uiu (бывший iwwwanow_ui-kit), owo её потребляет. Сейчас направление обратное: uiu копирует `web/static/css` через `scripts/sync-owo.sh`. При переезде источник истины переносится в uiu, `sync-owo.sh` там удаляется, а `web/static/css/index.css` в owo сводится к импорту `/static/uiu/owo.css`.

Подключение — git submodule в `web/static/uiu`, закреплённый на теге uiu (`v0.2.0` и т.д.). Сборку почти не меняет: Dockerfile уже делает `COPY . .` и копирует `web/static` в образ, так что сабмодуль попадает в статику сам. Нужно только, чтобы checkout в `.github/workflows/release.yml` (job `docker`) шёл с `submodules: true`, иначе папка в образе будет пустой. Локально после клонирования — `git submodule update --init`.

Разработка: docker-compose монтирует `./web/static`, поэтому сабмодуль сразу виден на `/static/uiu/`. Для одновременной правки uiu и owo — gitignored `docker-compose.override.yml`, который монтирует рабочую копию `../iwwwanow_uiu` поверх `/var/www/owo/static/uiu`. Правки видны вживую, без коммитов в сабмодуль. Когда uiu готов: тег в uiu → `git -C web/static/uiu checkout vX.Y.Z` → коммит указателя в owo.

Отброшено: CDN в рантайме (owo self-hosted, лишняя внешняя зависимость), npm-пакет (ради CSS тянет node-шаг в Go-сборку), subtree (смешивает историю, неудобно отправлять правки обратно).
