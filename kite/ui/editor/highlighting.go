package editor

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type tagSet struct {
	str     *gtk.TextTag
	keyword *gtk.TextTag
	parenth *gtk.TextTag
	op      *gtk.TextTag
	fun     *gtk.TextTag
	pseudo  *gtk.TextTag
	comment *gtk.TextTag
	field   *gtk.TextTag
}

func makeStyling(buffer *gtk.TextBuffer, bg *gdk.RGBA) (*gtk.CssProvider, *tagSet, error) {
	s, err := gtk.CssProviderNew()
	if err != nil {
		return nil, nil, err
	}
	s.LoadFromData(`
		GtkTextView {
		    font-family: monospace;
		}
		textview {
		    font-family: monospace;
		}
    `)

	var strTag, fun *gtk.TextTag
	if f := bg.Floats(); f[0] > 0.75 && f[1] > 0.75 && f[2] > 0.75 { // light background
		strTag = buffer.CreateTag("string", map[string]interface{}{
			"foreground": "#aa00aa",
		})
		fun = buffer.CreateTag("func", map[string]interface{}{
			"foreground": "#211fd4",
		})
	} else { // dark background / theme
		strTag = buffer.CreateTag("string", map[string]interface{}{
			"foreground": "#98c379",
		})
		fun = buffer.CreateTag("func", map[string]interface{}{
			"foreground": "#61afef",
		})
	}

	// function blue: #61afef
	// type green: #56b6c2
	// cool magenta: #c678dd
	// comment grey: #5c6370

	return s, &tagSet{
		str: strTag,
		keyword: buffer.CreateTag("keyword", map[string]interface{}{
			"foreground": "orange",
		}),
		parenth: buffer.CreateTag("parenth", map[string]interface{}{
			//"foreground": "cyan",
			//"weight": pango.WEIGHT_BOLD,
		}),
		op: buffer.CreateTag("op", map[string]interface{}{
			"foreground": "red",
		}),
		field: buffer.CreateTag("field", map[string]interface{}{
			"foreground": "red",
		}),
		fun: fun,
		pseudo: buffer.CreateTag("pseudo", map[string]interface{}{
			"foreground": "#c678dd",
		}),
		comment: buffer.CreateTag("comment", map[string]interface{}{
			"foreground": "#5c6370",
		}),
	}, nil
}
