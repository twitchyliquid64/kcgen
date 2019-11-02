package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui"
	"go.starlark.net/resolve"
)

func main() {
	gtk.Init(nil)
	flag.Parse()
	resolve.AllowFloat = true

	win, err := ui.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize GUI: %v\n", err)
		os.Exit(1)
	}

	// TODO: Tidy this spaghetti.
	if flag.Arg(0) != "" {
		win.Controller.LoadFromFile(flag.Arg(0))
	}
	gtk.Main()
}
