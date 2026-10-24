# Стили owo брать из uiu через сабмодуль

Идея: дизайн-систему ведёт uiu (бывший iwwwanow_ui-kit), owo её потребляет. Сейчас направление обратное: uiu копирует `web/static/css` через `scripts/sync-owo.sh`. При переезде источник истины переносится в uiu, `sync-owo.sh` там удаляется, а `web/static/css/index.css` в owo сводится к импорту `/static/uiu/owo.css`.

Подключение — git submodule в `web/static/uiu`, закреплённый на теге uiu (`v0.2.0` и т.д.). Сборку почти не меняет: Dockerfile уже делает `COPY . .` и копирует `web/static` в образ, так что сабмодуль попадает в статику сам. Нужно только, чтобы checkout в `.github/workflows/release.yml` (job `docker`) шёл с `submodules: true`, иначе папка в образе будет пустой. Локально после клонирования — `git submodule update --init`.

Разработка с одновременной правкой uiu — `2026-10-24_dev-override-mount-uiu.md`. Решение о переносе источника правды — `2026-10-24_uiu-source-of-truth.md`.

Отброшено: CDN в рантайме (owo self-hosted, лишняя внешняя зависимость), npm-пакет (ради CSS тянет node-шаг в Go-сборку), subtree (смешивает историю, неудобно отправлять правки обратно).
