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
	var size, dpiVal, scaleVal starlark.Value
	var at starlark.Value
	if err := starlark.UnpackArgs("TextPoly", args, kwargs,
		"font", &font, "content", &text, "at?", &at,
		"size?", &size, "dpi?", &dpiVal, "scale?", &scaleVal); err != nil {
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
	var scale float64 = 1
	if scaleVal != nil {
		if val, ok := starlark.AsFloat(scaleVal); ok {
			scale = val
		}
	}

	var offset pcb.XY
	if at != nil {
		p, ok := at.(*pcb.XY)
		if !ok {
			return starlark.None, fmt.Errorf("at is type %T, wanted XY", at)
		}
		offset = *p
	}

	tp, err := textpoly.NewVectorizer(string(font), fs, dpi)
	if err != nil {
		return starlark.None, err
	}
	if err := tp.DrawString(string(text), fixed.Point26_6{}); err != nil {
		return starlark.None, err
	}

	var out []starlark.Value
	for _, poly := range tp.Vectors() {
		var innerOut []starlark.Value
		for _, p := range poly {
			x := (p[0] * scale) + offset.X
			y := (p[1] * scale) + offset.Y
			innerOut = append(innerOut, &pcb.XY{X: x, Y: y})
		}
		out = append(out, starlark.NewList(innerOut))
	}
	return starlark.NewList(out), nil
})
