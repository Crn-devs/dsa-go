package recursionproblems

/**
someRecursive

Write a recursive function called someRecursive which accepts an array and a callback.
The function returns true if a single value in the array
returns true when passed to the callback. Otherwise it returns false.

**/

func SomeRecursive(arr []int, callBack func(i int) bool) bool {

	if len(arr) <= 0 {
		return false
	}

	if callBack(arr[0]) {
		return true
	} else {
		return SomeRecursive(arr[1:], callBack)
	}
}

func IsOdd(i int) bool {
	return i%2 != 0
}

// if a value passed to the call back returns true we return true otherwise we return false
// we need to call some with differing input which is a shorter and shorter array inside we need to call callback and return its result
