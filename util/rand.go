package util

import (
	"math/rand"
)

// RandRange returns an integer between lo and hi, inclusive.  That means that
// both lo and hi might are included in the possibilities.
func RandRange(lo, hi int) int {
	r := abs(hi-lo) + 1
	return min(lo, hi) + rand.Intn(r)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
