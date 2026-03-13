# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

**Build:**
```bash
go build -o ./tmp/main ./cmd/main.go
```

**Run (development):**
```bash
docker-compose up air
# or directly with air:
air -c .air.toml
```

**Format templates:**
```bash
pnpm prettier
```

**Production build:**
```bash
docker build -t owo .
docker-compose -f docker-compose.image.yml up
```

There are no Go tests in this project.

## Architecture

**owo** is a file-based portfolio web server written in Go. It reads a directory of uploaded files and renders them as a dynamic website.

### Layered structure (`internal/`)

- **Controller** (`controller.go`) — HTTP router; distinguishes static/resource routes and delegates to handler
- **Handler** (`handler.go`) — business logic; coordinates repository and renderer, converts markdown to HTML
- **Repository** (`repository.go`) — filesystem access; classifies files (image/text/directory/other), discovers metadata and static assets
- **Renderer** (`renderer.go`) — renders Go `html/template` pages from props

### Data flow

HTTP request → Controller → Handler → Repository (reads files) → Renderer (renders template) → HTTP response

### File system conventions

Portfolio items live under `PUBLIC_DIR` (production: `/var/www/owo/uploads`, dev: `./test/uploads`). Each directory can have a `.meta/` subfolder with:
- `index.html`, `index.css`, `index.js`, `index.md` — static assets injected into the page
- `cover.png` / `cover.jpg` / etc. — directory cover image

### Templates (`web/templates/`)

- `pages/resource` — main page template
- `fragments/` — head, header, content, footer partials
- `components/` — card, iframe, html, image, code, hr

In production, templates are loaded from `/var/www/owo/templates`. In dev, from `./web/templates` (symlink or direct path).

### Key environment variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3000` | HTTP server port |
| `PUBLIC_DIR` | — | Path to uploads directory |
| `TZ` | `Europe/Moscow` | Timezone |

### Deployment

Push to `master` triggers GitHub Actions: Semantic Release → Docker multi-platform build → push to GHCR → Dokku deploy to production.
