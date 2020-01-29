package preview

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/pcb"
)

// Preview encapsulates the UI widget for previewing a module.
type Preview struct {
	canvas *gtk.DrawingArea
	mod    *pcb.Module
	pcb    *pcb.PCB

	zoom             float64
	offsetX, offsetY float64

	dragStartX, dragStartY float64
	dragging               bool

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

	cr.Translate(float64(p.width/2)+p.offsetX, float64(p.height/2)+p.offsetY)
	cr.Scale(p.zoom, p.zoom)

	if p.mod != nil {
		if err := renderModule(p.mod, modRenderOptions{
			Width:  p.width,
			Height: p.height,
		}, da, cr); err != nil {
			fmt.Fprintf(os.Stderr, "Render failed: %v\n", err)
		}
	}
	return false
}

func (p *Preview) onMotionEvent(area *gtk.DrawingArea, event *gdk.Event) {
	evt := gdk.EventMotionNewFromEvent(event)
	if p.dragging {
		x, y := evt.MotionVal()
		p.offsetX = -(p.dragStartX - x)
		p.offsetY = -(p.dragStartY - y)
	}
	p.canvas.QueueDraw()
}

func (p *Preview) onPressEvent(area *gtk.DrawingArea, event *gdk.Event) {
	evt := gdk.EventButtonNewFromEvent(event)
	switch evt.Button() {
	case 2: // middle button
		p.dragging = true
		p.dragStartX, p.dragStartY = gdk.EventMotionNewFromEvent(event).MotionVal()
		p.dragStartX -= p.offsetX
		p.dragStartY -= p.offsetY
	}
}

func (p *Preview) onReleaseEvent(area *gtk.DrawingArea, event *gdk.Event) {
	evt := gdk.EventButtonNewFromEvent(event)
	switch evt.Button() {
	case 2: // middle button
		p.dragging = false
	}
}

func (p *Preview) onScrollEvent(area *gtk.DrawingArea, event *gdk.Event) {
	evt := gdk.EventScrollNewFromEvent(event)
	amt := evt.DeltaY()
	if amt == 0 {
		amt = 1
	}

	switch evt.Direction() {
	case gdk.SCROLL_DOWN:
		amt *= -1
	}

	p.zoom += amt
	p.canvas.QueueDraw()
}

// NewPreview creates a new preview UI widget.
func NewPreview(b *gtk.Builder) (*Preview, error) {
	out := &Preview{
		zoom: 25,
	}

	e, err := b.GetObject("kite_preview")
	if err != nil {
		return nil, err
	}
	out.canvas = e.(*gtk.DrawingArea)

	out.canvas.Connect("draw", out.onCanvasDrawEvent)
	out.canvas.Connect("configure-event", out.onCanvasConfigureEvent)

	out.canvas.Connect("motion-notify-event", out.onMotionEvent)
	out.canvas.Connect("button-press-event", out.onPressEvent)
	out.canvas.Connect("button-release-event", out.onReleaseEvent)
	out.canvas.Connect("scroll-event", out.onScrollEvent)
	out.canvas.SetEvents(int(gdk.POINTER_MOTION_MASK |
		gdk.BUTTON_PRESS_MASK |
		gdk.BUTTON_RELEASE_MASK |
		gdk.SCROLL_MASK)) // GDK_MOTION_NOTIFY

	return out, nil
}
