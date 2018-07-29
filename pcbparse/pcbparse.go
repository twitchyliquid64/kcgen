package pcbparse

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/nsf/sexp"
)

// Layer describes the attributes of a layer.
type Layer struct {
	Name string
	Type string
}

// Net represents a netlist.
type Net struct {
	Name string
}

// Via represents a via.
type Via struct {
	X, Y     float64
	Size     float64
	Drill    float64
	Layers   []string
	NetIndex int
}

// Zone represents a zone.
type Zone struct {
	NetNum  int
	NetName string
	Layer   string

	Polys [][][]float64
}

// PCB represents the parsed contents of a kicad_pcb file.
type PCB struct {
	FormatVersion int
	CreatedBy     struct {
		Tool    string
		Version string
	}

	LayersByName map[string]*Layer
	Layers       map[int]*Layer
	Nets         map[int]Net
	Zones        []Zone
}

// DecodeFile reads a .kicad_pcb file at fpath, returning a parsed representation.
func DecodeFile(fpath string) (*PCB, error) {
	f, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	ast, err := sexp.Parse(strings.NewReader(string(f)), nil)
	if err != nil {
		return nil, err
	}

	if !ast.IsList() {
		return nil, errors.New("invalid format: expected s-expression list at top level")
	}
	if ast.NumChildren() != 1 {
		return nil, errors.New("invalid format: top level list of size 1")
	}
	mainAST, _ := ast.Nth(0)
	if !mainAST.IsList() {
		return nil, errors.New("invalid format: expected s-expression list at 1st level")
	}

	if mainAST.NumChildren() < 5 {
		return nil, errors.New("invalid format: expected at least 5 nodes in main expression")
	}
	if mainAST.Children.Value != "kicad_pcb" {
		return nil, errors.New("invalid format: missing leading element kicad_pcb")
	}

	pcb := &PCB{Layers: map[int]*Layer{}, LayersByName: map[string]*Layer{}, Nets: map[int]Net{}}

	for i := 1; i < mainAST.NumChildren(); i++ {
		n := sexp.Help(mainAST).Child(i)
		if n.IsList() && n.Child(1).IsValid() {
			switch n.Child(0).MustString() {
			case "version":
				pcb.FormatVersion, err = n.Child(1).Int()
				if err != nil {
					return nil, errors.New("invalid format: version value must be an int")
				}
			case "host":
				pcb.CreatedBy.Tool, err = n.Child(1).String()
				if err != nil {
					return nil, errors.New("invalid format: host value[1] must be a string")
				}
				pcb.CreatedBy.Version, err = n.Child(2).String()
				if err != nil {
					return nil, errors.New("invalid format: host value[2] must be a string")
				}
			case "layers":
				for x := 1; x < n.MustNode().NumChildren(); x++ {
					c := n.Child(x)
					num, err2 := c.Child(0).Int()
					if err2 != nil {
						return nil, err
					}
					pcb.Layers[num] = &Layer{
						Name: c.Child(1).MustString(),
						Type: c.Child(2).MustString(),
					}
					pcb.LayersByName[c.Child(1).MustString()] = pcb.Layers[num]
				}
			case "net":
				num, err2 := n.Child(1).Int()
				if err2 != nil {
					return nil, err
				}
				pcb.Nets[num] = Net{Name: n.Child(2).MustString()}

			case "zone":
				z := &Zone{}
				for x := 1; x < n.MustNode().NumChildren(); x++ {
					c := n.Child(x)
					switch c.Child(0).MustString() {
					case "net":
						z.NetNum = c.Child(1).MustInt()
					case "net_name":
						z.NetName = c.Child(1).MustString()
					case "layer":
						z.Layer = c.Child(1).MustString()

					case "filled_polygon":
						var points [][]float64
						for y := 1; y < c.Child(1).MustNode().NumChildren(); y++ {
							pt := c.Child(1).Child(y)
							ptType, err2 := pt.Child(0).String()
							if err2 != nil || ptType != "xy" {
								return nil, errors.New("zone.filled_polygon point is not xy point")
							}
							points = append(points, []float64{pt.Child(1).MustFloat64(), pt.Child(2).MustFloat64()})
						}
						z.Polys = append(z.Polys, points)
					}
				}
				pcb.Zones = append(pcb.Zones, *z)
			}
		}
	}

	return pcb, nil
}
