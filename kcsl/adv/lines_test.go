package adv

import (
	"math"
	"testing"

	"github.com/twitchyliquid64/kcgen/pcb"
)

func TestMakeLine(t *testing.T) {
	tcs := []struct {
		name string
		in   pcb.Line
		out  yIntRay
	}{
		{
			name: "basic",
			in: pcb.Line{
				Start: pcb.XY{X: 0, Y: 0},
				End:   pcb.XY{X: 5, Y: 5},
			},
			out: yIntRay{
				grad: 1,
			},
		},
		{
			name: "neg",
			in: pcb.Line{
				Start: pcb.XY{X: 0, Y: -2},
				End:   pcb.XY{X: -5, Y: -2},
			},
			out: yIntRay{
				grad:      0,
				intercept: -2,
			},
		},
		{
			name: "horizontal",
			in: pcb.Line{
				Start: pcb.XY{X: 0, Y: 0},
				End:   pcb.XY{X: 5, Y: 0},
			},
			out: yIntRay{
				grad: 0,
			},
		},
		{
			name: "vertical",
			in: pcb.Line{
				Start: pcb.XY{X: 0, Y: 0},
				End:   pcb.XY{X: 0, Y: 5},
			},
			out: yIntRay{
				grad:      math.Inf(+0),
				intercept: math.NaN(),
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := makeYIntRay(&tc.in); got != tc.out {
				if math.IsNaN(got.intercept) && math.IsNaN(tc.out.intercept) && math.IsInf(got.grad, 1) && math.IsInf(tc.out.grad, 1) {
					// Its fine
				} else {
					t.Errorf("makeYLine(%+v) = %+v, want %+v", tc.in, got, tc.out)
				}
			}
		})
	}
}

func TestLinesIntersect(t *testing.T) {
	tcs := []struct {
		name       string
		l1, l2     pcb.Line
		intersects bool
	}{
		{
			name:       "zero",
			l1:         pcb.Line{Start: pcb.XY{}, End: pcb.XY{}},
			l2:         pcb.Line{Start: pcb.XY{}, End: pcb.XY{}},
			intersects: true,
		},
		{
			name:       "inf",
			l1:         pcb.Line{Start: pcb.XY{Y: -1}, End: pcb.XY{Y: 1}},
			l2:         pcb.Line{Start: pcb.XY{}, End: pcb.XY{X: 1}},
			intersects: true,
		},
		{
			name:       "inf 2",
			l1:         pcb.Line{Start: pcb.XY{Y: -1}, End: pcb.XY{Y: 1}},
			l2:         pcb.Line{Start: pcb.XY{X: 1}, End: pcb.XY{X: 2}},
			intersects: false,
		},
		{
			name:       "vert 1",
			l1:         pcb.Line{Start: pcb.XY{Y: 2}, End: pcb.XY{Y: 4}},
			l2:         pcb.Line{Start: pcb.XY{Y: 3}, End: pcb.XY{Y: 5}},
			intersects: true,
		},
		{
			name:       "vert 2",
			l1:         pcb.Line{Start: pcb.XY{Y: 3}, End: pcb.XY{Y: 5}},
			l2:         pcb.Line{Start: pcb.XY{Y: 2}, End: pcb.XY{Y: 4}},
			intersects: true,
		},
		{
			name:       "vert 3",
			l1:         pcb.Line{Start: pcb.XY{Y: 4}, End: pcb.XY{Y: 5}},
			l2:         pcb.Line{Start: pcb.XY{Y: 2}, End: pcb.XY{Y: 4}},
			intersects: true,
		},
		{
			name:       "vert 4",
			l1:         pcb.Line{Start: pcb.XY{Y: 5}, End: pcb.XY{Y: 4}},
			l2:         pcb.Line{Start: pcb.XY{Y: 2}, End: pcb.XY{Y: 4}},
			intersects: true,
		},
		{
			name:       "vert 5",
			l1:         pcb.Line{Start: pcb.XY{Y: 4.1}, End: pcb.XY{Y: 5}},
			l2:         pcb.Line{Start: pcb.XY{Y: 2}, End: pcb.XY{Y: 4}},
			intersects: false,
		},
		{
			name:       "vert 6",
			l1:         pcb.Line{Start: pcb.XY{Y: 4}, End: pcb.XY{Y: 2}},
			l2:         pcb.Line{Start: pcb.XY{Y: 3}, End: pcb.XY{Y: 5}},
			intersects: true,
		},
		{
			name:       "vert 7",
			l1:         pcb.Line{Start: pcb.XY{Y: 5}, End: pcb.XY{Y: 3}},
			l2:         pcb.Line{Start: pcb.XY{Y: 2}, End: pcb.XY{Y: 4}},
			intersects: true,
		},
		{
			name:       "cross 1",
			l1:         pcb.Line{Start: pcb.XY{}, End: pcb.XY{X: 3}},
			l2:         pcb.Line{Start: pcb.XY{X: 2}, End: pcb.XY{Y: -4}},
			intersects: true,
		},
		{
			name:       "cross 2",
			l1:         pcb.Line{Start: pcb.XY{Y: -1, X: 1.8}, End: pcb.XY{X: 3}},
			l2:         pcb.Line{Start: pcb.XY{X: 2}, End: pcb.XY{Y: -4}},
			intersects: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := linesIntersect(&tc.l1, &tc.l2); got != tc.intersects {
				t.Errorf("linesIntersect(%+v, %+v) = %v, want %v", tc.l1, tc.l2, got, tc.intersects)
			}
		})
	}
}
