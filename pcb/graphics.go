package pcb

import (
	"github.com/nsf/sexp"
)

// Line represents a graphical line.
type Line struct {
	Start XY
	End   XY
	Layer string
	Width float64

	order int
}

func parseGRLine(n sexp.Helper, ordering int) (Line, error) {
	l := Line{order: ordering}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "start":
			l.Start.X = c.Child(1).MustFloat64()
			l.Start.Y = c.Child(2).MustFloat64()
		case "end":
			l.End.X = c.Child(1).MustFloat64()
			l.End.Y = c.Child(2).MustFloat64()
		case "width":
			l.Width = c.Child(1).MustFloat64()
		case "layer":
			l.Layer = c.Child(1).MustString()
		}
	}
	return l, nil
}
