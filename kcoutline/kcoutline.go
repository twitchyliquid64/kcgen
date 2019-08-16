package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/twitchyliquid64/kcgen"
)

var (
	referenceOffset = flag.Float64("refY", 0, "Y-axis offset at which module designator is placed")
	radius          = flag.Float64("radius", 3, "Rounded edges radius")
	resolution      = flag.Int("resolution", 1, "How many interpolations to make per degree")
	output          = flag.String("o", "-", "Where output is written")
	mounts          = flag.Bool("make-mounts", false, "Generate mounting holes")
	rimming         = flag.Bool("make-rim", false, "Generate copper rim")
	rimWidth        = flag.Float64("rim-width", 0.5, "Width of copper rim")
)

func pointOnCircle(center [2]float64, radius float64, angle float64) [2]float64 {
	x := center[0] + (radius * math.Cos(angle))
	y := center[1] + (radius * math.Sin(angle))
	return [2]float64{x, y}
}

func pointOnCircleDegrees(center [2]float64, radius float64, angle float64) [2]float64 {
	return pointOnCircle(center, radius, angle*math.Pi/180)
}

func drawArc(m *kcgen.ModBuilder, radius, start, end, x, y float64) {
	var last *[2]float64
	for i := 0; i < int(end-start)*(*resolution); i++ {
		p := pointOnCircleDegrees([2]float64{x, y}, radius, start+float64(i/(*resolution)))
		if last != nil {
			l := kcgen.NewLine(kcgen.LayerEdgeCuts)
			l.Positions(last[0], last[1], p[0], p[1])
			m.AddLine(l)
		}
		last = &p
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <module-name> <width> <height>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	m := kcgen.NewModuleBuilder(flag.Arg(0), "The outline of the PCB.", kcgen.LayerFrontCopper)
	m.RefTextOffset(0, *referenceOffset)

	width, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(1), err)
		os.Exit(2)
	}

	height, err := strconv.ParseFloat(flag.Arg(2), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(2), err)
		os.Exit(2)
	}

	// Create the main edge-cut lines.
	halfWidth := width / 2.0
	halfHeight := height / 2.0
	rimOffset := *rimWidth / 2.0
	if *mounts {
		rimOffset += *radius / 15
	}
	// Horizontal.
	top := kcgen.NewLine(kcgen.LayerEdgeCuts)
	top.Positions(*radius-halfWidth, -halfHeight, halfWidth-*radius, -halfHeight)
	m.AddLine(top)
	bottom := kcgen.NewLine(kcgen.LayerEdgeCuts)
	bottom.Positions(*radius-halfWidth, halfHeight, halfWidth-*radius, halfHeight)
	m.AddLine(bottom)

	// Vertical.
	left := kcgen.NewLine(kcgen.LayerEdgeCuts)
	left.Positions(-halfWidth, *radius-halfHeight, -halfWidth, halfHeight-*radius)
	m.AddLine(left)
	right := kcgen.NewLine(kcgen.LayerEdgeCuts)
	right.Positions(halfWidth, *radius-halfHeight, halfWidth, halfHeight-*radius)
	m.AddLine(right)

	if *rimming {
		r1 := kcgen.NewPad(kcgen.SMDRect, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		r1.Size(*rimWidth, height-*radius*2)
		r1.Center(halfWidth-rimOffset, 0)
		m.AddPad(r1)

		r2 := kcgen.NewPad(kcgen.SMDRect, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		r2.Size(*rimWidth, height-*radius*2)
		r2.Center(rimOffset-halfWidth, 0)
		m.AddPad(r2)

		r3 := kcgen.NewPad(kcgen.SMDRect, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		r3.Size(width-*radius*2, *rimWidth)
		r3.Center(0, halfHeight-rimOffset)
		m.AddPad(r3)

		r4 := kcgen.NewPad(kcgen.SMDRect, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		r4.Size(width-*radius*2, *rimWidth)
		r4.Center(0, rimOffset-halfHeight)
		m.AddPad(r4)
	}

	// Radius arcs.
	if *radius > 0 {
		drawArc(m, *radius, -90, 0, halfWidth-*radius, *radius-halfHeight)
		drawArc(m, *radius, 0, 90, halfWidth-*radius, halfHeight-*radius)
		drawArc(m, *radius, 90, 180, *radius-halfWidth, halfHeight-*radius)
		drawArc(m, *radius, 180, 270, *radius-halfWidth, *radius-halfHeight)
	}

	if *radius < 2.8 {
		*radius = 2.8
	}
	if *mounts {
		tl := kcgen.NewPad(kcgen.TH, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		tl.Size(5.6, 5.6)
		tl.Center(*radius-halfWidth, *radius-halfHeight)
		tl.DrillSize(3.2, 3.2)
		m.AddPad(tl)

		bl := kcgen.NewPad(kcgen.TH, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		bl.Size(5.6, 5.6)
		bl.Center(*radius-halfWidth, halfHeight-*radius)
		bl.DrillSize(3.2, 3.2)
		m.AddPad(bl)

		tr := kcgen.NewPad(kcgen.TH, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		tr.Size(5.6, 5.6)
		tr.Center(halfWidth-*radius, *radius-halfHeight)
		tr.DrillSize(3.2, 3.2)
		m.AddPad(tr)

		br := kcgen.NewPad(kcgen.TH, "1", kcgen.LayerAllCopper, kcgen.LayerAllMask)
		br.Size(5.6, 5.6)
		br.Center(halfWidth-*radius, halfHeight-*radius)
		br.DrillSize(3.2, 3.2)
		m.AddPad(br)
	}

	// Render output.
	w := os.Stdout
	if *output != "" && *output != "-" {
		f, err := os.OpenFile(*output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open %q: %v\n", *output, err)
			os.Exit(3)
		}
		defer f.Close()
		w = f
	}

	if err := m.Write(w); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		if *output != "" && *output != "-" { //close the file if its not standard input
			w.Close()
		}
		os.Exit(4)
	}
}
