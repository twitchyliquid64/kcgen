package ui

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/twitchyliquid64/kcgen/kcsl"
	"github.com/twitchyliquid64/kcgen/kite/ui/editor"
	"github.com/twitchyliquid64/kcgen/kite/ui/preview"
)

// Controller routes UI events & operations on the backend.
// This object can be considered the nerve center for a KiTE
// window.
type Controller struct {
	win     *Win
	editor  *editor.Editor
	preview *preview.Preview
}

func (c *Controller) LoadFromFile(path string) error {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c.win.Model.scriptPath = path
	c.editor.SetContent(string(d))
	c.win.Model.dirty = false
	c.win.flushState()
	return nil
}

func (c *Controller) Render() {
	c.editor.Restyle()
	content := c.editor.GetContent()
	script, err := kcsl.NewScript([]byte(content), flag.Arg(0), false, &kcsl.WDLoader{}, flag.Args(), func(msg string) {
		glib.IdleAdd(func() {
			b, _ := c.win.console.GetBuffer()
			b.InsertAtCursor(msg + "\n")
		})
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Script initialization failed: %v\n", err)
		return
	}
	defer script.Close()

	if m := script.Mod(); m != nil {
		c.preview.Render(m)
	}
}

func (c *Controller) Save() {
	if c.win.Model.scriptPath != "" {
		content := c.editor.GetContent()
		if err := ioutil.WriteFile(c.win.Model.scriptPath, []byte(content), 0744); err != nil {
			fmt.Fprintf(os.Stderr, "Failed save: %v\n", err)
		} else {
			c.win.Model.dirty = false
			c.win.flushState()
		}
	}
}

func (c *Controller) onTextChange() {
	if c.win.Model.dirty {
		return
	}
	c.win.Model.dirty = true
	c.win.flushState()
}

func (c *Controller) ShowConsole() {
	c.win.tabs.SetCurrentPage(1)
}
func (c *Controller) ShowPreview() {
	c.win.tabs.SetCurrentPage(0)
}
