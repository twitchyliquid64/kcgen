// Package kcgen implements the legacy library for generating kicad footprints.
package kcgen

import (
	"github.com/twitchyliquid64/kcgen/pcb"
)

// Line represents a 2d graphical line in a Footprint.
type Line struct {
	l pcb.ModLine
}

// Width sets the thickness of the line.
func (l *Line) Width(thicc float64) {
	l.l.Width = thicc
}

// Start sets the starting point.
func (l *Line) Start(x, y float64) {
	l.l.Start = pcb.XY{X: x, Y: y}
}

// GetStart returns the starting point.
func (l *Line) GetStart() (x, y float64) {
	return l.l.Start.X, l.l.Start.Y
}

// End sets the finishing point.
func (l *Line) End(x, y float64) {
	l.l.End = pcb.XY{X: x, Y: y}
}

// GetEnd returns the finishing point.
func (l *Line) GetEnd() (x, y float64) {
	return l.l.End.X, l.l.End.Y
}

// Positions sets the starting and finishing point.
func (l *Line) Positions(x1, y1, x2, y2 float64) {
	l.l.Start = pcb.XY{X: x1, Y: y1}
	l.l.End = pcb.XY{X: x2, Y: y2}
}

// NewLine creates a new line element.
func NewLine(layer Layer) *Line {
	return &Line{
		l: pcb.ModLine{
			Width: 0.15,
			Layer: layer.Strictname(),
		},
	}
}

// Polygon represents a 2d polygon in a Footprint.
type Polygon struct {
	p pcb.ModPolygon
}

// NewPolygon creates a new polygon element.
func NewPolygon(points [][2]float64, width float64, layer Layer) *Polygon {
	p := &Polygon{
		p: pcb.ModPolygon{
			Layer: layer.Strictname(),
			Width: width,
		},
	}
	p.p.Points = make([]pcb.XY, len(points))
	for i := range points {
		p.p.Points[i] = pcb.XY{X: points[i][0], Y: points[i][1]}
	}
	return p
}

// Text represents a string of text drawn at a position.
type Text struct {
	t pcb.ModText
}

// Italic sets whether the text is italized or not.
func (t *Text) Italic(on bool) {
	t.t.Effects.Italic = on
}

// Bold sets whether the text is bolded or not.
func (t *Text) Bold(on bool) {
	t.t.Effects.Bold = on
}

// Position sets the position of the text.
func (t *Text) Position(x, y, z float64) {
	t.t.At = pcb.XYZ{X: x, Y: y, Z: z}
	if z != 0 {
		t.t.At.ZPresent = true
	}
}

// Hidden sets whether the text is hidden or not.
func (t *Text) Hidden(on bool) {
	t.t.Hidden = on
}

// FontSize sets the size of the text.
func (t *Text) FontSize(x, y float64) {
	t.t.Effects.FontSize = pcb.XY{X: x, Y: y}
}

// Thickness sets the thickness of the text.
func (t *Text) Thickness(thicc float64) {
	t.t.Effects.Thickness = thicc
}

// NewText creates a new text element.
func NewText(text string, layer Layer) *Text {
	return &Text{
		t: pcb.ModText{
			Kind:  pcb.UserText,
			Text:  text,
			Layer: layer.Strictname(),
			Effects: pcb.TextEffects{
				FontSize:  pcb.XY{X: 1, Y: 1},
				Thickness: 0.15,
			},
		},
	}
}
