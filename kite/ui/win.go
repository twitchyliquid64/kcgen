package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui/editor"
	"github.com/twitchyliquid64/kcgen/kite/ui/preview"
)

// Win encapsulates the UI state of the KiTE window.
type Win struct {
	win     *gtk.Window
	style   *gtk.StyleContext
	editor  *gtk.TextView
	preview *gtk.DrawingArea

	Model      WindowModel
	Controller Controller
}

func (w *Win) build() error {
	b, err := gtk.BuilderNewFromFile("kite/kite.glade")
	if err != nil {
		return err
	}

	win, err := b.GetObject("kite_win")
	if err != nil {
		return err
	}
	w.win = win.(*gtk.Window)

	w.style, err = w.win.GetStyleContext()
	if err != nil {
		return fmt.Errorf("GetStyleContext() failed: %v", err)
	}

	editor, err := editor.New(b)
	if err != nil {
		return err
	}
	w.Controller.editor = editor

	preview, err := preview.NewPreview(b)
	if err != nil {
		return err
	}
	w.Controller.preview = preview

	w.win.SetDefaultSize(1000, 700)
	w.win.Connect("destroy", gtk.MainQuit)
	w.win.SetResizable(true)
	w.win.ShowAll()

	return w.setupKeyBindings()
}

func (w *Win) setupKeyBindings() error {
	// TODO: Refactor this into some configurable mapping.
	w.win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if keyEvent.KeyVal() == gdk.KEY_r && keyEvent.State()&gdk.GDK_CONTROL_MASK != 0 {
			w.Controller.Render()
		}
		if keyEvent.KeyVal() == gdk.KEY_s && keyEvent.State()&gdk.GDK_CONTROL_MASK != 0 {
			w.Controller.Save()
		}
	})
	return nil
}

func (w *Win) flushState() {
	title := "KiTE"
	if w.Model.scriptPath != "" {
		title += " - " + w.Model.scriptPath
	}
	if w.Model.dirty {
		title = "*" + title
	}
	w.win.SetTitle(title)
}

// New creates and initializes a new KiTE window.
func New() (*Win, error) {
	out := Win{
		Controller: Controller{},
	}
	out.Controller.win = &out

	if err := out.build(); err != nil {
		return nil, err
	}
	return &out, nil
}
