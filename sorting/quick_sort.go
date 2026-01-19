package sorting

func partition(arr []int, low, high int) int {

	pivot := arr[high] // last element
	swapIdx := low - 1

	// loop the arr and check

	for j := low; j < high; j++ {
		// is the element we are on less than the pivot element
		if arr[j] < pivot {
			// if its less than update the swapidx by 1 increasing the less than space anything <= i is less than the pivot
			swapIdx++
			// now we have a window of less than elements but if we hit a greater than element its skipped
			// this causes less than elements further down the array to also now need a solution and that
			// is swapping the less than values to the place of the current smaller than window index at the end of the smaller
			// than window
			arr[swapIdx], arr[j] = arr[j], arr[swapIdx]
		}
	}
	// swap the pivot (last element)
	arr[swapIdx+1], arr[high] = arr[high], arr[swapIdx+1]
	return swapIdx + 1
}

func QuickSort(arr []int, low, high int) {
	// run this block once
	if low < high {
		// pivot is computed and used as the next argument of high/low
		pivot := partition(arr, low, high)

		// pivot being used as differing input left / right of the array
		// left
		QuickSort(arr, low, pivot-1)
		// right
		QuickSort(arr, pivot+1, high)
	}
}
