# Golang Distroless API (Sample)

Sample **Go (Golang)** API application consuming the [Go Golden Image](../README.md).

Demonstrates a native REST API using Go's standard library (`net/http`), packaged via **Multi-Stage Build**:
1. **Stage 1 (Builder)**: Compiles static binary using `golang:1.24-alpine` with `CGO_ENABLED=0`.
2. **Stage 2 (Runtime)**: Runs the static binary on top of the Google Distroless Static Golden Image.

---

## 🚀 API Endpoints

The application listens natively on port `8080`:

| Method | Path | Description | Example Response |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Application and runtime info | `{"message":"Hello from Go Golden Image","goVersion":"go1.24.x","environment":"production"}` |
| `GET` | `/health` | API health check | `{"status":"ok","environment":"production"}` |

---

## 🏷️ Image Naming Convention

Images follow the repository pattern:

$$\mathbf{alemos/distroless-api:go(version\_number|version-LTS)-(debug|regular)}$$

- **Debug Mode (Local default)**: `alemos/distroless-api:go1.24-debug`
- **Production Mode (Regular)**: `alemos/distroless-api:go1.24-regular`

---

## 🛠️ Makefile Commands

Execute commands from this directory (`golang/distroless-api/`):

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
make build-base-debug   # Builds go1.24-debug
make build-base-prod    # Builds go1.24-regular

# Clean containers and generated image
make clean
```

---

## ⚙️ Configuration Variables

| Variable | Default | Description |
| :--- | :--- | :--- |
| `DEBUG` | `true` | If `true`, consumes debug base and tags `-debug`. If `false`, uses `-regular`. |
| `LTS` | `false` | When `true`, appends `-lts` to version (e.g., `go1.24-lts-debug`). |
| `VERSION` | `1.24` | Target Go Golden Image version. |
| `PORT` | `8080` | Host port exposed by container. |
| `APP_ENV` | `production` | Environment variable passed to container. |
| `CONTAINER_NAME` | `distroless-api-go` | Managed Docker container name. |
