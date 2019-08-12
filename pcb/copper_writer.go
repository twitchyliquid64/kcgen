package pcb

import (
	"github.com/twitchyliquid64/kcgen/swriter"
)

// write generates an s-expression describing the via.
func (v *Via) write(sw *swriter.SExpWriter) error {
	sw.StartList(false)
	sw.StringScalar("via")
	if err := v.At.write("at", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("size")
	sw.StringScalar(f(v.Size))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("drill")
	sw.StringScalar(f(v.Drill))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layers")
	for _, l := range v.Layers {
		sw.StringScalar(l)
	}
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("net")
	sw.IntScalar(v.NetIndex)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	return sw.CloseList(false)
}

// write generates an s-expression describing the zone.
func (z *Zone) write(sw *swriter.SExpWriter) error {
	sw.StartList(false)
	sw.StringScalar("zone")

	sw.StartList(false)
	sw.StringScalar("net")
	sw.IntScalar(z.NetNum)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("net_name")
	sw.StringScalar(z.NetName)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(z.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("tstamp")
	sw.StringScalar(z.Tstamp)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("hatch")
	sw.StringScalar(z.Hatch.Mode)
	sw.StringScalar(f(z.Hatch.Size))
	if err := sw.CloseList(false); err != nil {
		return err
	}
	sw.Newlines(1)

	sw.StartList(false)
	sw.StringScalar("connect_pads")
	sw.StartList(false)
	sw.StringScalar("clearance")
	sw.StringScalar(f(z.ConnectPads.Clearance))
	if err := sw.CloseList(false); err != nil {
		return err
	}
	if err := sw.CloseList(false); err != nil {
		return err
	}
	sw.Newlines(1)

	sw.StartList(false)
	sw.StringScalar("min_thickness")
	sw.StringScalar(f(z.MinThickness))
	if err := sw.CloseList(false); err != nil {
		return err
	}
	sw.Newlines(1)

	sw.StartList(false)
	sw.StringScalar("fill")
	if z.Fill.Enabled {
		sw.StringScalar("yes")
	} else {
		sw.StringScalar("no")
	}
	sw.StartList(false)
	sw.StringScalar("arc_segments")
	sw.IntScalar(z.Fill.Segments)
	if err := sw.CloseList(false); err != nil {
		return err
	}
	sw.StartList(false)
	sw.StringScalar("thermal_gap")
	sw.StringScalar(f(z.Fill.ThermalGap))
	if err := sw.CloseList(false); err != nil {
		return err
	}
	sw.StartList(false)
	sw.StringScalar("thermal_bridge_width")
	sw.StringScalar(f(z.Fill.ThermalBridgeWidth))
	if err := sw.CloseList(false); err != nil {
		return err
	}
	if err := sw.CloseList(false); err != nil {
		return err
	}
	sw.Newlines(1)

	for _, p := range z.BasePolys {
		sw.StartList(false)
		sw.StringScalar("polygon")
		sw.Newlines(1)
		sw.StartList(false)
		sw.StringScalar("pts")
		sw.Newlines(1)

		for i, pts := range p {
			if err := pts.write("xy", sw); err != nil {
				return err
			}
			if i%5 == 4 {
				sw.Newlines(1)
			}
		}

		if err := sw.CloseList(true); err != nil {
			return err
		}
		if err := sw.CloseList(true); err != nil {
			return err
		}
	}
	sw.Newlines(1)

	for _, p := range z.Polys {
		sw.StartList(false)
		sw.StringScalar("filled_polygon")
		sw.Newlines(1)
		sw.StartList(false)
		sw.StringScalar("pts")
		sw.Newlines(1)

		for i, pts := range p {
			if err := pts.write("xy", sw); err != nil {
				return err
			}
			if i%5 == 4 {
				sw.Newlines(1)
			}
		}

		if err := sw.CloseList(true); err != nil {
			return err
		}
		if err := sw.CloseList(true); err != nil {
			return err
		}
	}

	return sw.CloseList(true)
}

// write generates an s-expression describing the track.
func (t *Track) write(sw *swriter.SExpWriter) error {
	sw.StartList(false)
	sw.StringScalar("segment")
	if err := t.Start.write("start", sw); err != nil {
		return err
	}
	if err := t.End.write("end", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("width")
	sw.StringScalar(f(t.Width))
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layer")
	sw.StringScalar(t.Layer)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("net")
	sw.IntScalar(t.NetIndex)
	if err := sw.CloseList(false); err != nil {
		return err
	}

	if t.Tstamp != "" {
		sw.StartList(false)
		sw.StringScalar("tstamp")
		sw.StringScalar(t.Tstamp)
		if err := sw.CloseList(false); err != nil {
			return err
		}
	}

	return sw.CloseList(false)
}
