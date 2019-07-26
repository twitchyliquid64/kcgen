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
