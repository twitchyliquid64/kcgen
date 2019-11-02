package editor

import (
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
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

	css   *gtk.CssProvider
	style *gtk.StyleContext
	tags  *tagSet
}

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

	css, tags, err := makeStyling(buffer)
	if err != nil {
		return nil, err
	}

	buffer.SetText("AAAAAAAAAA")
	buffer.ApplyTag(tags.str, buffer.GetStartIter(), buffer.GetEndIter())

	style, err := editor.GetStyleContext()
	if err != nil {
		return nil, err
	}
	style.AddProvider(css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	out := &Editor{
		editor: editor,
		buffer: buffer,
		css:    css,
		style:  style,
		tags:   tags,
	}

	if _, err := out.buffer.Connect("insert-text", func(tb *gtk.TextBuffer, loc *gtk.TextIter, ins string, len int, e *Editor) {
		e.onInsert(ins, tb, loc, len)
	}, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (e *Editor) onInsert(text string, tb *gtk.TextBuffer, loc *gtk.TextIter, l int) {
	// fmt.Printf("Insert: line %d, char %d\n", loc.GetLine(), loc.GetCharsInLine())
	line := loc.GetLine()
	start, end := tb.GetIterAtLine(line), tb.GetIterAtLine(line+1)

	tb.RemoveAllTags(start, end)
	tok, _ := lexer.Tokenise(nil, start.GetSlice(end)+text)

	var lineOffset int
	for t := tok(); t != chroma.EOF; t = tok() {
		switch t.Type {
		case chroma.Keyword:
			tb.ApplyTag(e.tags.keyword, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
		case chroma.Punctuation:
			if t.Value == "(" || t.Value == ")" {
				tb.ApplyTag(e.tags.str, tb.GetIterAtLineOffset(line, lineOffset), tb.GetIterAtLineOffset(line, lineOffset+len(t.Value)))
			}
		}
		lineOffset += len(t.Value)
	}
}

type tagSet struct {
	str     *gtk.TextTag
	keyword *gtk.TextTag
}

func makeStyling(buffer *gtk.TextBuffer) (*gtk.CssProvider, *tagSet, error) {
	s, err := gtk.CssProviderNew()
	if err != nil {
		return nil, nil, err
	}
	s.LoadFromData(`
    textview { color: green; }
    `)

	return s, &tagSet{
		str: buffer.CreateTag("string", map[string]interface{}{
			"foreground": "blue",
		}),
		keyword: buffer.CreateTag("keyword", map[string]interface{}{
			"foreground": "orange",
		}),
	}, nil
}
