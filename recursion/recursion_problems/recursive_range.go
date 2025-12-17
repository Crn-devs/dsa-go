package recursionproblems

/**

Write a function called recursiveRange which accepts a number and adds up
all the numbers from 0 to the number passed to the function

**/

func RecursiveRange(num int) int {

	if num == 0 {
		return 0
	}
	return num + RecursiveRange(num-1)
}

/**

num = 5

result = 5 + 4 + 3 + 2 + 1 + 0

**/
