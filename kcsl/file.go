package kcsl

import (
	"bufio"
	"os"

	"github.com/twitchyliquid64/kcgen/pcb"
	"go.starlark.net/starlark"
)

var fileLoadMod = starlark.NewBuiltin("load_mod", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var p starlark.String
	if err := starlark.UnpackArgs("load_mod", args, kwargs,
		"path", &p); err != nil {
		return starlark.None, err
	}

	f, err := os.Open(string(p))
	if err != nil {
		return starlark.None, err
	}
	defer f.Close()
	return pcb.ParseModule(bufio.NewReader(f))
})
