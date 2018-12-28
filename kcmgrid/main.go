package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	netIndex    = flag.Int("stride_net_start_index", 1, "Starting net index for strides")
	numStrides  = flag.Int("num_strides", 4, "Number of strides")
	numElements = flag.Int("num_elements", 16, "Number of grid elements")
	usableWidth = flag.Int("usable_width", 15, "Usable area in mm")

	viaDrill       = flag.Float64("via_drill", 0.4, "Size of the via drill hole in mm")
	viaSize        = flag.Float64("via_size", 0.6, "Size of the via annular ring in mm")
	minClearance   = flag.Float64("min_clearance", 0.2, "Minimum spacing between via and edges of zones")
	traceThickness = flag.Float64("trace_thicc", 0.18, "Trace thickness")
	bottomLayer    = flag.Bool("bottom_layer", false, "Put the grid on the bottom layer")
)

func viaSeparation() float64 {
	return *viaSize/2 + *minClearance
}
func wastedArea() float64 {
	return viaSeparation() * float64(*numStrides)
}
func elementSeparation() float64 {
	return math.Max(*traceThickness+*minClearance, viaSeparation())
}
func layer(isBottomLayer bool) string {
	if isBottomLayer {
		return "B.Cu"
	}
	return "F.Cu"
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
func emitVia(x, y float64, net int) {
	fmt.Printf("  (via (at %s %s) (size %s) (drill %s) (layers %s) (net %d))\n", f(x), f(y), f(*viaSize), f(*viaDrill), "F.Cu B.Cu", net)
}
func emitTrace(x1, y1, x2, y2 float64, net int, inverseLayer bool) {
	//(segment (start 68.326 128.016) (end 92.202 128.016) (width 0.25) (layer B.Cu) (net 1))
	fmt.Printf("  (segment (start %s %s) (end %s %s) (width %s) (layer %s) (net %d))\n", f(x1), f(y1), f(x2), f(y2), f(*traceThickness), layer(inverseLayer), net)
}

func main() {
	flag.Parse()

	marginSize := wastedArea()
	for i := 0; i < *numElements; i++ {
		stride := i % *numStrides

		viaX := viaSeparation() * float64(stride)
		Y := float64(i) * elementSeparation()
		emitVia(viaX, Y, *netIndex+stride)

		emitTrace(viaX, Y, marginSize, Y, *netIndex+stride, false)
		emitTrace(marginSize, Y, float64(*usableWidth)+marginSize, Y, *netIndex+stride, false)
		endX := float64(*usableWidth) + marginSize*2 - viaX
		emitTrace(float64(*usableWidth)+marginSize, Y, endX, Y, *netIndex+stride, false)
		emitVia(endX, Y, *netIndex+stride)

		// If there is another stride of the same net to come ...
		if (i + *numStrides) < *numElements {
			// Every 1/2 strides should connect the lower stride at the end/beginning.
			if ((i / (*numStrides)) % 2) == 0 {
				emitTrace(endX, Y, endX, Y+(float64(*numStrides)*elementSeparation()), *netIndex+stride, true)
			} else {
				emitTrace(viaX, Y, viaX, Y+(float64(*numStrides)*elementSeparation()), *netIndex+stride, true)
			}
		}
	}
}
