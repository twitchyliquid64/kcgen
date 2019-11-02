package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui"
)

func main() {
	gtk.Init(nil)
	flag.Parse()

	_, err := ui.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize GUI: %v\n", err)
		os.Exit(1)
	}
	gtk.Main()
}
