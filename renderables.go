// Package kcgen implements a minimal library for generating kicad footprints.
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

// Polygon represents a 2d polygon in a Footprint.
type Polygon struct {
	Layer  Layer
	Points []Point2D
	Width  float64
}

// Render generates output suitable for inclusion in a kicad_mod file.
func (l *Polygon) Render(w io.Writer) error {
	width := l.Width
	if width == 0 {
		width = 0.15
	}

	if _, err := fmt.Fprint(w, "\n  (fp_poly\n    (pts \n"); err != nil {
		return err
	}
	for i := range l.Points {
		if _, err := fmt.Fprintf(w, "      %s\n", l.Points[i].Sexp("xy")); err != nil {
			return err
		}
	}

	_, err := fmt.Fprintf(w, "    )\n    (layer %s)\n    (width %s)\n  )\n", l.Layer.Strictname(), f(width))
	return err
}

// Text represents a text string.
type Text struct {
	Layer     Layer
	Text      string
	Position  Point2D
	Thickness float64
}

// Render generates output suitable for inclusion in a kicad_mod file.
func (l *Text) Render(w io.Writer) error {
	thickness := l.Thickness
	if thickness == 0 {
		thickness = 0.15
	}
	_, err := fmt.Fprintf(w, `(fp_text user %q %s (layer %s)
    (effects (font (size 1 1) (thickness %s)))
  )`+"\n", l.Text, l.Position.Sexp("at"), l.Layer.Strictname(), f(thickness))
	return err
}
