# devopstech.site

Monorepo for devopstech.site.

## Components

- `site/`: Astro + Starlight website.
- `docs/`: Markdown content source.
- `bin/`: CLI tools (`devdocs`).
- `ssh-tui/`: SSH TUI server.

## Usage

See `Makefile` for commands.

### Quickstart

1. `make setup` - Links docs to site.
2. `make dev` - Runs website locally.
3. `make ssh-run` - Runs SSH TUI locally.
