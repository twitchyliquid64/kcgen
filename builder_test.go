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
