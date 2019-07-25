package pcb

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nsf/sexp"
)

func TestParseMod(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected Module
	}{
		{
			name: "basic",
			input: `
(module Gauge_50mm_Type2_SilkScreenTop (layer F.Cu)
  (at 0 0)
  (descr "Gauge, Massstab, 50mm, SilkScreenTop, Type 2,")
  (tags "Gauge Massstab 50mm SilkScreenTop Type 2")
  (attr virtual)
)
    `,
			expected: Module{
				Name:        "Gauge_50mm_Type2_SilkScreenTop",
				Layer:       "F.Cu",
				Description: "Gauge, Massstab, 50mm, SilkScreenTop, Type 2,",
				Tags:        []string{"Gauge", "Massstab", "50mm", "SilkScreenTop", "Type", "2"},
				Attrs:       []string{"virtual"},
			},
		},
		{
			name: "text and lines",
			input: `
(module Gauge_50mm_Type2_SilkScreenTop (layer F.Cu)
  (at 10 20)
  (descr "Gauge, Massstab, 50mm, SilkScreenTop, Type 2,")
  (tags "Gauge Massstab 50mm SilkScreenTop Type 2")
  (attr virtual)
  (fp_text reference REF** (at 20.50034 9.99998) (layer F.SilkS)
    (effects (font (size 1 1) (thickness 0.15)))
  )
  (fp_line (start 9.99998 0) (end 9.99998 1.99898) (layer F.SilkS) (width 0.15))
)
    `,
			expected: Module{
				Name:        "Gauge_50mm_Type2_SilkScreenTop",
				Placement:   ModPlacement{At: XYZ{X: 10, Y: 20}},
				Layer:       "F.Cu",
				Description: "Gauge, Massstab, 50mm, SilkScreenTop, Type 2,",
				Tags:        []string{"Gauge", "Massstab", "50mm", "SilkScreenTop", "Type", "2"},
				Attrs:       []string{"virtual"},
				Graphics: []ModGraphic{
					{
						Ident: "fp_text",
						Renderable: &ModText{
							Kind:  RefText,
							Text:  "REF**",
							At:    XYZ{X: 20.50034, Y: 9.99998},
							Layer: "F.SilkS",
							Effects: TextEffects{
								Thickness: 0.15,
								FontSize:  XY{X: 1, Y: 1},
							},
						},
					},
					{
						Ident: "fp_line",
						Renderable: &ModLine{
							Width: 0.15,
							Start: XY{X: 9.99998},
							End:   XY{X: 9.99998, Y: 1.99898},
							Layer: "F.SilkS",
						},
					},
				},
			},
		},
		{
			name: "model",
			input: `
(module Resistors_SMD:R_0805_HandSoldering (layer F.Cu) (tedit 5ADA758D) (tstamp 5AE3D8D3)
  (at 133.604 96.266 90)
  (descr "Resistor SMD 0805, hand soldering")
  (tags "resistor 0805")
  (path /5ADA7384)
  (attr smd)
  (model Resistors_SMD.3dshapes/R_0805_HandSoldering.wrl
    (at (xyz 0 0 0))
    (scale (xyz 1 1 1))
    (rotate (xyz 0 0 0))
  )
)
    `,
			expected: Module{
				Name:        "Resistors_SMD:R_0805_HandSoldering",
				Placement:   ModPlacement{At: XYZ{X: 133.604, Y: 96.266, Z: 90, ZPresent: true}},
				Tedit:       "5ADA758D",
				Tstamp:      "5AE3D8D3",
				Path:        "/5ADA7384",
				Layer:       "F.Cu",
				Description: "Resistor SMD 0805, hand soldering",
				Tags:        []string{"resistor", "0805"},
				Attrs:       []string{"smd"},
				Model: &ModModel{
					Path:   "Resistors_SMD.3dshapes/R_0805_HandSoldering.wrl",
					At:     XYZ{ZPresent: true},
					Scale:  XYZ{X: 1, Y: 1, Z: 1, ZPresent: true},
					Rotate: XYZ{ZPresent: true},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ast, err := sexp.Parse(strings.NewReader(tc.input), nil)
			if err != nil {
				t.Fatal(err)
			}

			mod, err := parseModule(sexp.Help(ast).Child(0), 0)
			if err != nil {
				t.Fatalf("parseModule() failed: %v", err)
			}
			if got, want := mod, &tc.expected; !reflect.DeepEqual(got, want) {
				t.Errorf("mod = %+v, want %+v", got, want)
			}
		})
	}
}
