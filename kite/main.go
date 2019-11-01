package main

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui"
)

func main() {
	gtk.Init(nil)

	_, err := ui.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize GUI: %v\n", err)
		os.Exit(1)
	}
	gtk.Main()
}
