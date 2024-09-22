ARG GO_VERSION=1.23
ARG GO_IMAGE=mcr.microsoft.com/devcontainers/go
ARG ALPINE_VERSION=bookworm

FROM ${GO_IMAGE}:${GO_VERSION}-${ALPINE_VERSION}

# Update non-root user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Switch to non-root user
USER $USERNAME

# Install packages and Go language server
RUN go install golang.org/x/tools/gopls@latest 2>&1

# Set up workspace
WORKDIR /workspace

# Expose ports
EXPOSE 9100