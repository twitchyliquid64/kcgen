package netparse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nsf/sexp"
)

// Component describes a component in the netlist.
type Component struct {
	Ref, Value string
	Footprint  string
	TStamp     string
}

// Net describes a net in the netlist.
type Net struct {
	Name string
	Code int
}

// Netlist represents the parsed contents of a netlist file.
type Netlist struct {
	Version    string
	Components []Component
	Nets       []Net
}

// DecodeFile reads a .sch file at fpath, returning a parsed representation.
func DecodeFile(fpath string) (*Netlist, error) {
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
	if mainAST.Children.Value != "export" {
		return nil, errors.New("invalid format: missing leading element export")
	}

	nl := &Netlist{}

	for i := 1; i < mainAST.NumChildren(); i++ {
		n := sexp.Help(mainAST).Child(i)
		if n.IsList() && n.Child(1).IsValid() {
			switch n.Child(0).MustString() {
			case "version":
				nl.Version, err = n.Child(1).String()
				if err != nil {
					return nil, errors.New("invalid format: version value[1] must be a string")
				}
			case "components":
				for x := 1; x < n.MustNode().NumChildren(); x++ {
					c := n.Child(x)
					comp, err := parseComponent(c)
					if err != nil {
						return nil, err
					}
					nl.Components = append(nl.Components, comp)
				}
			case "nets":
				for x := 1; x < n.MustNode().NumChildren(); x++ {
					c := n.Child(x)
					net, err := parseNet(c)
					if err != nil {
						return nil, err
					}
					nl.Nets = append(nl.Nets, net)
				}
			}
		}
	}

	return nl, nil
}

func parseComponent(c sexp.Helper) (Component, error) {
	ident, err := c.Child(0).String()
	if err != nil {
		return Component{}, err
	}
	if ident != "comp" {
		return Component{}, errors.New("invalid format: component must be 'comp'")
	}

	out := Component{}
	for x := 1; x < c.MustNode().NumChildren(); x++ {
		c2 := c.Child(x)
		switch c2.Child(0).MustString() {
		case "code":
			out.Ref, err = c2.Child(1).String()
			if err != nil {
				return Component{}, fmt.Errorf("ref value: %v", err)
			}
		case "value":
			out.Value, err = c2.Child(1).String()
			if err != nil {
				return Component{}, fmt.Errorf("value value: %v", err)
			}
		case "footprint":
			out.Footprint, err = c2.Child(1).String()
			if err != nil {
				return Component{}, fmt.Errorf("footprint value: %v", err)
			}
		case "tstamp":
			out.TStamp, err = c2.Child(1).String()
			if err != nil {
				return Component{}, fmt.Errorf("tstamp value: %v", err)
			}
		}
	}
	return out, nil
}

func parseNet(c sexp.Helper) (Net, error) {
	ident, err := c.Child(0).String()
	if err != nil {
		return Net{}, err
	}
	if ident != "net" {
		return Net{}, errors.New("invalid format: net must be 'net'")
	}

	out := Net{}
	for x := 1; x < c.MustNode().NumChildren(); x++ {
		c2 := c.Child(x)
		switch c2.Child(0).MustString() {
		case "code":
			out.Code, err = c2.Child(1).Int()
			if err != nil {
				return Net{}, fmt.Errorf("code value: %v", err)
			}
		case "name":
			out.Name, err = c2.Child(1).String()
			if err != nil {
				return Net{}, fmt.Errorf("name value: %v", err)
			}
		}
	}
	return out, nil
}
