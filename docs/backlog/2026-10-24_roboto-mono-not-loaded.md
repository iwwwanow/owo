# Roboto Mono указан у code, но не загружается

В `web/static/css/font.css` у `code` стоит `font-family: "Roboto Mono", monospace`, но Roboto Mono нигде не подключается: в `@import` есть только Roboto Condensed и Roboto. Поэтому код рендерится системным monospace, а он на разных ОС разный. Только у тех, у кого Roboto Mono установлен локально, код выглядит так, как задумано. Нужно добавить Roboto Mono (400) в `@import` шрифтов, либо осознанно перейти на системный monospace и убрать имя шрифта из правила.
