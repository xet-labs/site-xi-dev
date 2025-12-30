# Configuration System

The configuration system (`xi/internal/config`) is designed for flexibility, allowing hierarchal loading, environment variable injection, and hot-reloading.

## Features

### 1. Environment Variable Injection
You can inject environment variables directly into your JSON configuration files using Bash-like syntax.
*   **Syntax:** `${VAR_NAME}` or `${VAR_NAME:-default_value}`
*   **Fallback:** If the environment variable is not set, the default value is used.

**Example `config.json`:**
```json
{
  "database": {
    "host": "${DB_HOST:-localhost}",
    "port": "${DB_PORT:-3306}",
    "password": "${DB_PASS}" 
  }
}
```

### 2. Multi-File Loading & Precedence
The system allows loading configuration from multiple files to separate concerns (e.g., `default.json`, `production.json`).
*   **Order Matters:** Files are loaded sequentially.
*   **Override Rule:** Values in later files override values from earlier files. This allows you to have a base configuration and simpler environment-specific overrides.

**Example:**
`go run cmd/web/main.go --config=base.json --config=prod.json`

In this case, `prod.json` settings will take precedence over `base.json`.

### 3. Hot Reloading
The system uses `fsnotify` to watch loaded configuration files for changes.
*   **Behavior:** When a config file is modified (e.g., via `vim` or deployment update), the system detects the change.
*   **Action:** It automatically re-reads and re-applies the configuration **without restarting the application**.
*   **Zero Downtime:** Perfect for rotating credentials or tweaking feature flags in production.

### 4. Recursive Variable Resolution
Config values can reference other defined values within the configuration using `${key.path}` syntax.
*   **Example:**
    ```json
    {
      "app": {
        "domain": "example.com",
        "url": "https://${app.domain}/api"
      }
    }
    ```

## Usage in Code
Access configuration via the global `cfg` package:

```go
import "xi/internal/config/cfg"

func main() {
    // Access strongly-typed config
    port := cfg.Server.Port
    dbHost := cfg.Store.Db.Conn.Host
    
    // ...
}
```
