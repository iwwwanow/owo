# API / Routes

## Страницы ресурсов

```
GET /{path}
```

Рендерит HTML-страницу для ресурса по пути `{path}` относительно `uploads/`.

- Пустой путь `/` — корень uploads
- Путь к директории — карточки дочерних ресурсов
- Путь к файлу — отрисовка контента (изображение, markdown, код и пр.)

---

## Статика приложения

```
GET /static/{path}
```

Отдаёт файлы из директории статики приложения (`/var/www/owo/static`): CSS, JS, шрифты.

---

## Uploads: статические файлы

```
GET /{path}?static
```

Отдаёт файл из `uploads/` как есть (без рендера страницы).

Используется для изображений-обложек и вложений внутри шаблонов.

```
# пример
GET /projects/cover.jpg?static
```

---

## Uploads: ресайз изображений с кэшом

```
GET /{path}?static&width={W}&height={H}
```

Возвращает изображение в указанных размерах.

- Параметры `width` и `height` — в пикселях. Можно передать только один — второй будет подобран с сохранением пропорций.
- Результат кэшируется в `/var/www/owo/cache/{W}x{H}/{path}`. Повторный запрос отдаётся из кэша без пересборки.
- Поддерживаемые форматы: JPEG, PNG, GIF, BMP, WebP.

```
# примеры
GET /projects/cover.jpg?static&width=240&height=200
GET /avatar.png?static&width=96
GET /photo.jpg?static&height=480
```

---

## Бэкап uploads

```
GET /backup
```

Возвращает zip-архив всей директории `uploads/` для скачивания.

Имя файла в заголовке: `owo-backup-YYYYMMDD-HHMMSS.zip`.

```bash
# скачать архив через curl
curl https://your-host/backup -o backup.zip
```

> Роут не защищён аутентификацией — если сервер публичный, закройте `/backup` на уровне nginx/reverse proxy.

---

## Очистка кэша изображений

```
POST /cache/clear
```

Удаляет всю директорию кэша (`/var/www/owo/cache`). Ресайз будет пересчитан при следующем запросе.

- Возвращает `204 No Content` при успехе.

```bash
curl -X POST https://your-host/cache/clear
```

> Роут не защищён аутентификацией — закройте на уровне nginx/reverse proxy в продакшене.

**Очистить вручную (без API):**

```bash
# через Dokku (на сервере)
dokku run owo rm -rf /var/www/owo/cache

# или напрямую на хосте
rm -rf /home/dokku/owo/cache/*
```
