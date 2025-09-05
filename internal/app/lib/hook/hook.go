package hook

import (
	"fmt"
	"sort"
)

type HookFunc func(args ...any) (any, error)

type NamedHook struct {
	Name string
	Fn   HookFunc
}

type Hook struct {
	Pre  []NamedHook
	Core []NamedHook
	Post []NamedHook
}

func (h *Hook) AddPre(name string, fn HookFunc)  { h.Pre = append(h.Pre, NamedHook{name, fn}) }
func (h *Hook) AddCore(name string, fn HookFunc) { h.Core = append(h.Core, NamedHook{name, fn}) }
func (h *Hook) AddPost(name string, fn HookFunc) { h.Post = append(h.Post, NamedHook{name, fn}) }

func runHooks(hooks []NamedHook, args ...any) ([]any, []error) {
	sort.SliceStable(hooks, func(i, j int) bool {
		return hooks[i].Name < hooks[j].Name
	})

	var (
		results []any
		errs    []error
	)

	for _, hook := range hooks {
		func() {
			defer func() {
				if r := recover(); r != nil {
					errs = append(errs, fmt.Errorf("panic in hook %s: %v", hook.Name, r))
				}
			}()

			res, err := hook.Fn(args...)
			if err != nil {
				errs = append(errs, fmt.Errorf("hook %s failed: %w", hook.Name, err))
				return
			}
			if res != nil {
				results = append(results, res)
			}
		}()
	}
	return results, errs
}

func (h *Hook) RunPre(args ...any) ([]any, []error) {
	return runHooks(h.Pre, args...)
}

func (h *Hook) RunCore(args ...any) ([]any, []error) {
	return runHooks(h.Core, args...)
}

func (h *Hook) RunPost(args ...any) ([]any, []error) {
	return runHooks(h.Post, args...)
}
