package flabild

import (
	"math/rand"
	"strings"
	"time"
)

type Pair [2]rune

type Triplet [3]rune

func (t Triplet) String() string {
	var b strings.Builder
	for _, r := range t {
		b.WriteByte(byte(r))
	}
	return b.String()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
