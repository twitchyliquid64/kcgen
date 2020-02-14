package adv

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/twitchyliquid64/kcgen/pcb"
)

func TestCarveLines(t *testing.T) {
	re := makeRegion(pcb.XY{X: 10, Y: 10}, pcb.XY{X: 50, Y: 50})
	tcs := []struct {
		name     string
		line     *pcb.Line
		expected []newLinePair
	}{
		{
			name:     "no intersect 1",
			line:     &pcb.Line{Start: pcb.XY{X: 10, Y: 100}, End: pcb.XY{X: 50, Y: 110}},
			expected: nil,
		},
		{
			name:     "no intersect 2",
			line:     &pcb.Line{Start: pcb.XY{X: 10, Y: 100}, End: pcb.XY{X: 50, Y: 100}},
			expected: nil,
		},
		{
			name:     "top only 1",
			line:     &pcb.Line{Start: pcb.XY{X: 0, Y: 0}, End: pcb.XY{X: 30, Y: 10}},
			expected: []newLinePair{{End: pcb.XY{X: 30, Y: 10}}},
		},
		{
			name:     "top only 2",
			line:     &pcb.Line{End: pcb.XY{X: 0, Y: 0}, Start: pcb.XY{X: 30, Y: 10}},
			expected: []newLinePair{{Start: pcb.XY{X: 30, Y: 10}}},
		},
		{
			name:     "left only 1",
			line:     &pcb.Line{Start: pcb.XY{X: -50, Y: 30}, End: pcb.XY{X: 20, Y: 30}},
			expected: []newLinePair{{Start: pcb.XY{X: -50, Y: 30}, End: pcb.XY{X: 10, Y: 30}}},
		},
		{
			name:     "left only 2",
			line:     &pcb.Line{End: pcb.XY{X: -50, Y: 30}, Start: pcb.XY{X: 20, Y: 30}},
			expected: []newLinePair{{End: pcb.XY{X: -50, Y: 30}, Start: pcb.XY{X: 10, Y: 30}}},
		},
		{
			name:     "right only 1",
			line:     &pcb.Line{Start: pcb.XY{X: 60, Y: 30}, End: pcb.XY{X: 40, Y: 30}},
			expected: []newLinePair{{Start: pcb.XY{X: 60, Y: 30}, End: pcb.XY{X: 50, Y: 30}}},
		},
		{
			name:     "right only 2",
			line:     &pcb.Line{End: pcb.XY{X: 70, Y: 30}, Start: pcb.XY{X: 40, Y: 30}},
			expected: []newLinePair{{End: pcb.XY{X: 70, Y: 30}, Start: pcb.XY{X: 50, Y: 30}}},
		},
		{
			name:     "bottom only 1",
			line:     &pcb.Line{Start: pcb.XY{X: 30, Y: 80}, End: pcb.XY{X: 30, Y: 40}},
			expected: []newLinePair{{Start: pcb.XY{X: 30, Y: 80}, End: pcb.XY{X: 30, Y: 50}}},
		},
		{
			name:     "bottom only 2",
			line:     &pcb.Line{End: pcb.XY{X: 30, Y: 80}, Start: pcb.XY{X: 30, Y: 40}},
			expected: []newLinePair{{End: pcb.XY{X: 30, Y: 80}, Start: pcb.XY{X: 30, Y: 50}}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := carveLine(tc.line, re)
			if err != nil {
				t.Fatalf("carveLine(%v) failed: %v", tc.line, err)
			}
			if diff := cmp.Diff(got, tc.expected); diff != "" {
				t.Errorf("incorrect output (+got, -want): \n%s", diff)
			}
		})
	}
}
