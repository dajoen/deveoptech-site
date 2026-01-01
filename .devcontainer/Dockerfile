FROM archlinux:base-devel

# Update system and install dependencies
RUN pacman -Syu --noconfirm && \
    pacman -S --noconfirm \
    git \
    go \
    nodejs \
    npm \
    pre-commit \
    ripgrep \
    fzf \
    glow \
    openssh \
    vim \
    neovim \
    zsh \
    starship \
    && pacman -Scc --noconfirm

# Create a non-root user
RUN useradd -m -G wheel -s /bin/zsh vscode && \
    echo '%wheel ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

# Set up Go environment for the user
USER vscode
ENV GOPATH=/home/vscode/go
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p $GOPATH/bin && \
    go install golang.org/x/tools/cmd/goimports@latest

# Set working directory
WORKDIR /workspace
