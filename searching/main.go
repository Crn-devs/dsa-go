package main

import (
	"fmt"

	binarysearch "github.com/Crn-devs/dsa-go/searching/binary_search"
)

func main() {

	cmp := func(a, b int) int {
		return a - b
	}
	pos := binarysearch.BinarySearchFunc([]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, cmp)
	fmt.Println(pos)

	fmt.Println(stringSearch("ofgmomghimomomg", "mom"))

}

// quick brute force string search
func stringSearch(s1, s2 string) int {

	count := 0

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if string(s1[i+j]) != string(s2[j]) {
				break
			}
			if j == len(s2)-1 {
				count++
			}
		}
	}

	return count
}
