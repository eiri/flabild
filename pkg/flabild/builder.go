package flabild

import (
	"fmt"
	"plugin"
	"strings"

	"github.com/mroth/weightedrand/v2"
)

const PLUGIN_API_FUNC = "BuildChoicesMap"

type Pair [2]rune

func NewPair(a, b rune) Pair {
	return Pair{a, b}
}

type PairMap map[Pair][]weightedrand.Choice[rune, int]

type flabildPlugin interface {
	Lookup(string) (plugin.Symbol, error)
}

type Chooser struct {
	choices PairMap
}

func NewChooser(p flabildPlugin) (*Chooser, error) {
	f, err := p.Lookup(PLUGIN_API_FUNC)
	if err != nil {
		return nil, fmt.Errorf("can't load plugin API: %w", err)
	}

	buildChoicesMap, ok := f.(func() PairMap)
	if !ok {
		return nil, fmt.Errorf("invalid plugin API signature")
	}

	return &Chooser{choices: buildChoicesMap()}, nil
}

func (c *Chooser) Word() (string, error) {
	var wb strings.Builder
	pair := Pair{'_', '_'}
	for {
		chooser, err := weightedrand.NewChooser(c.choices[pair]...)
		if err != nil {
			return "", fmt.Errorf("can't create new chooser: %w", err)
		}
		l := chooser.Pick()
		if l == '|' {
			break
		}
		wb.WriteByte(byte(l))
		pair[0], pair[1] = pair[1], l
	}

	return wb.String(), nil
}
