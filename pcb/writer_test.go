package pcb

import (
	"bytes"
	// "io/ioutil"
	"testing"
)

func TestPCBWrite(t *testing.T) {
	tcs := []struct {
		name     string
		pcb      PCB
		expected string
	}{
		{
			name: "simple",
			pcb: PCB{
				FormatVersion: 4,
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n (general)\n\n (page A4)\n\n (layers\n )\n (setup\n  (zone_45_only no)\n  (uvias_allowed no)\n )\n)",
		},
		{
			name: "layers",
			pcb: PCB{
				FormatVersion: 4,
				Layers: []*Layer{
					{Name: "F.Cu", Type: "signal"},
					{Num: 31, Name: "B.Cu", Type: "signal"},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n (general)\n\n (page A4)\n\n (layers\n  (0 F.Cu signal)\n  (31 B.Cu signal)\n )\n (setup\n  (zone_45_only no)\n  (uvias_allowed no)\n )\n)",
		},
		{
			name: "nets",
			pcb: PCB{
				FormatVersion: 4,
				Nets: map[int]Net{
					0: {Name: ""},
					1: {Name: "+5C"},
					2: {Name: "GND"},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n (general)\n\n (page A4)\n\n (layers\n )\n (setup\n  (zone_45_only no)\n  (uvias_allowed no)\n )\n\n (net 0 \"\")\n (net 1 +5C)\n (net 2 GND)\n)",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			if err := tc.pcb.Write(&b); err != nil {
				t.Fatalf("pcb.Write() failed: %v", err)
			}
			if tc.expected != b.String() {
				t.Error("output mismatch")
				t.Logf("want = %q", tc.expected)
				t.Logf("got  = %q", b.String())
			}
			// ioutil.WriteFile("test.kicad_pcb", b.Bytes(), 0755)
		})
	}
}
