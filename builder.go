package kcgen

import (
	"fmt"
	"io"
	"time"

	"github.com/twitchyliquid64/kcgen/pcb"
)

// ModBuilder provides an easy interface to construct a new KiCad module.
type ModBuilder struct {
	mod pcb.Module
}

func (m *ModBuilder) Write(w io.Writer) error {
	return m.mod.WriteModule(w)
}

// NewModuleBuilder returns a builder for constructing a module.
func NewModuleBuilder(name, description string, layer Layer) *ModBuilder {
	return &ModBuilder{
		mod: pcb.Module{
			Name:        name,
			Layer:       layer.Strictname(),
			Tedit:       fmt.Sprintf("%8X", time.Now().Unix()),
			Description: description,
			ZoneConnect: pcb.ZoneConnectInherited,
		},
	}
}

// ZoneConnectMode sets the zone connect mode.
func (m *ModBuilder) ZoneConnectMode(mode pcb.ZoneConnectMode) {
	m.mod.ZoneConnect = mode
}

// Tags sets the tags.
func (m *ModBuilder) Tags(tags []string) {
	m.mod.Tags = tags
}

// Attributes sets the attributes.
func (m *ModBuilder) Attributes(attrs []string) {
	m.mod.Attrs = attrs
}

// AddText adds a text graphic to the module.
func (m *ModBuilder) AddText(t *Text) {
	m.mod.Graphics = append(m.mod.Graphics, pcb.ModGraphic{
		Ident:      "fp_text",
		Renderable: &t.t,
	})
}

// AddText adds a text graphic to the module.
func (m *ModBuilder) AddCircle(c *Circle) {
	m.mod.Graphics = append(m.mod.Graphics, pcb.ModGraphic{
		Ident:      "fp_circle",
		Renderable: &c.c,
	})
}

// AddLine adds a line graphic to the module.
func (m *ModBuilder) AddLine(l *Line) {
	m.mod.Graphics = append(m.mod.Graphics, pcb.ModGraphic{
		Ident:      "fp_line",
		Renderable: &l.l,
	})
}

// AddPolygon adds a polygon graphic to the module.
func (m *ModBuilder) AddPolygon(p *Polygon) {
	m.mod.Graphics = append(m.mod.Graphics, pcb.ModGraphic{
		Ident:      "fp_poly",
		Renderable: &p.p,
	})
}

// AddArc adds an arc graphic to the module.
func (m *ModBuilder) AddArc(a *Arc) {
	m.mod.Graphics = append(m.mod.Graphics, pcb.ModGraphic{
		Ident:      "fp_arc",
		Renderable: &a.a,
	})
}