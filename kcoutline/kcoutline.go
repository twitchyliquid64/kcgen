package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/twitchyliquid64/kcgen"
)

var (
	referenceOffset = flag.Float64("refY", 0, "Y-axis offset at which module designator is placed")
	radius          = flag.Float64("radius", 3, "Rounded edges radius")
	resolution      = flag.Int("resolution", 0, "How many interpolations to make per degree")
	output          = flag.String("o", "-", "Where output is written")
	mounts          = flag.Bool("make-mounts", false, "Generate mounting holes")
	rimming         = flag.Bool("make-rim", false, "Generate copper rim")
	rimWidth        = flag.Float64("rim-width", 0.5, "Width of copper rim")
)

func main() {
	flag.Parse()

	if flag.NArg() < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <module-name> <width> <height>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	fp := kcgen.Footprint{
		ModName:    flag.Arg(0),
		ReferenceY: *referenceOffset,
	}

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
	fp.Add(&kcgen.Line{
		Layer: kcgen.LayerEdgeCuts,
		Start: kcgen.Point2D{X: *radius - halfWidth, Y: -halfHeight},
		End:   kcgen.Point2D{X: halfWidth - *radius, Y: -halfHeight},
	})
	fp.Add(&kcgen.Line{
		Layer: kcgen.LayerEdgeCuts,
		Start: kcgen.Point2D{X: *radius - halfWidth, Y: halfHeight},
		End:   kcgen.Point2D{X: halfWidth - *radius, Y: halfHeight},
	})
	// Vertical.
	fp.Add(&kcgen.Line{
		Layer: kcgen.LayerEdgeCuts,
		Start: kcgen.Point2D{X: -halfWidth, Y: *radius - halfHeight},
		End:   kcgen.Point2D{X: -halfWidth, Y: halfHeight - *radius},
	})
	fp.Add(&kcgen.Line{
		Layer: kcgen.LayerEdgeCuts,
		Start: kcgen.Point2D{X: halfWidth, Y: *radius - halfHeight},
		End:   kcgen.Point2D{X: halfWidth, Y: halfHeight - *radius},
	})
	if *rimming {
		fp.Add(&kcgen.Pad{
			Number: 1,
			Size:   kcgen.Point2D{X: *rimWidth, Y: height - *radius*2},
			Center: kcgen.Point2D{X: halfWidth - rimOffset, Y: 0},
			Type:   "smd rect",
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
		})
		fp.Add(&kcgen.Pad{
			Number: 1,
			Size:   kcgen.Point2D{X: *rimWidth, Y: height - *radius*2},
			Center: kcgen.Point2D{X: rimOffset - halfWidth, Y: 0},
			Type:   "smd rect",
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
		})
		fp.Add(&kcgen.Pad{
			Number: 1,
			Size:   kcgen.Point2D{X: width - *radius*2, Y: *rimWidth},
			Center: kcgen.Point2D{X: 0, Y: halfHeight - rimOffset},
			Type:   "smd rect",
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
		})
		fp.Add(&kcgen.Pad{
			Number: 1,
			Size:   kcgen.Point2D{X: width - *radius*2, Y: *rimWidth},
			Center: kcgen.Point2D{X: 0, Y: rimOffset - halfHeight},
			Type:   "smd rect",
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
		})
	}

	// Radius arcs.
	if *radius > 0 {
		fp.Add(&kcgen.Arc{
			Layer:          kcgen.LayerEdgeCuts,
			Start:          -90,
			End:            0,
			Center:         kcgen.Point2D{X: halfWidth - *radius, Y: *radius - halfHeight},
			StepsPerDegree: *resolution,
			Radius:         *radius,
		})
		fp.Add(&kcgen.Arc{
			Layer:          kcgen.LayerEdgeCuts,
			Start:          0,
			End:            90,
			Center:         kcgen.Point2D{X: halfWidth - *radius, Y: halfHeight - *radius},
			StepsPerDegree: *resolution,
			Radius:         *radius,
		})
		fp.Add(&kcgen.Arc{
			Layer:          kcgen.LayerEdgeCuts,
			Start:          90,
			End:            180,
			Center:         kcgen.Point2D{X: *radius - halfWidth, Y: halfHeight - *radius},
			StepsPerDegree: *resolution,
			Radius:         *radius,
		})
		fp.Add(&kcgen.Arc{
			Layer:          kcgen.LayerEdgeCuts,
			Start:          180,
			End:            270,
			Center:         kcgen.Point2D{X: *radius - halfWidth, Y: *radius - halfHeight},
			StepsPerDegree: *resolution,
			Radius:         *radius,
		})
	}

	if *mounts {
		fp.Add(&kcgen.Pad{
			Type:   "thru_hole circle",
			Number: 1,
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
			Size:   kcgen.Point2D{X: 5.6, Y: 5.6},
			Center: kcgen.Point2D{X: *radius - halfWidth, Y: *radius - halfHeight},
			Drill:  3.2,
		})
		fp.Add(&kcgen.Pad{
			Type:   "thru_hole circle",
			Number: 1,
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
			Size:   kcgen.Point2D{X: 5.6, Y: 5.6},
			Center: kcgen.Point2D{X: *radius - halfWidth, Y: halfHeight - *radius},
			Drill:  3.2,
		})
		fp.Add(&kcgen.Pad{
			Type:   "thru_hole circle",
			Number: 1,
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
			Size:   kcgen.Point2D{X: 5.6, Y: 5.6},
			Center: kcgen.Point2D{X: halfWidth - *radius, Y: *radius - halfHeight},
			Drill:  3.2,
		})
		fp.Add(&kcgen.Pad{
			Type:   "thru_hole circle",
			Number: 1,
			Layers: []kcgen.Layer{kcgen.LayerAllCopper, kcgen.LayerAllMask},
			Size:   kcgen.Point2D{X: 5.6, Y: 5.6},
			Center: kcgen.Point2D{X: halfWidth - *radius, Y: halfHeight - *radius},
			Drill:  3.2,
		})
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

	if err := fp.Render(w); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		if *output != "" && *output != "-" { //close the file if its not standard input
			w.Close()
		}
		os.Exit(4)
	}
}
