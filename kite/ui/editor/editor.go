package editor

import (
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var lexer chroma.Lexer

func init() {
	lexer = lexers.Get("python")
}

// Editor manages a starlark code editor.
type Editor struct {
	editor *gtk.TextView
	buffer *gtk.TextBuffer

	bg    *gdk.RGBA
	css   *gtk.CssProvider
	style *gtk.StyleContext
	tags  *tagSet

	insertSig glib.SignalHandle
}

// New creates a new KiTE KCSL editing widget.
func New(b *gtk.Builder) (*Editor, error) {
	e, err := b.GetObject("kite_editor")
	if err != nil {
		return nil, err
	}
	editor := e.(*gtk.TextView)

	buffer, err := editor.GetBuffer()
	if err != nil {
		return nil, err
	}

	style, err := editor.GetStyleContext()
	if err != nil {
		return nil, err
	}

	bg, err := style.GetProperty("background-color", gtk.STATE_FLAG_BACKDROP)
	if err != nil {
		return nil, fmt.Errorf("GetProperty('background-color') failed: %v", err)
	}

	css, tags, err := makeStyling(buffer, bg.(*gdk.RGBA))
	if err != nil {
		return nil, err
	}
	style.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	buffer.SetText("# Welcome to KiTE! :D")

	out := &Editor{
		editor: editor,
		buffer: buffer,
		css:    css,
		style:  style,
		tags:   tags,
		bg:     bg.(*gdk.RGBA),
	}

	if out.insertSig, err = out.buffer.Connect("insert-text", func(tb *gtk.TextBuffer, loc *gtk.TextIter, ins string, len int, e *Editor) {
		e.onInsert(ins, tb, loc, len)
	}, out); err != nil {
		return nil, err
	}
	if _, err := out.editor.Connect("backspace", out.onBackspace, out); err != nil {
		return nil, err
	}
	if _, err := out.editor.Connect("paste-clipboard", func() {
		go func() {
			time.Sleep(time.Millisecond * 30)
			glib.IdleAdd(func() {
				out.Restyle()
			})
		}()
	}, out); err != nil {
		return nil, err
	}

	glib.IdleAdd(func() {
		out.Restyle()
	})

	return out, nil
}

// Restyle styles the editor from scratch.
func (e *Editor) Restyle() {
	for i := 0; i < e.buffer.GetLineCount(); i++ {
		start, end := e.buffer.GetIterAtLine(i), e.buffer.GetIterAtLine(i+1)
		content := start.GetSlice(end)
		tok, _ := lexer.Tokenise(nil, content)
		e.processLine(tok, content, e.buffer, start, end, i)
	}
}

// SetContent changes the contents of the editor.
func (e *Editor) SetContent(content string) {
	e.buffer.SetText(content)
	e.Restyle()
}

// GetContent returns the contents of the editor.
func (e *Editor) GetContent() string {
	t, _ := e.buffer.GetText(e.buffer.GetStartIter(), e.buffer.GetEndIter(), false)
	return t
}

func (e *Editor) processLine(tok func() chroma.Token, content string, tb *gtk.TextBuffer, start, end *gtk.TextIter, line int) {
	tb.RemoveAllTags(start, end)

	var (
		lineOffset      int
		lastType        chroma.Token
		lastStartOffset int
	)
	for t := tok(); t != chroma.EOF; t = tok() {
		// fmt.Println(t)

		switch t.Type {
		case chroma.Operator:
			if t.Value != "." {
				tb.ApplyTag(e.tags.op, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
			}
		case chroma.String:
			tb.ApplyTag(e.tags.str, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.Keyword:
			tb.ApplyTag(e.tags.keyword, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.NameFunction:
			tb.ApplyTag(e.tags.fun, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.NameBuiltinPseudo:
			tb.ApplyTag(e.tags.pseudo, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.CommentSingle:
			tb.ApplyTag(e.tags.comment, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.LiteralStringDouble, chroma.LiteralStringAffix:
			tb.ApplyTag(e.tags.str, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.Punctuation:
			switch t.Value {
			case "(":
				if lastType.Type == chroma.Name {
					tb.ApplyTag(e.tags.fun, tb.GetIterAtLineOffset(line, lastStartOffset), tb.GetIterAtLineOffset(line, lineOffset))
				}
				fallthrough
			case ")":
				tb.ApplyTag(e.tags.parenth, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
			}
		case chroma.Name:
			if lastType.Type == chroma.Operator && lastType.Value == "." {
				tb.ApplyTag(e.tags.field, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
			}
		}

		lastStartOffset = lineOffset
		lineOffset += len(t.Value)
		lastType = t
	}
}

func (e *Editor) onBackspace() {
	iter := e.buffer.GetIterAtMark(e.buffer.GetMark("insert"))
	line := iter.GetLine()
	start, end := e.buffer.GetIterAtLine(line), e.buffer.GetIterAtLine(line+1)
	content := start.GetSlice(end)

	// Do backspace
	if len(content) > 0 {
		content = content[:len(content)-1]
	}

	tok, _ := lexer.Tokenise(nil, content)
	e.processLine(tok, content, e.buffer, start, end, line)
}

func (e *Editor) onInsert(text string, tb *gtk.TextBuffer, loc *gtk.TextIter, l int) {
	//fmt.Printf("Insert: line %d, pos %d, char %q \n", loc.GetLine(), loc.GetCharsInLine(), text)
	line := loc.GetLine()

	// Handle newline case specially.
	if text == "\n" {
		start, end := tb.GetIterAtLine(line), tb.GetIterAtLine(line+1)
		content := start.GetSlice(end) + text
		tok, _ := lexer.Tokenise(nil, content)

		glib.IdleAdd(func() {
			nl := tb.GetIterAtLine(line + 1)
			t := tok()

			numSpaces := 0
			if strings.HasSuffix(strings.TrimRight(content, "\n"), ":") {
				numSpaces += 2
			}
			if t.Type == chroma.Text && len(t.Value) > 0 && t.Value[0] == ' ' && len(t.Value) < len(content)-1 {
				numSpaces += len(t.Value)
			}
			if numSpaces > 0 {
				tb.Insert(nl, strings.Repeat(" ", numSpaces))
			}
		})
		return
	}

	// Schedule a re-tag of the current line.
	glib.IdleAdd(func() {
		start, end := tb.GetIterAtLine(line), tb.GetIterAtLine(line+1)
		content := start.GetSlice(end) + text
		tok, _ := lexer.Tokenise(nil, content)
		e.processLine(tok, content, tb, start, end, line)
	})
}

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
