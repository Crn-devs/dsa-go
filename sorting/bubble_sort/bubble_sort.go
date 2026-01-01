package bubblesort

import (
	"cmp"
)

func BubbleSort[S ~[]E, E cmp.Ordered](s S) S {

	// optimise
	// using a bool no swap
	var noSwap bool
	// optimise

	if len(s) < 2 {
		return s
	}

	for i := len(s); i > 1; i-- {
		noSwap = true
		for j := 0; j < i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				noSwap = false
			}
		}
		if noSwap {
			return s
		}
	}
	return s
}
