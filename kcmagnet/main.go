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
	resolution   = flag.Int("resolution", 1, "How many interpolations to make per degree")
	output       = flag.String("o", "-", "Where output is written")
	skipWindings = flag.Float64("skip-windings", 1, "How many windings to skip on the inside")
)

func pointOnCircle(center kcgen.Point2D, radius float64, angle float64) *kcgen.Point2D {
	x := center.X + (radius * math.Cos(angle))
	y := center.Y + (radius * math.Sin(angle))
	return &kcgen.Point2D{X: x, Y: y}
}

func pointOnCircleDegrees(center kcgen.Point2D, radius float64, angle float64) *kcgen.Point2D {
	return pointOnCircle(center, radius, angle*math.Pi/180)
}

func main() {
	flag.Parse()

	if flag.NArg() < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <module-name> <trace-thickness> <trace-clearance> <windings>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	fp := kcgen.Footprint{
		ModName: flag.Arg(0),
	}

	thickness, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(1), err)
		os.Exit(2)
	}
	clearance, err := strconv.ParseFloat(flag.Arg(2), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(2), err)
		os.Exit(2)
	}
	windings, err := strconv.Atoi(flag.Arg(3))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(2), err)
		os.Exit(2)
	}

	startRadius := float64(*skipWindings) * (thickness + clearance)
	lastPoint := kcgen.Point2D{X: startRadius}
	radiusIncrement := (thickness + clearance) / 360 / float64(*resolution)

	for degree := 0; degree < 360*(*resolution)*windings; degree++ {
		next := pointOnCircleDegrees(kcgen.Point2D{}, startRadius+radiusIncrement*float64(degree), float64(degree)/float64(*resolution))
		fp.Add(&kcgen.Line{
			Layer: kcgen.LayerFrontCopper,
			Start: lastPoint,
			End:   *next,
			Width: thickness,
		})
		lastPoint = *next
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
