# Hooks System

An event-driven notification system (`pkg/hook`) to decouple modules.

## Architecture
Hooks manage lists of functions (Listeners) to run at specific points (Pre, Core, Post).

*   **Type:** `HookFn func(args ...any) (any, error)`
*   **Safety:** Executed with panic recovery.

## Usage

**Define:**
```go
MyHook := hook.NewHook(nil, nil, nil)
```

**Register:**
```go
MyHook.AddPre(func(args ...any) (any, error) {
    // Logic here
    return nil, nil
})
```

**Trigger:**
```go
MyHook.RunPre(context)
```
