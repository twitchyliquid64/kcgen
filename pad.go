package kcgen

import (
	"fmt"
	"io"
)

// Pad represents a connection point (pad/TH) in the module.
type Pad struct {
	Layers []Layer
	Center Point2D
	Size   Point2D
	Type   string
	Number int
	Drill  float64
}

// Render generates output suitable for inclusion in a kicad_mod file.
func (p *Pad) Render(w io.Writer) error {
	layers := ""
	for i := range p.Layers {
		layers += p.Layers[i].Strictname()
		if i < (len(p.Layers) - 1) {
			layers += " "
		}
	}

	if _, err := fmt.Fprintf(w, "  (pad %d %s %s %s (drill %s) (layers %s))\n", p.Number, p.Type, p.Center.Sexp("at"), p.Size.Sexp("size"), f(p.Drill), layers); err != nil {
		return err
	}
	return nil
}
