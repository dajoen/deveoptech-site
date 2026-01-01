---
id: task-5
title: Implement Arch-based Dev Container
status: Done
assignee:
  - '@copilot'
created_date: '2026-01-01 15:33'
updated_date: '2026-01-01 15:36'
labels: []
dependencies: []
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Set up a Development Container configuration to standardize the development environment using Arch Linux. This ensures all developers have the same tools and versions.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Create .devcontainer/devcontainer.json
- [x] #2 Create Dockerfile based on archlinux
- [x] #3 Install dependencies: go, nodejs, npm, git, pre-commit, ripgrep, fzf, glow
- [x] #4 Configure VS Code extensions (Go, Astro, Prettier)
- [x] #5 Verify 'make dev' and 'make ssh-build' work inside container
<!-- AC:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Created .devcontainer/Dockerfile (Arch Linux) and devcontainer.json. Configured dependencies and VS Code extensions. Updated README.md.
<!-- SECTION:NOTES:END -->
