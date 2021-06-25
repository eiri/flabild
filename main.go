package main

import (
	"log"
	"math/rand"
	"os"
	"time"
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
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")

	n := MIN_WORD_LEN + rand.Intn(MAX_WORD_LEN-MIN_WORD_LEN+1)
	s := make([]rune, n)
	for i := range s {
		s[i] = alphabet[rand.Intn(len(alphabet))]
	}

	log.Println(string(s))
}
