package preview

import (
	"fmt"
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"

	"github.com/twitchyliquid64/kcgen"
	"github.com/twitchyliquid64/kcgen/pcb"
)

// RenderMod is called to provide a module to the previewer and start the preview.
func (p *Preview) RenderMod(mod *pcb.Module) {
	p.mod = mod
	p.pcb = nil
	p.canvas.QueueDraw()
}

// RenderPCB is called to provide a pcb to the previewer and start the preview.
func (p *Preview) RenderPCB(pcb *pcb.PCB) {
	p.pcb = pcb
	p.mod = nil
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
	case kcgen.LayerEdgeCuts.Strictname():
		return 132.0 / 255, 132.0 / 255, 0
	}

	return 1, 1, 1
}

func (o modRenderOptions) GetColorFromLayers(l []string) (r, g, b float64) {
	return 132.0 / 255, 0, 0
}

func renderPCB(board *pcb.PCB, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	for _, graphic := range board.Drawings {
		switch g := graphic.(type) {
		case *pcb.Line:
			if err := renderLine(lineOpts{
				width: g.Width,
				start: g.Start,
				end:   g.End,
				layer: g.Layer,
			}, opts, da, cr); err != nil {
				return fmt.Errorf("rendering line: %v", err)
			}

		case *pcb.Arc:
			if err := renderArc(arcOpts{
				width: g.Width,
				start: g.Start,
				end:   g.End,
				angle: g.Angle,
				layer: g.Layer,
			}, opts, da, cr); err != nil {
				return fmt.Errorf("rendering arc: %v", err)
			}
		default:
			fmt.Printf("Cannot render: %v (%+v)\n", g, g)
		}
	}

	return nil
}

func renderModule(mod *pcb.Module, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	for _, graphic := range mod.Graphics {
		switch graphic.Ident {
		case "fp_line":
			ml := graphic.Renderable.(*pcb.ModLine)
			if err := renderLine(lineOpts{
				width: ml.Width,
				start: ml.Start,
				end:   ml.End,
				layer: ml.Layer,
			}, opts, da, cr); err != nil {
				return fmt.Errorf("rendering line: %v", err)
			}
		case "fp_poly":
			if err := renderPoly(graphic.Renderable.(*pcb.ModPolygon), opts, da, cr); err != nil {
				return fmt.Errorf("rendering polygon: %v", err)
			}
		case "fp_circle":
			mc := graphic.Renderable.(*pcb.ModCircle)

			if err := renderCircle(circleOpts{
				width:  mc.Width,
				center: mc.Center,
				end:    mc.End,
				layer:  mc.Layer,
			}, opts, da, cr); err != nil {
				return fmt.Errorf("rendering circle: %v", err)
			}
		case "fp_arc":
			ma := graphic.Renderable.(*pcb.ModArc)
			if err := renderArc(arcOpts{
				width: ma.Width,
				start: ma.Start,
				end:   ma.End,
				angle: ma.Angle,
				layer: ma.Layer,
			}, opts, da, cr); err != nil {
				return fmt.Errorf("rendering arc: %v", err)
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

type lineOpts struct {
	width float64
	start pcb.XY
	end   pcb.XY
	layer string
}

func renderLine(l lineOpts, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	cr.SetLineJoin(cairo.LINE_JOIN_ROUND)
	cr.SetLineCap(cairo.LINE_CAP_ROUND)
	cr.SetLineWidth(l.width)
	startX, startY := opts.ProjectXY(l.start)
	endX, endY := opts.ProjectXY(l.end)
	r, g, b := opts.GetLayerColor(l.layer)
	cr.SetSourceRGB(r, g, b)

	if curX, curY := cr.GetCurrentPoint(); startX != curX || startY != curY {
		cr.MoveTo(startX, startY)
	}
	cr.LineTo(endX, endY)

	cr.Stroke()
	return nil
}

type circleOpts struct {
	width  float64
	center pcb.XY
	end    pcb.XY
	layer  string
}

func renderCircle(c circleOpts, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	cr.SetLineJoin(cairo.LINE_JOIN_ROUND)
	cr.SetLineCap(cairo.LINE_CAP_ROUND)
	cr.SetLineWidth(c.width)

	centerX, centerY := opts.ProjectXY(c.center)
	radius := math.Sqrt(math.Pow(c.center.X-c.end.X, 2) + math.Pow(c.center.Y-c.end.Y, 2))
	r, g, b := opts.GetLayerColor(c.layer)
	cr.SetSourceRGB(r, g, b)
	cr.Arc(centerX, centerY, radius, 0, math.Pi*2)
	cr.Stroke()
	return nil
}

type arcOpts struct {
	width float64
	start pcb.XY
	end   pcb.XY
	angle float64
	layer string
}

func renderArc(a arcOpts, opts modRenderOptions, da *gtk.DrawingArea, cr *cairo.Context) error {
	cr.SetLineJoin(cairo.LINE_JOIN_ROUND)
	cr.SetLineCap(cairo.LINE_CAP_ROUND)
	cr.SetLineWidth(a.width)

	startX, startY := opts.ProjectXY(a.start)
	endX, endY := opts.ProjectXY(a.end)
	radius := math.Sqrt(math.Pow(startX-endX, 2) + math.Pow(startY-endY, 2))
	startAngle := math.Atan2(endY-startY, endX-startX)
	endAngle := startAngle + (a.angle * math.Pi / 180)

	r, g, b := opts.GetLayerColor(a.layer)
	cr.SetSourceRGB(r, g, b)
	cr.Arc(startX, startY, radius, startAngle, endAngle)
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
