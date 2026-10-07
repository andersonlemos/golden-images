# Node.js Golden Images

Hardened, minimal base container image (*Golden Image*) definitions for **Node.js 24**, built on top of **Google Distroless** (`gcr.io/distroless/nodejs24-debian13`).

---

## 🏷️ Tag Naming Convention

Images follow the repository specification:

$$\mathbf{node(version\_number|version-LTS)-(debug|regular)}$$

| Variant | Canonical Tag | Description |
| :--- | :--- | :--- |
| **Regular** | `alemos/golden_images:node24-regular` | Production. `nonroot` user (UID 65532), no shell, no package managers. |
| **Debug** | `alemos/golden_images:node24-debug` | Debugging. Includes `/busybox/sh` for diagnostics, maintaining `nonroot` user. |
| **LTS Regular** | `alemos/golden_images:node24-lts-regular` | Production with Long Term Support (LTS) release tag. |
| **LTS Debug** | `alemos/golden_images:node24-lts-debug` | Debugging with Long Term Support (LTS) release tag. |

---

## 🛠️ Makefile Commands

Run commands from this directory (`nodejs/`):

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
# Smoke test (checks Node version and NODE_ENV)
make test
make test-debug

# Open interactive shell in debug image
make run-debug
```

### Registry Push & Multi-Architecture (Buildx)

```bash
# Push images to the configured registry
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
| `LANGUAGE` | `node` | Language tag prefix |
| `VERSION` | `24` | Node.js major version |
| `LTS` | `false` | When `true`, appends `-lts` to version identifier |
| `NODE_ENV` | `production` | Default `NODE_ENV` passed to the image |
| `DISTROLESS_VARIANT` | `nonroot` | Distroless upstream variant for regular images |
| `DEBUG_VARIANT` | `debug-nonroot` | Distroless upstream variant for debug images |
| `PLATFORMS` | `linux/amd64,linux/arm64` | Target architectures for Buildx |

---

## 📂 Sample Project

To view a complete consumer application using this image, see the [`distroless-api/`](./distroless-api) directory.
