package pcb

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"go.starlark.net/starlark"
)

var MakeXY = starlark.NewBuiltin("XY", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.Float
		f1 starlark.Float
	)
	unpackErr := starlark.UnpackArgs(
		"XY",
		args,
		kwargs,
		"x?", &f0,
		"y?", &f1,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := XY{}

	out.X = float64(f0)
	out.Y = float64(f1)
	return &out, nil
})

func (p *XY) String() string {
	return fmt.Sprintf("XY{%v, %v}", p.X, p.Y)
}

// Type implements starlark.Value.
func (p *XY) Type() string {
	return "XY"
}

// Freeze implements starlark.Value.
func (p *XY) Freeze() {
}

// Truth implements starlark.Value.
func (p *XY) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *XY) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *XY) Attr(name string) (starlark.Value, error) {
	switch name {
	case "x":
		return starlark.Float(p.X), nil

	case "y":
		return starlark.Float(p.Y), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *XY) AttrNames() []string {
	return []string{"x", "y"}
}

// SetField implements starlark.HasSetField.
func (p *XY) SetField(name string, val starlark.Value) error {
	switch name {
	case "x":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to x using type %T", val)
		}
		p.X = float64(v)

	case "y":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to y using type %T", val)
		}
		p.Y = float64(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeXYZ = starlark.NewBuiltin("XYZ", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.Float
		f1 starlark.Float
		f2 starlark.Float
		f3 starlark.Bool
		f4 starlark.Bool
	)
	unpackErr := starlark.UnpackArgs(
		"XYZ",
		args,
		kwargs,
		"x?", &f0,
		"y?", &f1,
		"z?", &f2,
		"z_present?", &f3,
		"unlocked?", &f4,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := XYZ{}

	out.X = float64(f0)
	out.Y = float64(f1)
	out.Z = float64(f2)
	out.ZPresent = bool(f3)
	out.Unlocked = bool(f4)
	return &out, nil
})

func (p *XYZ) String() string {
	return fmt.Sprintf("XYZ{%v, %v, %v, %v, %v}", p.X, p.Y, p.Z, p.ZPresent, p.Unlocked)
}

// Type implements starlark.Value.
func (p *XYZ) Type() string {
	return "XYZ"
}

// Freeze implements starlark.Value.
func (p *XYZ) Freeze() {
}

// Truth implements starlark.Value.
func (p *XYZ) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *XYZ) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *XYZ) Attr(name string) (starlark.Value, error) {
	switch name {
	case "x":
		return starlark.Float(p.X), nil

	case "y":
		return starlark.Float(p.Y), nil

	case "z":
		return starlark.Float(p.Z), nil

	case "z_present":
		return starlark.Bool(p.ZPresent), nil

	case "unlocked":
		return starlark.Bool(p.Unlocked), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *XYZ) AttrNames() []string {
	return []string{"x", "y", "z", "z_present", "unlocked"}
}

// SetField implements starlark.HasSetField.
func (p *XYZ) SetField(name string, val starlark.Value) error {
	switch name {
	case "x":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to x using type %T", val)
		}
		p.X = float64(v)

	case "y":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to y using type %T", val)
		}
		p.Y = float64(v)

	case "z":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to z using type %T", val)
		}
		p.Z = float64(v)

	case "z_present":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to z_present using type %T", val)
		}
		p.ZPresent = bool(v)

	case "unlocked":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to unlocked using type %T", val)
		}
		p.Unlocked = bool(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeVia = starlark.NewBuiltin("Via", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 XY
		f1 starlark.Float
		f2 starlark.Float
		f3 *starlark.List
		f4 starlark.Int
		f5 starlark.String
	)
	unpackErr := starlark.UnpackArgs(
		"Via",
		args,
		kwargs,
		"at?", &f0,
		"size?", &f1,
		"drill?", &f2,
		"layers?", &f3,
		"net_index?", &f4,
		"status_flags?", &f5,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Via{}

	out.At = f0
	out.Size = float64(f1)
	out.Drill = float64(f2)
	if f3 != nil {
		for i := 0; i < f3.Len(); i++ {
			s, ok := f3.Index(i).(starlark.String)
			if !ok {
				return starlark.None, errors.New("layers is not a string")
			}
			out.Layers = append(out.Layers, string(s))
		}
	}

	if v, ok := f4.Int64(); ok {
		out.NetIndex = int(v)
	}
	out.StatusFlags = string(f5)
	return &out, nil
})

func (p *Via) String() string {
	return fmt.Sprintf("Via{%v, %v, %v, %v, %v, %v}", p.At, p.Size, p.Drill, p.Layers, p.NetIndex, p.StatusFlags)
}

// Type implements starlark.Value.
func (p *Via) Type() string {
	return "Via"
}

// Freeze implements starlark.Value.
func (p *Via) Freeze() {
}

// Truth implements starlark.Value.
func (p *Via) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Via) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Via) Attr(name string) (starlark.Value, error) {
	switch name {
	case "at":
		return &p.At, nil

	case "size":
		return starlark.Float(p.Size), nil

	case "drill":
		return starlark.Float(p.Drill), nil

	case "layers":
		l := starlark.NewList(nil)
		for _, e := range p.Layers {
			l.Append(starlark.String(e))
		}
		return l, nil

	case "net_index":
		return starlark.MakeInt(p.NetIndex), nil

	case "status_flags":
		return starlark.String(p.StatusFlags), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Via) AttrNames() []string {
	return []string{"at", "size", "drill", "layers", "net_index", "status_flags"}
}

// SetField implements starlark.HasSetField.
func (p *Via) SetField(name string, val starlark.Value) error {
	switch name {
	case "at":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v

	case "size":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to size using type %T", val)
		}
		p.Size = float64(v)

	case "drill":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to drill using type %T", val)
		}
		p.Drill = float64(v)

	case "layers":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to layers using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(starlark.String)
			if !ok {
				return errors.New("layers is not a string")
			}
			p.Layers = append(p.Layers, string(s))
		}

	case "net_index":
		v, ok := val.(*starlark.Int)
		if !ok {
			return fmt.Errorf("cannot assign to net_index using type %T", val)
		}
		i, ok := v.Int64()
		if !ok {
			return fmt.Errorf("cannot convert %v to int64", v)
		}
		p.NetIndex = int(i)

	case "status_flags":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to status_flags using type %T", val)
		}
		p.StatusFlags = string(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeTrack = starlark.NewBuiltin("Track", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 XY
		f1 XY
		f2 starlark.Float
		f3 starlark.String
		f4 starlark.Int
		f5 starlark.String
		f6 starlark.String
	)
	unpackErr := starlark.UnpackArgs(
		"Track",
		args,
		kwargs,
		"start?", &f0,
		"end?", &f1,
		"width?", &f2,
		"layer?", &f3,
		"net_index?", &f4,
		"tstamp?", &f5,
		"status_flags?", &f6,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Track{}

	out.Start = f0
	out.End = f1
	out.Width = float64(f2)
	out.Layer = string(f3)

	if v, ok := f4.Int64(); ok {
		out.NetIndex = int(v)
	}
	out.Tstamp = string(f5)
	out.StatusFlags = string(f6)
	return &out, nil
})

func (p *Track) String() string {
	return fmt.Sprintf("Track{%v, %v, %v, %v, %v, %v, %v}", p.Start, p.End, p.Width, p.Layer, p.NetIndex, p.Tstamp, p.StatusFlags)
}

// Type implements starlark.Value.
func (p *Track) Type() string {
	return "Track"
}

// Freeze implements starlark.Value.
func (p *Track) Freeze() {
}

// Truth implements starlark.Value.
func (p *Track) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Track) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Track) Attr(name string) (starlark.Value, error) {
	switch name {
	case "start":
		return &p.Start, nil

	case "end":
		return &p.End, nil

	case "width":
		return starlark.Float(p.Width), nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "net_index":
		return starlark.MakeInt(p.NetIndex), nil

	case "tstamp":
		return starlark.String(p.Tstamp), nil

	case "status_flags":
		return starlark.String(p.StatusFlags), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Track) AttrNames() []string {
	return []string{"start", "end", "width", "layer", "net_index", "tstamp", "status_flags"}
}

// SetField implements starlark.HasSetField.
func (p *Track) SetField(name string, val starlark.Value) error {
	switch name {
	case "start":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to start using type %T", val)
		}
		p.Start = *v

	case "end":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to end using type %T", val)
		}
		p.End = *v

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)

	case "net_index":
		v, ok := val.(*starlark.Int)
		if !ok {
			return fmt.Errorf("cannot assign to net_index using type %T", val)
		}
		i, ok := v.Int64()
		if !ok {
			return fmt.Errorf("cannot convert %v to int64", v)
		}
		p.NetIndex = int(i)

	case "tstamp":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tstamp using type %T", val)
		}
		p.Tstamp = string(v)

	case "status_flags":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to status_flags using type %T", val)
		}
		p.StatusFlags = string(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeLayer = starlark.NewBuiltin("Layer", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.Int
		f1 starlark.String
		f2 starlark.String
		f3 starlark.Bool
	)
	unpackErr := starlark.UnpackArgs(
		"Layer",
		args,
		kwargs,
		"num?", &f0,
		"name?", &f1,
		"typ?", &f2,
		"hidden?", &f3,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Layer{}

	if v, ok := f0.Int64(); ok {
		out.Num = int(v)
	}
	out.Name = string(f1)
	out.Typ = string(f2)
	out.Hidden = bool(f3)
	return &out, nil
})

func (p *Layer) String() string {
	return fmt.Sprintf("Layer{%v, %v, %v, %v}", p.Num, p.Name, p.Typ, p.Hidden)
}

// Type implements starlark.Value.
func (p *Layer) Type() string {
	return "Layer"
}

// Freeze implements starlark.Value.
func (p *Layer) Freeze() {
}

// Truth implements starlark.Value.
func (p *Layer) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Layer) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Layer) Attr(name string) (starlark.Value, error) {
	switch name {
	case "num":
		return starlark.MakeInt(p.Num), nil

	case "name":
		return starlark.String(p.Name), nil

	case "type":
		return starlark.String(p.Typ), nil

	case "hidden":
		return starlark.Bool(p.Hidden), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Layer) AttrNames() []string {
	return []string{"num", "name", "type", "hidden"}
}

// SetField implements starlark.HasSetField.
func (p *Layer) SetField(name string, val starlark.Value) error {
	switch name {
	case "num":
		v, ok := val.(*starlark.Int)
		if !ok {
			return fmt.Errorf("cannot assign to num using type %T", val)
		}
		i, ok := v.Int64()
		if !ok {
			return fmt.Errorf("cannot convert %v to int64", v)
		}
		p.Num = int(i)

	case "name":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to name using type %T", val)
		}
		p.Name = string(v)

	case "type":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to typ using type %T", val)
		}
		p.Typ = string(v)

	case "hidden":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to hidden using type %T", val)
		}
		p.Hidden = bool(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakePCBCreatedBy = starlark.NewBuiltin("PCBCreatedBy", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.String
		f1 starlark.String
	)
	unpackErr := starlark.UnpackArgs(
		"PCBCreatedBy",
		args,
		kwargs,
		"tool?", &f0,
		"version?", &f1,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := PCBCreatedBy{}

	out.Tool = string(f0)
	out.Version = string(f1)
	return &out, nil
})

func (p *PCBCreatedBy) String() string {
	return fmt.Sprintf("PCBCreatedBy{%v, %v}", p.Tool, p.Version)
}

// Type implements starlark.Value.
func (p *PCBCreatedBy) Type() string {
	return "PCBCreatedBy"
}

// Freeze implements starlark.Value.
func (p *PCBCreatedBy) Freeze() {
}

// Truth implements starlark.Value.
func (p *PCBCreatedBy) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *PCBCreatedBy) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *PCBCreatedBy) Attr(name string) (starlark.Value, error) {
	switch name {
	case "tool":
		return starlark.String(p.Tool), nil

	case "version":
		return starlark.String(p.Version), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *PCBCreatedBy) AttrNames() []string {
	return []string{"tool", "version"}
}

// SetField implements starlark.HasSetField.
func (p *PCBCreatedBy) SetField(name string, val starlark.Value) error {
	switch name {
	case "tool":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tool using type %T", val)
		}
		p.Tool = string(v)

	case "version":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to version using type %T", val)
		}
		p.Version = string(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeNetClass = starlark.NewBuiltin("NetClass", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0  starlark.String
		f1  starlark.String
		f2  starlark.Float
		f3  starlark.Float
		f4  starlark.Float
		f5  starlark.Float
		f6  starlark.Float
		f7  starlark.Float
		f8  starlark.Float
		f9  starlark.Float
		f10 *starlark.List
	)
	unpackErr := starlark.UnpackArgs(
		"NetClass",
		args,
		kwargs,
		"name?", &f0,
		"description?", &f1,
		"clearance?", &f2,
		"trace_width?", &f3,
		"via_diameter?", &f4,
		"via_drill?", &f5,
		"u_via_diameter?", &f6,
		"u_via_drill?", &f7,
		"diff_pair_width?", &f8,
		"diff_pair_gap?", &f9,
		"nets?", &f10,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := NetClass{}

	out.Name = string(f0)
	out.Description = string(f1)
	out.Clearance = float64(f2)
	out.TraceWidth = float64(f3)
	out.ViaDiameter = float64(f4)
	out.ViaDrill = float64(f5)
	out.UViaDiameter = float64(f6)
	out.UViaDrill = float64(f7)
	out.DiffPairWidth = float64(f8)
	out.DiffPairGap = float64(f9)
	if f10 != nil {
		for i := 0; i < f10.Len(); i++ {
			s, ok := f10.Index(i).(starlark.String)
			if !ok {
				return starlark.None, errors.New("nets is not a string")
			}
			out.Nets = append(out.Nets, string(s))
		}
	}
	return &out, nil
})

func (p *NetClass) String() string {
	return fmt.Sprintf("NetClass{%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v}", p.Name, p.Description, p.Clearance, p.TraceWidth, p.ViaDiameter, p.ViaDrill, p.UViaDiameter, p.UViaDrill, p.DiffPairWidth, p.DiffPairGap, p.Nets)
}

// Type implements starlark.Value.
func (p *NetClass) Type() string {
	return "NetClass"
}

// Freeze implements starlark.Value.
func (p *NetClass) Freeze() {
}

// Truth implements starlark.Value.
func (p *NetClass) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *NetClass) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *NetClass) Attr(name string) (starlark.Value, error) {
	switch name {
	case "name":
		return starlark.String(p.Name), nil

	case "description":
		return starlark.String(p.Description), nil

	case "clearance":
		return starlark.Float(p.Clearance), nil

	case "trace_width":
		return starlark.Float(p.TraceWidth), nil

	case "via_diameter":
		return starlark.Float(p.ViaDiameter), nil

	case "via_drill":
		return starlark.Float(p.ViaDrill), nil

	case "u_via_diameter":
		return starlark.Float(p.UViaDiameter), nil

	case "u_via_drill":
		return starlark.Float(p.UViaDrill), nil

	case "diff_pair_width":
		return starlark.Float(p.DiffPairWidth), nil

	case "diff_pair_gap":
		return starlark.Float(p.DiffPairGap), nil

	case "nets":
		l := starlark.NewList(nil)
		for _, e := range p.Nets {
			l.Append(starlark.String(e))
		}
		return l, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *NetClass) AttrNames() []string {
	return []string{"name", "description", "clearance", "trace_width", "via_diameter", "via_drill", "u_via_diameter", "u_via_drill", "diff_pair_width", "diff_pair_gap", "nets"}
}

// SetField implements starlark.HasSetField.
func (p *NetClass) SetField(name string, val starlark.Value) error {
	switch name {
	case "name":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to name using type %T", val)
		}
		p.Name = string(v)

	case "description":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to description using type %T", val)
		}
		p.Description = string(v)

	case "clearance":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to clearance using type %T", val)
		}
		p.Clearance = float64(v)

	case "trace_width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to trace_width using type %T", val)
		}
		p.TraceWidth = float64(v)

	case "via_diameter":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to via_diameter using type %T", val)
		}
		p.ViaDiameter = float64(v)

	case "via_drill":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to via_drill using type %T", val)
		}
		p.ViaDrill = float64(v)

	case "u_via_diameter":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to u_via_diameter using type %T", val)
		}
		p.UViaDiameter = float64(v)

	case "u_via_drill":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to u_via_drill using type %T", val)
		}
		p.UViaDrill = float64(v)

	case "diff_pair_width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to diff_pair_width using type %T", val)
		}
		p.DiffPairWidth = float64(v)

	case "diff_pair_gap":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to diff_pair_gap using type %T", val)
		}
		p.DiffPairGap = float64(v)

	case "nets":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to nets using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(starlark.String)
			if !ok {
				return errors.New("nets is not a string")
			}
			p.Nets = append(p.Nets, string(s))
		}
	}

	return errors.New("no such assignable field: " + name)
}

var MakeNet = starlark.NewBuiltin("Net", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.String
	)
	unpackErr := starlark.UnpackArgs(
		"Net",
		args,
		kwargs,
		"name?", &f0,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Net{}

	out.Name = string(f0)
	return &out, nil
})

func (p *Net) String() string {
	return fmt.Sprintf("Net{%v}", p.Name)
}

// Type implements starlark.Value.
func (p *Net) Type() string {
	return "Net"
}

// Freeze implements starlark.Value.
func (p *Net) Freeze() {
}

// Truth implements starlark.Value.
func (p *Net) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Net) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Net) Attr(name string) (starlark.Value, error) {
	switch name {
	case "name":
		return starlark.String(p.Name), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Net) AttrNames() []string {
	return []string{"name"}
}

// SetField implements starlark.HasSetField.
func (p *Net) SetField(name string, val starlark.Value) error {
	switch name {
	case "name":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to name using type %T", val)
		}
		p.Name = string(v)

	}

	return errors.New("no such assignable field: " + name)
}
