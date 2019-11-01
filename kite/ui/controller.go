package ui

import (
	"github.com/twitchyliquid64/kcgen/kite/ui/editor"
)

// Controller routes UI events & operations on the backend.
// This object can be considered the nerve center for a KiTE
// window.
type Controller struct {
	win    *Win
	editor *editor.Editor
}
