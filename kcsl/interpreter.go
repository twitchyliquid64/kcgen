package kcsl

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/twitchyliquid64/kcgen/kcsl/lib"
	"github.com/twitchyliquid64/kcgen/pcb"
	"go.starlark.net/starlark"
)

// WDLoader loads scripts relative to the working directory.
type WDLoader struct{}

func (l *WDLoader) resolveImport(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

// ScriptLoader provides a means for arbitrary imports to be resolved.
type ScriptLoader interface {
	resolveImport(name string) ([]byte, error)
}

func (s *Script) loadScript(script []byte, fname string, loader ScriptLoader) (*starlark.Thread, starlark.StringDict, error) {
	var moduleCache = map[string]starlark.StringDict{}
	var load func(_ *starlark.Thread, module string) (starlark.StringDict, error)

	load = func(_ *starlark.Thread, module string) (starlark.StringDict, error) {
		m, ok := moduleCache[module]
		if m == nil && ok {
			return nil, errors.New("cycle in dependency graph when loading " + module)
		}
		if m != nil {
			return m, nil
		}

		// loading in progress
		moduleCache[module] = nil
		d, err2 := loader.resolveImport(module)
		if err2 != nil {
			return nil, err2
		}
		thread := &starlark.Thread{
			Print: s.printFromSkylark,
			Load:  load,
		}
		mod, err2 := starlark.ExecFile(thread, module, d, builtins)
		if err2 != nil {
			return nil, err2
		}
		moduleCache[module] = mod
		return mod, nil
	}

	thread := &starlark.Thread{
		Print: s.printFromSkylark,
		Load:  load,
	}

	globals, err := starlark.ExecFile(thread, fname, script, builtins)
	if err != nil {
		return nil, nil, err
	}

	return thread, globals, nil
}

// Script represents a raspberry-box script.
type Script struct {
	loader  ScriptLoader
	printer func(string)

	args    []string
	verbose bool

	thread   *starlark.Thread
	globals  starlark.StringDict
	setupVal starlark.Value
}

// Close shuts down all resources associated with the script.
func (s *Script) Close() error {
	return nil
}

// NewScript initializes a new raspberry-box script environment.
func NewScript(data []byte, fname string, verbose bool, loader ScriptLoader, args []string, printer func(string)) (*Script, error) {
	return makeScript(data, fname, loader, args, verbose, nil, printer)
}

func makeScript(data []byte, fname string, loader ScriptLoader, args []string, verbose bool,
	testHook func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error),
	printer func(string)) (*Script, error) {
	out := &Script{
		loader:  loader,
		args:    args,
		verbose: verbose,
		printer: printer,
	}

	var err error
	out.thread, out.globals, err = out.loadScript(data, fname, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (s *Script) printFromSkylark(_ *starlark.Thread, msg string) {
	if s.printer != nil {
		s.printer(msg)
	} else {
		fmt.Println(msg)
	}
}

func (s *Script) resolveImport(path string) ([]byte, error) {
	d, exists := lib.Libs[path]
	if exists {
		return d, nil
	}
	if s.loader == nil {
		return nil, errors.New("no such import: " + path)
	}
	return s.loader.resolveImport(path)
}

func cvStrListToStarlark(in []string) *starlark.List {
	out := make([]starlark.Value, len(in))
	for i := range in {
		out[i] = starlark.String(in[i])
	}
	return starlark.NewList(out)
}

// CallFn calls an arbitrary function
func (s *Script) CallFn(fname string) (string, error) {
	fn, exists := s.globals[fname]
	if !exists {
		return "", fmt.Errorf("%s() function not present", fname)
	}
	ret, err := starlark.Call(s.thread, fn, starlark.Tuple{}, nil)
	if err != nil {
		return "", err
	}
	result, ok := ret.(starlark.String)
	if !ok {
		return "", fmt.Errorf("%s() returned type %T, want string", fname, ret)
	}
	return string(result), nil
}

// Mod returns a generated module, if applicable.
func (s *Script) Mod() *pcb.Module {
	if m, ok := s.globals["mod"]; ok {
		if mod, ok := m.(*pcb.Module); ok {
			return mod
		}
	}
	return nil
}

// Pcb returns a generated PCB, if applicable.
func (s *Script) Pcb() *pcb.PCB {
	if p, ok := s.globals["pcb"]; ok {
		if pcb, ok := p.(*pcb.PCB); ok {
			return pcb
		}
	}
	return nil
}
