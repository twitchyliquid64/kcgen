package kcgen

import (
	"github.com/twitchyliquid64/kcgen/pcb"
)

// PadType differenciates different types of pad.
type PadType uint8

// Valid pads.
const (
	TH PadType = iota
	SMDRect
	SMDCircle
	SMDOval
)

// Pad represents a pad as part of a module.
type Pad struct {
	p pcb.Pad
}

// Center sets the position of the pad.
func (p *Pad) Center(x, y float64) {
	p.p.At.X = x
	p.p.At.Y = y
}

// Center sets the rotation of the pad.
func (p *Pad) Rotation(z float64) {
	p.p.At.Z = z
}

// Size sets the size of the pad.
func (p *Pad) Size(x, y float64) {
	p.p.Size = pcb.XY{X: x, Y: y}
}

// DrillSize sets the size of the drill cutout.
func (p *Pad) DrillSize(x, y float64) {
	p.p.DrillSize = pcb.XY{X: x, Y: y}
}

func NewPad(t PadType, ident string, layers ...Layer) *Pad {
	out := Pad{
		p: pcb.Pad{
			Ident:       ident,
			ZoneConnect: pcb.ZoneConnectInherited,
		},
	}

	switch t {
	case TH:
		out.p.Surface = pcb.SurfaceTH
		out.p.Shape = pcb.ShapeOval
	case SMDRect:
		out.p.Surface = pcb.SurfaceSMD
		out.p.Shape = pcb.ShapeRect
	case SMDCircle:
		out.p.Surface = pcb.SurfaceSMD
		out.p.Shape = pcb.ShapeCircle
	case SMDOval:
		out.p.Surface = pcb.SurfaceSMD
		out.p.Shape = pcb.ShapeOval
	}

	out.p.Layers = make([]string, len(layers))
	for i := range layers {
		out.p.Layers[i] = layers[i].Strictname()
	}

	return &out
}
