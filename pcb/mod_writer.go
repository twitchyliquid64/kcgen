package pcb

import (
	"strings"

	"github.com/twitchyliquid64/kcgen/swriter"
)

func (m *Module) write(sw *swriter.SExpWriter, doPlacement bool) error {
	sw.StartList(false)
	sw.StringScalar("module")
	sw.StringScalar(m.Name)

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(m.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	if m.Tedit != "" {
		sw.StartList(false)
		sw.StringScalar("tedit")
		sw.StringScalar(m.Tedit)
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}
	if m.Tstamp != "" {
		sw.StartList(false)
		sw.StringScalar("tstamp")
		sw.StringScalar(m.Tstamp)
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}
	sw.Newlines(1)

	if doPlacement {
		if err := m.Placement.At.write("at", sw); err != nil {
			return err
		}
	}

	if m.Description != "" {
		sw.StartList(true)
		sw.StringScalar("descr")
		sw.StringScalar(m.Description)
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}

	if len(m.Tags) > 0 {
		sw.StartList(true)
		sw.StringScalar("tags")
		sw.StringScalar(strings.Join(m.Tags, " "))
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}

	if m.Path != "" {
		sw.StartList(true)
		sw.StringScalar("path")
		sw.StringScalar(m.Path)
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}

	if len(m.Attrs) > 0 {
		sw.StartList(true)
		sw.StringScalar("attr")
		for _, a := range m.Attrs {
			sw.StringScalar(a)
		}
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}

	for _, g := range m.Graphics {
		if err := g.Renderable.write(sw); err != nil {
			return err
		}
	}

	if m.Model != nil {
		sw.StartList(true)
		sw.StringScalar("model")
		sw.StringScalar(m.Model.Path)

		sw.StartList(true)
		sw.StringScalar("at")
		if err := m.Model.At.write("xyz", sw); err != nil {
			return err
		}
		if err := sw.CloseList(false); err != nil {
			return err
		}
		sw.StartList(true)
		sw.StringScalar("scale")
		if err := m.Model.Scale.write("xyz", sw); err != nil {
			return err
		}
		if err := sw.CloseList(false); err != nil {
			return err
		}
		sw.StartList(true)
		sw.StringScalar("rotate")
		if err := m.Model.Rotate.write("xyz", sw); err != nil {
			return err
		}
		if err := sw.CloseList(false); err != nil {
			return err
		}

		if err := sw.CloseList(true); err != nil {
			return err
		}
	}

	if err := sw.CloseList(true); err != nil {
		return err
	}
	return nil
}

func (l *ModLine) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("fp_line")
	if err := l.Start.write("start", sw); err != nil {
		return err
	}
	if err := l.End.write("end", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(l.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("width")
	sw.StringScalar(f(l.Width))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	return sw.CloseList(false)
}

func (a *ModArc) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("fp_arc")
	if err := a.Start.write("start", sw); err != nil {
		return err
	}
	if err := a.End.write("end", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("angle")
	sw.StringScalar(f(a.Angle))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(a.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("width")
	sw.StringScalar(f(a.Width))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	return sw.CloseList(false)
}

func (c *ModCircle) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("fp_circle")
	if err := c.Center.write("center", sw); err != nil {
		return err
	}
	if err := c.End.write("end", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(c.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("width")
	sw.StringScalar(f(c.Width))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	return sw.CloseList(false)
}

func (t *ModText) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("fp_text")
	sw.StringScalar(t.Kind.String())
	sw.StringScalar(t.Text)
	if err := t.At.write("at", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(t.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(true)
	sw.StringScalar("effects")
	sw.StartList(false)
	sw.StringScalar("font")
	if err := t.Effects.FontSize.write("size", sw); err != nil {
		return err
	}
	sw.StartList(false)
	sw.StringScalar("thickness")
	sw.StringScalar(f(t.Effects.Thickness))
	if err := sw.CloseList(false); err != nil {
		return err
	}
	if err := sw.CloseList(false); err != nil {
		return err
	}
	if err := sw.CloseList(false); err != nil {
		return err
	}

	if err := sw.CloseList(true); err != nil {
		return err
	}
	return nil
}

func (p *ModPolygon) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("fp_poly")

	sw.StartList(false)
	sw.StringScalar("pts")
	for i, pts := range p.Points {
		if err := pts.write("xy", sw); err != nil {
			return err
		}
		if i%4 == 3 {
			sw.Newlines(1)
		}
	}
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(p.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("width")
	sw.StringScalar(f(p.Width))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	return sw.CloseList(false)
}
