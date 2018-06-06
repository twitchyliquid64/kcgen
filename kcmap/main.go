package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	geo "github.com/kellydunn/golang-geo"
	geojson "github.com/paulmach/go.geojson"
	"github.com/twitchyliquid64/kcgen"
)

var (
	ref    = flag.String("reference", "map", "Module reference")
	output = flag.String("o", "-", "Where output is written")

	lat1   = flag.Float64("lat1", 37.79393471716305, "Bounding latitude")
	lng1   = flag.Float64("lng1", -122.40305337372274, "Bounding longitude")
	lat2   = flag.Float64("lat2", 37.78525259209928, "Bounding latitude")
	lng2   = flag.Float64("lng2", -122.38693866196127, "Bounding longitude")
	width  = flag.Float64("width", 98, "Output size")
	height = flag.Float64("height", 98, "Output size")
	angle  = flag.Float64("angle", 0, "Angle to rotate the map at")
)

func mapCoordinates(lat, lng float64) (float64, float64) {
	latDiff := -math.Abs(*lat1 - *lat2)
	latMin := math.Min(*lat1, *lat2)
	lngDiff := math.Abs(*lng1 - *lng2)
	lngMin := math.Min(*lng1, *lng2)

	latUnitScaled := (lat - latMin) / latDiff
	lat = (latUnitScaled * (*height)) + *height/2
	lngUnitScaled := (lng - lngMin) / lngDiff
	lng = (lngUnitScaled * (*width)) - *width/2
	radianAngle := (*angle) * math.Pi / 180
	return math.Sin(radianAngle)*lng + math.Cos(radianAngle)*lat, math.Cos(radianAngle)*lng - math.Sin(radianAngle)*lat
}

func withinBounds(lat, lng float64) bool {
	return geo.NewPolygon([]*geo.Point{
		geo.NewPoint(*lat1, *lng1),
		geo.NewPoint(*lat1, *lng2),
		geo.NewPoint(*lat2, *lng2),
		geo.NewPoint(*lat2, *lng1),
	}).Contains(geo.NewPoint(lat, lng))
}

func boundLine(line *kcgen.Line) {
	if line.Start.X == line.End.X {
		panic("infinite slope not yet supported")
	}
	slope := (line.End.Y - line.Start.Y) / (line.End.X - line.Start.X)
	b := line.End.Y - (slope * line.End.X) //y = mx + b which is equivalent to b = y - mx
	if line.Start.X < (-*width/2) {
		line.Start.Y = (slope * (-*width/2)) + b
		line.Start.X = -*width/2
	}
	if line.End.X < (-*width/2) {
		line.End.Y = (slope * (-*width/2)) + b
		line.End.X = -*width/2
	}
	if line.Start.X > (*width/2) {
		line.Start.Y = (slope * (*width/2)) + b
		line.Start.X = *width/2
	}
	if line.End.X > (*width/2) {
		line.End.Y = (slope * (*width/2)) + b
		line.End.X = *width/2
	}

	if line.Start.Y < (-*height/2) {
		line.Start.Y = -*height/2
		line.Start.X = ((-*height/2) - b) / slope //y = mx + b equiv. (y-b)/m = x
	}
	if line.End.Y < (-*height/2) {
		line.End.Y = -*height/2
		line.End.X = ((-*height/2) - b) / slope //y = mx + b equiv. (y-b)/m = x
	}
	if line.Start.Y > (*height/2) {
		line.Start.Y = *height/2
		line.Start.X = ((*height/2) - b) / slope //y = mx + b equiv. (y-b)/m = x
	}
	if line.End.Y > (*height/2) {
		line.End.Y = *height/2
		line.End.X = ((*height/2) - b) / slope //y = mx + b equiv. (y-b)/m = x
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <path to geojson file>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	d, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(1)
	}

	geo, err := geojson.UnmarshalFeatureCollection(d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding GeoJSON: %v\n", err)
		os.Exit(1)
	}

	fp := kcgen.Footprint{
		ModName: *ref,
	}

	for _, f := range geo.Features {
		if f.Properties["layer"] != "STREETS" {
			continue
		}
		width := 0.2
		if f.Properties["classcode"] == "4" {
			width = 0.35
		}
		if f.Properties["classcode"] == "6" {
			width = 0.1
		}
		for i, p := range f.Geometry.LineString {
			if i > 0 {
				lastPoint := f.Geometry.LineString[i-1]
				if !withinBounds(p[1], p[0]) && !withinBounds(lastPoint[1], lastPoint[0]) {
					continue
				}
				y1, x1 := mapCoordinates(lastPoint[1], lastPoint[0])
				y2, x2 := mapCoordinates(p[1], p[0])
				l := &kcgen.Line{
					Layer: kcgen.LayerFrontSilkscreen,
					Start: kcgen.Point2D{X: x1, Y: y1},
					End:   kcgen.Point2D{X: x2, Y: y2},
					Width: width,
				}
				boundLine(l)
				fp.Add(l)
			}
		}
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
