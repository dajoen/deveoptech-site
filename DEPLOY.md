# Deployment Guide

## Website (Cloudflare Pages)

1. Connect GitHub repo to Cloudflare Pages.
2. Project Settings:
   - **Framework Preset**: Astro
   - **Root Directory**: `site`
   - **Build Command**: `npm run build`
   - **Output Directory**: `dist`
3. Environment Variables:
   - None required for basic setup.

**Note**: The `package.json` has a `prebuild` script that symlinks `../../docs` to `src/content/docs`. This ensures the shared docs are included in the build.

## SSH TUI (VPS)

1. Provision a VPS (e.g., DigitalOcean, Hetzner).
2. DNS: Point `ssh.devopstech.site` (A record) to VPS IP.
3. Copy binary and service file:
   ```bash
   # Build locally
   make ssh-build
   
   # Copy to server
   scp ssh-tui/devdocs-ssh user@ssh.devopstech.site:/usr/local/bin/
   scp docs/ user@ssh.devopstech.site:/var/lib/devopstech/docs -r
   ```
4. Systemd Service (`/etc/systemd/system/devdocs-ssh.service`):
   ```ini
   [Unit]
   Description=DevOpsTech SSH Docs
   After=network.target

   [Service]
   Type=simple
   User=devdocs
   ExecStart=/usr/local/bin/devdocs-ssh
   Environment="DOCS_DIR=/var/lib/devopstech/docs"
   Restart=always
   
   # Hardening
   NoNewPrivileges=yes
   ProtectSystem=strict
   ProtectHome=yes
   PrivateTmp=yes

   [Install]
   WantedBy=multi-user.target
   ```
5. Enable and start:
   ```bash
   useradd -r -s /bin/false devdocs
   mkdir -p /var/lib/devopstech/docs
   chown -R devdocs:devdocs /var/lib/devopstech
   systemctl enable --now devdocs-ssh
   ```

## DNS Records

| Type | Name | Content | Proxy Status |
|------|------|---------|--------------|
| A    | @    | (Cloudflare Pages IP) | Proxied |
| CNAME| www  | devopstech.site | Proxied |
| A    | ssh  | (VPS IP) | DNS Only |
