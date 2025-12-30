# Configuration

The configuration system (`internal/config`) supports hierarchical loading, environment variables, and hot-reloading.

## Features & Examples

### 1. Environment Injection
Inject environment variables using `${VAR}` or `${VAR:-default}`.
*   **Automatic Loading:** Reads `.env` from the project root.

```json
{
  "database": {
    "host": "${DB_HOST:-localhost}",
    "password": "${DB_PASS}"
  }
}
```

### 2. Recursive References
Values can reference other config keys via `$(key.path)`.

```json
{
  "app": {
    "domain": "example.com",
    "url": "https://$(app.domain)"
  }
}
```

### 3. File Precedence
Load multiple files (e.g., `default.json`, `prod.json`). Later files override earlier ones.

### 4. Hot Reloading
Modifying a config file triggers an immediate reload without restarting.

## Usage
```go
import "xi/internal/config/cfg"

func main() {
    port := cfg.Server.Port
}
```
