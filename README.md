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

## Development

```sh
docker-compose up air
```

Build:

```sh
go build -o ./tmp/main ./cmd/main.go
```
