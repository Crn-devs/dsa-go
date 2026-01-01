package selectionsort

import (
	"cmp"
	"fmt"
)

func SelectionSort[S ~[]E, E cmp.Ordered](s S) S {

	var min int

	for i := 0; i < len(s)-1; i++ {
		// set min to be the first element
		min = i
		fmt.Println("min is:", min)

		for j := i + 1; j < len(s); j++ {
			// check the first element (min) againt the rest of the slice
			// if min > next element swap min
			if s[j] < s[min] {
				min = j
			}

			fmt.Println("min is:", min)
		}

		s[i], s[min] = s[min], s[i]
	}
	return s
}
