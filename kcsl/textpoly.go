package kcsl

import (
	"fmt"

	"github.com/twitchyliquid64/kcgen/kcsl/textpoly"
	"github.com/twitchyliquid64/kcgen/pcb"
	"go.starlark.net/starlark"
	"golang.org/x/image/math/fixed"
)

var makeTextPoly = starlark.NewBuiltin("TextPoly", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var font, text starlark.String
	var size, dpiVal starlark.Value
	var at starlark.Value
	if err := starlark.UnpackArgs("TextPoly", args, kwargs, "font", &font, "content", &text, "size?", &size, "at?", &at); err != nil {
		return starlark.None, err
	}

	var fs float64 = 12
	if size != nil {
		if val, ok := starlark.AsFloat(size); ok {
			fs = val
		}
	}
	var dpi float64 = 72
	if dpiVal != nil {
		if val, ok := starlark.AsFloat(dpiVal); ok {
			dpi = val
		}
	}

	var point fixed.Point26_6
	if at != nil {
		p, ok := at.(*pcb.XY)
		if !ok {
			return starlark.None, fmt.Errorf("at is type %T, wanted XY", at)
		}
		point = fixed.Point26_6{X: fixed.Int26_6(p.X * 64), Y: fixed.Int26_6(p.Y * 64)}
	}

	tp, err := textpoly.NewVectorizer(string(font), fs, dpi)
	if err != nil {
		return starlark.None, err
	}
	if err := tp.DrawString(string(text), point); err != nil {
		return starlark.None, err
	}

	var out []starlark.Value
	for _, poly := range tp.Vectors() {
		var innerOut []starlark.Value
		for _, p := range poly {
			innerOut = append(innerOut, &pcb.XY{X: p[0], Y: p[1]})
		}
		out = append(out, starlark.NewList(innerOut))
	}
	return starlark.NewList(out), nil
})
