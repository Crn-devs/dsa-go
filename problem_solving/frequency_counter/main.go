package main

import (
	"fmt"
)

func main() {

	fmt.Println("Same")
	fmt.Println(Same([]int{1, 2, 3, 2}, []int{9, 1, 4, 4}))

	fmt.Println("Same Improved")
	fmt.Println(SameFC([]int{1, 2, 3, 2}, []int{9, 1, 4, 4}))
	fmt.Println(SameFC([]int{4, 1, 8, 9}, []int{9, 2, 5, 7}))

	fmt.Println("CHECK ANAGRAM")
	fmt.Println(CheckAnagram("", ""))
	fmt.Println(CheckAnagram("timetwisttext", "texttwisttime"))
	fmt.Println(CheckAnagram("worldhelloday", "helloworld"))
	fmt.Println(CheckAnagram("cinema", "iceman"))

	fmt.Println("same freq = ", SameFrequency(122, 123))

	fmt.Println(AreThereDuplicates(1, 2, 3, 4, 5, 1))
	fmt.Println(AreThereDuplicates())
	fmt.Println(AreThereDuplicates(1, 2, 3, 4, 5, 0))

}
