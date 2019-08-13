// Package kcgen implements a minimal library for generating kicad footprints.
package kcgen

import (
	"github.com/twitchyliquid64/kcgen/pcb"
)

// Line represents a 2d graphical line in a Footprint.
type Line struct {
	Layer      Layer
	Start, End pcb.XY
	Width      float64
}

// Polygon represents a 2d polygon in a Footprint.
type Polygon struct {
	Layer  Layer
	Points []pcb.XY
	Width  float64
}

// Text represents a string of text drawn at a position.
type Text struct {
	t pcb.ModText
}

// Italic sets whether the text is italized or not.
func (t *Text) Italic(on bool) {
	t.t.Effects.Italic = on
}

// Italic sets whether the text is bolded or not.
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
