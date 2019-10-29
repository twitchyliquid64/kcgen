// Package toolbox provides parameterized generators for common components.
package toolbox

import (
	"fmt"
	"reflect"
)

// Parameter describes an input to a tool.
type Parameter struct {
	Name  string `json:"name"`
	Value interface{}
	Type  reflect.Kind
}

// Tool represents a generator.
type Tool struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}

// SetParameter sets the value of the parameter, returning an error if
// the named parameter does not exist, or the wrong type was provided.
func (t *Tool) SetParameter(name string, value interface{}) error {
	for i, p := range t.Parameters {
		if p.Name == name {
			if k := reflect.TypeOf(value).Kind(); k != p.Type {
				return fmt.Errorf("cannot set parameter of type %v using value of type %v", p.Type, k)
			}
			t.Parameters[i].Value = value
			break
		}
	}
	return fmt.Errorf("no such parameter: %v", name)
}
