package textpoly

import (
	"testing"

	"golang.org/x/image/math/fixed"
)

func TestBasic(t *testing.T) {
	v, err := NewVectorizer("RobotoMono-Bold.ttf", 12, 72)
	if err != nil {
		t.Fatal(err)
	}

	err = v.DrawString("C", fixed.Point26_6{})
	if err != nil {
		t.Error(err)
	}
	if len(v.Vectors()[0]) != 65 {
		t.Errorf("len(v.Vectors()) = %d, want 215", len(v.Vectors()))
	}
}
