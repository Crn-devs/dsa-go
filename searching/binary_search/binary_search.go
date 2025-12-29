package binarysearch

// a generic version of binary search: takes in a comparrsion function cmp to compare the elements
// this mirrors the slices.binarysearchfunc method in the slices package
func BinarySearchFunc[T any](arr []T, x T, cmp func(a, b T) int) int {

	if len(arr) <= 0 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		c := cmp(arr[mid], x)

		if c == 0 {
			// element found return it
			return mid
		}
		if c < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
