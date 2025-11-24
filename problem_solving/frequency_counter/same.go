package main

import "fmt"

// Write a function that accecpts 2 arrays as input, the function should return true if every value
// in the array has its value squared in the corrisponding array, the frequency should be the same

func Same(arr1, arr2 []int) bool {

	// check empty input

	if len(arr1) <= 0 || len(arr2) <= 0 {
		fmt.Println("please provide valid arrays with atleast one item")
		return false
	}

	// if lengths arent the same then automatically false
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		val := arr1[i] * arr1[i]
		correctIndex := indexOf(val, arr2)

		fmt.Println("CORRECT INDEX:", correctIndex)

		if correctIndex == -1 {
			fmt.Println("val not found", arr1, arr2, val)
			return false
		}
		arr2 = append(arr2[:correctIndex], arr2[correctIndex+1:]...)
		fmt.Println(arr2)
		fmt.Println(arr2)
	}

	return true
}

func indexOf(target int, arr []int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1 // not found, just like JS indexOf
}

func SameFC(arr1, arr2 []int) bool {

	// check empty input

	if len(arr1) <= 0 || len(arr2) <= 0 {
		fmt.Println("please provide valid arrays with atleast one item")
		return false
	}

	// if lengths arent the same then automatically false
	if len(arr1) != len(arr2) {
		return false
	}

	frequencyCounter1 := make(map[int]int, len(arr1))
	frequencyCounter2 := make(map[int]int, len(arr2))

	for i := 0; i < len(arr1); i++ {
		frequencyCounter1[i] = frequencyCounter1[i] + 1
	}

	for i := 0; i < len(arr2); i++ {
		frequencyCounter2[i] = frequencyCounter2[i] + 1
	}

	for key, count := range frequencyCounter1 {
		square := key * key
		if frequencyCounter2[square] != count {
			return false
		}
	}

	return true

}
