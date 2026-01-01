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

## Development

### Pre-commit Hooks

This project uses [pre-commit](https://pre-commit.com/) to ensure code quality.

1. Install pre-commit: `pip install pre-commit` (or via brew/apt).
2. Install hooks: `pre-commit install`.
3. Run manually: `pre-commit run --all-files`.

