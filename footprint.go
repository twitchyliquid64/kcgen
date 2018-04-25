package kcgen

import (
	"fmt"
	"io"
	"time"
)

// Footprint represents the mechanical layout of a part/thing in Kicad.
type Footprint struct {
	ModName    string
	ReferenceY float64

	Renderables []renderable
	Tname       string
}

// Point2D represents a point in 2-dimensional space.
type Point2D struct {
	X, Y float64
}

// Sexp returns a S-Expression tuple string, starting with the given name
func (p *Point2D) Sexp(name string) string {
	return fmt.Sprintf("(%s %s %s)", name, f(p.X), f(p.Y))
}

// Render generates output suitable for inclusion in a kicad_mod file.
func (f *Footprint) Render(w io.Writer) error {
	tname := f.Tname
	if tname == "" {
		tname = fmt.Sprintf("%8X", time.Now().Unix())
	}
	if _, err := fmt.Fprintf(w, "(module %s (layer F.Cu) (tedit %s)\n", f.ModName, tname); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, `  (fp_text reference REF** (at 0 %f) (layer F.SilkS)
      (effects (font (size 1 1) (thickness 0.15)))
    )`+"\n", f.ReferenceY); err != nil {
		return err
	}

	for i := range f.Renderables {
		if err := f.Renderables[i].Render(w); err != nil {
			return fmt.Errorf("renderable %d failed to render: %v", i, err)
		}
	}

	if _, err := fmt.Fprint(w, ")\n"); err != nil {
		return err
	}
	return nil
}

// Add inserts a renderable component into the design of the footprint.
func (f *Footprint) Add(r renderable) error {
	f.Renderables = append(f.Renderables, r)
	return nil //TODO: return errors for known cases of badly-formed renderables.
}
