package preview

import (
	"fmt"
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/twitchyliquid64/kcgen/pcb"
)

var (
	padClearance = 0.2 / 2 // also called mask_margin in KiCad source.
)

// somewhat loosely based off pcbnew/pad_print_functions.cpp.
func renderPad(pad pcb.Pad, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	newOpts := opts
	newOpts.PadClearance = padClearance

	switch pad.Shape {
	case pcb.ShapeRect, pcb.ShapeRoundRect:
		renderRectPad(pad, newOpts, da, cr)
	case pcb.ShapeCircle, pcb.ShapeOval:
		renderCirclePad(pad, newOpts, da, cr)
	default:
		fmt.Printf("Cannot render pad with shape %v: (%+v)\n", pad.Shape, pad)
	}

	if pad.DrillSize.X > 0 {
		cr.NewPath()
		drillCenter := pcb.XY{X: pad.At.X + pad.DrillOffset.X, Y: pad.At.Y + pad.DrillOffset.Y}
		centerX, centerY := opts.ProjectXY(drillCenter)
		r, g, b := opts.GetLayerColor("drill")
		cr.SetSourceRGB(r, g, b)
		cr.Arc(centerX, centerY, pad.DrillSize.X/2, 0, math.Pi*2)
		cr.Fill()
	}
	return nil
}

func renderCirclePad(pad pcb.Pad, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) {
	sizeX, sizeY := opts.ProjectXY(opts.ReduceByClearance(pad.Size))
	centerX, centerY, _ := opts.ProjectXYZ(pad.At)
	r, g, b := opts.GetColorFromLayers(pad.Layers)
	cr.SetSourceRGB(r, g, b)

	mat := cr.GetMatrix()
	cr.Scale(sizeX/sizeY, sizeY/sizeX)
	cr.NewPath()
	cr.Arc(centerX, centerY, sizeX/2, 0, math.Pi*2)
	cr.SetMatrix(mat)
	cr.Fill()
}

func renderRectPad(pad pcb.Pad, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) {
	sizeX, sizeY := opts.ProjectXY(opts.ReduceByClearance(pad.Size))
	centerX, centerY, _ := opts.ProjectXYZ(pad.At)
	r, g, b := opts.GetColorFromLayers(pad.Layers)
	cr.SetSourceRGB(r, g, b)
	if pad.Shape == pcb.ShapeRoundRect {
		rounding := math.Min(0.25, pad.RoundRectRRatio*math.Min(sizeX, sizeY))
		roundedRect(centerX-sizeX/2, centerY-sizeY/2, sizeX, sizeY, rounding, da, cr)
	} else {
		cr.Rectangle(centerX-sizeX/2, centerY-sizeY/2, sizeX, sizeY)
	}
	cr.Fill()
}

func roundedRect(atX, atY, sizeX, sizeY, rounding float64, da *gtk.DrawingArea, cr *cairo.Context) {
	cr.NewPath()
	cr.Arc(atX+sizeX-rounding, atY+rounding, rounding, -math.Pi/2, 0)
	cr.Arc(atX+sizeX-rounding, atY+sizeY-rounding, rounding, 0, math.Pi/2)
	cr.Arc(atX+rounding, atY+sizeY-rounding, rounding, math.Pi/2, math.Pi)
	cr.Arc(atX+rounding, atY+rounding, rounding, -math.Pi, -math.Pi/2)
	cr.ClosePath()
	cr.Fill()
}
