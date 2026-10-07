# Hardened Distroless Golden Images: Node.js & Go

> **Short Description (Docker Hub)**:
> `Hardened Distroless Golden Images for Node.js & Go. Minimal, nonroot, with regular & debug tags.`

---

Hardened, minimal, and optimized base container images (*Golden Images*) based on **Google Distroless** for production and debugging environments in **Node.js** and **Go (Golang)**.

---

## 🏷️ Available Tags

Tags adhere to the specification: `(node|go)(version_number|version-LTS)-(debug|regular)`

| Tag | Language | Distroless Base | User | Shell | Purpose |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `node24-regular` | Node.js 24 | `nodejs24-debian13:nonroot` | `nonroot (65532)` | ❌ None | **Production** (maximum security) |
| `node24-debug` | Node.js 24 | `nodejs24-debian13:debug-nonroot` | `nonroot (65532)` |  `/busybox/sh` | **Debugging / Local testing** |
| `node24-lts-regular` | Node.js 24 LTS | `nodejs24-debian13:nonroot` | `nonroot (65532)` | ❌ None | **Production (LTS)** |
| `node24-lts-debug` | Node.js 24 LTS | `nodejs24-debian13:debug-nonroot` | `nonroot (65532)` |  `/busybox/sh` | **Debugging (LTS)** |
| `go1.24-regular` | Go 1.24 | `static-debian12:nonroot` | `nonroot (65532)` | ❌ None | **Production** (static, no cgo) |
| `go1.24-debug` | Go 1.24 | `static-debian12:debug-nonroot` | `nonroot (65532)` |  `/busybox/sh` | **Debugging / Local testing** |
| `go1.24-lts-regular`| Go 1.24 LTS | `static-debian12:nonroot` | `nonroot (65532)` | ❌ None | **Production (LTS)** |
| `go1.24-lts-debug` | Go 1.24 LTS | `static-debian12:debug-nonroot` | `nonroot (65532)` |  `/busybox/sh` | **Debugging (LTS)** |

---

## 🚀 How to Use

### 🟢 1. In Node.js Applications

To build your application on top of the Node.js Golden Image:

```dockerfile
# Use "-regular" tag for final production images:
FROM alemos/golden_images:node24-regular

WORKDIR /app

# Copy application files with the unprivileged nonroot user
COPY --chown=nonroot:nonroot package*.json ./
COPY --chown=nonroot:nonroot src ./src

ENV NODE_ENV=production
ENV PORT=3000

EXPOSE 3000

CMD ["src/server.js"]
```

> **Development Tip**: Switch the base image to `alemos/golden_images:node24-debug` during development to allow interactive troubleshooting via `docker exec -it <container> /busybox/sh`.

---

### 🔵 2. In Go Applications (Multi-Stage Build)

To compile and run statically linked Go binaries with minimal size and zero runtime dependencies:

```dockerfile
# Stage 1: Build the static executable
FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-w -s" -o /app/server .

# Stage 2: Minimal Distroless runtime
FROM alemos/golden_images:go1.24-regular

WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /app/server /app/server

ENV APP_ENV=production
ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/app/server"]
```

---

## 🔒 Why Choose These Golden Images?

1. **Reduced Attack Surface**: Images contain only the language runtime and critical dependencies (CA certificates, tzdata). No package managers (`apt`, `apk`), compilers, or shell interpreters in `regular` builds.
2. **Safe Non-Root Execution**: Runs out-of-the-box as UID/GID `65532` (`nonroot`), satisfying security policies for Kubernetes and containerized enterprise workloads.
3. **Multi-Architecture**: Built for both `linux/amd64` (Intel/AMD) and `linux/arm64` (Apple Silicon and ARM cloud servers like AWS Graviton).
4. **Hassle-Free Troubleshooting**: The `-debug` variant provides `/busybox/sh` without altering permissions or compromising the nonroot security model.

---

## 💻 Source Code & Sample Apps

Complete source code, Makefiles, and sample applications are available on GitHub:  
👉 **[andersonlemos/golden-images](https://github.com/andersonlemos/golden-images)**
