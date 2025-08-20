package bigo

import "fmt"

// anotherFunChallenge demonstrates O(n) time.
// Operation count model: 4 + 7n  =>  O(n)
func AnotherFunChallenge(input int) {
	a := 5  // O(1)
	b := 10 // O(1)
	c := 50 // O(1)
	_ = a
	_ = b
	_ = c // avoid unused vars

	for i := 0; i < input; i++ { // O(n) iterations
		x := i + 1 // O(1) per iteration
		y := i + 2 // O(1) per iteration
		z := i + 3 // O(1) per iteration
		_ = x
		_ = y
		_ = z
	}

	for j := 0; j < input; j++ { // O(n) iterations
		p := j * 2 // O(1) per iteration
		q := j * 2 // O(1) per iteration
		_ = p
		_ = q
	}

	whoAmI := "I don't know" // O(1)
	_ = whoAmI
}

// funChallenge demonstrates a simple O(n) function
func FunChallenge(input []int) int {
	a := 10    // O(1)
	a = 50 + 3 // O(1)

	for i := 0; i < len(input); i++ { // O(n)
		anotherFunction() // O(1) each iteration
		stranger := true  // O(1) each iteration
		_ = stranger      // avoid unused variable
		a++               // O(1) each iteration
	}

	return a // O(1)
}

// Add it up
// O(1) + O(1) + O(n) + O(1) = O(n).
// Constants are dropped in Big-O, so it’s O(n) overall.

// anotherFunction()
func anotherFunction() {
	// do something trivial
	fmt.Println("another function called")
}
