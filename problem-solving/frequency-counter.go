package main

import (
	"fmt"
)

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

func CheckAnagram(str1, str2 string) bool {

	// if lengths arent the same then automatically false
	if len(str1) != len(str2) {
		return false
	}

	// create maps to hold string values and counts
	frequency1 := make(map[string]int, len(str1))
	frequency2 := make(map[string]int, len(str2))

	// write the values to the map and increment the count for each letter
	for _, r := range str1 {
		frequency1[string(r)] = frequency1[string(r)] + 1
	}

	// same again for map 2
	for _, r := range str2 {
		frequency2[string(r)] = frequency2[string(r)] + 1
	}

	// loop over one of the frequency counters and check to make sure that the counts match
	// if the counts of each letter in the map do not match then return false
	for key, count := range frequency1 {
		if frequency2[key] != count {
			return false
		}
	}

	/**
	return true - all letters and counts match so return true
		as the strings contain the same letters and the same amount of letters
	**/

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

func main() {

	fmt.Println(Same([]int{1, 2, 3, 2}, []int{9, 1, 4, 4}))

	fmt.Println("CHECK ANAGRAM")
	fmt.Println(CheckAnagram("", ""))
	fmt.Println(CheckAnagram("timetwisttext", "texttwisttime"))
	fmt.Println(CheckAnagram("worldhelloday", "helloworld"))
	fmt.Println(CheckAnagram("cinema", "iceman"))

}
