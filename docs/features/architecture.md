# Architecture

The project follows a modular, scalable architecture designed to bootstrap standard web applications quickly.

## Core Components

### 1. Controllers (`internal/handlers`)
Handle HTTP requests. We use a **Base Controller** pattern (`pkg/app/ctrl`) to standardize responses (JSON success, Error handling, View rendering).

### 2. Services (`internal/services`)
Contains the business logic (Auth, User management, etc.). Services are singletons initialized at startup (`services.init.go`).

### 3. Models (`internal/models`)
*   **Store:** Database entities (GORM models).
*   **Config:** Configuration structs.

### 4. Infrastructure (`internal/infra`)
Low-level components like Database connections (`store`) and Logging (`logger`).

## Key Patterns
*   **Plugin Router:** Routes are registered via plugins, keeping `main.go` clean.
*   **Hooks:** Event-driven architecture for decoupling components.
