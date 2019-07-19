// Package pcb understands KiCad PCB format.
package pcb

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/nsf/sexp"
)

// Layer describes the attributes of a layer.
type Layer struct {
	Name string
	Type string

	order int
}

// Net represents a netlist.
type Net struct {
	Name string

	order int
}

// NetClass represents a net class.
type NetClass struct {
	Name        string
	Description string

	Clearance    float64
	TraceWidth   float64
	ViaDiameter  float64
	ViaDrill     float64
	UViaDiameter float64
	UViaDrill    float64

	// Nets contains the names of nets which are part of this class.
	Nets []string

	order int
}

// Via represents a via.
type Via struct {
	X, Y     float64
	Size     float64
	Drill    float64
	Layers   []string
	NetIndex int

	order int
}

// Zone represents a zone.
type Zone struct {
	NetNum  int
	NetName string
	Layer   string

	Hatch struct {
		Mode string
		Size float64
	}

	ConnectPads struct {
		Clearance float64
	}

	Fill struct {
		Enabled            bool
		Segments           int
		ThermalGap         float64
		ThermalBridgeWidth float64
	}

	MinThickness float64

	Polys     [][][]float64
	BasePolys [][][]float64

	order int
}

// Track represents a PCB track.
type Track struct {
	StartX, StartY float64
	EndX, EndY     float64
	Width          float64
	Layer          string
	NetIndex       int

	order int
}

// PCB represents the parsed contents of a kicad_pcb file.
type PCB struct {
	FormatVersion int
	CreatedBy     struct {
		Tool    string
		Version string
	}

	EditorSetup EditorSetup

	LayersByName map[string]*Layer
	Layers       map[int]*Layer
	Tracks       []Track
	Vias         []Via
	Nets         map[int]Net
	NetClasses   []NetClass
	Zones        []Zone
}

// EditorSetup describes how the editor should be configured when
// editing this PCB.
type EditorSetup struct {
	LastTraceWidth  float64
	UserTraceWidths []float64
	TraceClearance  float64
	ZoneClearance   float64
	Zone45Only      bool
	TraceMin        float64

	TextWidth    float64
	TextSize     []float64
	SegmentWidth float64
	EdgeWidth    float64

	ViaSize      float64
	ViaMinSize   float64
	ViaDrill     float64
	ViaMinDrill  float64
	UViaSize     float64
	UViaMinSize  float64
	UViaDrill    float64
	UViaMinDrill float64
	AllowUVias   bool

	ModEdgeWidth       float64
	ModTextSize        []float64
	ModTextWidth       float64
	PadSize            []float64
	PadDrill           float64
	PadToMaskClearance float64

	Unrecognised map[string]sexp.Helper
	order        int
}

// DecodeFile reads a .kicad_pcb file at fpath, returning a parsed representation.
func DecodeFile(fpath string) (*PCB, error) {
	f, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	ast, err := sexp.Parse(strings.NewReader(string(f)), nil)
	if err != nil {
		return nil, err
	}

	if !ast.IsList() {
		return nil, errors.New("invalid format: expected s-expression list at top level")
	}
	if ast.NumChildren() != 1 {
		return nil, errors.New("invalid format: top level list of size 1")
	}
	mainAST, _ := ast.Nth(0)
	if !mainAST.IsList() {
		return nil, errors.New("invalid format: expected s-expression list at 1st level")
	}

	if mainAST.NumChildren() < 5 {
		return nil, errors.New("invalid format: expected at least 5 nodes in main expression")
	}
	if mainAST.Children.Value != "kicad_pcb" {
		return nil, errors.New("invalid format: missing leading element kicad_pcb")
	}

	pcb := &PCB{Layers: map[int]*Layer{}, LayersByName: map[string]*Layer{}, Nets: map[int]Net{}}
	var ordering int

	for i := 1; i < mainAST.NumChildren(); i++ {
		n := sexp.Help(mainAST).Child(i)
		if n.IsList() && n.Child(1).IsValid() {
			switch n.Child(0).MustString() {
			case "version":
				pcb.FormatVersion, err = n.Child(1).Int()
				if err != nil {
					return nil, errors.New("invalid format: version value must be an int")
				}
			case "host":
				pcb.CreatedBy.Tool, err = n.Child(1).String()
				if err != nil {
					return nil, errors.New("invalid format: host value[1] must be a string")
				}
				pcb.CreatedBy.Version, err = n.Child(2).String()
				if err != nil {
					return nil, errors.New("invalid format: host value[2] must be a string")
				}
			case "setup":
				s, err := parseSetup(n, ordering)
				if err != nil {
					return nil, err
				}
				pcb.EditorSetup = *s
			case "layers":
				for x := 1; x < n.MustNode().NumChildren(); x++ {
					c := n.Child(x)
					num, err2 := c.Child(0).Int()
					if err2 != nil {
						return nil, err
					}
					pcb.Layers[num] = &Layer{
						Name:  c.Child(1).MustString(),
						Type:  c.Child(2).MustString(),
						order: ordering,
					}
					pcb.LayersByName[c.Child(1).MustString()] = pcb.Layers[num]
					ordering++
				}
			case "net":
				num, err2 := n.Child(1).Int()
				if err2 != nil {
					return nil, err
				}
				pcb.Nets[num] = Net{Name: n.Child(2).MustString(), order: ordering}

			case "segment":
				t, err := parseSegment(n, ordering)
				if err != nil {
					return nil, err
				}
				pcb.Tracks = append(pcb.Tracks, t)

			case "via":
				v, err := parseVia(n, ordering)
				if err != nil {
					return nil, err
				}
				pcb.Vias = append(pcb.Vias, v)

			case "zone":
				z, err := parseZone(n, ordering)
				if err != nil {
					return nil, err
				}
				pcb.Zones = append(pcb.Zones, *z)

			case "net_class":
				c, err := parseNetClass(n, ordering)
				if err != nil {
					return nil, err
				}
				pcb.NetClasses = append(pcb.NetClasses, *c)
			}
		}
		ordering++
	}

	return pcb, nil
}

func parseSegment(n sexp.Helper, ordering int) (Track, error) {
	t := Track{order: ordering}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "width":
			t.Width = c.Child(1).MustFloat64()
		case "net":
			t.NetIndex = c.Child(1).MustInt()
		case "layer":
			t.Layer = c.Child(1).MustString()
		case "start":
			t.StartX = c.Child(1).MustFloat64()
			t.StartY = c.Child(2).MustFloat64()
		case "end":
			t.EndX = c.Child(1).MustFloat64()
			t.EndY = c.Child(2).MustFloat64()
		}
	}
	return t, nil
}

func parseVia(n sexp.Helper, ordering int) (Via, error) {
	v := Via{order: ordering}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "size":
			v.Size = c.Child(1).MustFloat64()
		case "drill":
			v.Drill = c.Child(1).MustFloat64()
		case "net":
			v.NetIndex = c.Child(1).MustInt()
		case "at":
			v.X = c.Child(1).MustFloat64()
			v.Y = c.Child(2).MustFloat64()
		case "layers":
			for j := 1; j < c.MustNode().NumChildren(); j++ {
				v.Layers = append(v.Layers, c.Child(j).MustString())
			}
		}
	}
	return v, nil
}

func parseZone(n sexp.Helper, ordering int) (*Zone, error) {
	z := Zone{order: ordering}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "net":
			z.NetNum = c.Child(1).MustInt()
		case "net_name":
			z.NetName = c.Child(1).MustString()
		case "layer":
			z.Layer = c.Child(1).MustString()

		case "hatch":
			z.Hatch.Mode = c.Child(1).MustString()
			z.Hatch.Size = c.Child(2).MustFloat64()
		case "min_thickness":
			z.MinThickness = c.Child(1).MustFloat64()

		case "connect_pads":
			for y := 1; y < c.MustNode().NumChildren(); y++ {
				c2 := c.Child(y)
				switch c2.Child(0).MustString() {
				case "clearance":
					z.ConnectPads.Clearance = c2.Child(1).MustFloat64()
				}
			}
		case "fill":
			z.Fill.Enabled = c.Child(1).MustString() == "yes"
			for y := 2; y < c.MustNode().NumChildren(); y++ {
				c2 := c.Child(y)
				switch c2.Child(0).MustString() {
				case "arc_segments":
					z.Fill.Segments = c2.Child(1).MustInt()
				case "thermal_gap":
					z.Fill.ThermalGap = c2.Child(1).MustFloat64()
				case "thermal_bridge_width":
					z.Fill.ThermalBridgeWidth = c2.Child(1).MustFloat64()
				}
			}

		case "polygon":
			var points [][]float64
			for y := 1; y < c.Child(1).MustNode().NumChildren(); y++ {
				pt := c.Child(1).Child(y)
				ptType, err2 := pt.Child(0).String()
				if err2 != nil || ptType != "xy" {
					return nil, errors.New("zone.polygon point is not xy point")
				}
				points = append(points, []float64{pt.Child(1).MustFloat64(), pt.Child(2).MustFloat64()})
			}
			z.BasePolys = append(z.BasePolys, points)

		case "filled_polygon":
			var points [][]float64
			for y := 1; y < c.Child(1).MustNode().NumChildren(); y++ {
				pt := c.Child(1).Child(y)
				ptType, err2 := pt.Child(0).String()
				if err2 != nil || ptType != "xy" {
					return nil, errors.New("zone.filled_polygon point is not xy point")
				}
				points = append(points, []float64{pt.Child(1).MustFloat64(), pt.Child(2).MustFloat64()})
			}
			z.Polys = append(z.Polys, points)
		}
	}
	return &z, nil
}

func parseNetClass(n sexp.Helper, ordering int) (*NetClass, error) {
	nc := NetClass{
		order:       ordering,
		Name:        n.Child(1).MustString(),
		Description: n.Child(2).MustString(),
	}
	for x := 3; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "clearance":
			nc.Clearance = c.Child(1).MustFloat64()
		case "trace_width":
			nc.TraceWidth = c.Child(1).MustFloat64()
		case "via_dia":
			nc.ViaDiameter = c.Child(1).MustFloat64()
		case "via_drill":
			nc.ViaDrill = c.Child(1).MustFloat64()
		case "uvia_dia":
			nc.UViaDiameter = c.Child(1).MustFloat64()
		case "uvia_drill":
			nc.UViaDrill = c.Child(1).MustFloat64()
		case "add_net":
			nc.Nets = append(nc.Nets, c.Child(1).MustString())
		}
	}
	return &nc, nil
}

func parseSetup(n sexp.Helper, ordering int) (*EditorSetup, error) {
	e := EditorSetup{
		order:        ordering,
		Unrecognised: map[string]sexp.Helper{},
	}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "last_trace_width":
			e.LastTraceWidth = c.Child(1).MustFloat64()
		case "user_trace_width":
			e.UserTraceWidths = append(e.UserTraceWidths, c.Child(1).MustFloat64())
		case "trace_clearance":
			e.TraceClearance = c.Child(1).MustFloat64()
		case "zone_clearance":
			e.ZoneClearance = c.Child(1).MustFloat64()
		case "zone_45_only":
			e.Zone45Only = c.Child(1).MustString() == "yes"
		case "trace_min":
			e.TraceMin = c.Child(1).MustFloat64()
		case "segment_width":
			e.SegmentWidth = c.Child(1).MustFloat64()
		case "edge_width":
			e.EdgeWidth = c.Child(1).MustFloat64()

		case "via_size":
			e.ViaSize = c.Child(1).MustFloat64()
		case "via_min_size":
			e.ViaMinSize = c.Child(1).MustFloat64()
		case "via_min_drill":
			e.ViaMinDrill = c.Child(1).MustFloat64()
		case "via_drill":
			e.ViaDrill = c.Child(1).MustFloat64()
		case "uvia_size":
			e.UViaSize = c.Child(1).MustFloat64()
		case "uvia_min_size":
			e.UViaMinSize = c.Child(1).MustFloat64()
		case "uvia_min_drill":
			e.UViaMinDrill = c.Child(1).MustFloat64()
		case "uvia_drill":
			e.UViaDrill = c.Child(1).MustFloat64()
		case "uvias_allowed":
			e.AllowUVias = c.Child(1).MustString() == "yes"

		case "pcb_text_width":
			e.TextWidth = c.Child(1).MustFloat64()
		case "pcb_text_size":
			for y := 1; y < c.MustNode().NumChildren(); y++ {
				e.TextSize = append(e.TextSize, c.Child(y).MustFloat64())
			}

		case "mod_edge_width":
			e.ModEdgeWidth = c.Child(1).MustFloat64()
		case "mod_text_size":
			for y := 1; y < c.MustNode().NumChildren(); y++ {
				e.ModTextSize = append(e.ModTextSize, c.Child(y).MustFloat64())
			}
		case "mod_text_width":
			e.ModTextWidth = c.Child(1).MustFloat64()

		case "pad_size":
			for y := 1; y < c.MustNode().NumChildren(); y++ {
				e.PadSize = append(e.PadSize, c.Child(y).MustFloat64())
			}
		case "pad_drill":
			e.PadDrill = c.Child(1).MustFloat64()
		case "pad_to_mask_clearance":
			e.PadToMaskClearance = c.Child(1).MustFloat64()

		default:
			e.Unrecognised[c.Child(0).MustString()] = c
		}
	}
	return &e, nil
}
