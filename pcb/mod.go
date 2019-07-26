package pcb

import (
	"fmt"
	"strings"

	"github.com/nsf/sexp"
	"github.com/twitchyliquid64/kcgen/swriter"
)

// Module describes a KiCad module.
type Module struct {
	Name      string
	Placement ModPlacement

	Layer string

	SolderMaskMargin  float64
	SolderPasteMargin float64
	SolderPasteRatio  float64
	Clearance         float64

	Tedit  string
	Tstamp string
	Path   string

	Description string
	Tags        []string
	Attrs       []string
	order       int

	Graphics []ModGraphic
	Pads     []Pad
	Model    *ModModel
}

// ModPlacement describes the positioning of a module on a PCB.
type ModPlacement struct {
	At XYZ
}

// ModGraphic represents a graphical feature in a module.
type ModGraphic struct {
	Ident      string
	Renderable modDrawable
}

type modDrawable interface {
	write(sw *swriter.SExpWriter) error
}

// ModPolygon represents a polygon drawn in a module.
type ModPolygon struct {
	At     XY
	Points []XY
	Layer  string
	Width  float64
}

// ModText represents text drawn in a module.
type ModText struct {
	Kind ModTextKind
	Text string
	At   XYZ

	Layer   string
	Effects TextEffects
}

// ModTextKind describes the type of text drawing.
type ModTextKind uint8

func (k ModTextKind) String() string {
	switch k {
	case RefText:
		return "reference"
	case UserText:
		return "user"
	case ValueText:
		return "value"
	}
	return "?????"
}

// Valid ModTextKind values.
const (
	RefText ModTextKind = iota
	ValueText
	UserText
)

// ModLine represents a line drawn in a module.
type ModLine struct {
	Start XY
	End   XY

	Layer string
	Width float64
}

// ModCircle represents a circle drawn in a module.
type ModCircle struct {
	Center XY
	End    XY
	Layer  string
	Width  float64
}

// ModArc represents an arc drawn in a module.
type ModArc struct {
	Start XY
	End   XY
	Layer string
	Angle float64
	Width float64
}

// ModModel describes configuration for rendering a 3d model of the part.
type ModModel struct {
	Path   string
	At     XYZ
	Scale  XYZ
	Rotate XYZ
}

type PadSurface uint8

type PadShape uint8

// Pad constants
const (
	ShapeInvalid PadShape = iota
	ShapeRect
	ShapeOval
	ShapeCircle
	ShapeTrapezoid
	ShapeRoundRect
	ShapeChamferedRect
	ShapeCustom
	ShapeDrillOblong

	SurfaceInvalid PadSurface = iota
	SurfaceSMD
	SurfaceTH
	SurfaceNPTH
	SurfaceConnect
)

// Pad represents a copper pad.
type Pad struct {
	Ident   string
	NetNum  int
	NetName string

	At     XYZ
	Size   XY
	Layers []string

	RectDelta XY

	DrillOffset XY
	DrillSize   XY
	DrillShape  PadShape

	DieLength              float64
	ZoneConnect            int
	ThermalWidth           float64
	ThermalGap             float64
	RoundRectRRatio        float64
	ChamferRatio           float64
	SolderMaskMargin       float64
	SolderPasteMargin      float64
	SolderPasteMarginRatio float64
	Clearance              float64

	Surface PadSurface
	Shape   PadShape
}

func parseModule(n sexp.Helper, ordering int) (*Module, error) {
	m := Module{
		Name:  n.Child(1).MustString(),
		order: ordering,
	}
	for x := 2; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "tedit":
			m.Tedit = c.Child(1).MustString()
		case "tstamp":
			m.Tstamp = c.Child(1).MustString()
		case "layer":
			m.Layer = c.Child(1).MustString()
		case "descr":
			m.Description = c.Child(1).MustString()
		case "path":
			m.Path = c.Child(1).MustString()

		case "attr":
			m.Attrs = strings.Split(c.Child(1).MustString(), " ")
		case "tags":
			m.Tags = strings.Split(c.Child(1).MustString(), " ")

		case "at":
			m.Placement.At.X = c.Child(1).MustFloat64()
			m.Placement.At.Y = c.Child(2).MustFloat64()
			if c.MustNode().NumChildren() >= 4 {
				m.Placement.At.Z = c.Child(3).MustFloat64()
				m.Placement.At.ZPresent = true
			}

		case "clearance":
			m.Clearance = c.Child(1).MustFloat64()
		case "solder_paste_margin":
			m.SolderPasteMargin = c.Child(1).MustFloat64()
		case "solder_mask_margin":
			m.SolderMaskMargin = c.Child(1).MustFloat64()
		case "solder_paste_ratio":
			m.SolderPasteRatio = c.Child(1).MustFloat64()

		case "fp_text":
			t, err := parseModText(c)
			if err != nil {
				return nil, err
			}
			m.Graphics = append(m.Graphics, ModGraphic{
				Ident:      c.Child(0).MustString(),
				Renderable: t,
			})

		case "fp_line":
			l, err := parseModLine(c)
			if err != nil {
				return nil, err
			}
			m.Graphics = append(m.Graphics, ModGraphic{
				Ident:      c.Child(0).MustString(),
				Renderable: l,
			})

		case "fp_arc":
			a, err := parseModArc(c)
			if err != nil {
				return nil, err
			}
			m.Graphics = append(m.Graphics, ModGraphic{
				Ident:      c.Child(0).MustString(),
				Renderable: a,
			})

		case "fp_circle":
			a, err := parseModCircle(c)
			if err != nil {
				return nil, err
			}
			m.Graphics = append(m.Graphics, ModGraphic{
				Ident:      c.Child(0).MustString(),
				Renderable: a,
			})

		case "fp_poly":
			a, err := parseModPolygon(c)
			if err != nil {
				return nil, err
			}
			m.Graphics = append(m.Graphics, ModGraphic{
				Ident:      c.Child(0).MustString(),
				Renderable: a,
			})

			// TODO: curve

		case "pad":
			pad, err := parseModPad(c)
			if err != nil {
				return nil, err
			}
			m.Pads = append(m.Pads, *pad)

		case "model":
			model, err := parseModModel(c)
			if err != nil {
				return nil, err
			}
			m.Model = model
		}
	}
	return &m, nil
}

func parseModText(n sexp.Helper) (*ModText, error) {
	t := ModText{
		Text: n.Child(2).MustString(),
	}

	switch n.Child(1).MustString() {
	case "reference":
		t.Kind = RefText
	case "value":
		t.Kind = ValueText
	case "user":
		t.Kind = UserText
	default:
		return nil, fmt.Errorf("unknown fp_text type: %v", n.Child(1).MustString())
	}

	for x := 3; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "at":
			t.At.X = c.Child(1).MustFloat64()
			t.At.Y = c.Child(2).MustFloat64()
			if c.MustNode().NumChildren() >= 4 {
				t.At.Z = c.Child(3).MustFloat64()
				t.At.ZPresent = true
			}
		case "layer":
			t.Layer = c.Child(1).MustString()
		case "effects":
			effects, err := parseTextEffects(c)
			if err != nil {
				return nil, err
			}
			t.Effects = effects
		}
	}

	return &t, nil
}

func parseModLine(n sexp.Helper) (*ModLine, error) {
	l := ModLine{}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "start":
			l.Start.X = c.Child(1).MustFloat64()
			l.Start.Y = c.Child(2).MustFloat64()
		case "end":
			l.End.X = c.Child(1).MustFloat64()
			l.End.Y = c.Child(2).MustFloat64()
		case "layer":
			l.Layer = c.Child(1).MustString()
		case "width":
			l.Width = c.Child(1).MustFloat64()
		}
	}

	return &l, nil
}

func parseModPolygon(n sexp.Helper) (*ModPolygon, error) {
	p := ModPolygon{}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "at":
			p.At.X = c.Child(1).MustFloat64()
			p.At.Y = c.Child(2).MustFloat64()
		case "pts":
			for j := 1; j < c.MustNode().NumChildren(); j++ {
				c := c.Child(j)
				if marker := c.Child(0).MustString(); marker != "xy" {
					return nil, fmt.Errorf("expected 'xy', got %q", marker)
				}
				p.Points = append(p.Points, XY{X: c.Child(1).MustFloat64(), Y: c.Child(2).MustFloat64()})
			}
		case "layer":
			p.Layer = c.Child(1).MustString()
		case "width":
			p.Width = c.Child(1).MustFloat64()
		}
	}

	return &p, nil
}

func parseModArc(n sexp.Helper) (*ModArc, error) {
	a := ModArc{}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "start":
			a.Start.X = c.Child(1).MustFloat64()
			a.Start.Y = c.Child(2).MustFloat64()
		case "end":
			a.End.X = c.Child(1).MustFloat64()
			a.End.Y = c.Child(2).MustFloat64()
		case "layer":
			a.Layer = c.Child(1).MustString()
		case "width":
			a.Width = c.Child(1).MustFloat64()
		case "angle":
			a.Angle = c.Child(1).MustFloat64()
		}
	}

	return &a, nil
}

func parseModCircle(n sexp.Helper) (*ModCircle, error) {
	a := ModCircle{}
	for x := 1; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "center":
			a.Center.X = c.Child(1).MustFloat64()
			a.Center.Y = c.Child(2).MustFloat64()
		case "end":
			a.End.X = c.Child(1).MustFloat64()
			a.End.Y = c.Child(2).MustFloat64()
		case "layer":
			a.Layer = c.Child(1).MustString()
		case "width":
			a.Width = c.Child(1).MustFloat64()
		}
	}

	return &a, nil
}

func parseModPad(n sexp.Helper) (*Pad, error) {
	p := Pad{
		Ident: n.Child(1).MustString(),
	}

	switch n.Child(2).MustString() {
	case "smd":
		p.Surface = SurfaceSMD
	case "thru_hole":
		p.Surface = SurfaceTH
	case "np_thru_hole":
		p.Surface = SurfaceNPTH
	case "connect":
		p.Surface = SurfaceConnect
	}

	switch n.Child(3).MustString() {
	case "rect":
		p.Shape = ShapeRect
	case "oval":
		p.Shape = ShapeOval
	case "circle":
		p.Shape = ShapeCircle
	case "trapezoid":
		p.Shape = ShapeTrapezoid
	case "roundrect":
		p.Shape = ShapeRoundRect
	case "custom":
		p.Shape = ShapeCustom
	}

	for x := 4; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "at":
			p.At.X = c.Child(1).MustFloat64()
			p.At.Y = c.Child(2).MustFloat64()
			if c.MustNode().NumChildren() >= 4 {
				p.At.Z = c.Child(3).MustFloat64()
				p.At.ZPresent = true
			}
		case "size":
			p.Size.X = c.Child(1).MustFloat64()
			p.Size.Y = c.Child(2).MustFloat64()
		case "layers":
			for j := 1; j < c.MustNode().NumChildren(); j++ {
				p.Layers = append(p.Layers, c.Child(j).MustString())
			}

		case "rect_delta":
			p.RectDelta.X = c.Child(1).MustFloat64()
			p.RectDelta.Y = c.Child(2).MustFloat64()

		case "drill":
			readWidth := false
			for z := 1; z < c.MustNode().NumChildren(); z++ {
				c := c.Child(z)
				if c.IsList() {
					switch c.Child(0).MustString() {
					case "offset":
						p.DrillOffset = XY{X: c.Child(1).MustFloat64(), Y: c.Child(2).MustFloat64()}
					}
				} else {
					switch {
					case c.MustString() == "oval":
						p.DrillShape = ShapeDrillOblong
					default:
						// Width or height
						if readWidth {
							p.DrillSize.Y = c.MustFloat64()
						} else {
							p.DrillSize.X = c.MustFloat64()
							readWidth = true
						}
					}
				}
			}

		case "net":
			p.NetNum = c.Child(1).MustInt()
			p.NetName = c.Child(2).MustString()

		case "clearance":
			p.Clearance = c.Child(1).MustFloat64()
		case "die_length":
			p.DieLength = c.Child(1).MustFloat64()
		case "solder_paste_margin":
			p.SolderPasteMargin = c.Child(1).MustFloat64()
		case "solder_mask_margin":
			p.SolderMaskMargin = c.Child(1).MustFloat64()
		case "solder_paste_margin_ratio":
			p.SolderPasteMarginRatio = c.Child(1).MustFloat64()
		case "zone_connect":
			p.ZoneConnect = c.Child(1).MustInt()
		case "thermal_width":
			p.ThermalWidth = c.Child(1).MustFloat64()
		case "thermal_gap":
			p.ThermalGap = c.Child(1).MustFloat64()
		case "roundrect_rratio":
			p.RoundRectRRatio = c.Child(1).MustFloat64()
		case "chamfer_ratio":
			p.ChamferRatio = c.Child(1).MustFloat64()
			if p.ChamferRatio > 0 {
				p.Shape = ShapeChamferedRect
			}

			// TODO: chamfer, options, primitives
		}
	}

	return &p, nil
}

func parseModModel(n sexp.Helper) (*ModModel, error) {
	m := ModModel{
		Path: n.Child(1).MustString(),
	}

	for x := 2; x < n.MustNode().NumChildren(); x++ {
		c := n.Child(x)
		switch c.Child(0).MustString() {
		case "at":
			m.At.X = c.Child(1).Child(1).MustFloat64()
			m.At.Y = c.Child(1).Child(2).MustFloat64()
			if c.Child(1).MustNode().NumChildren() >= 4 {
				m.At.Z = c.Child(1).Child(3).MustFloat64()
				m.At.ZPresent = true
			}
		case "scale":
			m.Scale.X = c.Child(1).Child(1).MustFloat64()
			m.Scale.Y = c.Child(1).Child(2).MustFloat64()
			if c.Child(1).MustNode().NumChildren() >= 4 {
				m.Scale.Z = c.Child(1).Child(3).MustFloat64()
				m.Scale.ZPresent = true
			}
		case "rotate":
			m.Rotate.X = c.Child(1).Child(1).MustFloat64()
			m.Rotate.Y = c.Child(1).Child(2).MustFloat64()
			if c.Child(1).MustNode().NumChildren() >= 4 {
				m.Rotate.Z = c.Child(1).Child(3).MustFloat64()
				m.Rotate.ZPresent = true
			}
		}
	}

	return &m, nil
}
