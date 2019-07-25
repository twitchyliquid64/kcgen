package pcb

import (
	"fmt"
	"strings"

	"github.com/nsf/sexp"
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
	Model    *ModModel
}

// ModPlacement describes the positioning of a module on a PCB.
type ModPlacement struct {
	At XYZ
}

// ModGraphic represents a graphical feature in a module.
type ModGraphic struct {
	Ident      string
	Renderable renderable
}

type renderable interface {
}

// ModText represents a text drawing in a module.
type ModText struct {
	Kind ModTextKind
	Text string
	At   XYZ

	Layer   string
	Effects TextEffects
}

// ModTextKind describes the type of text drawing.
type ModTextKind uint8

// Valid ModTextKind values.
const (
	RefText ModTextKind = iota
	ValueText
)

// ModLine represents a line drawing in a module.
type ModLine struct {
	Start XY
	End   XY

	Layer string
	Width float64
}

// ModModel describes configuration for rendering a 3d model of the part.
type ModModel struct {
	Path   string
	At     XYZ
	Scale  XYZ
	Rotate XYZ
}

// PadType enumerates valid kinds of copper pads.
type PadType uint8

// Valid PadTypes.
const (
	PadInvalid PadType = iota
	ThroughHole
	SurfaceMount
)

// Pad represents a copper pad.
type Pad interface {
	Type() PadType
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
	default:
		return nil, fmt.Errorf("unknown fp_line type: %v", n.Child(1).MustString())
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
