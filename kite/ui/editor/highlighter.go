package editor

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

type Highlighter struct {
	lines [1024 * 100]lineState
}

type lineState struct {
	leadingIndentation int
	declaration        string
}

func (l *lineState) Update(start, end *gtk.TextIter, newChars string) {
	line := start.GetSlice(end) + newChars
	// fmt.Println(line)

	var (
		leadingIndentation   int
		processedIndentation bool
		firstWord, currWord  string
	)

	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			if !processedIndentation {
				leadingIndentation++
			} else {
				if firstWord == "" {
					firstWord = currWord
				}
				currWord = ""
			}
		} else {
			processedIndentation = true
			currWord += string(line[i])
		}
	}

	l.leadingIndentation = leadingIndentation
	l.declaration = firstWord
	fmt.Printf("%q - %v\n", line, l)
}

// state of the lex scan when the beginning of this line was reached.
type scanContext struct {
}
