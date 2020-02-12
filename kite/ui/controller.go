package ui

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
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

func (c *Controller) Exit() bool {
	if c.win.Model.dirty {
		d, _ := gtk.DialogNewWithButtons("Save before exiting?", c.win.win, gtk.DIALOG_MODAL, "Save changes", "Discard changes")
		ca, _ := d.GetContentArea()
		l, _ := gtk.LabelNew("Your script has changes that have not been saved. Would you like to save your changes?")
		ca.PackStart(l, true, true, 2)
		l.Show()

		switch d.Run() {
		case 0: // Save changes
			c.Save()
		case gtk.RESPONSE_DELETE_EVENT: // Cancel
			return true
		}
	}

	gtk.MainQuit()
	return false
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
	logScriptErr := func(err error) {
		glib.IdleAdd(func() {
			b, _ := c.win.console.GetBuffer()
			b.InsertAtCursor(fmt.Sprintf("Error: %v\n", err))
		})
	}

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
		logScriptErr(err)
		return
	}
	defer script.Close()

	var buff bytes.Buffer
	if p := script.Pcb(); p != nil {
		c.preview.RenderPCB(p)
		if err := p.Write(&buff); err != nil {
			fmt.Fprintf(os.Stderr, "Write() failed: %v\n", err)
			logScriptErr(err)
			return
		}
	} else if m := script.Mod(); m != nil {
		c.preview.RenderMod(m)
		if err := m.WriteModule(&buff); err != nil {
			fmt.Fprintf(os.Stderr, "WriteModule() failed: %v\n", err)
			logScriptErr(err)
			return
		}
	}

	b, _ := c.win.output.GetBuffer()
	b.SetText(buff.String())
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
	} else {
		if err := c.SaveAs(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed save as: %v\n", err)
		}
	}
}

func (c *Controller) SaveAs() error {
	dl, err := gtk.FileChooserDialogNewWith2Buttons("Save As", c.win.win, gtk.FILE_CHOOSER_ACTION_SAVE, "Cancel", gtk.RESPONSE_CANCEL, "Save", gtk.RESPONSE_OK)
	if err != nil {
		return err
	}
	defer dl.Destroy()

	resp := dl.Run()
	switch resp {
	case gtk.RESPONSE_CANCEL:
		return nil
	case gtk.RESPONSE_OK:
		if fn := dl.GetFilename(); fn != "" {
			c.win.Model.scriptPath = fn
			c.Save()
		}
		return nil
	default:
		return fmt.Errorf("unknown dialog response: %v", resp)
	}
}

func (c *Controller) Open() error {
	dl, err := gtk.FileChooserDialogNewWith2Buttons("Open script", c.win.win, gtk.FILE_CHOOSER_ACTION_SAVE, "Cancel", gtk.RESPONSE_CANCEL, "Open", gtk.RESPONSE_OK)
	if err != nil {
		return err
	}
	defer dl.Destroy()

	resp := dl.Run()
	switch resp {
	case gtk.RESPONSE_CANCEL:
		return nil
	case gtk.RESPONSE_OK:
		if fn := dl.GetFilename(); fn != "" {
			return c.LoadFromFile(fn)
		}
		return nil
	default:
		return fmt.Errorf("unknown dialog response: %v", resp)
	}
}

func (c *Controller) onTextChange() {
	if c.win.Model.dirty {
		return
	}
	c.win.Model.dirty = true
	c.win.flushState()
}

func (c *Controller) ShowPreview() {
	c.win.tabs.SetCurrentPage(0)
}
func (c *Controller) ShowConsole() {
	c.win.tabs.SetCurrentPage(1)
}
func (c *Controller) ShowOutput() {
	c.win.tabs.SetCurrentPage(2)
}
