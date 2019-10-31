// Package textpoly renders text to a series of closed polygons.
package textpoly

import (
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Most of the code is derived from freetype-go.
//
// Copyright 2010 The Freetype-Go Authors. All rights reserved.
// Use of this source code is governed by your choice of either the
// FreeType License or the GNU General Public License version 2 (or
// any later version), further info can be found in the LICENSE file.

type TextVectorizer struct {
	f         *truetype.Font
	glyphBuf  truetype.GlyphBuf
	runOffset fixed.Point26_6

	// fontSize and dpi are used to calculate scale. scale is the number of
	// 26.6 fixed point units in 1 em.
	fontSize, dpi float64
	scale         fixed.Int26_6

	paths []path
}

// Vectors returns the set of polygons representing the text runs.
func (v *TextVectorizer) Vectors() [][][2]float64 {
	out := make([][][2]float64, 0, len(v.paths))

	for _, path := range v.paths {
		pOut := make([][2]float64, 0, len(path.segments)*4)
		offX, offY := float64(path.glyphOffset.X)/64, float64(path.glyphOffset.Y)/64

		for _, v := range path.Rasterize() {
			x, y := float64(v.X)/64, float64(v.Y)/64
			pOut = append(pOut, [2]float64{x + offX, y + offY})
		}
		out = append(out, pOut)
	}
	return out
}

// drawContour draws the given closed contour with the given offset.
func (v *TextVectorizer) drawContour(ps []truetype.Point, dx, dy fixed.Int26_6, glyphOffset fixed.Point26_6) {
	if len(ps) == 0 {
		return
	}

	// The low bit of each point's Flags value is whether the point is on the
	// curve. Truetype fonts only have quadratic Bézier curves, not cubics.
	// Thus, two consecutive off-curve points imply an on-curve point in the
	// middle of those two.
	//
	// See http://chanae.walon.org/pub/ttf/ttf_glyphs.htm for more details.

	// ps[0] is a truetype.Point measured in FUnits and positive Y going
	// upwards. start is the same thing measured in fixed point units and
	// positive Y going downwards, and offset by (dx, dy).
	start := fixed.Point26_6{
		X: dx + ps[0].X,
		Y: dy - ps[0].Y,
	}
	others := []truetype.Point(nil)
	if ps[0].Flags&0x01 != 0 {
		others = ps[1:]
	} else {
		last := fixed.Point26_6{
			X: dx + ps[len(ps)-1].X,
			Y: dy - ps[len(ps)-1].Y,
		}
		if ps[len(ps)-1].Flags&0x01 != 0 {
			start = last
			others = ps[:len(ps)-1]
		} else {
			start = fixed.Point26_6{
				X: (start.X + last.X) / 2,
				Y: (start.Y + last.Y) / 2,
			}
			others = ps
		}
	}

	pth := path{start: start, glyphOffset: glyphOffset, segments: make([]segment, 0, 12)}
	q0, on0 := start, true
	for _, p := range others {
		q := fixed.Point26_6{
			X: dx + p.X,
			Y: dy - p.Y,
		}
		on := p.Flags&0x01 != 0
		if on {
			if on0 {
				pth.segments = append(pth.segments, segment{kind: SegmentLinear, p1: q})
			} else {
				pth.segments = append(pth.segments, segment{kind: SegmentQuadratic, p1: q0, p2: q})
			}
		} else {
			if on0 {
				// No-op.
			} else {
				mid := fixed.Point26_6{
					X: (q0.X + q.X) / 2,
					Y: (q0.Y + q.Y) / 2,
				}
				pth.segments = append(pth.segments, segment{kind: SegmentQuadratic, p1: q0, p2: mid})
			}
		}
		q0, on0 = q, on
	}
	// Close the curve.
	if on0 {
		pth.segments = append(pth.segments, segment{kind: SegmentLinear, p1: start})
	} else {
		pth.segments = append(pth.segments, segment{kind: SegmentQuadratic, p1: q0, p2: start})
	}

	v.paths = append(v.paths, pth)
}

// DrawString draws the string at point p.
func (v *TextVectorizer) DrawString(s string, p fixed.Point26_6) error {
	prev, hasPrev := truetype.Index(0), false
	for _, rune := range s {
		index := v.f.Index(rune)
		if hasPrev {
			kern := v.f.Kern(v.scale, prev, index)
			p.X += kern
		}

		// Split p.X and p.Y into their fractional parts.
		fx, fy := p.X&0x3f, p.Y&0x3f

		if err := v.glyphBuf.Load(v.f, v.scale, index, font.HintingNone); err != nil {
			return err
		}
		// Calculate the integer-pixel bounds for the glyph.
		xmin := int(fx+v.glyphBuf.Bounds.Min.X) >> 6
		ymin := int(fy-v.glyphBuf.Bounds.Max.Y) >> 6
		xmax := int(fx+v.glyphBuf.Bounds.Max.X+0x3f) >> 6
		ymax := int(fy-v.glyphBuf.Bounds.Min.Y+0x3f) >> 6
		if xmin > xmax || ymin > ymax {
			return errors.New("negative sized glyph")
		}
		// A TrueType's glyph's nodes can have negative co-ordinates, but the
		// rasterizer clips anything left of x=0 or above y=0. xmin and ymin are
		// the pixel offsets, based on the font's FUnit metrics, that let a
		// negative co-ordinate in TrueType space be non-negative in rasterizer
		// space. xmin and ymin are typically <= 0.
		fx -= fixed.Int26_6(xmin << 6)
		fy -= fixed.Int26_6(ymin << 6)
		e0 := 0
		for _, e1 := range v.glyphBuf.Ends {
			v.drawContour(v.glyphBuf.Points[e0:e1], fx, fy, p)
			e0 = e1
		}

		p.X += v.glyphBuf.AdvanceWidth

		prev, hasPrev = index, true
	}
	return nil
}

// NewVectorizer creates a vectorizer using the truetype font at the given
// path and the given font size.
func NewVectorizer(fontPath string, fs, dpi float64) (*TextVectorizer, error) {
	d, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(d)
	if err != nil {
		return nil, err
	}

	return &TextVectorizer{
		f:        f,
		fontSize: fs,
		dpi:      dpi,
		scale:    fixed.Int26_6(fs * dpi * (64.0 / 72.0)),
		paths:    make([]path, 0, 6),
	}, nil
}

// maxAbs returns the maximum of abs(a) and abs(b).
func maxAbs(a, b fixed.Int26_6) fixed.Int26_6 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a < b {
		return b
	}
	return a
}

type path struct {
	start       fixed.Point26_6
	glyphOffset fixed.Point26_6
	segments    []segment
}

func (p *path) Rasterize() []fixed.Point26_6 {
	var (
		splitScale = 32
		points     = make([]fixed.Point26_6, 0, 64)
	)
	points = append(points, p.start)

	for _, seg := range p.segments {
		last := points[len(points)-1]

		switch seg.kind {
		case SegmentLinear:
			points = append(points, seg.p1)
		case SegmentQuadratic:
			// Calculate nSplit (the number of recursive decompositions) based on how
			// 'curvy' it is. Specifically, how much the middle point b deviates from
			// (a+c)/2.
			dev := maxAbs(last.X-2*seg.p1.X+seg.p2.X, last.Y-2*seg.p1.Y+seg.p2.Y) / fixed.Int26_6(splitScale)
			nsplit := 0
			for dev > 0 {
				dev /= 4
				nsplit++
			}
			// dev is 32-bit, and nsplit++ every time we shift off 2 bits, so maxNsplit
			// is 16.
			const maxNsplit = 16
			if nsplit > maxNsplit {
				panic("nsplit too large: " + strconv.Itoa(nsplit))
			}
			// Recursively decompose the curve nSplit levels deep.
			var (
				pStack [2*maxNsplit + 3]fixed.Point26_6
				sStack [maxNsplit + 1]int
				i      int
			)
			sStack[0] = nsplit
			pStack[0] = seg.p2
			pStack[1] = seg.p1
			pStack[2] = last
			for i >= 0 {
				s := sStack[i]
				p := pStack[2*i:]
				if s > 0 {
					// Split the quadratic curve p[:3] into an equivalent set of two
					// shorter curves: p[:3] and p[2:5]. The new p[4] is the old p[2],
					// and p[0] is unchanged.
					mx := p[1].X
					p[4].X = p[2].X
					p[3].X = (p[4].X + mx) / 2
					p[1].X = (p[0].X + mx) / 2
					p[2].X = (p[1].X + p[3].X) / 2
					my := p[1].Y
					p[4].Y = p[2].Y
					p[3].Y = (p[4].Y + my) / 2
					p[1].Y = (p[0].Y + my) / 2
					p[2].Y = (p[1].Y + p[3].Y) / 2
					// The two shorter curves have one less split to do.
					sStack[i] = s - 1
					sStack[i+1] = s - 1
					i++
				} else {
					// Replace the level-0 quadratic with a two-linear-piece
					// approximation.
					midx := (p[0].X + 2*p[1].X + p[2].X) / 4
					midy := (p[0].Y + 2*p[1].Y + p[2].Y) / 4
					points = append(points, fixed.Point26_6{X: midx, Y: midy})
					points = append(points, p[0])
					i--
				}
			}

		}
	}

	return points
}

type SegmentKind uint8

const (
	InvalidSegment SegmentKind = iota
	SegmentLinear
	SegmentQuadratic
)

type segment struct {
	kind   SegmentKind
	p1, p2 fixed.Point26_6
}
