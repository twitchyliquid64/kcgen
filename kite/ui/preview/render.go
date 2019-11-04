package preview

import (
	"fmt"
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/twitchyliquid64/kcgen"
	"github.com/twitchyliquid64/kcgen/pcb"
)

// Render is called to provide a module to the previewer and start the preview.
func (p *Preview) Render(mod *pcb.Module) {
	p.mod = mod
	p.canvas.QueueDraw()
}

type modRenderOptions struct {
	Width, Height int
	X, Y          float64
	Zoom          float64
	PadClearance  float64
}

func (o modRenderOptions) ProjectXY(pt pcb.XY) (x, y float64) {
	return pt.X, pt.Y
}

func (o modRenderOptions) ProjectXYZ(pt pcb.XYZ) (x, y, z float64) {
	if pt.ZPresent {
		return pt.X, pt.Y, pt.Z
	}
	return pt.X, pt.Y, 0
}

func (o modRenderOptions) ReduceByClearance(pt pcb.XY) pcb.XY {
	return pcb.XY{X: pt.X - o.PadClearance, Y: pt.Y - o.PadClearance}
}

func (o modRenderOptions) GetLayerColor(layer string) (r, g, b float64) {
	switch layer {
	case "drill":
		return 37.0 / 255, 37.0 / 255, 37.0 / 255
	case kcgen.LayerFrontCourtyard.Strictname():
		return 72.0 / 255, 72.0 / 255, 72.0 / 255
	case kcgen.LayerFrontCopper.Strictname():
		return 132.0 / 255, 0, 0
	case kcgen.LayerFrontSilkscreen.Strictname():
		return 0, 132.0 / 255, 132.0 / 255
	}

	return 1, 1, 1
}

func (o modRenderOptions) GetColorFromLayers(l []string) (r, g, b float64) {
	return 132.0 / 255, 0, 0
}

func renderModule(mod *pcb.Module, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	for _, graphic := range mod.Graphics {
		switch graphic.Ident {
		case "fp_line":
			if err := renderLine(graphic.Renderable.(*pcb.ModLine), opts, da, cr); err != nil {
				return fmt.Errorf("rendering line: %v", err)
			}
		case "fp_poly":
			if err := renderPoly(graphic.Renderable.(*pcb.ModPolygon), opts, da, cr); err != nil {
				return fmt.Errorf("rendering polygon: %v", err)
			}
		case "fp_circle":
			if err := renderCircle(graphic.Renderable.(*pcb.ModCircle), opts, da, cr); err != nil {
				return fmt.Errorf("rendering circle: %v", err)
			}
		default:
			fmt.Printf("Cannot render: %v (%+v)\n", graphic.Ident, graphic.Renderable)
		}
	}

	for _, pad := range mod.Pads {
		if err := renderPad(pad, opts, da, cr); err != nil {
			return fmt.Errorf("rendering pad: %v", err)
		}
	}
	return nil
}

func renderLine(line *pcb.ModLine, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	cr.SetLineJoin(cairo.LINE_JOIN_ROUND)
	cr.SetLineCap(cairo.LINE_CAP_ROUND)
	cr.SetLineWidth(line.Width)
	startX, startY := opts.ProjectXY(line.Start)
	endX, endY := opts.ProjectXY(line.End)
	r, g, b := opts.GetLayerColor(line.Layer)
	cr.SetSourceRGB(r, g, b)

	if curX, curY := cr.GetCurrentPoint(); startX != curX || startY != curY {
		cr.MoveTo(startX, startY)
	}
	cr.LineTo(endX, endY)

	cr.Stroke()
	return nil
}

func renderCircle(c *pcb.ModCircle, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	cr.SetLineJoin(cairo.LINE_JOIN_ROUND)
	cr.SetLineCap(cairo.LINE_CAP_ROUND)
	cr.SetLineWidth(c.Width)

	centerX, centerY := opts.ProjectXY(c.Center)
	radius := math.Sqrt(math.Pow(c.Center.X-c.End.X, 2) + math.Pow(c.Center.Y-c.End.Y, 2))
	r, g, b := opts.GetLayerColor(c.Layer)
	cr.SetSourceRGB(r, g, b)
	cr.Arc(centerX, centerY, radius, 0, math.Pi*2)
	cr.Stroke()
	return nil
}

func renderPoly(poly *pcb.ModPolygon, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	cr.SetLineJoin(cairo.LINE_JOIN_ROUND)
	cr.SetLineCap(cairo.LINE_CAP_ROUND)
	wdth := poly.Width
	if wdth < 0.01 {
		wdth = 0.01
	}
	cr.SetLineWidth(wdth)

	x, y := opts.ProjectXY(poly.At)
	cr.MoveTo(x, y)
	r, g, b := opts.GetLayerColor(poly.Layer)
	cr.SetSourceRGB(r, g, b)

	for i, pt := range poly.Points {
		x, y := opts.ProjectXY(pt)
		if i == 0 {
			cr.MoveTo(x, y)
		} else {
			cr.LineTo(x, y)
		}
	}

	cr.Fill()
	return nil
}
