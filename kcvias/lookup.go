package main

import (
	"math"

	pip "github.com/JamesMilnerUK/pip-go"
	"github.com/twitchyliquid64/kcgen/pcbparse"
)

func pointOnCircle(center pip.Point, radius float64, angle float64) *pip.Point {
	x := center.X + (radius * math.Cos(angle))
	y := center.Y + (radius * math.Sin(angle))
	return &pip.Point{X: x, Y: y}
}

func pointOnCircleDegrees(center pip.Point, radius float64, angle float64) *pip.Point {
	return pointOnCircle(center, radius, angle*math.Pi/180)
}

type zone struct {
	sections []pip.Polygon
	layer    string
}

func (p *zone) pointInPoly(point pip.Point) bool {
	for _, section := range p.sections {
		if pip.PointInPolygon(point, section) {
			return true
		}
	}
	return false
}

type geometry struct {
	zones    []zone
	smallest pip.Point
	largest  pip.Point
}

func (g *geometry) possiblePoints(strategy string, separation float64) (out [][]float64) {
	switch strategy {
	case "grid":
		for x := g.smallest.X; x < g.largest.X; x += separation {
			for y := g.smallest.Y; y < g.largest.Y; y += separation {
				out = append(out, []float64{x, y})
			}
		}
	}
	return
}

// testPoint returns true if:
//  - There is at least clearance distance on 12 cardinal points on all candidate zones
//  - There are at least two zones on different layers comprising the candidate point
func (g *geometry) testPoint(p pip.Point, clearance float64) (bool, []string) {
	seenLayers := map[string]bool{}
	var candidateZones []zone
	for _, zone := range g.zones {
		if zone.pointInPoly(p) {
			candidateZones = append(candidateZones, zone)
			seenLayers[zone.layer] = true
		}
	}
	if len(candidateZones) < 2 {
		return false, nil // at least two zones must be at the point to stitch at this point.
	}
	if len(seenLayers) < 2 {
		return false, nil // zones must stitch across layers, otherwise we aint stitching anything.
	}

	// at this stage, candidateZones contains at least two zones that have overlap
	// at point p, and are on at least two separate layers.
	// we test if there is still overlap at clearance distance away from the point
	// in 12 directions.
	seenLayers = map[string]bool{}
	var perfectZones []zone
	for _, candidate := range candidateZones {
		match := false
		for x := 0; x <= 360; x += (360 / 12) {
			p2 := pointOnCircleDegrees(p, clearance, float64(x))
			if !candidate.pointInPoly(*p2) {
				match = true
				break
			}
		}
		if !match {
			seenLayers[candidate.layer] = true
			perfectZones = append(perfectZones, candidate)
		}
	}

	// no longer across 2 layers.
	if len(seenLayers) < 2 {
		return false, nil
	}
	if len(perfectZones) < 2 {
		return false, nil
	}

	var layers []string
	for l := range seenLayers {
		layers = append(layers, l)
	}
	return true, layers
}

func buildGeometry(wantNet string, pcb *pcbparse.PCB) (*geometry, error) {
	g := &geometry{smallest: pip.Point{999999, 999999}, largest: pip.Point{-9999999, -9999999}}

	for _, z := range pcb.Zones {
		if z.NetName == wantNet {
			p := zone{layer: z.Layer}
			for _, section := range z.Polys {
				pg := pip.Polygon{}
				for _, point := range section {
					pg.Points = append(pg.Points, pip.Point{X: point[0], Y: point[1]})

					if point[0] < g.smallest.X {
						g.smallest.X = point[0]
					}
					if point[1] < g.smallest.Y {
						g.smallest.Y = point[1]
					}
					if point[0] > g.largest.X {
						g.largest.X = point[0]
					}
					if point[1] > g.largest.Y {
						g.largest.Y = point[1]
					}
				}
				p.sections = append(p.sections, pg)
			}
			g.zones = append(g.zones, p)
		}
	}

	return g, nil
}
