package bigo

import (
	"fmt"
)

/**

Big O - How an algorithms resources grow (time, space) relative to its input size (n)

O(1) - Constant time: Execution time does not grow with n (accessing an array element by index)

O(log n) - Logarithmic time: Grows slowly

O(n) - Linear time: Work grows proportionally to (n) (scanning a list)

O(n log(n)) - Linearithmic

O(n^2) - Quadratic: nested loops over the same data

Loop over all elements once → O(n)

Nested loops → O(n²), O(n³), etc.

Halving problem each step → O(log n)

Halving but touching all elements → O(n log n)

Recursive branching (recomputing stuff) → O(2ⁿ)

Generating permutations → O(n!)

**/

/**
Big O Rule Book


Rule 1. Worst Case: Big-O always describes the worst-case performance of an algorithm.
 For example, even if you add a return inside a linear search loop to exit early
 the Big-O is still O(n) because in the worst case the search must check every element.

Rule 2. Drop the constants: When calculating big o we discard the constants this is because
 constants dont change the growth trend as input becomes larger

Rule 3. Different terms for inputs: if an algorithm works on two separate inputs of different
 sizes then 2 different variables are used to express the complexity O(a+b)

Rule 4. Drop non dominants: when simplifying big o we keep only the dominant terms the term that grows
 fastest with n is kept and the rest are dropped eg, O(n² + n + 1) +n and +1 become insignificant compared
 to n² and are thus dropped leaving us with the big o O(n²).

**/

// O(n) - linear Time
// An algorithm with O(n) time complexity grows proportionally with its input size n.
// As n increases, the runtime increases linearly — doubling n roughly doubles the work the algorithm does.
func FindNemo(arr []string) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == "nemo" {
			fmt.Println("Found...")
			return
		}
	}
}

// O(1) - Constant Time
// An algorithm with o(1) time complexity does not grow with the input size
// Execution time is independent of n
// whether an algorithm has 10 inputs or 10 million the time it takes is the same
func GetFirstFish(arr []string) {
	fmt.Println(arr[0])
}
