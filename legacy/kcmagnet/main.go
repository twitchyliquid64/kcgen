package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	kcgen "github.com/twitchyliquid64/kcgen/legacy"
)

var (
	resolution   = flag.Int("resolution", 1, "How many interpolations to make per degree")
	output       = flag.String("o", "-", "Where output is written")
	skipWindings = flag.Float64("skip-windings", 1, "How many windings to skip on the inside")
)

func pointOnCircle(center [2]float64, radius float64, angle float64) [2]float64 {
	x := center[0] + (radius * math.Cos(angle))
	y := center[1] + (radius * math.Sin(angle))
	return [2]float64{x, y}
}

func pointOnCircleDegrees(center [2]float64, radius float64, angle float64) [2]float64 {
	return pointOnCircle(center, radius, angle*math.Pi/180)
}

func main() {
	flag.Parse()

	if flag.NArg() < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <module-name> <trace-thickness> <trace-clearance> <windings>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	m := kcgen.NewModuleBuilder(flag.Arg(0), "1D magnet grid.", kcgen.LayerFrontCopper)

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
	lastPoint := [2]float64{startRadius, 0}
	radiusIncrement := (thickness + clearance) / 360 / float64(*resolution)

	for degree := 0; degree < 360*(*resolution)*windings; degree++ {
		next := pointOnCircleDegrees([2]float64{}, startRadius+radiusIncrement*float64(degree), float64(degree)/float64(*resolution))
		l := kcgen.NewLine(kcgen.LayerFrontCopper)
		l.Start(lastPoint[0], lastPoint[1])
		l.End(next[0], next[1])
		l.Width(thickness)
		m.AddLine(l)
		lastPoint = next
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
