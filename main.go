package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	wr "github.com/mroth/weightedrand"
)

const (
	MIN_WORD_LEN = 4
	MAX_WORD_LEN = 12
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func main() {
	choices := makeChoices()
	chooser, _ := wr.NewChooser(choices...)

	n := MIN_WORD_LEN + rand.Intn(MAX_WORD_LEN-MIN_WORD_LEN+1)
	s := make([]rune, n)
	for i := range s {
		s[i] = chooser.Pick().(rune)
	}

	log.Println(string(s))
}

func makeChoices() []wr.Choice {
	// http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
	var weights = []uint{14810, 2715, 4943, 7874, 21912, 4200, 3693, 10795, 13318, 188, 1257, 7253, 4761, 12666, 14003, 3316, 205, 10977, 11450, 16587, 5246, 2019, 3819, 315, 3853, 128}

	c := make([]wr.Choice, len(alphabet))
	for i := range c {
		c[i] = wr.Choice{Item: alphabet[i], Weight: weights[i]}
	}

	return c
}
