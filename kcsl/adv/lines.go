package adv

import (
	"fmt"
	"math"

	"github.com/twitchyliquid64/kcgen/pcb"
)

type yIntRay struct {
	grad      float64
	intercept float64
}

func (y yIntRay) isVertical() bool {
	return math.IsNaN(y.grad) || math.IsInf(y.grad, 0)
}

func (y1 yIntRay) Y(x float64) float64 {
	return y1.grad*x + y1.intercept
}

func (y1 yIntRay) intersection(y2 yIntRay) pcb.XY {
	d := (y1.grad - y2.grad)
	x := (y2.intercept - y1.intercept) / d
	return pcb.XY{X: x, Y: y1.Y(x)}
}

func makeYIntRay(l *pcb.Line) yIntRay {
	grad := (l.End.Y - l.Start.Y) / (l.End.X - l.Start.X)
	return yIntRay{grad: grad, intercept: l.Start.Y - grad*l.Start.X}
}

func linesIntercept(l1, l2 *pcb.Line) pcb.XY {
	// Special case: identical points.
	switch {
	case l1.Start == l2.Start, l1.Start == l2.End:
		return l1.Start
	case l1.End == l2.End, l1.End == l2.Start:
		return l1.End
	}

	// Compute the intercepts based on the 0 = mx + b form.
	y1, y2 := makeYIntRay(l1), makeYIntRay(l2)
	//fmt.Println(l1, l2, y1, y2, y1.isVertical(), y2.isVertical())

	switch {
	case y1.isVertical() && y2.isVertical():
		l1Start, l1End := math.Min(l1.Start.Y, l1.End.Y), math.Max(l1.Start.Y, l1.End.Y)
		l2Start, l2End := math.Min(l2.Start.Y, l2.End.Y), math.Max(l2.Start.Y, l2.End.Y)
		// fmt.Fprintf(os.Stderr, "%v, %v, %v, %v\n", l1Start, l1End, l2Start, l2End)

		if l1.Start.X == l2.Start.X {
			switch {
			case l1Start >= l2Start && l1Start <= l2End:
				return pcb.XY{X: l1.Start.X, Y: l1Start}
			case l2Start >= l1Start && l2Start <= l1End:
				return pcb.XY{X: l2.Start.X, Y: l2Start}
			}
		}
		return pcb.XY{X: math.NaN(), Y: math.NaN()}
	case !y1.isVertical() && y2.isVertical():
		return pcb.XY{X: l2.Start.X, Y: y1.Y(l2.Start.X)}
	case y1.isVertical() && !y2.isVertical():
		return pcb.XY{X: l1.Start.X, Y: y2.Y(l1.Start.X)}
	}
	return y1.intersection(y2)
}

func linesIntersect(l1, l2 *pcb.Line) bool {
	// Special case: identical points.
	if l1.Start == l2.Start || l1.End == l2.End || l1.Start == l2.End || l1.End == l2.Start {
		return true
	}

	p := linesIntercept(l1, l2)
	// fmt.Printf("y1: %+v\ny2: %+v\np: %+v\n\n", l1, l2, p)
	// Test if the point is between the bounds of the line.
	return MakeRegion(l1.Start, l1.End).Within(p) && MakeRegion(l2.Start, l2.End).Within(p)
}

type newLinePair struct {
	Start, End pcb.XY
}

func carveLine(l *pcb.Line, re Region) (bool, []newLinePair, error) {
	if re.Within(l.Start) && re.Within(l.End) {
		return true, nil, nil
	}
	c := re.Center()

	cutPoints := make([]pcb.XY, 0, 4)
	for _, bound := range []*pcb.Line{
		re.TopBoundary(),
		re.BottomBoundary(),
		re.LeftBoundary(),
		re.RightBoundary()} {

		//fmt.Println(l, bound, linesIntersect(l, bound), linesIntercept(l, bound))
		if linesIntersect(l, bound) {
			cutPoints = append(cutPoints, linesIntercept(l, bound))
		}
	}
	if len(cutPoints) == 0 {
		return false, nil, nil
	}

	switch len(cutPoints) {
	case 1:
		if c.Distance(l.Start) < c.Distance(l.End) {
			return true, []newLinePair{{Start: cutPoints[0], End: l.End}}, nil
		}
		return true, []newLinePair{newLinePair{Start: l.Start, End: cutPoints[0]}}, nil

	case 2:
		p1, p2 := cutPoints[0], cutPoints[1]
		out := []newLinePair{newLinePair{Start: l.Start, End: p1}, newLinePair{Start: p2, End: l.End}}
		// TODO: Maybe the point ordering should be determined by distance to center point?
		if p1.Distance(l.Start) < p1.Distance(l.End) {
			out[0] = newLinePair{Start: p1, End: l.End}
		}
		return true, out, nil

	default:
		return false, nil, fmt.Errorf("cannot handle %d cutpoints", len(cutPoints))
	}
}
