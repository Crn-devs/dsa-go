package main

/**
Frequency Counter / Multiple Pointers - areThereDuplicates

Implement a function called, areThereDuplicates which accepts a variable number of arguments,
and checks whether there are any duplicates among the arguments passed in.
You can solve this using the frequency counter pattern OR the multiple pointers pattern.

**/

func AreThereDuplicates(args ...int) bool {

	// check for empty input
	if len(args) <= 0 {
		return false
	}

	// put each integer into a map with the integer as the key and a count as the value
	dup := make(map[int]int)
	for _, value := range args {
		dup[value]++
		if dup[value] > 1 {
			return true
		}
	}
	return false
}
