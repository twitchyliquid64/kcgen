package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui/editor"
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

	preview, err := b.GetObject("kite_preview")
	if err != nil {
		return err
	}
	w.preview = preview.(*gtk.DrawingArea)

	w.win.SetDefaultSize(1000, 700)
	w.win.Connect("destroy", gtk.MainQuit)
	w.win.SetResizable(true)
	w.win.ShowAll()

	return w.setupKeyBindings()
}

func (w *Win) setupKeyBindings() error {
	return nil
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
