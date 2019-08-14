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
	output          = flag.String("o", "-", "Where output is written")
)

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

	top := kcgen.NewLine(kcgen.LayerFrontSilkscreen)
	top.Positions(-width/2, height/2, width/2, height/2)
	m.AddLine(top)
	bottom := kcgen.NewLine(kcgen.LayerFrontSilkscreen)
	bottom.Positions(-width/2, -height/2, width/2, -height/2)
	m.AddLine(bottom)

	left := kcgen.NewLine(kcgen.LayerFrontSilkscreen)
	left.Positions(-width/2, height/2, -width/2, -height/2)
	m.AddLine(left)
	right := kcgen.NewLine(kcgen.LayerFrontSilkscreen)
	right.Positions(width/2, height/2, width/2, -height/2)
	m.AddLine(right)

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
