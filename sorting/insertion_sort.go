package sorting

import (
	"cmp"
)

func InsertionSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	// insert a value at its correct position in a collection by tracking a current value
	// think deck of cards being sorted in two piles one unsorted pile driving the sorted pile
	/**
	we take a card and place it down
	we then take the next card as our card to compare
	we check is the card down greater than the card in our hand
	if it is greater we essentially swap them but its not a swap its a shift so we take the card
	thats down into our hand instead place the card from our hand down and then finally place
	the card that was down back on the pile giving us the lowest on the bottom then on each
	loop we check against those cards placed down with the new card we have drawn
	**/

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		var j int

		for j = i - 1; j >= 0 && arr[j] > curr; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = curr
	}
	return arr
}

func InsertionSortGeneric[S ~[]E, E cmp.Ordered](s S) S {

	for i := 1; i < len(s); i++ {

		curr := s[i]
		var j int

		for j = i - 1; j >= 0 && s[j] > curr; j-- {
			s[j+1] = s[j]
		}
		s[j+1] = curr

	}
	return s
}
