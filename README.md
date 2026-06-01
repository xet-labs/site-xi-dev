<p align="center" style="margin: 0;">
  <a href="#" style="margin: 0;">
    <picture>
      <!-- Dark mode -->
      <source
        srcset="public/static/org/logo-wordmark-white.svg"
        media="(prefers-color-scheme: dark)"
      >
      <!-- Light mode -->
      <source
        srcset="public/static/org/logo-wordmark-black.svg"
        media="(prefers-color-scheme: light)"
      >
      <!-- Fallback -->
      <img
        src="public/static/org/logo-wordmark-black.svg"
        width="360"
        height="100"
      >
    </picture>
  </a>
</p>

<p align="center" style="margin-top: 0;"> A high-performance, modular web application to Bootstrap new web projects.</p>

---

## Getting Started

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

## Features

*   [**Hooks System:**](docs/features/hooks.md) **(Core)** Event-driven architecture. Decouples modules by allowing them to register listeners (`pkg/hook`).
*   [**Plugin-based Routes:**](docs/features/architecture.md) Controllers register their own routes via interfaces, keeping the main entry point clean.
*   [**SEO & Sitemaps:**](docs/features/seo.md) Dynamic, hook-based sitemap generation. Modules register their URLs.
*   [**Authentication:**](docs/features/auth.md) Stateless JWT auth with persistent sessions and auto-login.
*   [**Configuration:**](docs/features/config.md) Hot-reloadable; supports JSON configuration with environment variable injection, fallbacks, and `.env` files.
*   [**Database:**](docs/features/database.md) GORM-backed (MySQL/SQLite) with auto-migrations.
*   **Asset Management:** CSS, HTML minification/serving to save on bandwidth.

See [docs/](docs/) for details.

---

## Directory Structure

```graphql
.
├── cmd                 # entry points
├── config              # config files (JSON)
├── data                # app Data
│   └── default         # default app data
│       ├── config      # app config
│       ├── infra       # infra related files (docker configs etc)
│       └── seed        # seed data (databases etc)
├── docs                # documentation
├── internal            # private app code
│   ├── app             # core components
│   ├── config          # config logic
│   ├── constraints     # validation rules and constraints
│   ├── handlers        # HTTP Controllers, error handlers, etc.
│   ├── infra           # low-level infra (Logger, Store)
│   ├── models          # data/Config structs
│   └── services        # business logic
├── pkg                 # shared libraries
├── public              # publicly accessible static assets
└── ui                  # frontend templates
```

---

## Key Technologies

*   **Framework:** [Gin](https://github.com/gin-gonic/gin)
*   **Database:** [GORM](https://gorm.io/)
*   **Logging:** [Zerolog](https://github.com/rs/zerolog)
*   **Config:** [Koanf](https://github.com/knadh/koanf)
*   **Auth:** [Golang-JWT](https://github.com/golang-jwt/jwt)

---

<p align="center"> <a href="https://buymeacoffee.com/rishikeshprasad" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a> </p>

<p align="center"> If this project saved you time or helped you learn, consider supporting the development! ☕ </p>

<p align="center"> Licensed under the <a href="?tab=License-1-ov-file"> Apache License </a>. </p>
