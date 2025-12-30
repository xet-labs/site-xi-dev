# Hooks System

The Hooks system (`pkg/hook`) provides a flexible, event-driven architecture allowing disparate parts of the application to interact without tight coupling.

## Overview
Hooks allow you to register functions to be executed at specific lifecycles or events. Each hook can have **Pre**, **Core**, and **Post** execution phases.

## Structure
*   **Path:** `pkg/hook`
*   **Types:**
    *   `HookFn`: The function signature for a hook.
    *   `Hook`: The main struct managing lists of Pre/Core/Post functions.

## Usage

### defining a Hook
```go
// Create a hook
myHook := hook.NewHook(nil, nil, nil)
```

### Registering Listeners
```go
myHook.AddPre(func(args ...any) (any, error) {
    fmt.Println("Before event...")
    return nil, nil
})
```

### Executing
```go
// Run all 'Pre' hooks
myHook.RunPre(someArgs)
```

## Key Benefits
*   **Decoupling:** Modules can extend functionality (e.g., a Plugin registering a route) without modifying the core bootstrap logic.
*   **Safety:** Hooks run with panic recovery to prevent crashing the main application loop.
