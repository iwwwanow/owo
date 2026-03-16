# Roadmap owo → v2

> Основан на: `docs/dev/planning.md`, `docs/dev/backlog.md`, `docs/dev/desktop-planning.md`
> Текущее состояние: **v1.7.0**, semantic-release настроен, CI/CD работает

---

## v1.8 — Конфиг и файловая система

Цель: сделать поведение owo управляемым через конфиг и упорядочить файловую систему.

- [ ] `config.json` / `config.yaml` в корне uploads-директории
- [ ] параметр `hideItems` / `show-hidden-resources` в конфиге
- [ ] нейминг-конвенции:
  - `.` — директории и файлы скрыты
  - `!` — в самом верху, расширенная карточка
  - `#` — в самом низу, укороченная карточка
- [ ] визуальное разделение директорий и файлов (иконка / отступ / HR)
- [ ] render 404 и страниц ошибок
- [ ] render error message (например, при git pull failed)
- [ ] возможность расширять футер через `.meta/` контент

---

## v1.9 — Медиа и Git-интеграция

Цель: лучшая работа с изображениями и первичная поддержка git-директорий.

- [ ] если в `.meta` нет обложки — выбирать первую картинку из директории
- [ ] urlParam для получения изображения в нужном размере
- [ ] кеширование изображений с периодической чисткой (~раз в 5 дней)
- [ ] поддержка видео
- [ ] иконки / nerd fonts для обложек файлов (md, css, js и пр.)
- [ ] если директория — git-репозиторий: визуально выделять (карточка + страница)
- [ ] git pull кешировать, не чаще раза в 5 минут
- [ ] `originUrl` в конфиге директории — git remote
- [ ] символические ссылки / нейминг-alias для ресурсов без дублирования места

---

## v1.10 — Соцсети, API и UX

Цель: owo как полноценный публичный портфолио-хостинг.

- [ ] metainfo (og/meta-теги)
- [ ] social media preview
- [ ] API: ручка на получение всех айтемов со страницы
- [ ] перелинковка изображений и директорий
- [ ] web components для отрисовки айтемов по категориям (ref: timurnovikov.com/en)
- [ ] extended-page: полный рендер страницы-файла (cover, контент, прочие метаданные)
- [ ] витрина с перелинковкой (нужна витрина с золотыми рамами)
- [ ] перенести сайт на домен `iwwwanow.ru`
- [ ] опубликовать резюме и CV на owo

---

## v1.11 — CI/CD и инфра

Цель: автономная и отказоустойчивая инфраструктура.

- [ ] telegram-уведомления при релизе (semantic-release-telegram)
- [ ] FEAT stage branch + stage subdomain
- [ ] автообновление образов на хосте (auto-updater, ref: doco-cd)
- [ ] fast install: curl/скрипт для быстрой установки на сервере и localhost
- [ ] uploads backup (curl → archive, named docker volumes)
- [ ] save host keys при пересборке образа (постоянный ключ в образе)
- [ ] scripts.sh для infra (pull, clean, deploy)
- [ ] images list в infra config.yaml

---

## v2.0 — UI Kit + Desktop

Цель: вынести дизайн-систему в самостоятельный пакет и выпустить desktop-клиент.

### @owo-ui — design system

Монорепо на pnpm workspaces:

```
@owo-ui/
  tokens/   — CSS-переменные + JS-экспорт (colors, spaces, layout)
  css/       — reset, layout, components, fonts — чистый CSS
  svelte/    — Svelte-компоненты (Card, Hr, Header, Footer...)
  react/     — React-компоненты с теми же именами
```

- [ ] вынести `web/static/css/` в пакет `@owo-ui/css`
- [ ] сконвертировать `colors.css`, `spaces.css`, `layout.css` в `tokens.css` + `tokens.js`
- [ ] Svelte-компоненты: Card, Hr, Header, Footer
- [ ] React-компоненты с теми же именами
- [ ] сборка: Vite + `@sveltejs/package` / Vite lib mode
- [ ] опубликовать пакеты (npm / GitHub Packages)

### Desktop-клиент (Windows)

Иконка в трее + директория в проводнике через SFTP (аналог Google Drive / Яндекс.Диск).

**Рекомендованный путь:** WinFSP (user-mode, не нужна подпись драйвера) или WebDAV-прокси поверх SFTP как быстрый старт.

- [ ] трей + UI (Tauri предпочтительно, C#/WinForms как fallback)
- [ ] SFTP-подключение (SSH.NET / libssh2)
- [ ] виртуальная FS через WinFSP или WebDAV-прокси
- [ ] использует `@owo-ui/css` без адаптаций (Tauri — WebView)

### Архитектура (рефакторинг к v2)

- [ ] DI на entrypoint: `repo → useCase → handler` (см. `di-on-entrypoint.md`)
- [ ] архитектурная схема: domain / app / infra (adapters, presentation)
- [ ] возможное разделение репо на `site` и `system`

---

## Постепенно / без версии

- [ ] fluxCd / k3s / microk8s — если переход на оркестрацию
- [ ] обновить структуру uploads (scope_year_name, добавить обложки)
- [ ] чёрный/инвертированный шрифт на карточках без цветовых зависимостей
