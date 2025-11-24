package main

import (
	"fmt"
)

/**
Frequency Counter - sameFrequency

Write a function called sameFrequency. Given two positive integers, find out if the two numbers have the same frequency of digits.

Your solution MUST have the following complexities:

Time: O(N)
**/

// restate: check if 2 numbers contain the same numbers if a number in the numbers differs then return false

// what are inputs: 2 numbers
// outputs: not specified but likely a bool to show the options

// can outputs be determined from inputs?: yes a decision on a condition so ill use a bool

// test run
// [123, 12] false
// [123, 321] true
// [1, 1] true

func SameFrequency(num1, num2 int) bool {

	n1 := fmt.Sprintf("%d", num1)
	n2 := fmt.Sprintf("%d", num2)

	if len(n1) != len(n2) {
		return false
	}

	freq1 := make(map[string]int, len(n1))
	freq2 := make(map[string]int, len(n2))

	for i := 0; i < len(n1); i++ {
		freq1[string(n1[i])] = freq1[string(n1[i])] + 1
	}

	for i := 0; i < len(n1); i++ {
		freq2[string(n2[i])] = freq2[string(n2[i])] + 1
	}

	for key, count := range freq1 {

		if freq2[key] != count {
			return false
		}
	}

	return true
}

// additional notes

// can be impoved with a single map and increment-decrement flow of the count
// could use runes in the map to be more memory efficient

// improved version
// using single map, decrement increment pattern and runes
func SameFrequency2(num1, num2 int) bool {

	n1 := fmt.Sprintf("%d", num1)
	n2 := fmt.Sprintf("%d", num2)

	if len(n1) != len(n2) {
		return false
	}

	freq1 := make(map[rune]int, len(n1))

	for _, character := range n1 {
		freq1[character]++
	}

	for _, character := range n2 {
		freq1[character]--
		if freq1[character] < 0 {
			return false
		}
	}

	return true
}
