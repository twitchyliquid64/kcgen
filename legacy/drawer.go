package kcgen

import (
	"math"
)

type transformMatrix struct {
	XX, YX           float64
	XY, YY           float64
	XOffset, YOffset float64
}

// Apply returns the point after it has been transformed by the matrix.
func (m transformMatrix) Apply(x, y float64) (float64, float64) {
	x1 := (m.XX * x) + (m.XY * y) + m.XOffset
	y1 := (m.YX * x) + (m.YY * y) + m.YOffset
	return x1, y1
}

func (m transformMatrix) Translate(x, y float64) transformMatrix {
	return transformMatrix{
		XX:      1,
		YX:      0,
		XY:      0,
		YY:      1,
		XOffset: x,
		YOffset: y,
	}.Multiply(m)
}

func (m transformMatrix) Rotate(angle float64) transformMatrix {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return transformMatrix{
		XX: c,
		YX: s,
		XY: -s,
		YY: c,
	}.Multiply(m)
}

func (m transformMatrix) Multiply(in transformMatrix) transformMatrix {
	return transformMatrix{
		XX:      m.XX*in.XX + m.YX*in.XY,
		YX:      m.XX*in.YX + m.YX*in.YY,
		XY:      m.XY*in.XX + m.YY*in.XY,
		YY:      m.XY*in.YX + m.YY*in.YY,
		XOffset: m.XOffset*in.XX + m.YOffset*in.XY + in.XOffset,
		YOffset: m.XOffset*in.YX + m.YOffset*in.YY + in.YOffset,
	}
}

func newTransformMatrix() transformMatrix {
	return transformMatrix{XX: 1, YY: 1}
}

// DrawContext provides an easy API for drawing lines.
// The caller should call Draw() to flush drawing commands
// to the underlying module.
type DrawContext struct {
	mb *ModBuilder

	m          transformMatrix
	hasCurrent bool
	start      [2]float64
	current    [2]float64
	segments   [][2][2]float64

	width float64
	layer Layer
}

// Translate applies a translation to future drawing commands.
func (c *DrawContext) Translate(x, y float64) {
	c.m = c.m.Translate(x, y)
}

// Rotate applies a rotation to future drawing commands. The angle
// is given in radians.
func (c *DrawContext) Rotate(r float64) {
	c.m = c.m.Rotate(r)
}

// Draw finalizes the drawing and commits it to the module.
func (c *DrawContext) Draw() {
	for _, s := range c.segments {
		l := NewLine(c.layer)
		l.Width(c.width)
		l.Positions(s[0][0], s[0][1], s[1][0], s[1][1])
		c.mb.AddLine(l)
	}
}

// LineTo draws a line from the current position.
func (c *DrawContext) LineTo(x, y float64) {
	if !c.hasCurrent {
		c.MoveTo(x, y)
		return
	}
	x, y = c.m.Apply(x, y)

	p := [2]float64{x, y}
	c.segments = append(c.segments, [2][2]float64{c.current, p})
	c.current = p
}

// MoveTo moves the cursor to the specified position.
func (c *DrawContext) MoveTo(x, y float64) {
	x, y = c.m.Apply(x, y)

	c.current = [2]float64{x, y}
	c.start = c.current
	c.hasCurrent = true
}

// DrawContext returns a context for drawing operations.
func (m *ModBuilder) DrawContext(width float64, layer Layer) *DrawContext {
	return &DrawContext{
		mb:    m,
		layer: layer,
		width: width,
		m:     newTransformMatrix(),
	}
}
