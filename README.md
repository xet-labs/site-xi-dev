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

2.  **Run:**
    Ensure `config/config.json` exists.
    ```bash
    go run cmd/web/main.go
    ```
    Starts on port `5000` by default.

---

## 🛠 Features

*   [**Hooks System:**](docs/features/hooks.md) **(Core)** Event-driven architecture. Decouples modules by allowing them to register listeners (`pkg/hook`).
*   [**Plugin-based Routes:**](docs/features/architecture.md) Controllers register their own routes via interfaces, keeping the main entry point clean.
*   [**SEO & Sitemaps:**](docs/features/seo.md) Dynamic, hook-based sitemap generation. Modules auto-register their URLs.
*   [**Authentication:**](docs/features/auth.md) Stateless JWT auth with persistent sessions and auto-login.
*   [**Configuration:**](docs/features/config.md) Hot-reloadable; supports JSON configuration with environment variable injection, fallbacks, and `.env` files.
*   [**Database:**](docs/features/database.md) GORM-backed (MySQL/SQLite) with auto-migrations.
*   **Asset Management:** Automated minification/serving.

See [docs/](docs/) for details.

---

## 📂 Directory Structure

```graphql
.
├── cmd                 # Entry points
├── config              # Config files (JSON)
├── docs                # Documentation
├── internal            # Private app code
│   ├── app             # Core components
│   ├── config          # Config logic
│   ├── handlers        # HTTP Controllers
│   ├── infra           # Low-level infra (Logger, Store)
│   ├── models          # Data/Config structs
│   └── services        # Business logic
├── pkg                 # Shared libraries
│   ├── hook            # Hook system
│   ├── lib             # Core libs
│   └── util            # Utilities
├── public              # Static assets
└── ui                  # Frontend templates
```

---

## ⚙️ Key Technologies

*   **Framework:** [Gin](https://github.com/gin-gonic/gin)
*   **Database:** [GORM](https://gorm.io/)
*   **Logging:** [Zerolog](https://github.com/rs/zerolog)
*   **Config:** [Koanf](https://github.com/knadh/koanf)
*   **Auth:** [Golang-JWT](https://github.com/golang-jwt/jwt)

---

<a href="https://buymeacoffee.com/rishikeshprasad" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

If this project saved you time or helped you learn, consider supporting the development! ☕

Licensed under the [Apache License](LICENSE).
