package pcb

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"go.starlark.net/starlark"
)

var MakeXY = starlark.NewBuiltin("XY", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.Value
		f1 starlark.Value
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

	out.X, _ = starlark.AsFloat(f0)
	out.Y, _ = starlark.AsFloat(f1)
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
		var ok bool
		p.X, ok = starlark.AsFloat(val)
		if !ok {
			return fmt.Errorf("cannot assign to x using type %T", val)
		}
		return nil

	case "y":
		var ok bool
		p.Y, ok = starlark.AsFloat(val)
		if !ok {
			return fmt.Errorf("cannot assign to y using type %T", val)
		}
		return nil

	}

	return errors.New("no such assignable field: " + name)
}

var MakeXYZ = starlark.NewBuiltin("XYZ", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.Value
		f1 starlark.Value
		f2 starlark.Value
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

	out.X, _ = starlark.AsFloat(f0)
	out.Y, _ = starlark.AsFloat(f1)
	var hasZ bool
	out.Z, hasZ = starlark.AsFloat(f2)

	out.ZPresent = hasZ || bool(f3)
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
		return nil

	case "y":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to y using type %T", val)
		}
		p.Y = float64(v)
		return nil

	case "z":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to z using type %T", val)
		}
		p.Z = float64(v)
		return nil

	case "z_present":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to z_present using type %T", val)
		}
		p.ZPresent = bool(v)
		return nil

	case "unlocked":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to unlocked using type %T", val)
		}
		p.Unlocked = bool(v)
		return nil
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

var MakeModPolygon = starlark.NewBuiltin("ModPolygon", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 *starlark.List
		f2 starlark.String
		f3 starlark.Float
	)
	unpackErr := starlark.UnpackArgs(
		"ModPolygon",
		args,
		kwargs,
		"at?", &f0,
		"points?", &f1,
		"layer?", &f2,
		"width?", &f3,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModPolygon{}

	if f0 != nil {
		out.At = *f0
	}
	if f1 != nil {
		for i := 0; i < f1.Len(); i++ {
			s, ok := f1.Index(i).(*XY)
			if !ok {
				return starlark.None, fmt.Errorf("point[%d] is not an XY", i)
			}
			out.Points = append(out.Points, *s)
		}
	}
	out.Layer = string(f2)
	out.Width = float64(f3)
	return &out, nil
})

func (p *ModPolygon) String() string {
	return fmt.Sprintf("ModPolygon{%v, %v, %v, %v}", p.At, p.Points, p.Layer, p.Width)
}

// Type implements starlark.Value.
func (p *ModPolygon) Type() string {
	return "ModPolygon"
}

// Freeze implements starlark.Value.
func (p *ModPolygon) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModPolygon) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModPolygon) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModPolygon) Attr(name string) (starlark.Value, error) {
	switch name {
	case "at":
		return &p.At, nil

	case "points":
		l := starlark.NewList(nil)
		for _, e := range p.Points {
			dupe := e
			l.Append(&dupe)
		}
		return l, nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "width":
		return starlark.Float(p.Width), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModPolygon) AttrNames() []string {
	return []string{"at", "points", "layer", "width"}
}

// SetField implements starlark.HasSetField.
func (p *ModPolygon) SetField(name string, val starlark.Value) error {
	switch name {
	case "at":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v
		return nil

	case "points":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to points using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*XY)
			if !ok {
				return errors.New("points is not a XY")
			}
			p.Points = append(p.Points, *s)
		}
		return nil

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)
		return nil

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)
		return nil
	}

	return errors.New("no such assignable field: " + name)
}

var MakeTextEffects = starlark.NewBuiltin("TextEffects", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 starlark.Float
		f2 TextJustify
		f3 starlark.Bool
		f4 starlark.Bool
	)
	unpackErr := starlark.UnpackArgs(
		"TextEffects",
		args,
		kwargs,
		"font_size?", &f0,
		"thickness?", &f1,
		"justify?", &f2,
		"bold?", &f3,
		"italic?", &f4,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := TextEffects{}

	if f0 != nil {
		out.FontSize = *f0
	}
	out.Thickness = float64(f1)
	out.Justify = f2
	out.Bold = bool(f3)
	out.Italic = bool(f4)
	return &out, nil
})

func (p *TextEffects) String() string {
	return fmt.Sprintf("TextEffects{%v, %v, %v, %v, %v}", p.FontSize, p.Thickness, p.Justify, p.Bold, p.Italic)
}

// Type implements starlark.Value.
func (p *TextEffects) Type() string {
	return "TextEffects"
}

// Freeze implements starlark.Value.
func (p *TextEffects) Freeze() {
}

// Truth implements starlark.Value.
func (p *TextEffects) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *TextEffects) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *TextEffects) Attr(name string) (starlark.Value, error) {
	switch name {
	case "font_size":
		return &p.FontSize, nil

	case "thickness":
		return starlark.Float(p.Thickness), nil

	case "justify":
		return &p.Justify, nil

	case "bold":
		return starlark.Bool(p.Bold), nil

	case "italic":
		return starlark.Bool(p.Italic), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *TextEffects) AttrNames() []string {
	return []string{"font_size", "thickness", "justify", "bold", "italic"}
}

// SetField implements starlark.HasSetField.
func (p *TextEffects) SetField(name string, val starlark.Value) error {
	switch name {
	case "font_size":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to font_size using type %T", val)
		}
		p.FontSize = *v
		return nil

	case "thickness":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to thickness using type %T", val)
		}
		p.Thickness = float64(v)
		return nil

	case "justify":
		v, ok := val.(*TextJustify)
		if !ok {
			return fmt.Errorf("cannot assign to justify using type %T", val)
		}
		p.Justify = *v
		return nil

	case "bold":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to bold using type %T", val)
		}
		p.Bold = bool(v)
		return nil

	case "italic":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to italic using type %T", val)
		}
		p.Italic = bool(v)
		return nil
	}

	return errors.New("no such assignable field: " + name)
}

var MakeModText = starlark.NewBuiltin("ModText", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *ModTextKind
		f1 starlark.Bool
		f2 starlark.String
		f3 *XYZ
		f4 starlark.String
		f5 *TextEffects
	)
	unpackErr := starlark.UnpackArgs(
		"ModText",
		args,
		kwargs,
		"kind?", &f0,
		"hidden?", &f1,
		"text?", &f2,
		"at?", &f3,
		"layer?", &f4,
		"effects?", &f5,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModText{}

	out.Kind = *f0
	out.Hidden = bool(f1)
	out.Text = string(f2)
	out.At = *f3
	out.Layer = string(f4)
	if f5 != nil {
		out.Effects = *f5
	} else {
		out.Effects = TextEffects{FontSize: XY{X: 1, Y: 1}, Thickness: 0.15}
	}
	return &out, nil
})

func (p *ModText) String() string {
	return fmt.Sprintf("ModText{%v, %v, %v, %v, %v, %v}", p.Kind, p.Hidden, p.Text, p.At, p.Layer, p.Effects)
}

// Type implements starlark.Value.
func (p *ModText) Type() string {
	return "ModText"
}

// Freeze implements starlark.Value.
func (p *ModText) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModText) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModText) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModText) Attr(name string) (starlark.Value, error) {
	switch name {
	case "kind":
		return &p.Kind, nil

	case "hidden":
		return starlark.Bool(p.Hidden), nil

	case "text":
		return starlark.String(p.Text), nil

	case "at":
		return &p.At, nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "effects":
		return &p.Effects, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModText) AttrNames() []string {
	return []string{"kind", "hidden", "text", "at", "layer", "effects"}
}

// SetField implements starlark.HasSetField.
func (p *ModText) SetField(name string, val starlark.Value) error {
	switch name {
	case "kind":
		v, ok := val.(*ModTextKind)
		if !ok {
			return fmt.Errorf("cannot assign to kind using type %T", val)
		}
		p.Kind = *v
		return nil

	case "hidden":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to hidden using type %T", val)
		}
		p.Hidden = bool(v)
		return nil

	case "text":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to text using type %T", val)
		}
		p.Text = string(v)
		return nil

	case "at":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v
		return nil

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)
		return nil

	case "effects":
		v, ok := val.(*TextEffects)
		if !ok {
			return fmt.Errorf("cannot assign to effects using type %T", val)
		}
		p.Effects = *v
		return nil
	}

	return errors.New("no such assignable field: " + name)
}

var MakeModLine = starlark.NewBuiltin("ModLine", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 *XY
		f2 starlark.String
		f3 starlark.Float
	)
	unpackErr := starlark.UnpackArgs(
		"ModLine",
		args,
		kwargs,
		"start?", &f0,
		"end?", &f1,
		"layer?", &f2,
		"width?", &f3,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModLine{}

	out.Start = *f0
	out.End = *f1
	out.Layer = string(f2)
	out.Width = float64(f3)
	return &out, nil
})

func (p *ModLine) String() string {
	return fmt.Sprintf("ModLine{%v, %v, %v, %v}", p.Start, p.End, p.Layer, p.Width)
}

// Type implements starlark.Value.
func (p *ModLine) Type() string {
	return "ModLine"
}

// Freeze implements starlark.Value.
func (p *ModLine) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModLine) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModLine) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModLine) Attr(name string) (starlark.Value, error) {
	switch name {
	case "start":
		return &p.Start, nil

	case "end":
		return &p.End, nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "width":
		return starlark.Float(p.Width), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModLine) AttrNames() []string {
	return []string{"start", "end", "layer", "width"}
}

// SetField implements starlark.HasSetField.
func (p *ModLine) SetField(name string, val starlark.Value) error {
	switch name {
	case "start":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to start using type %T", val)
		}
		p.Start = *v
		return nil

	case "end":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to end using type %T", val)
		}
		p.End = *v
		return nil

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)
		return nil

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)
		return nil
	}

	return errors.New("no such assignable field: " + name)
}

var MakeModCircle = starlark.NewBuiltin("ModCircle", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 *XY
		f2 starlark.String
		f3 starlark.Float
	)
	unpackErr := starlark.UnpackArgs(
		"ModCircle",
		args,
		kwargs,
		"center?", &f0,
		"end?", &f1,
		"layer?", &f2,
		"width?", &f3,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModCircle{}

	if f0 != nil {
		out.Center = *f0
	}
	if f1 != nil {
		out.End = *f1
	}
	out.Layer = string(f2)
	out.Width = float64(f3)
	return &out, nil
})

func (p *ModCircle) String() string {
	return fmt.Sprintf("ModCircle{%v, %v, %v, %v}", p.Center, p.End, p.Layer, p.Width)
}

// Type implements starlark.Value.
func (p *ModCircle) Type() string {
	return "ModCircle"
}

// Freeze implements starlark.Value.
func (p *ModCircle) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModCircle) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModCircle) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModCircle) Attr(name string) (starlark.Value, error) {
	switch name {
	case "center":
		return &p.Center, nil

	case "end":
		return &p.End, nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "width":
		return starlark.Float(p.Width), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModCircle) AttrNames() []string {
	return []string{"center", "end", "layer", "width"}
}

// SetField implements starlark.HasSetField.
func (p *ModCircle) SetField(name string, val starlark.Value) error {
	switch name {
	case "center":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to center using type %T", val)
		}
		p.Center = *v
		return nil

	case "end":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to end using type %T", val)
		}
		p.End = *v
		return nil

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)
		return nil

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)
		return nil

	}

	return errors.New("no such assignable field: " + name)
}

var MakeModArc = starlark.NewBuiltin("ModArc", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 *XY
		f2 starlark.String
		f3 starlark.Float
		f4 starlark.Float
	)
	unpackErr := starlark.UnpackArgs(
		"ModArc",
		args,
		kwargs,
		"start?", &f0,
		"end?", &f1,
		"layer?", &f2,
		"angle?", &f3,
		"width?", &f4,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModArc{}

	if f0 != nil {
		out.Start = *f0
	}
	if f1 != nil {
		out.End = *f1
	}
	out.Layer = string(f2)
	out.Angle = float64(f3)
	out.Width = float64(f4)
	return &out, nil
})

func (p *ModArc) String() string {
	return fmt.Sprintf("ModArc{%v, %v, %v, %v, %v}", p.Start, p.End, p.Layer, p.Angle, p.Width)
}

// Type implements starlark.Value.
func (p *ModArc) Type() string {
	return "ModArc"
}

// Freeze implements starlark.Value.
func (p *ModArc) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModArc) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModArc) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModArc) Attr(name string) (starlark.Value, error) {
	switch name {
	case "start":
		return &p.Start, nil

	case "end":
		return &p.End, nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "angle":
		return starlark.Float(p.Angle), nil

	case "width":
		return starlark.Float(p.Width), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModArc) AttrNames() []string {
	return []string{"start", "end", "layer", "angle", "width"}
}

// SetField implements starlark.HasSetField.
func (p *ModArc) SetField(name string, val starlark.Value) error {
	switch name {
	case "start":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to start using type %T", val)
		}
		p.Start = *v
		return nil

	case "end":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to end using type %T", val)
		}
		p.End = *v
		return nil

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)
		return nil

	case "angle":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to angle using type %T", val)
		}
		p.Angle = float64(v)
		return nil

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)
		return nil

	}

	return errors.New("no such assignable field: " + name)
}

var MakeModModel = starlark.NewBuiltin("ModModel", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.String
		f1 *XYZ
		f2 *XYZ
		f3 *XYZ
		f4 *XYZ
	)
	unpackErr := starlark.UnpackArgs(
		"ModModel",
		args,
		kwargs,
		"path?", &f0,
		"at?", &f1,
		"offset?", &f2,
		"scale?", &f3,
		"rotate?", &f4,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModModel{}

	out.Path = string(f0)
	if f1 != nil {
		out.At = *f1
	}
	if f2 != nil {
		out.Offset = *f2
	}
	if f3 != nil {
		out.Scale = *f3
	}
	if f4 != nil {
		out.Rotate = *f4
	}
	return &out, nil
})

func (p *ModModel) String() string {
	return fmt.Sprintf("ModModel{%v, %v, %v, %v, %v}", p.Path, p.At, p.Offset, p.Scale, p.Rotate)
}

// Type implements starlark.Value.
func (p *ModModel) Type() string {
	return "ModModel"
}

// Freeze implements starlark.Value.
func (p *ModModel) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModModel) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModModel) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModModel) Attr(name string) (starlark.Value, error) {
	switch name {
	case "path":
		return starlark.String(p.Path), nil

	case "at":
		return &p.At, nil

	case "offset":
		return &p.Offset, nil

	case "scale":
		return &p.Scale, nil

	case "rotate":
		return &p.Rotate, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModModel) AttrNames() []string {
	return []string{"path", "at", "offset", "scale", "rotate"}
}

// SetField implements starlark.HasSetField.
func (p *ModModel) SetField(name string, val starlark.Value) error {
	switch name {
	case "path":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to path using type %T", val)
		}
		p.Path = string(v)
		return nil

	case "at":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v
		return nil

	case "offset":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to offset using type %T", val)
		}
		p.Offset = *v
		return nil

	case "scale":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to scale using type %T", val)
		}
		p.Scale = *v
		return nil

	case "rotate":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to rotate using type %T", val)
		}
		p.Rotate = *v
		return nil

	}

	return errors.New("no such assignable field: " + name)
}

var MakePadOptions = starlark.NewBuiltin("PadOptions", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.String
		f1 starlark.String
	)
	unpackErr := starlark.UnpackArgs(
		"PadOptions",
		args,
		kwargs,
		"clearance?", &f0,
		"anchor?", &f1,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := PadOptions{}

	out.Clearance = string(f0)
	out.Anchor = string(f1)
	return &out, nil
})

func (p *PadOptions) String() string {
	return fmt.Sprintf("PadOptions{%v, %v}", p.Clearance, p.Anchor)
}

// Type implements starlark.Value.
func (p *PadOptions) Type() string {
	return "PadOptions"
}

// Freeze implements starlark.Value.
func (p *PadOptions) Freeze() {
}

// Truth implements starlark.Value.
func (p *PadOptions) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *PadOptions) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *PadOptions) Attr(name string) (starlark.Value, error) {
	switch name {
	case "clearance":
		return starlark.String(p.Clearance), nil

	case "anchor":
		return starlark.String(p.Anchor), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *PadOptions) AttrNames() []string {
	return []string{"clearance", "anchor"}
}

// SetField implements starlark.HasSetField.
func (p *PadOptions) SetField(name string, val starlark.Value) error {
	switch name {
	case "clearance":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to clearance using type %T", val)
		}
		p.Clearance = string(v)

	case "anchor":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to anchor using type %T", val)
		}
		p.Anchor = string(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeModGraphic = starlark.NewBuiltin("ModGraphic", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.String
		f1 modDrawable
	)
	unpackErr := starlark.UnpackArgs(
		"ModGraphic",
		args,
		kwargs,
		"ident?", &f0,
		"renderable?", &f1,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModGraphic{}

	out.Ident = string(f0)
	out.Renderable = f1
	return &out, nil
})

func (p *ModGraphic) String() string {
	return fmt.Sprintf("ModGraphic{%v, %v}", p.Ident, p.Renderable)
}

// Type implements starlark.Value.
func (p *ModGraphic) Type() string {
	return "ModGraphic"
}

// Freeze implements starlark.Value.
func (p *ModGraphic) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModGraphic) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModGraphic) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v %v", p, p.Renderable)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModGraphic) Attr(name string) (starlark.Value, error) {
	switch name {
	case "ident":
		return starlark.String(p.Ident), nil

	case "renderable":
		return p.Renderable.(starlark.Value), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModGraphic) AttrNames() []string {
	return []string{"ident", "renderable"}
}

// SetField implements starlark.HasSetField.
func (p *ModGraphic) SetField(name string, val starlark.Value) error {
	switch name {
	case "ident":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to ident using type %T", val)
		}
		p.Ident = string(v)
		return nil

	case "renderable":
		v, ok := val.(modDrawable)
		if !ok {
			return fmt.Errorf("cannot assign to renderable using type %T", val)
		}
		p.Renderable = v
		return nil
	}

	return errors.New("no such assignable field: " + name)
}

var MakePad = starlark.NewBuiltin("Pad", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0  starlark.String
		f1  starlark.Int
		f2  starlark.String
		f3  starlark.Value
		f4  *XY
		f5  *starlark.List
		f6  *XY
		f7  *XY
		f8  *XY
		f9  *PadShape
		f10 starlark.Float
		f11 *ZoneConnectMode
		f12 starlark.Float
		f13 starlark.Float
		f14 starlark.Float
		f15 starlark.Float
		f16 starlark.Float
		f17 starlark.Float
		f18 starlark.Float
		f19 starlark.Float
		f20 *PadSurface
		f21 *PadShape
		f22 *PadOptions
		f23 *starlark.List
	)
	unpackErr := starlark.UnpackArgs(
		"Pad",
		args,
		kwargs,
		"ident?", &f0,
		"net_num?", &f1,
		"net_name?", &f2,
		"at?", &f3,
		"size?", &f4,
		"layers?", &f5,
		"rect_delta?", &f6,
		"drill_offset?", &f7,
		"drill_size?", &f8,
		"drill_shape?", &f9,
		"die_length?", &f10,
		"zone_connect?", &f11,
		"thermal_width?", &f12,
		"thermal_gap?", &f13,
		"round_rect_r_ratio?", &f14,
		"chamfer_ratio?", &f15,
		"solder_mask_margin?", &f16,
		"solder_paste_margin?", &f17,
		"solder_paste_margin_ratio?", &f18,
		"clearance?", &f19,
		"surface?", &f20,
		"shape?", &f21,
		"options?", &f22,
		"primitives?", &f23,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Pad{}

	out.Ident = string(f0)

	if v, ok := f1.Int64(); ok {
		out.NetNum = int(v)
	}
	out.NetName = string(f2)
	if f3 != nil {
		if xy, ok := f3.(*XY); ok {
			out.At = XYZ{X: xy.X, Y: xy.Y}
		} else if xyz, ok := f3.(*XYZ); ok {
			out.At = *xyz
		} else {
			return starlark.None, fmt.Errorf("cannot convert %T to XYZ", f3)
		}
	}
	if f4 != nil {
		out.Size = *f4
	}
	if f5 != nil {
		for i := 0; i < f5.Len(); i++ {
			s, ok := f5.Index(i).(starlark.String)
			if !ok {
				return starlark.None, errors.New("layers is not a string")
			}
			out.Layers = append(out.Layers, string(s))
		}
	}
	if f6 != nil {
		out.RectDelta = *f6
	}
	if f7 != nil {
		out.DrillOffset = *f7
	}
	if f8 != nil {
		out.DrillSize = *f8
	}
	if f9 != nil {
		out.DrillShape = *f9
	}
	out.DieLength = float64(f10)
	if f11 != nil {
		out.ZoneConnect = *f11
	} else {
		out.ZoneConnect = ZoneConnectInherited
	}
	out.ThermalWidth = float64(f12)
	out.ThermalGap = float64(f13)
	out.RoundRectRRatio = float64(f14)
	out.ChamferRatio = float64(f15)
	out.SolderMaskMargin = float64(f16)
	out.SolderPasteMargin = float64(f17)
	out.SolderPasteMarginRatio = float64(f18)
	out.Clearance = float64(f19)
	if f20 != nil {
		out.Surface = *f20
	} else {
		out.Surface = SurfaceTH
	}
	if f21 != nil {
		out.Shape = *f21
	} else {
		out.Shape = ShapeOval
	}
	out.Options = f22
	if f23 != nil {
		for i := 0; i < f23.Len(); i++ {
			s, ok := f23.Index(i).(*ModGraphic)
			if !ok {
				return starlark.None, errors.New("primitives is not a ModGraphic")
			}
			out.Primitives = append(out.Primitives, *s)
		}
	}
	return &out, nil
})

func (p *Pad) String() string {
	return fmt.Sprintf("Pad{%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v}", p.Ident, p.NetNum, p.NetName, p.At, p.Size, p.Layers, p.RectDelta, p.DrillOffset, p.DrillSize, p.DrillShape, p.DieLength, p.ZoneConnect, p.ThermalWidth, p.ThermalGap, p.RoundRectRRatio, p.ChamferRatio, p.SolderMaskMargin, p.SolderPasteMargin, p.SolderPasteMarginRatio, p.Clearance, p.Surface, p.Shape, p.Options, p.Primitives)
}

// Type implements starlark.Value.
func (p *Pad) Type() string {
	return "Pad"
}

// Freeze implements starlark.Value.
func (p *Pad) Freeze() {
}

// Truth implements starlark.Value.
func (p *Pad) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Pad) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Pad) Attr(name string) (starlark.Value, error) {
	switch name {
	case "ident":
		return starlark.String(p.Ident), nil

	case "net_num":
		return starlark.MakeInt(p.NetNum), nil

	case "net_name":
		return starlark.String(p.NetName), nil

	case "at":
		return &p.At, nil

	case "size":
		return &p.Size, nil

	case "layers":
		l := starlark.NewList(nil)
		for _, e := range p.Layers {
			l.Append(starlark.String(e))
		}
		return l, nil

	case "rect_delta":
		return &p.RectDelta, nil

	case "drill_offset":
		return &p.DrillOffset, nil

	case "drill_size":
		return &p.DrillSize, nil

	case "drill_shape":
		return &p.DrillShape, nil

	case "die_length":
		return starlark.Float(p.DieLength), nil

	case "zone_connect":
		return &p.ZoneConnect, nil

	case "thermal_width":
		return starlark.Float(p.ThermalWidth), nil

	case "thermal_gap":
		return starlark.Float(p.ThermalGap), nil

	case "round_rect_r_ratio":
		return starlark.Float(p.RoundRectRRatio), nil

	case "chamfer_ratio":
		return starlark.Float(p.ChamferRatio), nil

	case "solder_mask_margin":
		return starlark.Float(p.SolderMaskMargin), nil

	case "solder_paste_margin":
		return starlark.Float(p.SolderPasteMargin), nil

	case "solder_paste_margin_ratio":
		return starlark.Float(p.SolderPasteMarginRatio), nil

	case "clearance":
		return starlark.Float(p.Clearance), nil

	case "surface":
		return &p.Surface, nil

	case "shape":
		return &p.Shape, nil

	case "options":
		return p.Options, nil

	case "primitives":
		l := starlark.NewList(nil)
		for _, e := range p.Primitives {
			l.Append(&e)
		}
		return l, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Pad) AttrNames() []string {
	return []string{"ident", "net_num", "net_name", "at", "size", "layers", "rect_delta", "drill_offset", "drill_size", "drill_shape", "die_length", "zone_connect", "thermal_width", "thermal_gap", "round_rect_r_ratio", "chamfer_ratio", "solder_mask_margin", "solder_paste_margin", "solder_paste_margin_ratio", "clearance", "surface", "shape", "options", "primitives"}
}

// SetField implements starlark.HasSetField.
func (p *Pad) SetField(name string, val starlark.Value) error {
	switch name {
	case "ident":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to ident using type %T", val)
		}
		p.Ident = string(v)

	case "net_num":
		v, ok := val.(*starlark.Int)
		if !ok {
			return fmt.Errorf("cannot assign to net_num using type %T", val)
		}
		i, ok := v.Int64()
		if !ok {
			return fmt.Errorf("cannot convert %v to int64", v)
		}
		p.NetNum = int(i)

	case "net_name":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to net_name using type %T", val)
		}
		p.NetName = string(v)

	case "at":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v

	case "size":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to size using type %T", val)
		}
		p.Size = *v

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

	case "rect_delta":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to rect_delta using type %T", val)
		}
		p.RectDelta = *v

	case "drill_offset":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to drill_offset using type %T", val)
		}
		p.DrillOffset = *v

	case "drill_size":
		v, ok := val.(*XY)
		if !ok {
			return fmt.Errorf("cannot assign to drill_size using type %T", val)
		}
		p.DrillSize = *v

	case "drill_shape":
		v, ok := val.(*PadShape)
		if !ok {
			return fmt.Errorf("cannot assign to drill_shape using type %T", val)
		}
		p.DrillShape = *v

	case "die_length":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to die_length using type %T", val)
		}
		p.DieLength = float64(v)

	case "zone_connect":
		v, ok := val.(*ZoneConnectMode)
		if !ok {
			return fmt.Errorf("cannot assign to zone_connect using type %T", val)
		}
		p.ZoneConnect = *v

	case "thermal_width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to thermal_width using type %T", val)
		}
		p.ThermalWidth = float64(v)

	case "thermal_gap":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to thermal_gap using type %T", val)
		}
		p.ThermalGap = float64(v)

	case "round_rect_r_ratio":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to round_rect_r_ratio using type %T", val)
		}
		p.RoundRectRRatio = float64(v)

	case "chamfer_ratio":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to chamfer_ratio using type %T", val)
		}
		p.ChamferRatio = float64(v)

	case "solder_mask_margin":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to solder_mask_margin using type %T", val)
		}
		p.SolderMaskMargin = float64(v)

	case "solder_paste_margin":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to solder_paste_margin using type %T", val)
		}
		p.SolderPasteMargin = float64(v)

	case "solder_paste_margin_ratio":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to solder_paste_margin_ratio using type %T", val)
		}
		p.SolderPasteMarginRatio = float64(v)

	case "clearance":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to clearance using type %T", val)
		}
		p.Clearance = float64(v)

	case "surface":
		v, ok := val.(*PadSurface)
		if !ok {
			return fmt.Errorf("cannot assign to surface using type %T", val)
		}
		p.Surface = *v

	case "shape":
		v, ok := val.(*PadShape)
		if !ok {
			return fmt.Errorf("cannot assign to shape using type %T", val)
		}
		p.Shape = *v

	case "options":
		v, ok := val.(*PadOptions)
		if !ok {
			return fmt.Errorf("cannot assign to options using type %T", val)
		}
		p.Options = v

	case "primitives":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to primitives using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*ModGraphic)
			if !ok {
				return errors.New("primitives is not a ModGraphic")
			}
			p.Primitives = append(p.Primitives, *s)
		}
	}

	return errors.New("no such assignable field: " + name)
}

var MakeModPlacement = starlark.NewBuiltin("ModPlacement", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XYZ
	)
	unpackErr := starlark.UnpackArgs(
		"ModPlacement",
		args,
		kwargs,
		"at?", &f0,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := ModPlacement{}

	if f0 != nil {
		out.At = *f0
	}
	return &out, nil
})

func (p *ModPlacement) String() string {
	return fmt.Sprintf("ModPlacement{%v}", p.At)
}

// Type implements starlark.Value.
func (p *ModPlacement) Type() string {
	return "ModPlacement"
}

// Freeze implements starlark.Value.
func (p *ModPlacement) Freeze() {
}

// Truth implements starlark.Value.
func (p *ModPlacement) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *ModPlacement) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *ModPlacement) Attr(name string) (starlark.Value, error) {
	switch name {
	case "at":
		return &p.At, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *ModPlacement) AttrNames() []string {
	return []string{"at"}
}

// SetField implements starlark.HasSetField.
func (p *ModPlacement) SetField(name string, val starlark.Value) error {
	switch name {
	case "at":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v

	}

	return errors.New("no such assignable field: " + name)
}

var MakeModule = starlark.NewBuiltin("Module", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0  starlark.String
		f1  ModPlacement
		f2  starlark.Bool
		f3  starlark.Bool
		f4  starlark.String
		f5  *ZoneConnectMode
		f6  starlark.Float
		f7  starlark.Float
		f8  starlark.Float
		f9  starlark.Float
		f10 starlark.String
		f11 starlark.String
		f12 starlark.String
		f13 starlark.String
		f14 *starlark.List
		f15 *starlark.List
		f17 *starlark.List
		f18 *starlark.List
		f19 *starlark.List
	)
	unpackErr := starlark.UnpackArgs(
		"Module",
		args,
		kwargs,
		"name?", &f0,
		"placement?", &f1,
		"placed?", &f2,
		"locked?", &f3,
		"layer?", &f4,
		"zone_connect?", &f5,
		"solder_mask_margin?", &f6,
		"solder_paste_margin?", &f7,
		"solder_paste_ratio?", &f8,
		"clearance?", &f9,
		"tedit?", &f10,
		"tstamp?", &f11,
		"path?", &f12,
		"description?", &f13,
		"tags?", &f14,
		"attrs?", &f15,
		"graphics?", &f17,
		"pads?", &f18,
		"models?", &f19,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Module{}

	out.Name = string(f0)
	out.Placement = f1
	out.Placed = bool(f2)
	out.Locked = bool(f3)
	out.Layer = string(f4)
	if f5 != nil {
		out.ZoneConnect = *f5
	} else {
		out.ZoneConnect = ZoneConnectInherited
	}
	out.SolderMaskMargin = float64(f6)
	out.SolderPasteMargin = float64(f7)
	out.SolderPasteRatio = float64(f8)
	out.Clearance = float64(f9)
	out.Tedit = string(f10)
	out.Tstamp = string(f11)
	out.Path = string(f12)
	out.Description = string(f13)
	if f14 != nil {
		for i := 0; i < f14.Len(); i++ {
			s, ok := f14.Index(i).(starlark.String)
			if !ok {
				return starlark.None, errors.New("tags is not a string")
			}
			out.Tags = append(out.Tags, string(s))
		}
	}
	if f15 != nil {
		for i := 0; i < f15.Len(); i++ {
			s, ok := f15.Index(i).(starlark.String)
			if !ok {
				return starlark.None, errors.New("attrs is not a string")
			}
			out.Attrs = append(out.Attrs, string(s))
		}
	}
	if f17 != nil {
		for i := 0; i < f17.Len(); i++ {
			s, ok := f17.Index(i).(*ModGraphic)
			if !ok {
				return starlark.None, errors.New("graphics is not a ModGraphic")
			}
			out.Graphics = append(out.Graphics, *s)
		}
	}
	if f18 != nil {
		for i := 0; i < f18.Len(); i++ {
			s, ok := f18.Index(i).(*Pad)
			if !ok {
				return starlark.None, errors.New("pads is not a Pad")
			}
			out.Pads = append(out.Pads, *s)
		}
	}
	if f19 != nil {
		for i := 0; i < f19.Len(); i++ {
			s, ok := f19.Index(i).(*ModModel)
			if !ok {
				return starlark.None, errors.New("models is not a ModModel")
			}
			out.Models = append(out.Models, *s)
		}
	}
	return &out, nil
})

func (p *Module) String() string {
	return fmt.Sprintf("Module{%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v}", p.Name, p.Placement, p.Placed, p.Locked, p.Layer, p.ZoneConnect, p.SolderMaskMargin, p.SolderPasteMargin, p.SolderPasteRatio, p.Clearance, p.Tedit, p.Tstamp, p.Path, p.Description, p.Tags, p.Attrs, p.Graphics, p.Pads, p.Models)
}

// Type implements starlark.Value.
func (p *Module) Type() string {
	return "Module"
}

// Freeze implements starlark.Value.
func (p *Module) Freeze() {
}

// Truth implements starlark.Value.
func (p *Module) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Module) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Module) Attr(name string) (starlark.Value, error) {
	switch name {
	case "name":
		return starlark.String(p.Name), nil

	case "placement":
		return &p.Placement, nil

	case "placed":
		return starlark.Bool(p.Placed), nil

	case "locked":
		return starlark.Bool(p.Locked), nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "zone_connect":
		return &p.ZoneConnect, nil

	case "solder_mask_margin":
		return starlark.Float(p.SolderMaskMargin), nil

	case "solder_paste_margin":
		return starlark.Float(p.SolderPasteMargin), nil

	case "solder_paste_ratio":
		return starlark.Float(p.SolderPasteRatio), nil

	case "clearance":
		return starlark.Float(p.Clearance), nil

	case "tedit":
		return starlark.String(p.Tedit), nil

	case "tstamp":
		return starlark.String(p.Tstamp), nil

	case "path":
		return starlark.String(p.Path), nil

	case "description":
		return starlark.String(p.Description), nil

	case "tags":
		l := starlark.NewList(nil)
		for _, e := range p.Tags {
			l.Append(starlark.String(e))
		}
		return l, nil

	case "attrs":
		l := starlark.NewList(nil)
		for _, e := range p.Attrs {
			l.Append(starlark.String(e))
		}
		return l, nil

	case "graphics":
		l := starlark.NewList(nil)
		for _, e := range p.Graphics {
			dupeDrawable := e.Renderable
			l.Append(&ModGraphic{
				Ident:      e.Ident,
				Renderable: dupeDrawable,
			})
		}
		return l, nil

	case "pads":
		l := starlark.NewList(nil)
		for _, e := range p.Pads {
			dupe := e
			l.Append(&dupe)
		}
		return l, nil

	case "models":
		l := starlark.NewList(nil)
		for _, e := range p.Models {
			dupe := e
			l.Append(&dupe)
		}
		return l, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Module) AttrNames() []string {
	return []string{"name", "placement", "placed", "locked", "layer", "zone_connect", "solder_mask_margin", "solder_paste_margin", "solder_paste_ratio", "clearance", "tedit", "tstamp", "path", "description", "tags", "attrs", "graphics", "pads", "models"}
}

// SetField implements starlark.HasSetField.
func (p *Module) SetField(name string, val starlark.Value) error {
	switch name {
	case "name":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to name using type %T", val)
		}
		p.Name = string(v)
		return nil

	case "placement":
		v, ok := val.(*ModPlacement)
		if !ok {
			return fmt.Errorf("cannot assign to placement using type %T", val)
		}
		p.Placement = *v
		return nil

	case "placed":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to placed using type %T", val)
		}
		p.Placed = bool(v)
		return nil

	case "locked":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to locked using type %T", val)
		}
		p.Locked = bool(v)
		return nil

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)
		return nil

	case "zone_connect":
		v, ok := val.(*ZoneConnectMode)
		if !ok {
			return fmt.Errorf("cannot assign to zone_connect using type %T", val)
		}
		p.ZoneConnect = *v
		return nil

	case "solder_mask_margin":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to solder_mask_margin using type %T", val)
		}
		p.SolderMaskMargin = float64(v)
		return nil

	case "solder_paste_margin":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to solder_paste_margin using type %T", val)
		}
		p.SolderPasteMargin = float64(v)
		return nil

	case "solder_paste_ratio":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to solder_paste_ratio using type %T", val)
		}
		p.SolderPasteRatio = float64(v)
		return nil

	case "clearance":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to clearance using type %T", val)
		}
		p.Clearance = float64(v)
		return nil

	case "tedit":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tedit using type %T", val)
		}
		p.Tedit = string(v)
		return nil

	case "tstamp":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tstamp using type %T", val)
		}
		p.Tstamp = string(v)
		return nil

	case "path":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to path using type %T", val)
		}
		p.Path = string(v)
		return nil

	case "description":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to description using type %T", val)
		}
		p.Description = string(v)
		return nil

	case "tags":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to tags using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(starlark.String)
			if !ok {
				return errors.New("tags is not a string")
			}
			p.Tags = append(p.Tags, string(s))
		}
		return nil

	case "attrs":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to attrs using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(starlark.String)
			if !ok {
				return errors.New("attrs is not a string")
			}
			p.Attrs = append(p.Attrs, string(s))
		}
		return nil

	case "graphics":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to graphics using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*ModGraphic)
			if !ok {
				return errors.New("graphics is not a ModGraphic")
			}
			p.Graphics = append(p.Graphics, *s)
		}
		return nil

	case "pads":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to pads using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*Pad)
			if !ok {
				return errors.New("pads is not a Pad")
			}
			p.Pads = append(p.Pads, *s)
		}
		return nil

	case "models":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to models using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*ModModel)
			if !ok {
				return errors.New("models is not a ModModel")
			}
			p.Models = append(p.Models, *s)
		}
		return nil
	}

	return errors.New("no such assignable field: " + name)
}

var MakePCB = starlark.NewBuiltin("PCB", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		segments *starlark.List
		drawings *starlark.List
		modules  *starlark.List
	)
	unpackErr := starlark.UnpackArgs(
		"PCB",
		args,
		kwargs,
		"segments?",
		&segments,
		"drawings?",
		&drawings,
		"modules?",
		&modules,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := EmptyPCB()

	if segments != nil {
		for i := 0; i < segments.Len(); i++ {
			s, ok := segments.Index(i).(NetSegment)
			if !ok {
				return starlark.None, errors.New("segments element is not a NetSegment")
			}
			out.Segments = append(out.Segments, NetSegment(s))
		}
	}
	if drawings != nil {
		for i := 0; i < drawings.Len(); i++ {
			d, ok := drawings.Index(i).(Drawing)
			if !ok {
				return starlark.None, errors.New("drawings element is not a NetSegment")
			}
			out.Drawings = append(out.Drawings, d)
		}
	}
	if modules != nil {
		for i := 0; i < modules.Len(); i++ {
			m, ok := modules.Index(i).(*Module)
			if !ok {
				return starlark.None, errors.New("modules element is not a NetSegment")
			}
			out.Modules = append(out.Modules, *m)
		}
	}

	return out, nil
})

func (p *PCB) String() string {
	return fmt.Sprintf("PCB{%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v}", p.FormatVersion, p.CreatedBy, p.TitleInfo, p.EditorSetup, p.LayersByName, p.Layers, p.Segments, p.Drawings, p.Nets, p.NetClasses, p.Zones, p.Modules)
}

// Type implements starlark.Value.
func (p *PCB) Type() string {
	return "PCB"
}

// Freeze implements starlark.Value.
func (p *PCB) Freeze() {
}

// Truth implements starlark.Value.
func (p *PCB) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *PCB) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *PCB) Attr(name string) (starlark.Value, error) {
	switch name {

	case "layers":
		l := starlark.NewList(nil)
		for _, e := range p.Layers {
			l.Append(e)
		}
		return l, nil

	case "segments":
		l := starlark.NewList(nil)
		for _, e := range p.Segments {
			switch s := e.(type) {
			case *Track:
				l.Append(s)
			case *Via:
				l.Append(s)
			default:
				return nil, fmt.Errorf("cannot process segment of type %T", s)
			}
		}
		return l, nil

	case "drawings":
		l := starlark.NewList(nil)
		for _, e := range p.Drawings {
			switch d := e.(type) {
			case *Text:
				l.Append(d)
			case *Line:
				l.Append(d)
			case *Arc:
				l.Append(d)
			default:
				return nil, fmt.Errorf("cannot process drawing of type %T", d)
			}
		}
		return l, nil

	// case "zones":
	// 	l := starlark.NewList(nil)
	// 	for _, e := range p.Zones {
	// 		l.Append(e)
	// 	}
	// 	return l, nil

	case "modules":
		l := starlark.NewList(nil)
		for _, e := range p.Modules {
			l.Append(&e)
		}
		return l, nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *PCB) AttrNames() []string {
	return []string{"layers", "segments", "drawings", "modules"}
}

// SetField implements starlark.HasSetField.
func (p *PCB) SetField(name string, val starlark.Value) error {
	switch name {

	case "layers":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to layers using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*Layer)
			if !ok {
				return errors.New("layers is not a *Layer")
			}
			p.Layers = append(p.Layers, s)
		}

	case "segments":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to segments using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(NetSegment)
			if !ok {
				return errors.New("segments is not a NetSegment")
			}
			p.Segments = append(p.Segments, NetSegment(s))
		}

	case "drawings":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to drawings using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(Drawing)
			if !ok {
				return errors.New("drawings is not a Drawing")
			}
			p.Drawings = append(p.Drawings, Drawing(s))
		}

	case "modules":
		v, ok := val.(*starlark.List)
		if !ok {
			return fmt.Errorf("cannot assign to modules using type %T", val)
		}

		for i := 0; i < v.Len(); i++ {
			s, ok := v.Index(i).(*Module)
			if !ok {
				return errors.New("modules is not a Module")
			}
			p.Modules = append(p.Modules, *s)
		}
	}

	return errors.New("no such assignable field: " + name)
}

var MakeLine = starlark.NewBuiltin("Line", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 *XY
		f2 starlark.Float
		f3 starlark.String
		f4 starlark.Float
		f5 starlark.String
	)
	unpackErr := starlark.UnpackArgs(
		"Line",
		args,
		kwargs,
		"start?", &f0,
		"end?", &f1,
		"angle?", &f2,
		"layer?", &f3,
		"width?", &f4,
		"tstamp?", &f5,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Line{}

	if f0 != nil {
		out.Start = *f0
	}
	if f1 != nil {
		out.End = *f1
	}
	out.Angle = float64(f2)
	out.Layer = string(f3)
	out.Width = float64(f4)
	out.Tstamp = string(f5)
	return &out, nil
})

func (p *Line) String() string {
	return fmt.Sprintf("Line{%v, %v, %v, %v, %v, %v}", p.Start, p.End, p.Angle, p.Layer, p.Width, p.Tstamp)
}

// Type implements starlark.Value.
func (p *Line) Type() string {
	return "Line"
}

// Freeze implements starlark.Value.
func (p *Line) Freeze() {
}

// Truth implements starlark.Value.
func (p *Line) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Line) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Line) Attr(name string) (starlark.Value, error) {
	switch name {
	case "start":
		return &p.Start, nil

	case "end":
		return &p.End, nil

	case "angle":
		return starlark.Float(p.Angle), nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "width":
		return starlark.Float(p.Width), nil

	case "tstamp":
		return starlark.String(p.Tstamp), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Line) AttrNames() []string {
	return []string{"start", "end", "angle", "layer", "width", "tstamp"}
}

// SetField implements starlark.HasSetField.
func (p *Line) SetField(name string, val starlark.Value) error {
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

	case "angle":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to angle using type %T", val)
		}
		p.Angle = float64(v)

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)

	case "tstamp":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tstamp using type %T", val)
		}
		p.Tstamp = string(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeArc = starlark.NewBuiltin("Arc", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 *XY
		f1 *XY
		f2 starlark.Float
		f3 starlark.String
		f4 starlark.String
		f5 starlark.Float
	)
	unpackErr := starlark.UnpackArgs(
		"Arc",
		args,
		kwargs,
		"start?", &f0,
		"end?", &f1,
		"angle?", &f2,
		"tstamp?", &f3,
		"layer?", &f4,
		"width?", &f5,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Arc{}

	if f0 != nil {
		out.Start = *f0
	}
	if f1 != nil {
		out.End = *f1
	}
	out.Angle = float64(f2)
	out.Tstamp = string(f3)
	out.Layer = string(f4)
	out.Width = float64(f5)
	return &out, nil
})

func (p *Arc) String() string {
	return fmt.Sprintf("Arc{%v, %v, %v, %v, %v, %v}", p.Start, p.End, p.Angle, p.Tstamp, p.Layer, p.Width)
}

// Type implements starlark.Value.
func (p *Arc) Type() string {
	return "Arc"
}

// Freeze implements starlark.Value.
func (p *Arc) Freeze() {
}

// Truth implements starlark.Value.
func (p *Arc) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Arc) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Arc) Attr(name string) (starlark.Value, error) {
	switch name {
	case "start":
		return &p.Start, nil

	case "end":
		return &p.End, nil

	case "angle":
		return starlark.Float(p.Angle), nil

	case "tstamp":
		return starlark.String(p.Tstamp), nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "width":
		return starlark.Float(p.Width), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Arc) AttrNames() []string {
	return []string{"start", "end", "angle", "tstamp", "layer", "width"}
}

// SetField implements starlark.HasSetField.
func (p *Arc) SetField(name string, val starlark.Value) error {
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

	case "angle":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to angle using type %T", val)
		}
		p.Angle = float64(v)

	case "tstamp":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tstamp using type %T", val)
		}
		p.Tstamp = string(v)

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)

	case "width":
		v, ok := val.(starlark.Float)
		if !ok {
			return fmt.Errorf("cannot assign to width using type %T", val)
		}
		p.Width = float64(v)

	}

	return errors.New("no such assignable field: " + name)
}

var MakeText = starlark.NewBuiltin("Text", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		f0 starlark.String
		f1 starlark.String
		f2 starlark.String
		f3 *XYZ
		f4 starlark.Bool
		f5 *TextEffects
		f6 starlark.Bool
	)
	unpackErr := starlark.UnpackArgs(
		"Text",
		args,
		kwargs,
		"text?", &f0,
		"layer?", &f1,
		"tstamp?", &f2,
		"at?", &f3,
		"unlocked?", &f4,
		"effects?", &f5,
		"hidden?", &f6,
	)
	if unpackErr != nil {
		return starlark.None, unpackErr
	}
	out := Text{}

	out.Text = string(f0)
	out.Layer = string(f1)
	out.Tstamp = string(f2)
	if f3 != nil {
		out.At = *f3
	}
	out.Unlocked = bool(f4)
	if f5 != nil {
		out.Effects = *f5
	}
	out.Hidden = bool(f6)
	return &out, nil
})

func (p *Text) String() string {
	return fmt.Sprintf("Text{%v, %v, %v, %v, %v, %v, %v}", p.Text, p.Layer, p.Tstamp, p.At, p.Unlocked, p.Effects, p.Hidden)
}

// Type implements starlark.Value.
func (p *Text) Type() string {
	return "Text"
}

// Freeze implements starlark.Value.
func (p *Text) Freeze() {
}

// Truth implements starlark.Value.
func (p *Text) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *Text) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *Text) Attr(name string) (starlark.Value, error) {
	switch name {
	case "text":
		return starlark.String(p.Text), nil

	case "layer":
		return starlark.String(p.Layer), nil

	case "tstamp":
		return starlark.String(p.Tstamp), nil

	case "at":
		return &p.At, nil

	case "unlocked":
		return starlark.Bool(p.Unlocked), nil

	case "effects":
		return &p.Effects, nil

	case "hidden":
		return starlark.Bool(p.Hidden), nil
	}

	return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *Text) AttrNames() []string {
	return []string{"text", "layer", "tstamp", "at", "unlocked", "effects", "hidden"}
}

// SetField implements starlark.HasSetField.
func (p *Text) SetField(name string, val starlark.Value) error {
	switch name {
	case "text":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to text using type %T", val)
		}
		p.Text = string(v)

	case "layer":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to layer using type %T", val)
		}
		p.Layer = string(v)

	case "tstamp":
		v, ok := val.(starlark.String)
		if !ok {
			return fmt.Errorf("cannot assign to tstamp using type %T", val)
		}
		p.Tstamp = string(v)

	case "at":
		v, ok := val.(*XYZ)
		if !ok {
			return fmt.Errorf("cannot assign to at using type %T", val)
		}
		p.At = *v

	case "unlocked":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to unlocked using type %T", val)
		}
		p.Unlocked = bool(v)

	case "effects":
		v, ok := val.(*TextEffects)
		if !ok {
			return fmt.Errorf("cannot assign to effects using type %T", val)
		}
		p.Effects = *v

	case "hidden":
		v, ok := val.(starlark.Bool)
		if !ok {
			return fmt.Errorf("cannot assign to hidden using type %T", val)
		}
		p.Hidden = bool(v)

	}

	return errors.New("no such assignable field: " + name)
}
