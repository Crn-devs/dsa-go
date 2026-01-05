package main

import (
	"cmp"
	"fmt"
	"slices"

	bubblesort "github.com/Crn-devs/dsa-go/sorting/bubble_sort"
	insertionsort "github.com/Crn-devs/dsa-go/sorting/insertion_sort"
	mergesort "github.com/Crn-devs/dsa-go/sorting/merge_sort"
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
	SelecNums = selectionsort.SelectionSort(SelecNums)
	fmt.Println(SelecNums)

	// Insertion sort

	InsertNums := []int{3, 7, 4, 2, 1, 5, 7, 9, 5, 7, 8, 9, 2}
	fmt.Println(InsertNums)
	InsertNums = insertionsort.InsertionSort(InsertNums)
	fmt.Println(InsertNums)

	a := []int{10010, 20, 30, 40, 50, 60, 60, 60, 65, 70, 71}

	fmt.Println(mergesort.MergeSort(a))

}
