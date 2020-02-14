// Package adv implements advanced modifications to PCBs.
package adv

import (
	"fmt"
	"math"

	"github.com/twitchyliquid64/kcgen/pcb"
)

type Region struct {
	From pcb.XY
	To   pcb.XY
}

func (r Region) Size() pcb.XY {
	return pcb.XY{
		X: r.To.X - r.From.X,
		Y: r.To.Y - r.From.Y,
	}
}

func (r Region) LeftBoundary() *pcb.Line {
	return &pcb.Line{
		Start: r.From,
		End:   pcb.XY{X: r.From.X, Y: r.To.Y},
	}
}

func (r Region) RightBoundary() *pcb.Line {
	return &pcb.Line{
		Start: pcb.XY{X: r.To.X, Y: r.From.Y},
		End:   r.To,
	}
}

func (r Region) TopBoundary() *pcb.Line {
	return &pcb.Line{
		Start: r.From,
		End:   pcb.XY{X: r.To.X, Y: r.From.Y},
	}
}

func (r Region) BottomBoundary() *pcb.Line {
	return &pcb.Line{
		Start: r.To,
		End:   pcb.XY{X: r.From.X, Y: r.To.Y},
	}
}

func (r Region) Center() pcb.XY {
	return pcb.XY{X: r.To.X - r.From.X, Y: r.To.Y - r.From.Y}
}

func (r Region) Within(p pcb.XY) bool {
	return r.From.X <= p.X && r.To.X >= p.X &&
		r.From.Y <= p.Y && r.To.Y >= p.Y
}

func MakeRegion(p1, p2 pcb.XY) Region {
	return Region{
		From: pcb.XY{X: math.Min(p1.X, p2.X), Y: math.Min(p1.Y, p2.Y)},
		To:   pcb.XY{X: math.Max(p1.X, p2.X), Y: math.Max(p1.Y, p2.Y)},
	}
}

// Carve cuts out all drawings which are within the given region.
func Carve(p *pcb.PCB, region Region) error {
	var (
		removeIdx   []int
		newDrawings []pcb.Drawing
		err         error
	)

	for i, d := range p.Drawings {
		var remove bool
		switch g := d.(type) {
		case *pcb.Line:
			var newLines []newLinePair
			remove, newLines, err = carveLine(g, region)
			// fmt.Println(g, region, newLines)
			for _, line := range newLines {
				dupe := *g
				dupe.Start = line.Start
				dupe.End = line.End
				newDrawings = append(newDrawings, &dupe)
			}
		default:
			return fmt.Errorf("cannot carve pcb element of type %T", g)
		}

		if err != nil {
			return err
		}
		if remove {
			removeIdx = append(removeIdx, i)
		}
	}

	alreadyRemoved := 0
	for _, idx := range removeIdx {
		idx -= alreadyRemoved
		p.Drawings = append(p.Drawings[:idx], p.Drawings[idx+1:]...)
		alreadyRemoved++
	}
	p.Drawings = append(p.Drawings, newDrawings...)

	return nil
}
