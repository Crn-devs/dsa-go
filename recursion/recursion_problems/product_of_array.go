package recursionproblems

/**
Write a function called productOfArray which takes in an array
of numbers and returns the product of them all.
**/

func ProductOfArray(arr []int) int {
	if len(arr) <= 0 {
		return 1
	}
	return arr[0] * ProductOfArray(arr[1:])
}

/**
example:

arr[3,2,10]

return 3 * ProdOfArr([2,10])
				return 2 * 10

result: 3 * 2 * 10

**/
