// Package kcsl implements a DSL for specifying operations on kicad PCBs / Mods.
package kcsl

import (
	"errors"
	"math"

	"github.com/twitchyliquid64/kcgen"
	"github.com/twitchyliquid64/kcgen/pcb"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var (
	builtins = starlark.StringDict{
		// basics
		"XY":          pcb.MakeXY,
		"XYZ":         pcb.MakeXYZ,
		"Layer":       pcb.MakeLayer,
		"Net":         pcb.MakeNet,
		"TextEffects": pcb.MakeTextEffects,
		// modules
		"Mod":          pcb.MakeModule,
		"ModPolygon":   pcb.MakeModPolygon,
		"ModText":      pcb.MakeModText,
		"ModLine":      pcb.MakeModLine,
		"ModCircle":    pcb.MakeModCircle,
		"ModArc":       pcb.MakeModArc,
		"ModModel":     pcb.MakeModModel,
		"ModPlacement": pcb.MakeModPlacement,
		"ModGraphic":   pcb.MakeModGraphic,
		"PadOptions":   pcb.MakePadOptions,
		"Pad":          pcb.MakePad,
		// builtins in own namespace
		"math":         starlarkstruct.FromStringDict(starlarkstruct.Default, mathBuiltins),
		"layers":       starlarkstruct.FromStringDict(starlarkstruct.Default, layers),
		"zone_connect": starlarkstruct.FromStringDict(starlarkstruct.Default, zc),
		"text":         starlarkstruct.FromStringDict(starlarkstruct.Default, txt),
		"pad":          starlarkstruct.FromStringDict(starlarkstruct.Default, pad),
		"shape":        starlarkstruct.FromStringDict(starlarkstruct.Default, shape),
		"defaults":     starlarkstruct.FromStringDict(starlarkstruct.Default, defaults),
		"struct":       starlark.NewBuiltin("struct", starlarkstruct.Make),
		// aux
		"crash": starlark.NewBuiltin("crash", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			return nil, errors.New("soft crash: " + args[0].String())
		}),
	}
	mathBuiltins = starlark.StringDict{
		"pi": starlark.Float(math.Pi),
		"sin": starlark.NewBuiltin("sin", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("sin", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Sin(float64(theta))), nil
		}),
		"sinh": starlark.NewBuiltin("sinh", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("sinh", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Sinh(float64(theta))), nil
		}),
		"asin": starlark.NewBuiltin("asin", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("asin", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Asin(float64(theta))), nil
		}),
		"asinh": starlark.NewBuiltin("asinh", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("asinh", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Asinh(float64(theta))), nil
		}),
		"cos": starlark.NewBuiltin("cos", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("cos", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Cos(float64(theta))), nil
		}),
		"tan": starlark.NewBuiltin("tan", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("tan", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Tan(float64(theta))), nil
		}),
		"tanh": starlark.NewBuiltin("tanh", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("tanh", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Tanh(float64(theta))), nil
		}),
		"atan": starlark.NewBuiltin("atan", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var theta starlark.Float
			if err := starlark.UnpackArgs("atan", args, kwargs, "theta", &theta); err != nil {
				return starlark.None, err
			}
			return starlark.Float(math.Atan(float64(theta))), nil
		}),

		"shl": starlark.NewBuiltin("shl", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var base starlark.Int
			var shiftAmount int
			if err := starlark.UnpackArgs("shl", args, kwargs, "base", &base, "shift", &shiftAmount); err != nil {
				return starlark.None, err
			}
			b, ok := base.Uint64()
			if !ok {
				return starlark.None, errors.New("cannot represent base as unsigned integer")
			}
			return starlark.MakeUint64(b << uint(shiftAmount)), nil
		}),
		"shr": starlark.NewBuiltin("shr", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var base starlark.Int
			var shiftAmount int
			if err := starlark.UnpackArgs("shr", args, kwargs, "base", &base, "shift", &shiftAmount); err != nil {
				return starlark.None, err
			}
			b, ok := base.Uint64()
			if !ok {
				return starlark.None, errors.New("cannot represent base as unsigned integer")
			}
			return starlark.MakeUint64(b >> uint(shiftAmount)), nil
		}),
		"_not": starlark.NewBuiltin("_not", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var base starlark.Int
			if err := starlark.UnpackArgs("_not", args, kwargs, "base", &base); err != nil {
				return starlark.None, err
			}
			b, ok := base.Uint64()
			if !ok {
				return starlark.None, errors.New("cannot represent base as unsigned integer")
			}
			return starlark.MakeUint64(^b), nil
		}),
		"_and": starlark.NewBuiltin("_and", func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var o1, o2 starlark.Int
			if err := starlark.UnpackArgs("_and", args, kwargs, "op1", &o1, "op2", &o2); err != nil {
				return starlark.None, err
			}
			i1, ok := o1.Uint64()
			if !ok {
				return starlark.None, errors.New("cannot represent op1 as unsigned integer")
			}
			i2, ok := o2.Uint64()
			if !ok {
				return starlark.None, errors.New("cannot represent op2 as unsigned integer")
			}
			return starlark.MakeUint64(i1 & i2), nil
		}),
	}

	layers = starlark.StringDict{
		"front": starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
			"copper":     starlark.String(kcgen.LayerFrontCopper.Strictname()),
			"fab":        starlark.String(kcgen.LayerFrontFab.Strictname()),
			"silkscreen": starlark.String(kcgen.LayerFrontSilkscreen.Strictname()),
			"paste":      starlark.String(kcgen.LayerFrontPaste.Strictname()),
			"smd": starlark.NewList([]starlark.Value{
				starlark.String(kcgen.LayerFrontCopper.Strictname()),
				starlark.String(kcgen.LayerFrontPaste.Strictname()),
				starlark.String(kcgen.LayerFrontMask.Strictname()),
			}),
		}),
		"back": starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
			"copper":     starlark.String(kcgen.LayerBackCopper.Strictname()),
			"fab":        starlark.String(kcgen.LayerBackFab.Strictname()),
			"silkscreen": starlark.String(kcgen.LayerBackSilkscreen.Strictname()),
			"paste":      starlark.String(kcgen.LayerBackPaste.Strictname()),
			"smd": starlark.NewList([]starlark.Value{
				starlark.String(kcgen.LayerBackCopper.Strictname()),
				starlark.String(kcgen.LayerBackPaste.Strictname()),
				starlark.String(kcgen.LayerBackMask.Strictname()),
			}),
		}),
		"th": starlark.NewList([]starlark.Value{
			starlark.String(kcgen.LayerAllCopper.Strictname()),
			starlark.String(kcgen.LayerAllMask.Strictname()),
		}),
	}

	zci = pcb.ZoneConnectInherited
	zcn = pcb.ZoneConnectNone
	zct = pcb.ZoneConnectThermal
	zc  = starlark.StringDict{
		"inherited": &zci,
		"none":      &zcn,
		"thermal":   &zct,
	}

	rt  = pcb.RefText
	ut  = pcb.UserText
	vt  = pcb.ValueText
	txt = starlark.StringDict{
		"reference": &rt,
		"user":      &ut,
		"value":     &vt,
		"vertical":  starlark.Float(90),
	}

	defaults = starlark.StringDict{
		"width":     starlark.Float(0.15),
		"clearance": starlark.Float(0.20),
		"thickness": starlark.Float(0.2),
	}

	sth   = pcb.SurfaceTH
	snpth = pcb.SurfaceNPTH
	ssmd  = pcb.SurfaceSMD
	pad   = starlark.StringDict{
		"through_hole":    &sth,
		"np_through_hole": &snpth,
		"smd":             &ssmd,
	}

	soval  = pcb.ShapeOval
	scirc  = pcb.ShapeCircle
	srect  = pcb.ShapeRect
	srrect = pcb.ShapeRoundRect
	shape  = starlark.StringDict{
		"oval":       &soval,
		"circle":     &scirc,
		"rect":       &srect,
		"round_rect": &srrect,
	}
)
