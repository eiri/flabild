package main

import (
	"bufio"
	"log"
	"os"
	"text/template"

	"github.com/eiri/flabild/pkg/flabild"
)

const (
	MOD_SRC = `// THIS MODULE WAS GENERATED, DO NOT EDIT MANUALLY
package main

import (
	"github.com/mroth/weightedrand/v2"

	"github.com/eiri/flabild/pkg/flabild"
)

func MakeChoices() map[flabild.Pair][]weightedrand.Choice[rune, int] {
	m := make(map[flabild.Pair][]weightedrand.Choice[rune, int])
{{ range $p, $m := . }}
	m[flabild.Pair{ {{index $p 0}}, {{index $p 1}} }] = []weightedrand.Choice[rune, int]{
	{{- range $l, $w := $m }}
		{Item: rune({{ $l }}), Weight: {{ $w }}},
	{{- end }}
	}
{{ end }}
	return m
}
`
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	dictFile := os.Args[1]
	if _, err := os.Stat(dictFile); os.IsNotExist(err) {
		log.Fatalf("Error: missing dictionary file %q", dictFile)
	}

	reduceMap, err := generateMap(dictFile)
	if err != nil {
		log.Fatalf("Error: can't generate triplets: %s", err)
	}

	tmpl, err := template.New("mod").Parse(MOD_SRC)
	if err != nil {
		log.Fatalf("Error: can't parse template: %s", err)
	}

	if err := tmpl.Execute(os.Stdout, reduceMap); err != nil {
		log.Fatalf("Error: can't compile template: %s", err)
	}

	os.Exit(0)
}

// TODO: fun little sub-proj would be to convert this to goroutines
func generateMap(dictFile string) (map[flabild.Pair]map[rune]int, error) {
	file, err := os.Open(dictFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Map
	mapSink := make([]flabild.Triplet, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		triplet := flabild.Triplet{'_', '_', '_'}
		word := scanner.Text() + "|"
		for _, letter := range word {
			triplet[0], triplet[1] = triplet[1], triplet[2]
			triplet[2] = letter
			mapSink = append(mapSink, triplet)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Reduce
	reduceMap := make(map[flabild.Pair]map[rune]int)
	for _, t := range mapSink {
		p := flabild.Pair{t[0], t[1]}
		l := t[2]
		if m, ok := reduceMap[p]; ok {
			if r, ok := m[l]; ok {
				m[l] = r + 1
			} else {
				m[l] = 1
			}
		} else {
			reduceMap[p] = map[rune]int{l: 1}
		}
	}

	return reduceMap, nil
}
