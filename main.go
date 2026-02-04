package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/mroth/weightedrand/v2"

	"github.com/eiri/flabild/pkg/flabild"
)

//go:generate go tool generator english.txt plugin/en/en.go
//go:generate go fmt plugin/en/en.go

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
}

func main() {
	var n int
	var lang string
	flag.IntVar(&n, "number", 1, "number of words to generate")
	flag.IntVar(&n, "n", 1, "number of words to generate")
	flag.StringVar(&lang, "lang", "en", "language to load")
	flag.Parse()

	pluginName := fmt.Sprintf("libflabild-%s.so", lang)
	path := filepath.Join(filepath.Dir(os.Getenv("LD_LIBRARY_PATH")), pluginName)
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("MakeChoices")
	if err != nil {
		log.Fatal(err)
	}

	makeChoices, ok := f.(func() map[flabild.Pair][]weightedrand.Choice[rune, int])
	if !ok {
		log.Fatal("Unexpected type for MakeChoices function")
	}

	choices := makeChoices()

	for {
		var wb strings.Builder
		pair := flabild.Pair{'_', '_'}
		for {
			chooser, err := weightedrand.NewChooser(choices[pair]...)
			if err != nil {
				log.Fatalf("Error: Can't create new chooser: %s", err)
			}
			l := chooser.Pick()
			if l == '|' {
				break
			}
			wb.WriteByte(byte(l))
			pair[0], pair[1] = pair[1], l
		}
		fmt.Fprintln(os.Stdout, wb.String())

		if n--; n <= 0 {
			break
		}
	}
}
