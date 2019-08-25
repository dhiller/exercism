package hamming

import (
	"fmt"
)

// Distance calculates the Hamming distance between two strings.
func Distance(a, b string) (hamming int, err error) {
	if len(a) != len(b) {
		return hamming, fmt.Errorf("a and b have different length (%d != %d)", len(a), len(b))
	}
	runesA := []rune(a)
	runesB := []rune(b)
	for index, runeFromA := range runesA {
		if runesB[index] != runeFromA {
			hamming++
		}
	}
	return hamming, nil
}
