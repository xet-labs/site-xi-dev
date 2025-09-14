package hook

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
)

type (
	HookFn   func(args ...any) (any, error)
	HookPre  HookFn
	HookCore HookFn
	HookPost HookFn
)

type NamedHook struct {
	Name string
	Fn   HookFn
}

type Hook struct {
	Pre  []NamedHook
	Core []NamedHook
	Post []NamedHook

	PreCheck  func(obj any) (HookFn, bool)
	CoreCheck func(obj any) (HookFn, bool)
	PostCheck func(obj any) (HookFn, bool)
}

// Factory for a Hook configured with interface signatures
func NewHook(
	preCheck  func(obj any) (HookFn, bool),
	coreCheck func(obj any) (HookFn, bool),
	postCheck func(obj any) (HookFn, bool),
) *Hook {
	return &Hook{
		PreCheck:  preCheck,
		CoreCheck: coreCheck,
		PostCheck: postCheck,
	}
}

func FnName(fn any) string { return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name() }

// Standard Add methods
func (h *Hook) AddPre(fns ...HookFn) {
	for _, fn := range fns { h.Pre = append(h.Pre, NamedHook{"pre:" + FnName(fn), fn}) }
}
func (h *Hook) AddCore(fns ...HookFn) {
	for _, fn := range fns { h.Core = append(h.Core, NamedHook{"core:" + FnName(fn), fn}) }
}
func (h *Hook) AddPost(fns ...HookFn) {
	for _, fn := range fns { h.Post = append(h.Post, NamedHook{"post:" + FnName(fn), fn}) }
}

func (h *Hook) RunPre(args ...any)  ([]any, []error) { return runHooks(h.Pre, args...)  }
func (h *Hook) RunCore(args ...any) ([]any, []error) { return runHooks(h.Core, args...) }
func (h *Hook) RunPost(args ...any) ([]any, []error) { return runHooks(h.Post, args...) }

func (h *Hook) Add(fn HookPre) {
	h.Pre = append(h.Pre, NamedHook{"func_" + FnName(fn), HookFn(fn)})
}

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
					errs = append(errs, fmt.Errorf("panic in hook:%s: %v", hook.Name, r))
				}
			}()

			res, err := hook.Fn(args...)
			if err != nil {
				errs = append(errs, fmt.Errorf("hook:%s failed: %w", hook.Name, err))
				return
			}
			if res != nil {
				results = append(results, res)
			}
		}()
	}
	return results, errs
}
