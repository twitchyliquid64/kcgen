package kcgen

type transformMatrix struct {
	XX, YX           float64
	XY, YY           float64
	XOffset, YOffset float64
}

func (m transformMatrix) Apply(x, y float64) (float64, float64) {
	x1 := (m.XX * x) + (m.XY * y) + m.XOffset
	y1 := (m.YX * x) + (m.YY * y) + m.YOffset
	return x1, y1
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
