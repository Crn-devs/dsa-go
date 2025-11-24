package main

import "slices"

func AreThereDuplicatesPointer(args ...int) bool {

	// check for empty input
	if len(args) <= 0 {
		return false
	}

	// sort the input for double pointer to work
	// **creates adjecent pairs**

	slices.Sort(args)

	for i := 1; i < len(args); i++ {
		if args[i] == args[i-1] {
			return true
		}
	}

	return false
}
