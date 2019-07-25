package pcb

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"

	diff "github.com/sergi/go-diff/diffmatchpatch"
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
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n \n)\n",
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
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers\n    (0 F.Cu signal)\n    (31 B.Cu signal)\n  )\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n \n)\n",
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
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n  (net 0 \"\")\n  (net 1 +5C)\n  (net 2 GND)\n\n \n)\n",
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
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n  (net_class Default \"This is the default net class.\"\n    (clearance 0.2)\n    (trace_width 0.25)\n    (add_net +5C)\n    (add_net GND)\n  )\n\n \n)\n",
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
			expected: "(kicad_pcb (version 0) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n    (pad_drill 0.762)\n    (pcbplotparams\n      (layerselection 0x010f0_80000001)\n      (scaleselection 1)\n      (usegerberextensions true))\n  )\n\n \n)\n",
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
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n \n\n  (via (at 100 32.5) (size 0) (drill 0) (layers F.Cu B.Cu) (net 2))\n  (via (at 10 32.5) (size 0) (drill 0) (layers F.Cu B.Cu) (net 2))\n)\n",
		},
		{
			name: "tracks",
			pcb: PCB{
				FormatVersion: 4,
				Tracks: []Track{
					{Start: XY{X: 100, Y: 32.5}, End: XY{X: 10, Y: 32.5}, Layer: "F.Cu", NetIndex: 2},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n \n  (segment (start 100 32.5) (end 10 32.5) (width 0) (layer F.Cu) (net 2))\n)\n",
		},
		{
			name: "lines",
			pcb: PCB{
				FormatVersion: 4,
				Lines: []Line{
					{Start: XY{X: 100, Y: 32.5}, End: XY{X: 10, Y: 32.5}, Layer: "Edge.Cuts", Width: 2},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n \n  (gr_line (start 100 32.5) (end 10 32.5) (layer Edge.Cuts) (width 2))\n\n \n)\n",
		},
		{
			name: "text",
			pcb: PCB{
				FormatVersion: 4,
				Texts: []Text{
					{At: XYZ{X: 100, Y: 32.5}, Text: "Oops", Layer: "F.SilkS", Effects: struct {
						FontSize  XY
						Thickness float64
					}{
						FontSize:  XY{X: 1.5, Y: 1.5},
						Thickness: 0.3,
					}},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n  (gr_text Oops (at 100 32.5) (layer F.SilkS)\n    (effects (font (size 1.5 1.5) (thickness 0.3)))\n  )\n)\n",
		},
		{
			name: "zones",
			pcb: PCB{
				FormatVersion: 4,
				Zones: []Zone{
					{NetNum: 42, Tstamp: "0", Layer: "F.Cu", NetName: "DBUS", MinThickness: 0.254,
						BasePolys: [][]XY{
							[]XY{{X: 11, Y: 22}, {X: 11.1, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}},
						},
						Polys: [][]XY{
							[]XY{{X: 11, Y: 22}, {X: 11.1, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}, {X: 11, Y: 22}},
						},
					},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n \n\n  (zone (net 42) (net_name DBUS) (layer F.Cu) (tstamp 0) (hatch \"\" 0)\n    (connect_pads (clearance 0))\n    (min_thickness 0.254)\n    (fill no (arc_segments 0) (thermal_gap 0) (thermal_bridge_width 0))\n    (polygon\n      (pts\n        (xy 11 22) (xy 11.1 22) (xy 11 22) (xy 11 22) (xy 11 22)\n        (xy 11 22) (xy 11 22)\n      )\n    )\n    (filled_polygon\n      (pts\n        (xy 11 22) (xy 11.1 22) (xy 11 22) (xy 11 22) (xy 11 22)\n        (xy 11 22) (xy 11 22)\n      )\n    )\n  )\n)\n",
		},
		{
			name: "dimensions",
			pcb: PCB{
				FormatVersion: 4,
				Dimensions: []Dimension{
					{
						CurrentMeasurement: 12.446,
						Width:              0.3,
						Layer:              "F.Fab",
						Text: Text{
							At:    XYZ{X: 125.396, Y: 93.853, Z: 90, ZPresent: true},
							Text:  "12.446 mm",
							Layer: "F.Fab", Effects: struct {
								FontSize  XY
								Thickness float64
							}{
								FontSize:  XY{X: 1.5, Y: 1.5},
								Thickness: 0.3,
							}},
						Features: []DimensionFeature{
							{
								Feature: "feature1",
								Points:  []XY{{X: 173.736, Y: 100.076}, {X: 173.736, Y: 106.586}},
							},
							{
								Feature: "feature2",
								Points:  []XY{{X: 132.08, Y: 100.076}, {X: 132.08, Y: 106.586}},
							},
						},
					},
				},
			},
			expected: "(kicad_pcb (version 4) (host kcgen 0.0.1)\n\n  (general)\n\n  (page A4)\n  (layers)\n\n  (setup\n    (zone_45_only no)\n    (uvias_allowed no)\n  )\n\n  (dimension 12.446 (width 0.3) (layer F.Fab)\n    (gr_text \"12.446 mm\" (at 125.396 93.853 90) (layer F.Fab)\n      (effects (font (size 1.5 1.5) (thickness 0.3)))\n    )\n    (feature1 (pts (xy 173.736 100.076) (xy 173.736 106.586)))\n    (feature2 (pts (xy 132.08 100.076) (xy 132.08 106.586)))\n  )\n)\n",
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

func TestDecodeThenSerializeMatches(t *testing.T) {
	tcs := []struct {
		name  string
		fname string
	}{
		{
			name:  "simple",
			fname: "simple_equality.kicad_pcb",
		},
		{
			name:  "zone",
			fname: "zone_equality.kicad_pcb",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			pcb, err := DecodeFile(path.Join("testdata", tc.fname))
			if err != nil {
				t.Fatalf("DecodeFile(%q) failed: %v", tc.fname, err)
			}
			var serialized bytes.Buffer
			if err := pcb.Write(&serialized); err != nil {
				t.Fatalf("Write() failed: %v", err)
			}

			d, err := ioutil.ReadFile(path.Join("testdata", tc.fname))
			if err != nil {
				t.Fatal(err)
			}

			// ioutil.WriteFile("test.kicad_pcb", serialized.Bytes(), 0755)
			if !bytes.Equal(d, serialized.Bytes()) {
				t.Error("outputs differ")
				diffs := diff.New()
				dm := diffs.DiffMain(string(d), serialized.String(), false)
				// t.Log(diffs.DiffPrettyText(dm))
				// t.Log(diffs.DiffToDelta(dm))
				t.Log(diffs.PatchToText(diffs.PatchMake(dm)))
			}
		})
	}
}
