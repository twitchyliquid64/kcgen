// Package lib contains libraries written in starlark.
package lib

// Libs exposes libraries which can be imported.
var Libs = map[string][]byte{
	"mod.lib": modLib,
}
