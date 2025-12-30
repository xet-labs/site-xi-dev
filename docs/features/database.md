# Database & Store

The application uses **GORM** for ORM capabilities, supporting distinct profiles for modular database management.

## Features
*   **Multi-Database Support:** Configure SQLite, MySQL, or PostgreSQL via config.
*   **Auto-Migration:** Automatically syncs Go structs in `internal/models/store` with the database schema on startup.
*   **Connection Pooling:** Configurable pool settings for high performance.

## Usage
Access the database instance via the global `Store` service:

```go
import "xi/internal/infra/store"

// Run a query
var users []User
store.Db.Cli().Find(&users)
```
