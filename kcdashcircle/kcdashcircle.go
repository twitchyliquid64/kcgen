package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/twitchyliquid64/kcgen"
)

var (
	referenceOffset = flag.Float64("refY", 2, "Y-axis offset at which module designator is placed")
	centerX         = flag.Float64("centerX", 0, "Center point X")
	centerY         = flag.Float64("centerY", 0, "Center point Y")
	resolution      = flag.Int("resolution", 0, "How many interpolations to make per degree")
	output          = flag.String("o", "-", "Where output is written")
)

func main() {
	flag.Parse()

	if flag.NArg() < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <module-name> <radius> <num-sections>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	fp := kcgen.Footprint{
		ModName:    flag.Arg(0),
		ReferenceY: *referenceOffset,
	}

	radius, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(1), err)
		os.Exit(2)
	}

	sections, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", flag.Arg(2), err)
		os.Exit(2)
	}

	segmentAngle := 360 / float64(sections) / 2
	for x := 0; x < sections; x++ {
		fp.Add(&kcgen.Arc{
			Layer:          kcgen.LayerFrontSilkscreen,
			Start:          segmentAngle * float64(x) * 2,
			End:            segmentAngle * (float64(x)*2 + 1),
			Center:         kcgen.Point2D{X: *centerX, Y: *centerY},
			StepsPerDegree: *resolution,
			Radius:         radius,
		})
	}

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
	}
}
