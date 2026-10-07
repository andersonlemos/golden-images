# Golden Images

Secure, hardened, and minimal container base images (*Golden Images*) based on **Google Distroless** for **Node.js** and **Go (Golang)**.

This repository includes both base image definitions and consumer sample applications (`distroless-api`) demonstrating how to use them in development and production environments.

---

## 🎯 Tag Naming Convention

All images in this repository strictly adhere to the following naming pattern:

$$\mathbf{(node|go)(version\_number|version-LTS)-(debug|regular)}$$

### Components:
- **Language**: `node` or `go`
- **Version**: version number (e.g., `24`, `1.24`) or with an LTS suffix (e.g., `24-lts`, `1.24-lts`)
- **Variant**:
  - `regular`: Production-ready minimal image (`nonroot` user, no shell, minimal attack surface).
  - `debug`: Debugging image (includes `/busybox/sh` and diagnostic tools, retaining the `nonroot` user).

### Generated Tags Reference:

| Language | Version | Variant | Golden Image Tag | Sample API Tag |
| :--- | :--- | :--- | :--- | :--- |
| **Node.js** | 24 | Regular | `alemos/golden_images:node24-regular` | `alemos/distroless-api:node24-regular` |
| **Node.js** | 24 | Debug | `alemos/golden_images:node24-debug` | `alemos/distroless-api:node24-debug` |
| **Node.js** | 24 (LTS) | Regular | `alemos/golden_images:node24-lts-regular` | `alemos/distroless-api:node24-lts-regular` |
| **Node.js** | 24 (LTS) | Debug | `alemos/golden_images:node24-lts-debug` | `alemos/distroless-api:node24-lts-debug` |
| **Go** | 1.24 | Regular | `alemos/golden_images:go1.24-regular` | `alemos/distroless-api:go1.24-regular` |
| **Go** | 1.24 | Debug | `alemos/golden_images:go1.24-debug` | `alemos/distroless-api:go1.24-debug` |
| **Go** | 1.24 (LTS) | Regular | `alemos/golden_images:go1.24-lts-regular` | `alemos/distroless-api:go1.24-lts-regular` |
| **Go** | 1.24 (LTS) | Debug | `alemos/golden_images:go1.24-lts-debug` | `alemos/distroless-api:go1.24-lts-debug` |

---

## 📁 Repository Structure

```text
golden-images/
├── nodejs/
│   ├── dockerfile.nodejs24       # Node.js 24 Golden Image base
│   ├── Makefile                   # Build and publish base Node.js image
│   └── distroless-api/            # Node.js sample API consuming the Golden Image
│       ├── dockerfile
│       ├── Makefile
│       ├── package.json
│       └── src/server.js
├── golang/
│   ├── dockerfile.golang124       # Go 1.24 Golden Image base
│   ├── Makefile                   # Build and publish base Go image
│   └── distroless-api/            # Go sample API consuming the Golden Image
│       ├── dockerfile
│       ├── Makefile
│       ├── go.mod
│       └── main.go
├── DOCKERHUB.md                   # Docker Hub repository description
└── README.md
```

---

## 🚀 Quick Start

### 1. Node.js

#### Building the Base Golden Image (`nodejs/`)

```bash
cd nodejs

# Build regular production image
make build

# Build debug image (with shell)
make build-debug

# Build with LTS tag
make build LTS=true
make build-debug LTS=true

# Build both simultaneously
make build-all

# Push images to the configured registry
make push-all REGISTRY=your-username
```

#### Running the Sample API (`nodejs/distroless-api/`)

The API Makefile automatically builds the local base image if not already cached:

```bash
cd nodejs/distroless-api

# 1. Start the API in debug mode (port 3000)
make start

# 2. Test HTTP endpoints (/ and /health)
make test

# 3. Open interactive busybox shell for inspection
make shell

# 4. View container logs
make logs

# 5. Stop and remove container
make stop

# Run in production mode (regular, no shell):
make start DEBUG=false
```

---

### 2. Golang

#### Building the Base Golden Image (`golang/`)

```bash
cd golang

# Build regular production image
make build

# Build debug image (with shell)
make build-debug

# Build with LTS tag
make build LTS=true
make build-debug LTS=true

# Build both simultaneously
make build-all

# Push images to the configured registry
make push-all REGISTRY=your-username
```

#### Running the Sample API (`golang/distroless-api/`)

```bash
cd golang/distroless-api

# 1. Start the API in debug mode (port 8080)
make start

# 2. Test HTTP endpoints (/ and /health)
make test

# 3. Open interactive busybox shell for inspection
make shell

# 4. View container logs
make logs

# 5. Stop and remove container
make stop

# Run in production mode (regular, no shell):
make start DEBUG=false
```

---

## ⚙️ Makefile Configuration Variables

All `Makefile`s accept environment variables or CLI overrides:

| Variable | Default | Description |
| :--- | :--- | :--- |
| `REGISTRY` | `alemos` | Docker registry or organization prefix |
| `DEBUG` | `true` | Toggle between debug (`true`) and production (`false`) variant |
| `LTS` | `false` | When `true`, appends `-lts` to the version tag |
| `NODE_ENV` | `production` | Sets `process.env.NODE_ENV` in the Node.js image |
| `APP_ENV` | `production` | Sets `APP_ENV` environment variable in the Go image |
| `PORT` | `3000` (Node) / `8080` (Go) | Host port exposed by the container |
| `PLATFORMS`| `linux/amd64,linux/arm64` | Target platforms for multi-arch Docker Buildx builds |

Usage example:

```bash
make build REGISTRY=myregistry NODE_ENV=development LTS=true
```

---

## 🌐 Multi-Architecture Builds (Docker Buildx)

To build images supporting multiple architectures (e.g., Intel/AMD `x86_64` and Apple Silicon / ARM64):

```bash
# In nodejs/ or golang/:
make buildx-all-push REGISTRY=your-username PLATFORMS=linux/amd64,linux/arm64
```

---

## 🔒 Security Principles & Best Practices

1. **Nonroot by Default**: Images natively run with the unprivileged user `nonroot` (UID/GID `65532`), preventing privilege escalation vulnerabilities.
2. **Minimal Attack Surface**: `regular` images contain no shell, package managers, or extra system utilities. For local debugging or staging diagnostics, use the `debug` variant.
3. **Static Go Binaries**: The Go sample application is compiled with `CGO_ENABLED=0` and runs on top of `distroless/static-debian12`, eliminating libc or dynamic library dependencies.
4. **Automatic Local Resolution**: The application Makefile builds the required local base image before running the API container, ensuring Docker resolves local images without unexpected registry pulls.

---

## 📄 License

See [LICENSE](LICENSE) for terms and details.
