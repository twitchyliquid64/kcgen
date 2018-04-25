package kcgen

import (
	"fmt"
	"io"
)

type renderable interface {
	Render(w io.Writer) error
}

// Line represents a 2d graphical line in a Footprint.
type Line struct {
	Layer      Layer
	Start, End Point2D
	Width      float64
}

// Render generates output suitable for inclusion in a kicad_mod file.
func (l *Line) Render(w io.Writer) error {
	width := l.Width
	if width == 0 {
		width = 0.15
	}
	_, err := fmt.Fprintf(w, "  (fp_line %s %s (layer %s) (width %s))\n", l.Start.Sexp("start"), l.End.Sexp("end"), l.Layer.Strictname(), f(width))
	return err
}

func f(f float64) string {
	t := fmt.Sprintf("%f", f)
	if t[len(t)-1] != '0' {
		return t
	}

	for i := len(t) - 1; i >= 0; i-- {
		if t[i] != '0' {
			if t[i] == '.' {
				return t[:i]
			}
			return t[:i+1]
		}
	}
	return t
}
