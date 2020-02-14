package adv

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/twitchyliquid64/kcgen/pcb"
)

func TestCarveLines(t *testing.T) {
	re := MakeRegion(pcb.XY{X: 10, Y: 10}, pcb.XY{X: 50, Y: 50})
	altRegion1 := MakeRegion(pcb.XY{X: -21.9, Y: 3}, pcb.XY{X: 25, Y: 50})
	tcs := []struct {
		name           string
		overrideRegion *Region
		line           *pcb.Line
		expected       []newLinePair
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
		{
			name:     "cut 1",
			line:     &pcb.Line{End: pcb.XY{X: 60, Y: 20}},
			expected: []newLinePair{{End: pcb.XY{X: 30, Y: 10}}, {Start: pcb.XY{X: 50, Y: 16.666666666666664}, End: pcb.XY{X: 60, Y: 20}}},
		},
		{
			name:     "cut 2",
			line:     &pcb.Line{Start: pcb.XY{X: 60, Y: 20}},
			expected: []newLinePair{{Start: pcb.XY{X: 60, Y: 20}, End: pcb.XY{X: 30, Y: 10}}, {Start: pcb.XY{X: 50, Y: 16.666666666666664}}},
		},
		{
			name:           "within 1",
			overrideRegion: &altRegion1, // XY(-21.9, 3), XY(25, 50)
			line:           &pcb.Line{Start: pcb.XY{X: 22.5, Y: 8.5}, End: pcb.XY{X: -22.5, Y: 8.5}},
			expected:       []newLinePair{{Start: pcb.XY{X: -21.9, Y: 8.5}, End: pcb.XY{X: -22.5, Y: 8.5}}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			region := re
			if tc.overrideRegion != nil {
				region = *tc.overrideRegion
			}
			_, got, err := carveLine(tc.line, region)
			if err != nil {
				t.Fatalf("carveLine(%v) failed: %v", tc.line, err)
			}
			if diff := cmp.Diff(got, tc.expected); diff != "" {
				t.Errorf("incorrect output (+got, -want): \n%s", diff)
			}
		})
	}
}
