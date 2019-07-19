package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	pip "github.com/JamesMilnerUK/pip-go"
	"github.com/twitchyliquid64/kcgen/pcb"
)

var (
	viaDrill     = flag.Float64("via-drill-mm", 0.4, "Size of the via drill hole in mm")
	viaSize      = flag.Float64("via-size-mm", 0.6, "Size of the via annular ring in mm")
	minClearance = flag.Float64("min-clearance", 0.2, "Minimum spacing between via and edges of zones")
	separation   = flag.Float64("separation", 1.5, "Space between stitching vias")
	netName      = flag.String("net-name", "GND", "Net name to stitch vias for")
	strategy     = flag.String("placement-strategy", "grid", "grid/alternating")
)

func checkPCB(pcb *pcb.PCB) {
	if len(pcb.Zones) < 2 {
		fmt.Fprintf(os.Stderr, "Error: Must have at least 2 zones to via stitch between zones!\n")
		os.Exit(1)
	}

	// check there are at least two zones, two zones for us to stitch
	// and that the zones are filled.
	foundMatchingName := 0
	for _, zone := range pcb.Zones {
		if zone.NetName == *netName {
			foundMatchingName++
			if len(zone.Polys) == 0 {
				fmt.Fprintf(os.Stderr, "Error: Zone on layer %s has not been filled! Make sure you fill all zones and save before running this tool.\n", zone.Layer)
				os.Exit(1)
			}
		}
	}
	if foundMatchingName < 2 {
		fmt.Fprintf(os.Stderr, "Error: Must have at least 2 zones on the net %q for stitching!\n", *netName)
		os.Exit(1)
	}
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

func serialize(vias []pcb.Via) string {
	out := ""
	for _, v := range vias {
		out += fmt.Sprintf("  (via (at %s %s) (size %s) (drill %s) (layers %s) (net %d))\n", f(v.X), f(v.Y), f(v.Size), f(v.Drill), strings.Join(v.Layers, " "), v.NetIndex)
	}
	return out
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <kicad-mod file>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	pcbF, err := pcb.DecodeFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}

	checkPCB(pcbF)

	netIndex := 0
	for ind := range pcbF.Nets {
		if pcbF.Nets[ind].Name == *netName {
			netIndex = ind
			break
		}
	}

	geo, err := buildGeometry(*netName, pcbF)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Print("Generating possible points...")
	pts := geo.possiblePoints(*strategy, *separation)
	fmt.Println("DONE.")

	fmt.Print("Filtering invalid points....")
	var finalVias []pcb.Via
	for _, pt := range pts {
		ok, layers := geo.testPoint(pip.Point{X: pt[0], Y: pt[1]}, *minClearance+*viaSize)
		if ok {
			finalVias = append(finalVias, pcb.Via{
				X:        pt[0],
				Y:        pt[1],
				Size:     *viaSize,
				Drill:    *viaDrill,
				Layers:   layers,
				NetIndex: netIndex,
			})
		}
	}
	fmt.Println("DONE.")

	fmt.Println(serialize(finalVias))
}
