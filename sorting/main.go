package main

import (
	"cmp"
	"fmt"
	"slices"

	bubblesort "github.com/Crn-devs/dsa-go/sorting/bubble_sort"
	selectionsort "github.com/Crn-devs/dsa-go/sorting/selection_sort"
)

func main() {

	words := []string{"hi", "Bye", "hello", "Good Morning", "Good Day", "Ey-Up", "Whatsup", "👋"}

	fmt.Println(words)

	byLength := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(words, byLength)

	fmt.Println(words)

	// Bubble Sort
	nums := []int{1, 4, 2, 7, 78, 2}
	fmt.Println(nums)
	bubblesort.BubbleSort(nums)
	fmt.Println(nums)

	// Selection Sort
	SelecNums := []int{4, 3, 2, 1}
	fmt.Println(SelecNums)
	selectionsort.SelectionSort(SelecNums)
	fmt.Println(SelecNums)
}
