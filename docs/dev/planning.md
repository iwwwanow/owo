# owo — планирование

- [ ] meta preview

## CI/CD и деплой

- [x] настроить версионирование при мердже в мастер (semantic-release)
- [x] выпустить v1
- [ ] telegram-уведомления при релизе
  - ref: https://github.com/pustovitDmytro/semantic-release-telegram
  - ref: https://github.com/skoropadas/semantic-release-telegram-bot
- [ ] сформировать roadmap до v2
- [ ] FEAT stage branch; stage subdomain

### Инфра / автообновление

- [ ] автообновление образов на хосте (auto-updater-image)
  - ref: https://github.com/kimdre/doco-cd
- [ ] github curl → owo: setup app
- [ ] github curl → infra: launch all apps
- [ ] fast install: curl/script для быстрой установки на сервере и localhost
- [ ] images list on infra (config.yaml)
- [ ] scripts.sh on infra (pull, clean & deploy.sh)
- [ ] fluxCd / k3s / microk8s (если переход на оркестрацию)
- [ ] uploads backup
  - [ ] curl backup to archive
  - [ ] переносимые данные в named docker named-volumes

---

## Архитектура

- [ ] архитектурная схема: domain, app, infra (adapters, presentation/view)

---

## Конфиг

- [ ] config.(json/yaml) в корне директории uploads
- [ ] config on page (отображать текущие настройки)
- [ ] параметр `hideItems` / `show-hidden-resources`

---

## Файловая система / нейминг

- [ ] `.` — директории и файлы скрыты
- [ ] `!` — директории и файлы в самом верху (расширенная карточка)
- [ ] `#` — директории и файлы в самом низу (укороченная карточка)
- [ ] разделять директории и файлы визуально (иконка / отступ / HR)
- [ ] ярлыки/символические ссылки или специфический нейминг для alias без дублирования места

---

## Git / SSH

- [ ] добавить в конфиг `originUrl` для директорий — git remote
- [ ] git pull кешировать, не чаще раз в 5 минут
- [ ] если директория — git-репозиторий, выделять стилистически (карточка + страница)
- [ ] save host keys при пересборке образа (не терять known_hosts)
- [ ] доработать SSH-подключение

---

## Изображения и медиа

- [ ] urlParam для получения изображения в специфичном размере
- [ ] кеширование изображений с периодической чисткой (раз в ~5 дней)
- [ ] если в `.meta` не задана обложка — выбирать первую картинку из директории
- [ ] FEAT cover для md, css и других файлов (nerd fonts / иконки)
- [ ] videos

---

## Рендеринг и фичи

- [ ] render error message (например, при git pull failed)
- [ ] render 404 и другие страницы ошибок
- [ ] render hidden resources / render hidden folders
- [x] FEAT html → iframe
- [ ] FEAT отрисовка страниц-файлов (extended-page)
- [ ] добавить перелинковку изображений и директорий
- [ ] нужна возможность расширять футер (пользователь кладёт контент в `.meta`, он встраивается в футер страницы)
- [ ] нужна витрина с перелинковкой и золотыми рамами

---

## API

- [ ] ручка на получение всех айтемов со страницы

---

## UI и стили

- [x] clean css
- [ ] просто чёрный/инвертированный шрифт на карточках (без цветовых зависимостей)
- [ ] web components для отрисовки айтемов по категориям
  - ref: https://timurnovikov.com/en
- [x] identical layout on mobile — 2 колонки, без hover-эффектов
- [ ] metainfo (og/meta-теги)
- [ ] social media preview
- [x] github link preview

---

## UI Kit

Идея: вынести дизайн приложения в отдельный пакет, применимый к Svelte, React и десктопу (Tauri).

### Структура монорепо

```
@owo-ui/
  tokens/   — CSS-переменные + JS-экспорт (colors, spaces, layout)
  css/       — reset, layout, components, fonts — чистый CSS
  svelte/    — Svelte-компоненты (Card, Hr, Header, Footer...)
  react/     — React-компоненты с теми же именами
```

### Задачи

- [ ] вынести `web/static/css/` в отдельный пакет `@owo-ui/css`
- [ ] сконвертировать `colors.css`, `spaces.css`, `layout.css` в `tokens.css` + `tokens.js`
- [ ] реализовать Svelte-компоненты (Card, Hr, Header, Footer)
- [ ] реализовать React-компоненты
- [ ] настроить сборку: Vite + `@sveltejs/package` / Vite lib mode
- [ ] опубликовать пакеты (npm / GitHub Packages)

### Инструменты

- Monorepo: pnpm workspaces
- Сборка: Vite
- Десктоп: Tauri — WebView-based, `@owo-ui/css` работает без адаптаций

---

## Контент и домен

- [ ] перебери структуру своего сайта (сначала правки на ЖД, потом на хост)
  - не у всех директорий есть обложки
  - структура `scope_year_name`
- [ ] запаблить резюме и cv на owo
- [ ] перенести на iwwwanow.ru домен
