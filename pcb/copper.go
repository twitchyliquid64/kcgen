package pcb

import (
	"errors"

	"github.com/nsf/sexp"
)

// XY represents a point in space.
type XY struct {
	X, Y float64
}

// Via represents a via.
type Via struct {
	At       XY
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

	Polys     [][]XY
	BasePolys [][]XY

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
			v.At.X = c.Child(1).MustFloat64()
			v.At.Y = c.Child(2).MustFloat64()
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
			var points []XY
			for y := 1; y < c.Child(1).MustNode().NumChildren(); y++ {
				pt := c.Child(1).Child(y)
				ptType, err2 := pt.Child(0).String()
				if err2 != nil || ptType != "xy" {
					return nil, errors.New("zone.polygon point is not xy point")
				}
				points = append(points, XY{X: pt.Child(1).MustFloat64(), Y: pt.Child(2).MustFloat64()})
			}
			z.BasePolys = append(z.BasePolys, points)

		case "filled_polygon":
			var points []XY
			for y := 1; y < c.Child(1).MustNode().NumChildren(); y++ {
				pt := c.Child(1).Child(y)
				ptType, err2 := pt.Child(0).String()
				if err2 != nil || ptType != "xy" {
					return nil, errors.New("zone.filled_polygon point is not xy point")
				}
				points = append(points, XY{X: pt.Child(1).MustFloat64(), Y: pt.Child(2).MustFloat64()})
			}
			z.Polys = append(z.Polys, points)
		}
	}
	return &z, nil
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
