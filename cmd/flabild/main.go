package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/eiri/flabild"
)

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

	chooser, err := flabild.NewChooser(p)
	if err != nil {
		log.Fatal(err)
	}

	for {
		word, err := chooser.Word()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(os.Stdout, word)

		if n--; n <= 0 {
			break
		}
	}
}
