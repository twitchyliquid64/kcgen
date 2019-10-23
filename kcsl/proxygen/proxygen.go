// Binary proxygen generates proxy structures for starlark.
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var (
	filterStruct = flag.String("struct", "", "")
)

var preamble = `package pcb

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/nsf/sexp"
	"go.starlark.net/starlark"
)`

func findStructs(f *ast.File) []*ast.TypeSpec {
	out := make([]*ast.TypeSpec, 0, 6)
	for _, d := range f.Decls {
		if gd, ok := d.(*ast.GenDecl); ok {
			for _, spec := range gd.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := ts.Type.(*ast.StructType); ok {
						out = append(out, ts)
					}
				}
			}
		}
	}
	return out
}

// translatedField represents how a field in a struct should be represented
// and processed by the proxy object of its parent structure.
type translatedField struct {
	Index int

	FieldName    string
	StarlarkName string
	Type         string
	IsPtr        bool
	Proxy        string

	ConstructorType     string // Type the argument should be unpacked to.
	ConstructorStrategy string // How the value should initialize the field.

	// Only set for array/slice types.
	ConstructorRepeatedType       string // Type when decoding from starlark argument.
	RepeatedType                  string // Type which is repeated.
	IsFixedLen                    bool   // true for arrays, false for slices.
	ArrayLen                      int
	NestedRepeatedType            string // Set only if 2D array.
	NestedConstructorRepeatedType string // Set only if 2D array.
	NestedProxy                   string // Set only if 2D array.

	// Only set for map types.
	MapValType string
	MapKeyType string

	Field *ast.Field
}

// translatedStruct represents the result of processing a struct and its
// fields for translation.
type translatedStruct struct {
	Name   string
	Fields []translatedField
	Struct *ast.StructType
}

func translateStruct(name string, st *ast.StructType) (translatedStruct, error) {
	out := translatedStruct{
		Name:   name,
		Struct: st,
		Fields: make([]translatedField, 0, len(st.Fields.List)),
	}

	for i, f := range st.Fields.List {
		if isExportedIdent(f.Names[0].Name) {
			field := translatedField{
				Index:           i,
				Field:           f,
				FieldName:       f.Names[0].Name,
				StarlarkName:    starlarkName(f.Names[0].Name),
				Type:            typeStr(f.Type),
				ConstructorType: constructionTypeStr(f.Type),
			}
			_, field.IsPtr = f.Type.(*ast.StarExpr)

			switch {
			case typeStr(f.Type) == "int":
				field.ConstructorStrategy = "int"

			case isPrimitiveType(f.Type):
				field.ConstructorStrategy = "primitive"

			case isIdentOrIdentPtr(f.Type):
				field.ConstructorStrategy = "ident"

			case isMap(f.Type):
				field.ConstructorStrategy = "map"
				field.ConstructorType = "*starlark.Dict"
				field.MapKeyType = typeStr(f.Type.(*ast.MapType).Key)
				field.MapValType = typeStr(f.Type.(*ast.MapType).Value) + "Proxy"

			case isArrayType(f.Type):
				field.ConstructorStrategy = "array"
				field.ConstructorType = "*starlark.List"
				field.RepeatedType = typeStr(f.Type.(*ast.ArrayType).Elt)
				field.ConstructorRepeatedType = constructionTypeStr(f.Type.(*ast.ArrayType).Elt)
				if field.IsFixedLen = f.Type.(*ast.ArrayType).Len != nil; field.IsFixedLen {
					field.ArrayLen, _ = strconv.Atoi(f.Type.(*ast.ArrayType).Len.(*ast.BasicLit).Value)
				}

				if n, isNested := f.Type.(*ast.ArrayType).Elt.(*ast.ArrayType); isNested {
					field.NestedRepeatedType = typeStr(n.Elt)
					field.NestedConstructorRepeatedType = constructionTypeStr(n.Elt)
					field.ConstructorStrategy = "nested array"
				}
			}

			out.Fields = append(out.Fields, field)
		}
	}
	// fmt.Printf("Fields: %+v\n\n", out.Fields)
	return out, nil
}

var tmpl string = `
{{define "fieldInit" -}}
  {{- if eq .ConstructorStrategy "int" -}}
    {{- "\n  " -}}
    {{- "\n  " -}}if v, ok := f{{.Index}}.Int64(); ok {
    {{- "\n    " -}}out.{{.FieldName}} = {{.Type}}(v)
    {{- "\n  " -}}}
    {{- "\n" -}}

  {{- else if eq .ConstructorStrategy "primitive" -}}
    {{- "  " -}}out.{{.FieldName}} = {{.Field.Type}}(f{{.Index}})
    {{- "\n" -}}

  {{- else if eq .ConstructorStrategy "ident" -}}
    {{- "  " -}}out.{{.FieldName}} = {{if .IsPtr}}&{{end}}f{{.Index}}{{.Proxy}}
    {{- "\n" -}}

	{{- else if eq .ConstructorStrategy "map" -}}
		{{- "\n  " -}}if f{{.Index}} != nil {
		{{- "\n    " -}}for _, k := range f{{.Index}}.Keys() {
		{{- "\n      " -}}v, _, _ := f{{.Index}}.Get(k)
		{{- "\n      " -}}out.{{.FieldName}}[{{.MapKeyType}}(k)] = v.({{.MapValType}})
		{{- "\n    " -}}}
		{{- "\n  " -}}}
    {{- "\n" -}}


  {{- else if eq .ConstructorStrategy "array" -}}
    {{- "  " -}}if f{{.Index}} != nil {
    {{-  if .IsFixedLen -}}
    {{- "\n    " -}}if f{{.Index}}.Len() > {{.ArrayLen}} {
    {{- "\n      " -}}return starlark.None, fmt.Errorf("{{.StarlarkName}} contains %d elements, expected len <= {{.ArrayLen}}", f{{.Index}}.Len())
    {{- "\n    " -}}}
    {{- end -}}
    {{- "\n    " -}}for i := 0; i < f{{.Index}}.Len(); i++ {
		{{- "\n      " -}}s, ok := f{{.Index}}.Index(i).({{.ConstructorRepeatedType}})
		{{- "\n      " -}}if !ok {
		{{- "\n        " -}}return starlark.None, errors.New("{{.StarlarkName}} is not a {{.RepeatedType}}")
		{{- "\n      " -}}}
    {{-  if .IsFixedLen -}}
			{{-  if eq .RepeatedType "int" -}}
			{{- "\n      " -}}if v, ok := s{{.Proxy}}.Int64(); ok {
			{{- "\n        " -}}out.{{.FieldName}}[i] = {{.RepeatedType}}(v)
			{{- "\n      " -}}}
			{{- else -}}
				{{- "\n      " -}}out.{{.FieldName}}[i] = {{.RepeatedType}}(s{{.Proxy}})
			{{- end -}}
    {{- else -}}
    	{{- "\n      " -}}out.{{.FieldName}} = append(out.{{.FieldName}}, {{.RepeatedType}}(s{{.Proxy}}))
    {{- end -}}
		{{- "\n    " -}}}
    {{- "\n  " -}}}
    {{- "\n" -}}

	{{- else if eq .ConstructorStrategy "nested array" -}}
    {{- "  " -}}if f{{.Index}} != nil {
    {{-  if .IsFixedLen -}}
    {{- "\n    " -}}if f{{.Index}}.Len() > {{.ArrayLen}} {
    {{- "\n      " -}}return starlark.None, fmt.Errorf("{{.StarlarkName}} contains %d elements, expected len <= {{.ArrayLen}}", f{{.Index}}.Len())
    {{- "\n    " -}}}
    {{- end -}}
    {{- "\n    " -}}for x := 0; x < f{{.Index}}.Len(); x++ {
		{{- "\n      " -}}s, ok := f{{.Index}}.Index(x).(*starlark.List)
		{{- "\n      " -}}if !ok {
		{{- "\n        " -}}return starlark.None, errors.New("{{.StarlarkName}} is not a list")
		{{- "\n      " -}}}
		{{- "\n      " -}}
		{{- "\n      " -}}var tmp {{.RepeatedType}}
		{{- "\n      " -}}for i := 0; i < s.Len(); i++ {
		{{- "\n        " -}}v, ok := s.Index(i).(*{{.NestedConstructorRepeatedType}})
		{{- "\n        " -}}if !ok {
		{{- "\n          " -}}return starlark.None, fmt.Errorf("{{.StarlarkName}}[%d] is not a {{.NestedRepeatedType}}", x)
		{{- "\n        " -}}}
		{{- "\n        " -}}tmp = append(tmp, v{{.NestedProxy}})
		{{- "\n      " -}}}
		{{- "\n      " -}}out.{{.FieldName}} = append(out.{{.FieldName}}, tmp)
		{{- "\n    " -}}}
    {{- "\n  " -}}}
    {{- "\n" -}}

  {{- end -}}
{{- end -}}

var Make{{.Name}} = starlark.NewBuiltin("{{.Name}}", func(t *starlark.Thread, f *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
  var (
    {{- range $f := .Fields -}}
      {{- "\n    " -}}
      f{{$f.Index}} {{"" -}}
      {{$f.ConstructorType -}}
    {{- end -}}
  {{- "\n  )\n  " -}}


  unpackErr := starlark.UnpackArgs(
    "{{.Name}}",
    args,
    kwargs,
    {{range $f := .Fields -}}
      "{{$f.StarlarkName}}?", &f{{$f.Index}},
    {{end -}})
  if unpackErr != nil {
    return starlark.None, unpackErr
  }
  out := {{.Name}}{}
  {{- "\n\n" -}}

  {{- range $f := .Fields -}}
    {{template "fieldInit" $f}}
  {{- end -}}
  {{- "  " -}}

  return &out, nil
})

func (p *{{.Name}}) String() string {
	return fmt.Sprintf("{{.Name}}{{- "{" -}}
	{{- $fl := len .Fields -}}
	{{- range $i, $f := .Fields -}}
		%v
		{{- if renderSeparator $i $fl -}}
			{{- ", " -}}
		{{- end -}}
	{{- end -}}}"

	{{- ", " -}}

	{{- range $i, $f := .Fields -}}
		p.{{- $f.FieldName -}}
		{{- if renderSeparator $i $fl -}}
			{{- ", " -}}
		{{- end -}}
	{{- end -}})
}

// Type implements starlark.Value.
func (p *{{.Name}}) Type() string {
	return "{{.Name}}"
}

// Freeze implements starlark.Value.
func (p *{{.Name}}) Freeze() {
}

// Truth implements starlark.Value.
func (p *{{.Name}}) Truth() starlark.Bool {
	return starlark.Bool(true)
}

// Hash implements starlark.Value.
func (p *{{.Name}}) Hash() (uint32, error) {
	h := sha256.Sum256([]byte(fmt.Sprintf("%+v", p)))
	return uint32(uint32(h[0]) + uint32(h[1])<<8 + uint32(h[2])<<16 + uint32(h[3])<<24), nil
}

// Attr implements starlark.Value.
func (p *{{.Name}}) Attr(name string) (starlark.Value, error) {
  switch name {
		{{- $fl := len .Fields -}}
		{{- range $i, $f := .Fields -}}
		{{- "\n    " -}}case "{{$f.StarlarkName}}":
		{{- "\n      " -}}

		{{- if eq $f.ConstructorStrategy "primitive" -}}
			{{- if eq $f.Type "string" -}}
				return starlark.String(p.{{$f.FieldName}}), nil
			{{- else if eq $f.Type "float64" -}}
				return starlark.Float(p.{{$f.FieldName}}), nil
			{{- else if eq $f.Type "bool" -}}
				return starlark.Bool(p.{{$f.FieldName}}), nil
			{{- end -}}

		{{- else if eq $f.ConstructorStrategy "int" -}}
			return starlark.MakeInt(p.{{$f.FieldName}}), nil

		{{- else if eq $f.ConstructorStrategy "ident" -}}
			return &p.{{$f.FieldName}}, nil

		{{- else if eq $f.ConstructorStrategy "array" -}}
		l := starlark.NewList(nil)
    {{- "\n      " -}}for _, e := range p.{{$f.FieldName}} {
		{{- if eq $f.RepeatedType "string" -}}
		{{- "\n        " -}}l.Append(starlark.String(e))
		{{- else -}}
    {{- "\n        " -}}l.Append(e)
		{{- end -}}
    {{- "\n      " -}}}
    {{- "\n      " -}}return l, nil

		{{- end -}}

		{{- "\n  " -}}
		{{- end}}}

  return nil, starlark.NoSuchAttrError(fmt.Sprintf("%s has no attribute %s", p.Type(), name))
}

// AttrNames implements starlark.Value.
func (p *{{.Name}}) AttrNames() []string {
	return []string{
		{{- $fl := len .Fields -}}
		{{- range $i, $f := .Fields -}}
			"{{.StarlarkName}}"
			{{- if renderSeparator $i $fl -}}
				{{- ", " -}}
			{{- end -}}
		{{- end -}}
	}
}

// SetField implements starlark.HasSetField.
func (p *{{.Name}}) SetField(name string, val starlark.Value) error {
  switch name {

	{{- $fl := len .Fields -}}
	{{- range $i, $f := .Fields -}}
	{{- "\n    " -}}case "{{$f.StarlarkName}}":
	{{- "\n      " -}}

	{{ if eq $f.ConstructorStrategy "ident" -}}
		v, ok := val.(*{{$f.ConstructorType}})
		{{- "\n      " -}}
		if !ok {
		{{- "\n        " -}}
			return fmt.Errorf("cannot assign to {{$f.StarlarkName}} using type %T", val){{- "\n      " -}}
		}
		{{- "\n      " -}}
		p.{{$f.FieldName}} = *v
		{{- "\n      " -}}
	{{ else if eq $f.ConstructorStrategy "int" -}}
		v, ok := val.(*{{$f.ConstructorType}})
		{{- "\n      " -}}
		if !ok {
		{{- "\n        " -}}
			return fmt.Errorf("cannot assign to {{$f.StarlarkName}} using type %T", val){{- "\n      " -}}
		}
		{{- "\n      " -}}
		i, ok := v.Int64()
		{{- "\n      " -}}
		if !ok {
		{{- "\n        " -}}
			return fmt.Errorf("cannot convert %v to int64", v){{- "\n      " -}}
		}
		{{- "\n      " -}}
		p.{{$f.FieldName}} = {{$f.Type}}(i)
		{{- "\n      " -}}

  {{ else if eq $f.ConstructorStrategy "array" -}}
  	v, ok := val.({{$f.ConstructorType}})
  	{{- "\n      " -}}
  	if !ok {
  	{{- "\n        " -}}
  		return fmt.Errorf("cannot assign to {{$f.StarlarkName}} using type %T", val){{- "\n      " -}}
  	}
	  {{- "\n      " -}}
		{{-  if $f.IsFixedLen -}}
    {{- "\n    " -}}if v.Len() > {{$f.ArrayLen}} {
    {{- "\n      " -}}return starlark.None, fmt.Errorf("{{$f.StarlarkName}} contains %d elements, expected len <= {{$f.ArrayLen}}", v.Len())
    {{- "\n    " -}}}
    {{- end -}}
    {{- "\n      " -}}for i := 0; i < v.Len(); i++ {
		{{- "\n        " -}}s, ok := v.Index(i).({{$f.ConstructorRepeatedType}})
		{{- "\n        " -}}if !ok {
		{{- "\n          " -}}return errors.New("{{$f.StarlarkName}} is not a {{$f.RepeatedType}}")
		{{- "\n        " -}}}
		{{-  if $f.IsFixedLen -}}
		{{- "\n        " -}}p.{{$f.FieldName}}[i] = {{$f.RepeatedType}}(s)
    {{- else -}}
    {{- "\n        " -}}p.{{$f.FieldName}} = append(p.{{$f.FieldName}}, {{$f.RepeatedType}}(s))
    {{- end -}}
		{{- "\n      " -}}}

	{{- else -}}
		v, ok := val.({{$f.ConstructorType}})
		{{- "\n      " -}}
		if !ok {
		{{- "\n        " -}}
			return fmt.Errorf("cannot assign to {{$f.StarlarkName}} using type %T", val){{- "\n      " -}}
		}
		{{- "\n      " -}}
		p.{{$f.FieldName}} = {{$f.Type}}(v)
		{{- "\n      " -}}
	{{- end -}}

	{{- "\n  " -}}
	{{- end}}}

  return errors.New("no such assignable field: " + name)
}
`

func isPrimitiveType(in interface{}) bool {
	if ident, ok := in.(*ast.Ident); ok {
		for _, t := range types.Typ {
			if t.Name() == ident.Name {
				return true
			}
		}
	}
	return false
}

// isExportedIdent returns true if the identifier is exported.
func isExportedIdent(in string) bool {
	return strings.ToUpper(string(in[0])) == string(in[0])
}

func isIdentOrIdentPtr(in interface{}) bool {
	_, ok := in.(*ast.Ident)
	if ok {
		return true
	}

	// Check if pointer to identifier.
	if se, ok := in.(*ast.StarExpr); ok {
		_, ok := se.X.(*ast.Ident)
		return ok
	}
	return false
}

func isMap(in interface{}) bool {
	_, ok := in.(*ast.MapType)
	return ok
}

// starlarkName converts the upper camel case identifier into a lower snake
// case identifier.
func starlarkName(in string) string {
	var out strings.Builder
	for i, c := range in {
		if strings.ToUpper(string(c)) != string(c) {
			out.WriteString(string(c))
		} else {
			if i > 0 && i < len(in)-1 {
				out.WriteByte('_')
			}
			out.WriteString(strings.ToLower(string(c)))
		}
	}
	return out.String()
}

func isArrayType(in interface{}) bool {
	_, ok := in.(*ast.ArrayType)
	return ok
}

func typeStr(in interface{}) string {
	if ident, ok := in.(*ast.Ident); ok {
		// for _, t := range types.Typ {
		// 	if t.Name() == ident.Name {
		// 		return ident.Name
		// 	}
		// }
		// return "pcb." + ident.Name
		return ident.Name
	}
	if at, ok := in.(*ast.ArrayType); ok {
		if at.Len == nil {
			return "[]" + typeStr(at.Elt)
		}
		return fmt.Sprintf("[%s]", at.Len.(*ast.BasicLit).Value) + typeStr(at.Elt)
	}
	if mt, ok := in.(*ast.MapType); ok {
		return fmt.Sprintf("map[%s]%s", typeStr(mt.Key), typeStr(mt.Value))
	}
	if s, ok := in.(*ast.StarExpr); ok {
		return "*" + typeStr(s.X)
	}
	if _, ok := in.(*ast.StructType); ok {
		panic("cannot generate proxies for anonymous structs")
	}
	if sel, ok := in.(*ast.SelectorExpr); ok {
		return fmt.Sprintf("%s.%s", sel.X.(*ast.Ident).Name, sel.Sel.Name)
	}

	return fmt.Sprintf("%T", in)
}

func constructionTypeStr(in interface{}) string {
	t := typeStr(in)
	switch t {
	case "string", "*string":
		return "starlark.String"
	case "float64", "float32", "*float64", "*float32":
		return "starlark.Float"
	case "bool", "*bool":
		return "starlark.Bool"
	case "int", "*int":
		return "starlark.Int"
	}
	return t
}

func makeTemplate() (*template.Template, error) {
	return template.New("").Funcs(template.FuncMap{
		"isExportedIdent":     isExportedIdent,
		"isPrimitiveType":     isPrimitiveType,
		"typeStr":             typeStr,
		"constructionTypeStr": constructionTypeStr,

		"makeStarlarkStr": starlarkName,
		"renderSeparator": func(i int, length int) bool {
			return i < (length - 1)
		},
	}).Parse(tmpl)
}

func main() {
	flag.Parse()
	fset := token.NewFileSet()

	t, err := makeTemplate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Template initialization failed: %v", err)
		return
	}

	os.Stdout.WriteString(preamble)

	var structs []translatedStruct
	for i := 0; i < flag.NArg(); i++ {
		f, err := parser.ParseFile(fset, flag.Arg(i), nil, parser.AllErrors)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse failed: %v", err)
			return
		}

		for _, st := range findStructs(f) {
			if *filterStruct != "" && *filterStruct != st.Name.Name { //TODO: Remove
				continue
			}

			s, err := translateStruct(st.Name.Name, st.Type.(*ast.StructType))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Translation failed for struct %q: %v", st.Name.Name, err)
				return
			}
			structs = append(structs, s)
		}
	}

	if err := translateAdjacent(structs); err != nil {
		fmt.Fprintf(os.Stderr, "translateAdjacent failed: %v", err)
		return
	}

	for _, s := range structs {
		if err := t.Execute(os.Stdout, s); err != nil {
			fmt.Fprintf(os.Stdout, "Failed to generate %v: %v\n", s.Name, err)
		}
	}
}

// translateAdjacent makes structs which reference other structs use their
// proxy objects.
func translateAdjacent(structs []translatedStruct) error {
	proxies := map[string]*translatedStruct{}
	for _, s := range structs {
		proxies[s.Name] = &s
	}

	for i := range structs {
		for x := range structs[i].Fields {
			if structs[i].Fields[x].IsPtr {
				// Pointer messes up the detection of the structure name. Handle specially.
				if _, ok := proxies[structs[i].Fields[x].ConstructorType[1:]]; ok {
					structs[i].Fields[x].Proxy = "." + structs[i].Fields[x].ConstructorType[1:]
					structs[i].Fields[x].ConstructorType = structs[i].Fields[x].ConstructorType + "Proxy"
				}
			} else {
				if _, ok := proxies[structs[i].Fields[x].ConstructorType]; ok {
					structs[i].Fields[x].Proxy = "." + structs[i].Fields[x].ConstructorType
					structs[i].Fields[x].ConstructorType = structs[i].Fields[x].ConstructorType + "Proxy"
				}
			}

			for _, innerType := range []string{"string", "int"} {
				if strings.HasPrefix(structs[i].Fields[x].ConstructorType, "map["+innerType+"]") {
					final := structs[i].Fields[x].ConstructorType[len("map["+innerType+"]"):]
					if _, ok := proxies[final]; ok {
						structs[i].Fields[x].ConstructorType += "Proxy"
						structs[i].Fields[x].Type += "Proxy"
						structs[i].Fields[x].Proxy = "." + final
					}
				}
			}

			if structs[i].Fields[x].ConstructorStrategy == "array" {
				if _, ok := proxies[structs[i].Fields[x].ConstructorRepeatedType]; ok {
					structs[i].Fields[x].Proxy = "." + structs[i].Fields[x].RepeatedType
					structs[i].Fields[x].ConstructorRepeatedType = "*" + structs[i].Fields[x].ConstructorRepeatedType + "Proxy"
					structs[i].Fields[x].RepeatedType = "pcb." + structs[i].Fields[x].RepeatedType
				} else if _, ok := proxies[structs[i].Fields[x].ConstructorRepeatedType[1:]]; ok {
					structs[i].Fields[x].Proxy = "." + structs[i].Fields[x].RepeatedType[1:]
					structs[i].Fields[x].ConstructorRepeatedType = "*" + structs[i].Fields[x].ConstructorRepeatedType[1:] + "Proxy"
					structs[i].Fields[x].RepeatedType = "&"
				}
			}
			if _, ok := proxies[structs[i].Fields[x].NestedConstructorRepeatedType]; ok {
				structs[i].Fields[x].NestedProxy = "." + structs[i].Fields[x].NestedConstructorRepeatedType
				structs[i].Fields[x].RepeatedType = "[]pcb." + structs[i].Fields[x].NestedConstructorRepeatedType
				structs[i].Fields[x].NestedConstructorRepeatedType = structs[i].Fields[x].NestedConstructorRepeatedType + "Proxy"
			}

			// Special cases for enum values.
			switch structs[i].Fields[x].ConstructorType {
			case "PadShape", "TextJustify", "ModTextKind", "PadSurface", "ZoneConnectMode":
				structs[i].Fields[x].ConstructorType = "pcb." + structs[i].Fields[x].ConstructorType
				structs[i].Fields[x].Type = "pcb." + structs[i].Fields[x].Type
			}
		}
	}
	return nil
}
