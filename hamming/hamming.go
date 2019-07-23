package hamming

import (
	"fmt"
)

// Distance calculates the Hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("a and b have different length (%d != %d)", len(a), len(b))
	}
	hamming := 0
	runesA := []rune(a)
	runesB := []rune(b)
	for index, runeFromA := range runesA {
		if runesB[index] != runeFromA {
			hamming++
		}
	}
	return hamming, nil
}
