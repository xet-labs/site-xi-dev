# Configuration System

The configuration system (`xi/internal/config`) is designed for flexibility, allowing hierarchal loading, environment variable injection, and hot-reloading.

## Features

### 1. Environment Variable Injection
You can inject environment variables directly into your JSON configuration files.
*   **Syntax:** `${VAR_NAME}` or `${VAR_NAME:-default_value}`
*   **Fallback:** Uses the default value if the variable is unset.

**Example `config.json`:**
```json
{
  "database": {
    "host": "${DB_HOST:-localhost}",
    "password": "${DB_PASS}" 
  }
}
```

### 2. .env File Support
The system automatically loads variables from a `.env` file in the project root if present, making it easy to manage local secrets.

### 3. Multi-File Loading & Precedence
Load configuration from multiple files (e.g., `default.json`, `production.json`).
*   **Rule:** Files are loaded sequentially; later files override earlier ones.

### 4. Hot Reloading
The system watches loaded config files and automatically reloads changes without restarting the application, enabling zero-downtime updates.

### 5. Recursive Variable Resolution
Config values can reference other defined values using `${key.path}` syntax.
*   **Example:** `"url": "https://${app.domain}/api"`

## Usage
Access configuration via the global `cfg` package:

```go
import "xi/internal/config/cfg"

func main() {
    port := cfg.Server.Port
    // ...
}
```
