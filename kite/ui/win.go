package ui

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui/editor"
	"github.com/twitchyliquid64/kcgen/kite/ui/preview"
)

// Win encapsulates the UI state of the KiTE window.
type Win struct {
	win          *gtk.Window
	isFullscreen bool
	style        *gtk.StyleContext
	editor       *gtk.TextView

	console *gtk.TextView
	output  *gtk.TextView
	preview *gtk.DrawingArea
	tabs    *gtk.Notebook

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

	tabs, err := b.GetObject("kite_output")
	if err != nil {
		return err
	}
	w.tabs = tabs.(*gtk.Notebook)

	w.style, err = w.win.GetStyleContext()
	if err != nil {
		return fmt.Errorf("GetStyleContext() failed: %v", err)
	}

	editor, err := editor.New(b, w.Controller.onTextChange)
	if err != nil {
		return err
	}
	w.Controller.editor = editor
	console, err := b.GetObject("kite_console")
	if err != nil {
		return err
	}
	w.console = console.(*gtk.TextView)
	output, err := b.GetObject("sexp_output")
	if err != nil {
		return err
	}
	w.output = output.(*gtk.TextView)

	preview, err := preview.NewPreview(b)
	if err != nil {
		return err
	}
	w.Controller.preview = preview

	open, err := b.GetObject("open_button")
	if err != nil {
		return err
	}
	open.(*gtk.ToolButton).Connect("clicked", func() {
		if err := w.Controller.Open(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		}
	})
	save, err := b.GetObject("save_button")
	if err != nil {
		return err
	}
	save.(*gtk.ToolButton).Connect("clicked", func() {
		w.Controller.Save()
	})
	render, err := b.GetObject("render_button")
	if err != nil {
		return err
	}
	render.(*gtk.ToolButton).Connect("clicked", func() {
		w.Controller.Render()
	})
	fullscreen, err := b.GetObject("fullscreen_button")
	if err != nil {
		return err
	}
	fullscreen.(*gtk.ToolButton).Connect("clicked", func() {
		if w.isFullscreen {
			w.win.Unfullscreen()
		} else {
			w.win.Fullscreen()
		}
		w.isFullscreen = !w.isFullscreen
	})

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
		if keyEvent.KeyVal() == gdk.KEY_1 && keyEvent.State()&gdk.GDK_CONTROL_MASK != 0 {
			w.Controller.ShowPreview()
		}
		if keyEvent.KeyVal() == gdk.KEY_2 && keyEvent.State()&gdk.GDK_CONTROL_MASK != 0 {
			w.Controller.ShowConsole()
		}
		if keyEvent.KeyVal() == gdk.KEY_3 && keyEvent.State()&gdk.GDK_CONTROL_MASK != 0 {
			w.Controller.ShowOutput()
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
