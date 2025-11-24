package main

import "fmt"

func main() {

	// fmt.Println(SumZeroQuadratic([]int{-3, -2, -1, 0, 1, 2, 3}))
	// fmt.Println(SumZero([]int{-3, -2, -1, 0, 1, 2, 3}))

	// fmt.Println(SumZero([]int{-1, -3, -5, 10, 5, 11, 22, 3}))

	// fmt.Println(SumZero([]int{-1, -2, 1, -1}))

	// fmt.Println(CountUniqueVals([]int{1, 2, 3, 3, 5, 6, 14, 16, 17, 19, 22, 22, 25}))
	// fmt.Println(CountUniqueVals([]int{1, 2, 3, 3}))

	// fmt.Println(TargetPairs([]int{1, 2, 3}, 5))
	// fmt.Println(TargetPairs([]int{1, 2, 3, 4, 5, 6, 7}, 8))

	// fmt.Println(ClosestPair([]int{-5, -2, -1, 3, 6, 10}, 8))
	// fmt.Println(ClosestPair([]int{-5, 10}, 5))
	fmt.Println(maxSubarraySum([]int{2, 6, 9, 2, 1, 8, 5, 6, 3}, 3))

	fmt.Println(AreThereDuplicatesPointer(1, 2, 3, 4, 5, 1))
	fmt.Println(AreThereDuplicatesPointer())
	fmt.Println(AreThereDuplicatesPointer(1, 2, 3, 4, 5, 0))

}
