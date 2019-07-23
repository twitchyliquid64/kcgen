package pcb

import (
	"fmt"
	"io"
	"sort"

	"github.com/twitchyliquid64/kcgen/swriter"
)

// Write produces the file on disk. IMPLEMENTATION IS NOT YET COMPLETE.
func (p *PCB) Write(w io.Writer) error {
	sw, err := swriter.NewSExpWriter(w)
	if err != nil {
		return err
	}
	sw.StartList(false)
	sw.StringScalar("kicad_pcb")

	// Version
	sw.StartList(false)
	sw.StringScalar("version")
	sw.IntScalar(p.FormatVersion)
	if err := sw.CloseList(); err != nil {
		return err
	}

	// EG: host pcbnew 4.0.7
	sw.StartList(false)
	sw.StringScalar("host")
	sw.StringScalar("kcgen")
	sw.StringScalar("0.0.1")
	if err := sw.CloseList(); err != nil {
		return err
	}
	sw.Newlines(2)

	// EG: general (no_connects 0) ...
	sw.StartList(false)
	sw.StringScalar("general")
	// sw.StartList(true)
	// sw.StringScalar("no_connects")
	// sw.IntScalar(0)
	// if err := sw.CloseList(); err != nil {
	// 	return err
	// }
	// sw.Newlines(1)
	if err := sw.CloseList(); err != nil {
		return err
	}
	sw.Newlines(1)

	// EG: page A4
	sw.StartList(true)
	sw.StringScalar("page")
	sw.StringScalar("A4")
	if err := sw.CloseList(); err != nil {
		return err
	}
	sw.Newlines(1)

	// Layers
	sw.StartList(true)
	sw.StringScalar("layers")
	sw.Newlines(1)
	for _, layer := range p.Layers {
		if err := layer.write(sw); err != nil {
			return err
		}
		sw.Newlines(1)
	}
	if err := sw.CloseList(); err != nil {
		return err
	}
	sw.Newlines(1)

	// Setup
	if err := p.EditorSetup.write(sw); err != nil {
		return err
	}

	// Nets
	if err := p.writeNets(sw); err != nil {
		return err
	}

	// Net classes
	for _, nc := range p.NetClasses {
		if err := nc.write(sw); err != nil {
			return err
		}
	}

	// Vias
	for _, v := range p.Vias {
		if err := v.write(sw); err != nil {
			return err
		}
	}

	return sw.CloseList()
}

type netPair struct {
	num int
	net Net
}

func (p *PCB) writeNets(sw *swriter.SExpWriter) error {
	var nets []netPair
	for num, net := range p.Nets {
		nets = append(nets, netPair{num: num, net: net})
	}
	sort.Slice(nets, func(i, j int) bool {
		return nets[i].num < nets[j].num
	})

	for _, n := range nets {
		sw.StartList(true)
		sw.StringScalar("net")
		sw.IntScalar(n.num)
		sw.StringScalar(n.net.Name)
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	if len(nets) > 0 {
		sw.Newlines(1)
	}
	return nil
}

// write generates an s-expression describing the layer.
func (l *Layer) write(sw *swriter.SExpWriter) error {
	sw.StartList(false)
	sw.IntScalar(l.Num)
	sw.StringScalar(l.Name)
	sw.StringScalar(l.Type)
	return sw.CloseList()
}

func f(f float64) string {
	t := fmt.Sprintf("%f", f)
	if t[len(t)-1] != '0' {
		return t
	}

	for i := len(t) - 1; i >= 0; i-- {
		if t[i] != '0' {
			if t[i] == '.' {
				return t[:i]
			}
			return t[:i+1]
		}
	}
	return t
}

// write generates an s-expression describing the point.
func (p *XY) write(prefix string, sw *swriter.SExpWriter) error {
	sw.StartList(false)
	sw.StringScalar(prefix)
	sw.StringScalar(f(p.X))
	sw.StringScalar(f(p.Y))
	return sw.CloseList()
}

// write generates an s-expression describing the via.
func (v *Via) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("via")
	if err := v.At.write("at", sw); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("size")
	sw.StringScalar(f(v.Size))
	if err := sw.CloseList(); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("drill")
	sw.StringScalar(f(v.Drill))
	if err := sw.CloseList(); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("layers")
	for _, l := range v.Layers {
		sw.StringScalar(l)
	}
	if err := sw.CloseList(); err != nil {
		return err
	}

	sw.StartList(false)
	sw.StringScalar("net")
	sw.IntScalar(v.NetIndex)
	if err := sw.CloseList(); err != nil {
		return err
	}

	return sw.CloseList()
}

// write generates an s-expression describing the layer.
func (l *EditorSetup) write(sw *swriter.SExpWriter) error {
	sw.StartList(false)
	sw.StringScalar("setup")

	if l.LastTraceWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("last_trace_width")
		sw.StringScalar(f(l.LastTraceWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	for _, w := range l.UserTraceWidths {
		sw.StartList(true)
		sw.StringScalar("user_trace_width")
		sw.StringScalar(f(w))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.TraceClearance > 0 {
		sw.StartList(true)
		sw.StringScalar("trace_clearance")
		sw.StringScalar(f(l.TraceClearance))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.ZoneClearance > 0 {
		sw.StartList(true)
		sw.StringScalar("zone_clearance")
		sw.StringScalar(f(l.ZoneClearance))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	sw.StartList(true)
	sw.StringScalar("zone_45_only")
	if l.Zone45Only {
		sw.StringScalar("yes")
	} else {
		sw.StringScalar("no")
	}
	if err := sw.CloseList(); err != nil {
		return err
	}
	if l.TraceMin > 0 {
		sw.StartList(true)
		sw.StringScalar("trace_min")
		sw.StringScalar(f(l.TraceMin))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.SegmentWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("segment_width")
		sw.StringScalar(f(l.SegmentWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.EdgeWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("edge_width")
		sw.StringScalar(f(l.EdgeWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	if l.ViaSize > 0 {
		sw.StartList(true)
		sw.StringScalar("via_size")
		sw.StringScalar(f(l.ViaSize))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.ViaMinSize > 0 {
		sw.StartList(true)
		sw.StringScalar("via_min_size")
		sw.StringScalar(f(l.ViaMinSize))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.ViaMinDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("via_min_drill")
		sw.StringScalar(f(l.ViaMinDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.ViaDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("via_drill")
		sw.StringScalar(f(l.ViaDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.UViaSize > 0 {
		sw.StartList(true)
		sw.StringScalar("uvia_size")
		sw.StringScalar(f(l.UViaSize))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.UViaMinSize > 0 {
		sw.StartList(true)
		sw.StringScalar("uvia_min_size")
		sw.StringScalar(f(l.UViaMinSize))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.UViaMinDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("uvia_min_drill")
		sw.StringScalar(f(l.UViaMinDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.UViaDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("uvia_drill")
		sw.StringScalar(f(l.UViaDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	sw.StartList(true)
	sw.StringScalar("uvias_allowed")
	if l.AllowUVias {
		sw.StringScalar("yes")
	} else {
		sw.StringScalar("no")
	}
	if err := sw.CloseList(); err != nil {
		return err
	}

	if l.TextWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("pcb_text_width")
		sw.StringScalar(f(l.TextWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if len(l.TextSize) > 0 {
		sw.StartList(true)
		sw.StringScalar("pcb_text_size")
		for _, w := range l.TextSize {
			sw.StringScalar(f(w))
		}
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	if l.ModEdgeWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("mod_edge_width")
		sw.StringScalar(f(l.ModEdgeWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if len(l.ModTextSize) > 0 {
		sw.StartList(true)
		sw.StringScalar("mod_text_size")
		for _, w := range l.ModTextSize {
			sw.StringScalar(f(w))
		}
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.ModTextWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("mod_text_width")
		sw.StringScalar(f(l.ModTextWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	if len(l.PadSize) > 0 {
		sw.StartList(true)
		sw.StringScalar("pad_size")
		for _, w := range l.PadSize {
			sw.StringScalar(f(w))
		}
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.PadDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("pad_drill")
		sw.StringScalar(f(l.PadDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if l.PadToMaskClearance > 0 {
		sw.StartList(true)
		sw.StringScalar("pad_to_mask_clearance")
		sw.StringScalar(f(l.PadToMaskClearance))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	if len(l.PlotParams) > 0 {
		sw.StartList(true)
		sw.StringScalar("pcbplotparams")
		var pps []PlotParam
		for _, pp := range l.PlotParams {
			pps = append(pps, pp)
		}
		sort.Slice(pps, func(i, j int) bool {
			return pps[i].order < pps[j].order
		})

		for _, pp := range pps {
			sw.StartList(true)
			sw.StringScalar(pp.name)
			for _, v := range pp.values {
				sw.StringScalar(v)
			}
			if err := sw.CloseList(); err != nil {
				return err
			}
		}
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	sw.Newlines(1)
	if err := sw.CloseList(); err != nil {
		return err
	}
	sw.Newlines(1)
	return nil
}

// write generates an s-expression describing the layer.
func (c *NetClass) write(sw *swriter.SExpWriter) error {
	sw.StartList(true)
	sw.StringScalar("net_class")
	sw.StringScalar(c.Name)
	sw.StringScalar(c.Description)

	if c.Clearance > 0 {
		sw.StartList(true)
		sw.StringScalar("clearance")
		sw.StringScalar(f(c.Clearance))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if c.TraceWidth > 0 {
		sw.StartList(true)
		sw.StringScalar("trace_width")
		sw.StringScalar(f(c.TraceWidth))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if c.ViaDiameter > 0 {
		sw.StartList(true)
		sw.StringScalar("via_dia")
		sw.StringScalar(f(c.ViaDiameter))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if c.ViaDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("via_drill")
		sw.StringScalar(f(c.ViaDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if c.UViaDiameter > 0 {
		sw.StartList(true)
		sw.StringScalar("uvia_dia")
		sw.StringScalar(f(c.UViaDiameter))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	if c.UViaDrill > 0 {
		sw.StartList(true)
		sw.StringScalar("uvia_drill")
		sw.StringScalar(f(c.UViaDrill))
		if err := sw.CloseList(); err != nil {
			return err
		}
	}

	for _, net := range c.Nets {
		sw.StartList(true)
		sw.StringScalar("add_net")
		sw.StringScalar(net)
		if err := sw.CloseList(); err != nil {
			return err
		}
	}
	sw.Newlines(1)
	return sw.CloseList()
}
