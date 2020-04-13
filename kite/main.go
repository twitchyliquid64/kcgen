package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twitchyliquid64/kcgen/kite/ui"
	"go.starlark.net/resolve"
)

func findGlade() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	for _, p := range []string{
		"kite.glade",
		"kite/kite.glade",
		"~/.kite/kite.glade",
		"/usr/share/kite/kite.glade",
	} {
		p = strings.Replace(p, "~", u.HomeDir, -1)
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", errors.New("could not find resource kite.glade")
}

func main() {
	gtk.Init(nil)
	flag.Parse()
	resolve.AllowFloat = true

	gladePath, err := findGlade()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	win, err := ui.New(gladePath)
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
