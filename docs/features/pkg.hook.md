# pkg.hook — Minimal Generic Hook System (Go)

This package provides a **type-safe hook system** with:

* priorities
* named hooks
* panic isolation
* support for normal, no-arg, and side-effect hooks

Hooks are executed in **priority order (low → high)**.

---

## Core Types

### 1. Regular Hook

```go
Hook[T, R]
```

* Input: `T`
* Output: `R`
* Signature: `func(T) (R, error)`

### 2. No-Arg Hook

```go
NoArgHook[R]
```

* No input
* Output: `R`
* Signature: `func() (R, error)`

### 3. Effect Hook

```go
EffectHook[T]
```

* Input: `T`
* No return value
* Signature: `func(T) error`

---

## Creating Hooks

```go
h := hook.New[int, string]()
nh := hook.NewNoArg[string]()
eh := hook.NewEffectHook[int]()
```

---

## Registering Hooks

### Default (auto name, priority = 0)

```go
h.Add(func(v int) (string, error) {
	return strconv.Itoa(v), nil
})
```

### Named

```go
h.AddNamed("to-string", fn)
```

### Named + Priority

```go
h.AddWithPriority("validate", -10, fn) // runs earlier
```

Lower priority runs first.

---

## No-Arg Hooks

```go
nh.Add(func() (string, error) {
	return "hello", nil
})
```

---

## Effect Hooks (side-effects only)

```go
eh.Add(func(v int) error {
	log.Println(v)
	return nil
})
```

---

## Running Hooks

```go
results, errs := h.Run(42)
```

* Hooks **do not stop on error**
* Panics are recovered and returned as errors
* Successful results are collected in order

---

## Execution Order

Hooks are executed by:

1. **Priority (ascending)**
2. **Name (lexicographical, stable sort)**

---

## Error Handling

* Returned errors are wrapped with hook name
* Panics are caught and reported as:

  ```
  hook '<name>' panicked: <value>
  ```

---

## Notes

* Function names are auto-detected via `runtime.FuncForPC`
* Safe for plugin systems, lifecycle hooks, middleware-like flows
