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
2. Install goimports: `go install golang.org/x/tools/cmd/goimports@latest` (ensure `$HOME/go/bin` is in your PATH).
3. Install hooks: `pre-commit install`.
4. Run manually: `pre-commit run --all-files`.

### Dev Container

This project includes a `.devcontainer` configuration for VS Code.
It provides an Arch Linux environment with all necessary tools (`go`, `node`, `pre-commit`, `rg`, `fzf`, `glow`).

1. Open the project in VS Code.
2. Click "Reopen in Container" when prompted.

