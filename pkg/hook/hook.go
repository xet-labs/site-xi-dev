package hook

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
)

// Hook function signature
type ArgFn[T, R any] func(T) (R, error)

// NoArgHook Hook container, takes no args
type NoArgFn[R any] func() (R, error)

// EffectHook Hook container, no return value
type EffectFn[T any] func(T) error

// Internal representation
type NamedHook[T, R any] struct {
	Name     string
	Priority int
	Fn       ArgFn[T, R]
}

// Hook container
type Hook[T, R any] struct {
	hooks []NamedHook[T, R]
}

// NoArgHook Hook container, takes no args
type NoArgHook[R any] struct {
	*Hook[struct{}, R]
}

// EffectHook Hook container, no return value
type EffectHook[T any] struct {
	*Hook[T, struct{}]
}

// Constructor
func New[T, R any]() *Hook[T, R] {
	return &Hook[T, R]{}
}
func NewNoArg[R any]() *NoArgHook[R] {
	return &NoArgHook[R]{
		Hook: New[struct{}, R](),
	}
}
func NewEffectHook[T any]() *EffectHook[T] {
	return &EffectHook[T]{Hook: New[T, struct{}]()}
}

// Add hooks

// Add registers a hook using the function name and default priority (0)
func (h *Hook[T, R]) Add(fns ...ArgFn[T, R]) {
	for _, fn := range fns {
		h.hooks = append(h.hooks, NamedHook[T, R]{
			Name:     FnName(fn),
			Priority: 0,
			Fn:       fn,
		})
	}
}

// AddNamed registers a hook with an explicit name
func (h *Hook[T, R]) AddNamed(name string, fn ArgFn[T, R]) {
	h.hooks = append(h.hooks, NamedHook[T, R]{
		Name:     name,
		Priority: 0,
		Fn:       fn,
	})
}

// AddWithPriority registers a hook with explicit name and priority
func (h *Hook[T, R]) AddWithPriority(name string, priority int, fn ArgFn[T, R]) {
	h.hooks = append(h.hooks, NamedHook[T, R]{
		Name:     name,
		Priority: priority,
		Fn:       fn,
	})
}

// Add no arg hook
func (h *NoArgHook[R]) Add(fns ...NoArgFn[R]) {
	for _, fn := range fns {
		h.Hook.Add(func(_ struct{}) (R, error) {
			return fn()
		})
	}
}

func (h *NoArgHook[R]) AddNamed(name string, fn NoArgFn[R]) {
	h.Hook.AddNamed(name, func(_ struct{}) (R, error) {
		return fn()
	})
}

func (h *NoArgHook[R]) AddWithPriority(name string, priority int, fn NoArgFn[R]) {
	wrapped := func(_ struct{}) (R, error) {
		return fn()
	}
	h.Hook.AddWithPriority(name, priority, wrapped)
}

// Add effect hook
func (h *EffectHook[T]) Add(fns ...EffectFn[T]) {
	for _, fn := range fns {
		h.Hook.Add(func(t T) (struct{}, error) {
			return struct{}{}, fn(t)
		})
	}
}

func (h *EffectHook[T]) AddNamed(name string, fn EffectFn[T]) {
	h.Hook.AddNamed(name, func(t T) (struct{}, error) {
		return struct{}{}, fn(t)
	})
}

func (h *EffectHook[T]) AddWithPriority(name string, priority int, fn EffectFn[T]) {
	wrapped := func(t T) (struct{}, error) {
		return struct{}{}, fn(t)
	}
	h.Hook.AddWithPriority(name, priority, wrapped)
}

// Execute hooks
func (h *Hook[T, R]) Run(arg T) ([]R, []error) {
	// stable sort: lower priority runs first
	sort.SliceStable(h.hooks, func(i, j int) bool {
		if h.hooks[i].Priority == h.hooks[j].Priority {
			return h.hooks[i].Name < h.hooks[j].Name
		}
		return h.hooks[i].Priority < h.hooks[j].Priority
	})

	var (
		results []R
		errs    []error
	)

	for _, hook := range h.hooks {
		func() {
			defer func() {
				if r := recover(); r != nil {
					errs = append(errs,
						fmt.Errorf("hook '%s' panicked: %v", hook.Name, r),
					)
				}
			}()

			res, err := hook.Fn(arg)
			if err != nil {
				errs = append(errs,
					fmt.Errorf("hook '%s' failed: %w", hook.Name, err),
				)
				return
			}

			results = append(results, res)
		}()
	}

	return results, errs
}

// Utilities
func FnName(fn any) string {
	return runtime.FuncForPC(
		reflect.ValueOf(fn).Pointer(),
	).Name()
}
