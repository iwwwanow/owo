# Источник правды стилей owo переносится в uiu

Решение: дизайн-систему ведёт uiu (`iwwwanow/uiu`), owo её только потребляет. Сейчас направление обратное: uiu копирует `web/static/css` скриптом `scripts/sync-owo.sh`, а свои отличия держит в `scripts/owo-deviations.css`. После переноса `owo.css` правится прямо в uiu, а `sync-owo.sh` и отдельный файл расхождений удаляются: расхождения становятся обычными правками в uiu. В owo `web/static/css` сводится к импорту стилей uiu. Заметки про найденные баги стилей (`code-font-size-compounds`, `roboto-mono-not-loaded`, `roboto-imported-unused`, `hr-standalone-component`) после переноса чинятся уже в uiu.

Как подключать — `2026-10-24_styles-from-uiu-submodule.md`, dev-режим — `2026-10-24_dev-override-mount-uiu.md`.
