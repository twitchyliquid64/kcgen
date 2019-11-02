package preview

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/pcb"
)

type Preview struct {
	canvas *gtk.DrawingArea
	mod    *pcb.Module

	width, height int
}

func (p *Preview) onCanvasConfigureEvent(da *gtk.DrawingArea, event *gdk.Event) bool {
	ce := gdk.EventConfigureNewFromEvent(event)
	p.width = ce.Width()
	p.height = ce.Height()
	return false
}

func (p *Preview) onCanvasDrawEvent(da *gtk.DrawingArea, cr *cairo.Context) bool {
	cr.SetSourceRGB(0, 0, 0)
	cr.Rectangle(0, 0, float64(p.width), float64(p.height))
	cr.Fill()
	return false
}

func NewPreview(b *gtk.Builder) (*Preview, error) {
	out := &Preview{}

	e, err := b.GetObject("kite_preview")
	if err != nil {
		return nil, err
	}
	out.canvas = e.(*gtk.DrawingArea)

	out.canvas.Connect("draw", out.onCanvasDrawEvent)
	out.canvas.Connect("configure-event", out.onCanvasConfigureEvent)

	return out, nil
}
