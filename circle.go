package kcgen

import (
	"github.com/twitchyliquid64/kcgen/pcb"
)

// func pointOnCircle(center Point2D, radius float64, angle float64) *Point2D {
// 	x := center.X + (radius * math.Cos(angle))
// 	y := center.Y + (radius * math.Sin(angle))
// 	return &Point2D{X: x, Y: y}
// }
//
// func pointOnCircleDegrees(center Point2D, radius float64, angle float64) *Point2D {
// 	return pointOnCircle(center, radius, angle*math.Pi/180)
// }

// Circle represents a circle drawn on a module or PCB.
type Circle struct {
	Layer  Layer
	Center pcb.XY
	Radius float64
	Width  float64
}

// Arc represents an arc drawn on a module or PCB.
type Arc struct {
	Layer          Layer
	Center         pcb.XY
	Start, End     float64
	StepsPerDegree int
	Width, Radius  float64
}
