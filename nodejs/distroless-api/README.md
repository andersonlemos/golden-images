# Node.js Distroless API (Sample)

Sample **Node.js** API application consuming the [Node.js Golden Image](../README.md).

Demonstrates a native REST API with zero external runtime dependencies, packaged into a hardened Distroless container image.

---

## 🚀 API Endpoints

The application listens natively on port `3000`:

| Method | Path | Description | Example Response |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Application and runtime info | `{"message":"Hello from Node.js Golden Image","nodeVersion":"v24.x.x","environment":"production"}` |
| `GET` | `/health` | API health check | `{"status":"ok","environment":"production"}` |

---

## 🏷️ Image Naming Convention

Images follow the repository pattern:

$$\mathbf{alemos/distroless-api:node(version\_number|version-LTS)-(debug|regular)}$$

- **Debug Mode (Local default)**: `alemos/distroless-api:node24-debug`
- **Production Mode (Regular)**: `alemos/distroless-api:node24-regular`

---

## 🛠️ Makefile Commands

Execute commands from this directory (`nodejs/distroless-api/`):

### Application Lifecycle

```bash
# 1. Start container in background (automatically builds base and API if needed)
make start

# 2. Test HTTP endpoints (/ and /health)
make test

# 3. Open interactive debug shell (/busybox/sh)
make shell

# 4. View container logs
make logs

# 5. Stop and remove container
make stop
```

### Alternative Modes

```bash
# Run production version (no shell)
make start DEBUG=false

# Run with LTS version tag
make start LTS=true

# Run in foreground (attached)
make run
```

### Image Management & Compilation

```bash
# Build API image
make build

# Explicitly build base image in parent directory
make build-base-debug   # Builds node24-debug
make build-base-prod    # Builds node24-regular

# Clean containers and generated image
make clean
```

---

## ⚙️ Configuration Variables

| Variable | Default | Description |
| :--- | :--- | :--- |
| `DEBUG` | `true` | If `true`, consumes debug base and tags `-debug`. If `false`, uses `-regular`. |
| `LTS` | `false` | When `true`, appends `-lts` to version (e.g., `node24-lts-debug`). |
| `VERSION` | `24` | Target Node.js Golden Image version. |
| `PORT` | `3000` | Host port exposed by container. |
| `NODE_ENV` | `production` | Environment variable passed to container. |
| `CONTAINER_NAME` | `distroless-api-node` | Managed Docker container name. |
