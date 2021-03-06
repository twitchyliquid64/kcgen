// Binary kcgen scripts the generation of kicad modules and PCBs.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/twitchyliquid64/kcgen/kcsl"
	"go.starlark.net/resolve"
)

var (
	verbose = flag.Bool("verbose", false, "Enables verbose logging.")
	out     = flag.String("o", "-", "Where to write output.")
)

func loadScript(p string) ([]byte, error) {
	d, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if d.IsDir() {
		return nil, fmt.Errorf("%v is a directory", p)
	}
	return ioutil.ReadFile(p)
}

func main() {
	flag.Parse()
	resolve.AllowFloat = true
	sData, err := loadScript(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load script: %v\n", err)
		os.Exit(1)
	}

	script, err := kcsl.NewScript(sData, flag.Arg(0), *verbose, &kcsl.WDLoader{}, flag.Args()[1:], nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Initialization failed: %v\n", err)
		os.Exit(1)
	}

	if err := run(script); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(s *kcsl.Script) error {
	defer s.Close()
	outF := os.Stdout
	if *out != "-" {
		var err error
		outF, err = os.OpenFile(*out, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer outF.Close()
	}

	if m := s.Mod(); m != nil {
		m.WriteModule(outF)
	}

	if *out == "-" {
		fmt.Println()
	}
	return nil
}
