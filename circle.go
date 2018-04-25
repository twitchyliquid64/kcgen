package kcgen

import (
	"fmt"
	"io"
	"math"
)

func pointOnCircle(center Point2D, radius float64, angle float64) *Point2D {
	x := center.X + (radius * math.Cos(angle))
	y := center.Y + (radius * math.Sin(angle))
	return &Point2D{X: x, Y: y}
}

func pointOnCircleDegrees(center Point2D, radius float64, angle float64) *Point2D {
	return pointOnCircle(center, radius, angle*math.Pi/180)
}

// Arc represents a 2d graphical arc in a Footprint.
type Arc struct {
	Layer          Layer
	Center         Point2D
	Start, End     float64
	StepsPerDegree int
	Width, Radius  float64
}

// Render generates output suitable for inclusion in a kicad_mod file.
func (a *Arc) Render(w io.Writer) error {
	width := a.Width
	if width == 0 {
		width = 0.15
	}
	spd := a.StepsPerDegree
	if spd == 0 {
		spd = 20
	}

	diff := a.End - a.Start
	for x := 0; x < spd; x++ {
		lStart := pointOnCircleDegrees(a.Center, a.Radius, a.Start+(diff*(float64(x)/float64(spd))))
		lEnd := pointOnCircleDegrees(a.Center, a.Radius, a.Start+(diff*(float64(x+1)/float64(spd))))
		if _, err := fmt.Fprintf(w, "  (fp_line %s %s (layer %s) (width %s))\n", lStart.Sexp("start"), lEnd.Sexp("end"), a.Layer.Strictname(), f(width)); err != nil {
			return err
		}
	}

	return nil
}
