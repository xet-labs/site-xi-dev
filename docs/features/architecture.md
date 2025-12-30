# Architecture & Routing

Modular, logic-first structure powered by a **Plugin-Based Router**.

## Components
*   **`cmd/`**: Entry points.
*   **`internal/services/`**: Singleton business logic (Auth, User).
*   **`internal/handlers/`**: Http Controllers (`BaseController`).
*   **`pkg/`**: Public libraries (`hook`, `lib/router`).

## Plugin-Based Routing
Routes are not defined in a central file. Instead, each Controller implements the `CoreRouter` interface to register its own routes. This keeps the application modular.

### Usage
Implement the `RouterCore` method in your controller:

```go
type MyCtrl struct{}

// Register routes
func (c *MyCtrl) RouterCore(r *gin.Engine) {
    r.GET("/api/my-endpoint", c.MyHandler)
}

func (c *MyCtrl) MyHandler(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"status": "ok"})
}
```

The system automatically detects this method and registers the routes at startup.
