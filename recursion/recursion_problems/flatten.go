package recursionproblems

/**
flatten

Write a recursive function called flatten which accepts an array of arrays
and returns a new array with all values flattened.

**/

func FlattenArr(arr []any) []int {

	result := []int{}
	// type switch needed for nested arrays in go no dynamic checking
	// v becomes the first element
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			result = append(result, val)
		case []any:
			// if the element is an array then recall with the inner array passed in if that array contains arrays it will all be flattened before moving on
			result = append(result, FlattenArr(val)...)
		default:
			panic("unsupported type")
		}
	}
	return result
}

// on each call append the contents of the array at position 0 to the new
// array if the element is an array then we need to re-call with that inner array
