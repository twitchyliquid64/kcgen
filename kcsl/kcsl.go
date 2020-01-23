// Package kcsl implements a DSL for specifying operations on kicad PCBs / Mods.
package kcsl

import (
	"errors"

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
		// PCB
		"PCB":   pcb.MakePCB,
		"Line":  pcb.MakeLine,
		"Arc":   pcb.MakeArc,
		"Text":  pcb.MakeText,
		"Track": pcb.MakeTrack,
		"Via":   pcb.MakeVia,
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
		// textpoly
		"TextPoly": makeTextPoly,
		// file manipulation
		"file": starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
			"load_mod": fileLoadMod,
		}),
	}

	layers = starlark.StringDict{
		"front": starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
			"copper":     starlark.String(kcgen.LayerFrontCopper.Strictname()),
			"fab":        starlark.String(kcgen.LayerFrontFab.Strictname()),
			"silkscreen": starlark.String(kcgen.LayerFrontSilkscreen.Strictname()),
			"paste":      starlark.String(kcgen.LayerFrontPaste.Strictname()),
			"courtyard":  starlark.String(kcgen.LayerFrontCourtyard.Strictname()),
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
			"courtyard":  starlark.String(kcgen.LayerBackCourtyard.Strictname()),
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
