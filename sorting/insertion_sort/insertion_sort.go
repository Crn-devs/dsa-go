package insertionsort

import (
	"cmp"
)

func InsertionSort[S ~[]E, E cmp.Ordered](s S) S {

	for i := 1; i < len(s); i++ {

		curr := s[i]
		var j int

		for j = i - 1; j >= 0 && s[j] > curr; j-- {
			s[j+1] = s[j]
		}
		s[j+1] = curr

	}
	return s
}
