package main

import (
	"flag"
	"log"
	"os"
	"plugin"
	"strings"

	wr "github.com/mroth/weightedrand"

	fb "github.com/eiri/flabild/pkg/flabild"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func main() {
	var n int
	flag.IntVar(&n, "number", 1, "number of words to generate")
	flag.IntVar(&n, "n", 1, "number of words to generate")
	flag.Parse()

	p, err := plugin.Open("en.so")
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("MakeChoices")
	if err != nil {
		log.Fatal(err)
	}

	makeChoices, ok := f.(func() map[fb.Pair][]wr.Choice)
	if !ok {
		log.Fatalf("Unexpected type for MakeChoices function")
	}

	choices := makeChoices()

	for {
		var wb strings.Builder
		pair := fb.Pair{'_', '_'}
		for {
			chooser, err := wr.NewChooser(choices[pair]...)
			if err != nil {
				log.Fatalf("Error: Can't create new chooser: %s", err)
			}
			l := chooser.Pick().(rune)
			if l == '|' {
				break
			}
			wb.WriteByte(byte(l))
			pair[0], pair[1] = pair[1], l
		}
		log.Println(wb.String())

		if n--; n <= 0 {
			break
		}
	}
}
