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
		{
			name: "net classes",
			pcb: PCB{
				FormatVersion: 4,
				NetClasses: []NetClass{
					{Name: "Default", Description: "This is the default net class.",
						Clearance: 0.2, TraceWidth: 0.25, Nets: []string{"+5C", "GND"}},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n (general)\n\n (page A4)\n\n (layers\n )\n (setup\n  (zone_45_only no)\n  (uvias_allowed no)\n )\n\n (net_class Default \"This is the default net class.\"\n  (clearance 0.2)\n  (trace_width 0.25)\n  (add_net +5C)\n  (add_net GND)\n ))",
		},
		{
			name: "plot params",
			pcb: PCB{
				EditorSetup: EditorSetup{
					PadDrill: 0.762,
					PlotParams: map[string]PlotParam{
						"usegerberextensions": PlotParam{name: "usegerberextensions", values: []string{"true"}, order: 11},
						"scaleselection":      PlotParam{name: "scaleselection", values: []string{"1"}, order: 10},
						"layerselection":      PlotParam{name: "layerselection", values: []string{"0x010f0_80000001"}},
					},
				},
			},
			expected: "(kicad_pcb (version 0) (host kcgen 0.0.1)\n\n (general)\n\n (page A4)\n\n (layers\n )\n (setup\n  (zone_45_only no)\n  (uvias_allowed no)\n  (pad_drill 0.762)\n  (pcbplotparams\n   (layerselection 0x010f0_80000001)\n   (scaleselection 1)\n   (usegerberextensions true))\n )\n)",
		},
		{
			name: "vias",
			pcb: PCB{
				FormatVersion: 4,
				Vias: []Via{
					{At: XY{X: 100, Y: 32.5}, Layers: []string{"F.Cu", "B.Cu"}, NetIndex: 2},
					{At: XY{X: 10, Y: 32.5}, Layers: []string{"F.Cu", "B.Cu"}, NetIndex: 2},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n (general)\n\n (page A4)\n\n (layers\n )\n (setup\n  (zone_45_only no)\n  (uvias_allowed no)\n )\n)",
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
