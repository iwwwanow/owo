# owo — мастер-промпт для ИИ агентов

Этот файл — точка входа для работы с репозиторием. Скинь его агенту перед началом работы.

---

## Что такое owo

**owo** — self-hosted веб-сервер на Go. Читает директорию файлов и рендерит их как портфолио-сайт. Никакой БД, никакого CMS — файловая система и есть модель данных.

Продакшн: `iwwwanow.ru`. Автор: Кирилл Иванов (художник + Go-разработчик).

---

## Стек

- **Backend:** Go, `html/template`, `gomarkdown`
- **Frontend:** Go-шаблоны + чистый CSS + vanilla JS (без фреймворков)
- **Деплой:** Docker → Dokku → GHCR; push в `master` → GitHub Actions → semantic release → deploy
- **Dev:** `docker-compose up air` (hot reload через air)

---

## Архитектура (`internal/`)

```
HTTP → Controller → Handler → Repository → Renderer → HTTP response
```

| Файл | Роль |
|---|---|
| `controller.go` | HTTP роутер, различает статику / ресурсы |
| `handler.go` | Бизнес-логика, координирует repo + renderer |
| `repository.go` | Работа с FS: классификация файлов, мета-данные |
| `renderer.go` | Рендер Go html/template |

---

## Шаблоны (`web/templates/`)

- `pages/resource.page.html.tmpl` — главный шаблон страницы
- `fragments/content.fragment.html.tmpl` — грид карточек, секции, скрытые элементы
- `components/card.component.html.tmpl` — карточка ресурса (`<a class="card__wrapper">`)

**Карточка в DOM:**
```html
<a id="card-{path}" class="card__wrapper border_light" href="/{path}">...</a>
```

---

## URL и query-параметры

| Паттерн | Описание |
|---|---|
| `GET /{path}` | Рендер owo-страницы ресурса |
| `GET /{path}?static` | Отдать файл из uploads как есть (без рендера) |
| `GET /{path}?static&width=W` | Ресайз изображения по ширине |
| `GET /{path}?static&height=H` | Ресайз изображения по высоте |
| `GET /{path}?static&width=W&height=H` | Ресайз в точные размеры |
| `GET /backup` | ZIP-архив всего uploads |
| `POST /cache/clear` | Сбросить кэш ресайза |

Кэш ресайза: `/var/www/owo/cache/{W}x{H}/{path}`. Форматы: JPEG, PNG, GIF, BMP, WebP.

---

## Файловая система uploads

```
uploads/
└── my-project/
    ├── photo.jpg
    ├── video.mp4
    └── .meta/
        ├── index.md      # markdown, рендерится над гридом
        ├── index.css     # кастомные стили (только на этой странице)
        ├── index.js      # кастомный скрипт (defer, только на этой странице)
        └── cover.jpg     # обложка карточки в родительской директории
```

**Важно для JS в `.meta/index.js`:**
- Карточки рендерятся как `<a href="/{path}/filename.ext">` — это owo-роут, не файл
- Чтобы получить файл: добавь `?static` → `href + '?static'`
- Fallback через `window.location.pathname` + имя файла + `?static`

---

## Нейминг-конвенции директорий и файлов

| Префикс | Поведение |
|---|---|
| `.` | Скрытый ресурс (идёт в самый низ, отделён divider'ом "hidden") |
| `_directory-name` | Создаёт разделитель секций; дочерние блоки рендерятся на родительской странице |

---

## Переменные окружения

| Переменная | По умолчанию | Описание |
|---|---|---|
| `PORT` | `3000` | Порт HTTP-сервера |
| `PUBLIC_DIR` | — | Абсолютный путь к директории uploads |
| `TZ` | `Europe/Moscow` | Временная зона |
| `SSH_PUBLIC_KEY` | — | Публичный SSH-ключ для загрузки файлов |

---

## Команды

```bash
# Разработка
docker-compose up air

# Сборка
go build -o ./tmp/main ./cmd/main.go

# Форматирование шаблонов
pnpm prettier

# Продакшн-образ
docker build -t owo .
```

Тестов нет.

---

## Деплой

`master` → GitHub Actions → semantic release (conventional commits) → Docker multi-platform build → push GHCR → Dokku deploy.

Конвенция коммитов: `feat:`, `fix:`, `chore:` и т.д.

---

## Что сейчас в работе

- Roadmap: `docs/roadmap.md`
- Стилистика: `docs/styling-guide.md`
- Архитектура (dev): `docs/dev/`
