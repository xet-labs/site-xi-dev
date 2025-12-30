# XetIndustries
A high-performance, modular web application to **Bootstrap new web projects**.

---

## 🚀 Getting Started

### Prerequisites
*   [Go](https://go.dev/dl/) 1.20+
*   [Git](https://git-scm.com/)

### Quick Setup

1.  **Clone & Configure:**
    ```bash
    git clone https://github.com/xet-labs/site-xi-dev new-proj
    cd new-proj
    git remote rename origin base
    git checkout main
    ```

2.  **Run the Application:**
    Copy the example config (if available) or ensure `config/config.json` exists.
    ```bash
    go run cmd/web/main.go
    ```
    The server will start on port `5000` (by default).

---

## 🛠 Features

This project is packed with features designed for scalability, performance, and developer experience.

### Core Systems
*   [**Modular Architecture:**](docs/features/architecture.md) Clean separation of concerns with a custom plugin-based router.
*   [**Authentication:**](docs/features/auth.md) Secure, stateless JWT-based authentication with auto-login on signup.
*   [**Configuration:**](docs/features/config.md) Hot-reloadable configuration via JSON and environment variable injection.
*   [**Hooks System:**](docs/features/hooks.md) Event-driven hooks (`pkg/hook`) for decoupled feature integration.
*   [**Database:**](docs/features/database.md) GORM-backed storage with support for MySQL/SQLite and connection pooling.

### Additional Features
*   **Asset Management:** Automatic minification and serving of static assets.
*   **Blog System:** Full-featured blog module (inprogress).

See [docs/](docs/) for detailed documentation on all features.

---

## 📂 Directory Structure

```graphql
.
├── cmd                 # Application entry points
├── config              # Configuration files (config.json)
├── docs                # Project documentation
│   └── features        # Detailed feature guides
├── internal            # Private application code
│   ├── app             # Core app components and wiring logic
│   ├── config          # Configuration logic
│   ├── handlers        # HTTP handlers (Controllers), error handlers etc
│   ├── infra           # Infrastructure (Logger, Store, etc.)
│   ├── models          # Data models and structs
│   └── services        # Business logic services (Auth, Mail, etc.)
├── pkg                 # Public shared libraries
│   ├── hook            # Event hook system
│   ├── lib             # Core libraries *will be removed in future*
│   └── util            # Utility functions
├── public              # Static assets (CSS, JS, Images)
├── ui                  # Frontend resources
└── Makefile            # Build and run commands
```

---

## ⚙️ Key Technologies

*   **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
*   **ORM:** [GORM](https://gorm.io/)
*   **Logging:** [Zerolog](https://github.com/rs/zerolog)
*   **Config:** [Koanf](https://github.com/knadh/koanf) & [fsnotify](https://github.com/fsnotify/fsnotify)
*   **Auth:** [Golang-JWT](https://github.com/golang-jwt/jwt)
*   **Minification:** [Minify](https://github.com/tdewolff/minify)

---

If this project saved you time or helped you learn, consider supporting the development! ☕
<br/>
<a href="https://buymeacoffee.com/rishikeshprasad" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

**Licensed under the [Apache License](LICENSE).**
