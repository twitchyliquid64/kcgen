package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	geo "github.com/kellydunn/golang-geo"
	geojson "github.com/paulmach/go.geojson"
	kcgen "github.com/twitchyliquid64/kcgen/legacy"
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
	startX, startY := line.GetStart()
	endX, endY := line.GetEnd()
	if startX == endX {
		panic("infinite slope not yet supported")
	}
	slope := (endY - startY) / (endX - startX)
	b := endY - (slope * endX) //y = mx + b which is equivalent to b = y - mx

	// Handle X upper and lower bounds.
	if startX < (-*width / 2) {
		startY = (slope * (-*width / 2)) + b
		startX = -*width / 2
		line.Start(startX, startY)
	}
	if endX < (-*width / 2) {
		endY = (slope * (-*width / 2)) + b
		endX = -*width / 2
		line.End(endX, endY)
	}
	if startX > (*width / 2) {
		startY = (slope * (*width / 2)) + b
		startX = *width / 2
		line.Start(startX, startY)
	}
	if endX > (*width / 2) {
		endY = (slope * (*width / 2)) + b
		endX = *width / 2
		line.End(endX, endY)
	}

	// Handle Y upper and lower bounds.
	if startY < (-*height / 2) {
		startY = -*height / 2
		startX = ((-*height / 2) - b) / slope //y = mx + b equiv. (y-b)/m = x
		line.Start(startX, startY)
	}
	if endY < (-*height / 2) {
		endY = -*height / 2
		endX = ((-*height / 2) - b) / slope //y = mx + b equiv. (y-b)/m = x
		line.End(endX, endY)
	}
	if startY > (*height / 2) {
		startY = *height / 2
		startX = ((*height / 2) - b) / slope //y = mx + b equiv. (y-b)/m = x
		line.Start(startX, startY)
	}
	if endY > (*height / 2) {
		endY = *height / 2
		endX = ((*height / 2) - b) / slope //y = mx + b equiv. (y-b)/m = x
		line.End(endX, endY)
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

	m := kcgen.NewModuleBuilder(*ref, "A map.", kcgen.LayerFrontCopper)

	for _, f := range geo.Features {
		if f.Properties["layer"] != "STREETS" { // && f.Properties["layer"] != "PSEUDO" {
			//fmt.Fprintf(os.Stderr, "Layer: %q\n", f.Properties)
			continue
		}
		width := 0.38
		if f.Properties["classcode"] == "4" {
			width = 0.55
		}
		if f.Properties["classcode"] == "6" {
			width = 0.2
		}
		for i, p := range f.Geometry.LineString {
			if i > 0 {
				lastPoint := f.Geometry.LineString[i-1]
				if !withinBounds(p[1], p[0]) && !withinBounds(lastPoint[1], lastPoint[0]) {
					continue
				}
				y1, x1 := mapCoordinates(lastPoint[1], lastPoint[0])
				y2, x2 := mapCoordinates(p[1], p[0])
				l := kcgen.NewLine(kcgen.LayerFrontSilkscreen)
				l.Positions(x1, y1, x2, y2)
				l.Width(width)
				boundLine(l)
				m.AddLine(l)
			}
		}
	}

	// fp.Add(&kcgen.Text{
	// 	Layer:    kcgen.LayerFrontFab,
	// 	Position: kcgen.Point2D{X: 0, Y: 2},
	// 	Text:     fmt.Sprintf("Min Latitude: %v", *lat1),
	// })
	// fp.Add(&kcgen.Text{
	// 	Layer:    kcgen.LayerFrontFab,
	// 	Position: kcgen.Point2D{X: 0, Y: 4},
	// 	Text:     fmt.Sprintf("Max Latitude: %v", *lat2),
	// })
	// fp.Add(&kcgen.Text{
	// 	Layer:    kcgen.LayerFrontFab,
	// 	Position: kcgen.Point2D{X: 0, Y: 6},
	// 	Text:     fmt.Sprintf("Min Longitude: %v", *lng1),
	// })
	// fp.Add(&kcgen.Text{
	// 	Layer:    kcgen.LayerFrontFab,
	// 	Position: kcgen.Point2D{X: 0, Y: 8},
	// 	Text:     fmt.Sprintf("Mac Longitude: %v", *lng2),
	// })

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
