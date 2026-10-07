# Golang Golden Images

Hardened, minimal base container image (*Golden Image*) definitions for **Go (Golang) 1.24**, built on top of **Google Distroless** (`gcr.io/distroless/static-debian12`).

Designed specifically for statically compiled Go binaries (`CGO_ENABLED=0`), packing only what is strictly necessary: CA certificates (`ca-certificates`), timezone data (`tzdata`), and the unprivileged `nonroot` user (UID `65532`).

---

## 🏷️ Tag Naming Convention

Images follow the repository specification:

$$\mathbf{go(version\_number|version-LTS)-(debug|regular)}$$

| Variant | Canonical Tag | Description |
| :--- | :--- | :--- |
| **Regular** | `alemos/golden_images:go1.24-regular` | Production. `nonroot` user (UID 65532), no shell, no extra binaries. |
| **Debug** | `alemos/golden_images:go1.24-debug` | Debugging. Includes `/busybox/sh` for diagnostics, maintaining `nonroot` user. |
| **LTS Regular** | `alemos/golden_images:go1.24-lts-regular` | Production with Long Term Support (LTS) release tag. |
| **LTS Debug** | `alemos/golden_images:go1.24-lts-debug` | Debugging with Long Term Support (LTS) release tag. |

---

## 🛠️ Makefile Commands

Run commands from this directory (`golang/`):

### Local Compilation

```bash
# Build regular production image
make build

# Build debug image (with shell)
make build-debug

# Build with LTS flag
make build LTS=true
make build-debug LTS=true

# Build without cache
make build-nc
make build-debug-nc

# Build both simultaneously
make build-all
```

### Testing and Execution

```bash
# Inspect regular image
make test

# Smoke test debug shell
make test-debug

# Open interactive shell in debug image
make run-debug
```

### Registry Push & Multi-Architecture (Buildx)

```bash
# Push images to configured registry
make push
make push-debug
make push-all

# Build and push multi-arch (amd64 and arm64)
make buildx-all-push PLATFORMS=linux/amd64,linux/arm64
```

### Cleanup

```bash
make clean
```

---

## ⚙️ Configuration Variables

| Variable | Default | Description |
| :--- | :--- | :--- |
| `REGISTRY` | `alemos` | Docker registry or organization prefix |
| `IMAGE_NAME` | `golden_images` | Base repository name |
| `LANGUAGE` | `go` | Language tag prefix |
| `VERSION` | `1.24` | Go version |
| `LTS` | `false` | When `true`, appends `-lts` to version identifier |
| `APP_ENV` | `production` | Default application environment (`APP_ENV`) |
| `DISTROLESS_IMAGE` | `static-debian12` | Upstream static Distroless base image |
| `DISTROLESS_VARIANT` | `nonroot` | Distroless upstream variant for regular images |
| `DEBUG_VARIANT` | `debug-nonroot` | Distroless upstream variant for debug images |
| `PLATFORMS` | `linux/amd64,linux/arm64` | Target architectures for Buildx |

---

## 📂 Sample Project

To view a complete consumer Go application using this image via multi-stage build, see the [`distroless-api/`](./distroless-api) directory.
