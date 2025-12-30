# Database

ORM layer powered by **GORM**, supporting MySQL, SQLite, and PostgreSQL.

## Core Features
*   **Auto-Migration:** Syncs registered Go models (`internal/models/store`) with DB schema on startup.
*   **Connection Pooling:** Configurable idle/max connections.
*   **Profile Support:** Switch DB drivers via config.

## Usage
```go
import "xi/internal/infra/store"

store.Db.Cli().Model(&User{}).First(&user)
```
