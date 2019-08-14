package kcgen

import (
	"bytes"
	"testing"
)

func TestModule(t *testing.T) {
	m := NewModuleBuilder("test:label", "A test module", LayerFrontCopper)
	m.Attributes([]string{"virtual"})

	txt := NewText("kek", LayerFrontSilkscreen)
	txt.Position(0, 4, 0)
	m.AddText(txt)
	m.mod.Tedit = ""

	var out bytes.Buffer
	expect := "(module test:label (layer F.Cu)\n \n  (descr \"A test module\")\n  (attr virtual)\n  (fp_text user kek (at 0 4) (layer F.SilkS)\n    (effects (font (size 1 1) (thickness 0.15)))\n  )\n)"
	if err := m.Write(&out); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}
	if !bytes.Equal(out.Bytes(), []byte(expect)) {
		t.Error("Output mismatch")
		t.Logf("got  = %q", string(out.Bytes()))
		t.Logf("want = %q", expect)
	}
}

func TestModuleCircle(t *testing.T) {
	m := NewModuleBuilder("test:circle", "A test module", LayerFrontCopper)
	m.Attributes([]string{"virtual"})

	c := NewCircle(2, LayerFrontFab)
	c.Center(2, 4)
	m.AddCircle(c)
	m.mod.Tedit = ""

	var out bytes.Buffer
	expect := "(module test:circle (layer F.Cu)\n \n  (descr \"A test module\")\n  (attr virtual)\n  (fp_circle (center 2 4) (end 4 4) (layer F.Fab) (width 0.15))\n)"
	if err := m.Write(&out); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}
	if !bytes.Equal(out.Bytes(), []byte(expect)) {
		t.Error("Output mismatch")
		t.Logf("got  = %q", string(out.Bytes()))
		t.Logf("want = %q", expect)
	}
}

func TestModuleLine(t *testing.T) {
	m := NewModuleBuilder("test:line", "A test module", LayerFrontCopper)
	m.Attributes([]string{"virtual"})

	l := NewLine(LayerFrontSilkscreen)
	l.Width(0.2)
	l.Positions(2, 4, 4, 8)
	l.End(4, 5)
	m.AddLine(l)
	m.mod.Tedit = ""

	var out bytes.Buffer
	expect := "(module test:line (layer F.Cu)\n \n  (descr \"A test module\")\n  (attr virtual)\n  (fp_line (start 2 4) (end 4 5) (layer F.SilkS) (width 0.2))\n)"
	if err := m.Write(&out); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}
	if !bytes.Equal(out.Bytes(), []byte(expect)) {
		t.Error("Output mismatch")
		t.Logf("got  = %q", string(out.Bytes()))
		t.Logf("want = %q", expect)
	}
}

func TestModulePoly(t *testing.T) {
	m := NewModuleBuilder("test:poly", "A test module", LayerFrontCopper)
	m.Attributes([]string{"virtual"})

	p := NewPolygon([][2]float64{{1, 1}, {3, 3}, {1, 1}}, 0.2, LayerFrontFab)
	m.AddPolygon(p)
	m.mod.Tedit = ""

	var out bytes.Buffer
	expect := "(module test:poly (layer F.Cu)\n \n  (descr \"A test module\")\n  (attr virtual)\n  (fp_poly (pts (xy 1 1) (xy 3 3) (xy 1 1)) (layer F.Fab) (width 0.2))\n)"
	if err := m.Write(&out); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}
	if !bytes.Equal(out.Bytes(), []byte(expect)) {
		t.Error("Output mismatch")
		t.Logf("got  = %q", string(out.Bytes()))
		t.Logf("want = %q", expect)
	}
}

func TestModuleArc(t *testing.T) {
	m := NewModuleBuilder("test:arc", "A test module", LayerFrontCopper)
	m.Attributes([]string{"virtual"})

	a := NewArc(LayerFrontFab)
	a.Width(0.3)
	a.Positions(0, 0, 2, 1)
	a.Angle(60)
	m.AddArc(a)
	m.mod.Tedit = ""

	var out bytes.Buffer
	expect := "(module test:arc (layer F.Cu)\n \n  (descr \"A test module\")\n  (attr virtual)\n  (fp_arc (start 0 0) (end 2 1) (angle 600) (layer F.Fab) (width 0.3))\n)"
	if err := m.Write(&out); err != nil {
		t.Fatalf("Write() failed: %v", err)
	}
	if !bytes.Equal(out.Bytes(), []byte(expect)) {
		t.Error("Output mismatch")
		t.Logf("got  = %q", string(out.Bytes()))
		t.Logf("want = %q", expect)
	}
}
