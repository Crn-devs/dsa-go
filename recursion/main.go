package main

import (
	"fmt"

	recursionproblems "github.com/Crn-devs/dsa-go/recursion/recursion_problems"
)

// recursion - a function that calls itself

// - recursion is everywhere
// -

// CallStack Concept
//   - When a function is called its "pushed" to the call stack
//   - When that function ends it is "popped" from the call stack
//   - When we call a recursive function it would seem that it should constantly push the same functions
//     onto the call stack over and over again with no end thats where a base case comes in to stop the recursion

// recursive essentials
// - base case
// - differing input throughout the calls

func main() {

	// CountDown(50)

	// fmt.Println(sumRange(4))

	// fmt.Println(Factorial(3))

	// fmt.Println(FactorialIteratively(4))

	// fmt.Println(CollectOddNumbers([]int{1, 2, 3, 4, 5, 6, 7}))

	// fmt.Println(recursionproblems.Power(2, 5))

	// fmt.Println(recursionproblems.ProductOfArray([]int{2, 3, 10}))

	// fmt.Println(recursionproblems.RecursiveRange(6))

	// fmt.Println(recursionproblems.Fib(10))

	// fmt.Println(recursionproblems.Reverse("backwards"))

	// fmt.Println(recursionproblems.IsPalindrome("heeh"))
	// fmt.Println(recursionproblems.IsPalindrome("heehs"))

	// fmt.Println(recursionproblems.SomeRecursive([]int{2, 2, 3, 4}, recursionproblems.IsOdd))
	// fmt.Println(recursionproblems.SomeRecursive([]int{4, 6, 8}, recursionproblems.IsOdd))
	// fmt.Println(recursionproblems.SomeRecursive([]int{2, 2, 2, 2, 2, 19}, recursionproblems.IsOdd))

	// fmt.Println(recursionproblems.FlattenArr([]any{[]any{1, 2, 3}, 1, 2, 3, []any{5, 5, 5}}))

	// fmt.Println(recursionproblems.CapitaliseFirst([]string{"hello", "hi", "hey"}))

	fmt.Println(recursionproblems.StringifyNumbers(map[string]any{
		"a": 1,
		"b": map[string]any{
			"c": 2,
		},
	}))

	fmt.Println(recursionproblems.CollectStrings(map[string]any{
		"a": "hello",
		"b": map[string]any{
			"c": "hi again",
		},
	}))
}

func CountDown(num int) {

	// base case is the number less than or equal to zero
	if num <= 0 {
		fmt.Println("All Done")
		return
	}
	fmt.Println(num)
	num--
	// differing input -> we pass an updated num this is the differing input
	CountDown(num) // calls itself with differing input
}

// the above recursive function can be visualised as
/**

countDown(5)
is num <= 0 -> no -> --num
Call Countdown(num)
Repeat the flow

**/

// another example with explaination
func sumRange(num int) int {
	if num == 1 {
		return 1
	}
	return num + sumRange(num-1)
}

/**
num = 3

return 3 + sumRange(2)
				return 2 + sumRange(1)
								return 1
6

num = 4

return 4 + sumRange(3)
				return 3 + sumRange(2)
							return 2 + sumRange(1)
											return 1

**/

func Factorial(num int) int {
	// base case allowing recursion to stop / avoiding stack overflow
	if num == 1 {
		return 1
	}
	// differing input passed back to the function
	return num * Factorial(num-1)
}

func FactorialIteratively(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func CollectOddNumbers(arr []int) []int {

	result := []int{}
	var Inner func([]int)

	Inner = func(helperInput []int) {
		if len(helperInput) == 0 {
			return
		}

		if helperInput[0]%2 != 0 {
			result = append(result, helperInput[0])
		}
		Inner(helperInput[1:])
	}
	Inner(arr)
	return result
}
