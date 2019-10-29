package kcgen

import (
	"github.com/twitchyliquid64/kcgen/pcb"
)

// Circle represents a circle drawn on a module or PCB.
type Circle struct {
	radius float64
	c      pcb.ModCircle
}

// Center sets the position of the circle.
func (c *Circle) Center(x, y float64) {
	c.c.Center = pcb.XY{X: x, Y: y}
	c.c.End = pcb.XY{X: x + c.radius, Y: y}
}

// Width sets the stroke width of the circle.
func (c *Circle) Width(thicc float64) {
	c.c.Width = thicc
}

// NewCircle creates a new circle graphic.
func NewCircle(radius float64, layer Layer) *Circle {
	return &Circle{
		radius: radius,
		c: pcb.ModCircle{
			End:   pcb.XY{X: radius},
			Width: 0.15,
			Layer: layer.Strictname(),
		},
	}
}

// Arc represents an arc drawn on a module or PCB.
type Arc struct {
	a pcb.ModArc
}

// Start sets the starting point.
func (a *Arc) Start(x, y float64) {
	a.a.Start = pcb.XY{X: x, Y: y}
}

// End sets the finishing point.
func (a *Arc) End(x, y float64) {
	a.a.End = pcb.XY{X: x, Y: y}
}

// Positions sets the starting and finishing point.
func (a *Arc) Positions(x1, y1, x2, y2 float64) {
	a.a.Start = pcb.XY{X: x1, Y: y1}
	a.a.End = pcb.XY{X: x2, Y: y2}
}

// Width sets the stroke width of the arc.
func (a *Arc) Width(thicc float64) {
	a.a.Width = thicc
}

// Angle sets the angle of the arc.
func (a *Arc) Angle(ang float64) {
	a.a.Angle = ang * 10
}

// NewArc creates a new arc graphic.
func NewArc(layer Layer) *Arc {
	return &Arc{
		a: pcb.ModArc{
			Angle: 3600,
			End:   pcb.XY{X: 2},
			Width: 0.15,
			Layer: layer.Strictname(),
		},
	}
}
