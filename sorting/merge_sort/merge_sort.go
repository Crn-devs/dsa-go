package mergesort

func merge(a, b []int) []int {

	// create pointers (start of each array)
	p1, p2 := 0, 0
	// create empty array
	arr := []int{}

	// loop over both arrays
	for p1 < len(a) && p2 < len(b) {

		// find the smaller element of the two -> place it into the created array
		if a[p1] <= b[p2] {
			arr = append(arr, a[p1])
			p1++
		} else {
			arr = append(arr, b[p2])
			p2++
		}
	}
	arr = append(arr, b[p2:]...)
	arr = append(arr, a[p1:]...)
	return arr
}

func MergeSort(arr []int) []int {

	// base recursion case if the arr has one element its sorted
	if len(arr) <= 1 {
		return arr
	}

	// grab the mid point
	mid := len(arr) / 2

	// recursive call to mergeSort
	left := MergeSort(arr[mid:])
	right := MergeSort(arr[:mid])

	return merge(left, right)
}
