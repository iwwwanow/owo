# owo

file-based portfolio builder

<img src="assets/output/cover-cropped.png" width="480" alt="owo" />

_Kirill Ivanov 2026 Moscow_

---

**owo** is a self-hosted web server that turns a directory of files into a portfolio website. Drop in images, text, markdown, or any other files — owo renders them automatically with a clean UI, no configuration required.

## Features

- **File-first** — no database, no CMS. The filesystem is the content model
- **Markdown support** — `.md` files are rendered as HTML
- **Per-directory metadata** — add `index.md`, `index.css`, `index.js`, or a cover image via `.meta/` subfolder
- **Docker-native** — single container, minimal footprint
- **SSH key auth** — upload files via `scp` or `rsync` with your own SSH key

## Quick start

Initialize owo data directory on your host:

```sh
curl -sSL https://raw.githubusercontent.com/iwwwanow/owo/master/scripts/owo-init.sh | sh -s -- 'YOUR_SSH_PUBLIC_KEY'
```

Then run with Docker:

```sh
docker run -d \
  -p 3000:3000 \
  -e PUBLIC_DIR=/var/www/owo/uploads \
  -v /var/www/owo:/var/www/owo \
  ghcr.io/iwwwanow/owo:latest
```

## Environment variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3000` | HTTP server port |
| `PUBLIC_DIR` | — | Path to uploads directory |
| `TZ` | `Europe/Moscow` | Timezone |

## Directory structure

Each portfolio item is a directory under `PUBLIC_DIR`. Optionally add a `.meta/` subfolder:

```
uploads/
└── my-project/
    ├── photo.jpg
    ├── video.mp4
    └── .meta/
        ├── index.md      # description, rendered as HTML
        ├── index.css     # custom styles
        ├── index.js      # custom scripts
        └── cover.jpg     # directory cover image
```

## URL and query parameters

| URL pattern | Description |
|---|---|
| `GET /{path}` | Render owo page for resource at path |
| `GET /{path}?static` | Serve raw file from uploads (no page render) |
| `GET /{path}?static&width=W` | Serve resized image (preserves aspect ratio) |
| `GET /{path}?static&height=H` | Serve resized image (preserves aspect ratio) |
| `GET /{path}?static&width=W&height=H` | Serve resized image (exact dimensions) |
| `GET /backup` | Download zip archive of all uploads |
| `POST /cache/clear` | Clear image resize cache |

Resize cache is stored at `/var/www/owo/cache/{W}x{H}/{path}`. Supported formats: JPEG, PNG, GIF, BMP, WebP.

```sh
# examples
GET /projects/cover.jpg?static
GET /projects/cover.jpg?static&width=240&height=200
GET /avatar.png?static&width=96
```

## .meta/ customization

Each directory can override its page layout and behavior via a `.meta/` subfolder:

| File | Description |
|---|---|
| `index.md` | Markdown description injected above the file grid |
| `index.html` | Raw HTML injected into the page (iframe) |
| `index.css` | Custom CSS loaded for this directory's page |
| `index.js` | Custom JS loaded for this directory's page (deferred) |
| `cover.jpg` / `cover.png` / etc. | Cover image shown on the parent directory's card |

Scripts and styles are injected only on the directory's own page, not on parent pages.

Inside `index.js`, card links are available as `a[href$=".mp4"]`, `a[href$=".png"]` etc. — append `?static` to the href to get the raw file URL.

## Development

```sh
docker-compose up air
```

Build:

```sh
go build -o ./tmp/main ./cmd/main.go
```
