package pcb

import (
	"github.com/nsf/sexp"
)

// Text represents some text to be rendered.
type Text struct {
	Text  string
	Layer string
	At    XY

	Effects struct {
		FontSize  XY
		Thickness float64
	}

	order int
}

// Line represents a graphical line.
type Line struct {
	Start XY
	End   XY
	Layer string
	Width float64

	order int
}

func parseGRText(n sexp.Helper, ordering int) (Text, error) {
	t := Text{
		Text:  n.Child(1).MustString(),
		order: ordering,
	}
	for x := 2; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "at":
			t.At.X = c.Child(1).MustFloat64()
			t.At.Y = c.Child(2).MustFloat64()
		case "layer":
			t.Layer = c.Child(1).MustString()
		case "effects":
			for y := 1; y < c.MustNode().NumChildren(); y++ {
				c := c.Child(y)
				switch c.Child(0).MustString() {
				case "font":
					for z := 1; z < c.MustNode().NumChildren(); z++ {
						c := c.Child(z)
						switch c.Child(0).MustString() {
						case "size":
							t.Effects.FontSize.X = c.Child(1).MustFloat64()
							t.Effects.FontSize.Y = c.Child(2).MustFloat64()
						case "thickness":
							t.Effects.Thickness = c.Child(1).MustFloat64()
						}
					}
				}
			}
		}
	}
	return t, nil
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
