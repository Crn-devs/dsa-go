package recursionproblems

import "strconv"

/**
Write a function called stringifyNumbers which takes in an
object and finds all of the values which are numbers
and converts them to strings. Recursion would be a great way to solve this!

The exercise intends for you to create a
new object with the numbers converted to strings,
and not modify the original. Keep the original object unchanged.

**/

func StringifyNumbers(obj map[string]any) map[string]any {

	nm := make(map[string]any, len(obj))

	for key, val := range obj {
		switch v := val.(type) {
		case int:
			nm[key] = strconv.Itoa(v)
		case map[string]any:
			nm[key] = StringifyNumbers(v)
		default:
			nm[key] = v
		}
	}
	return nm
}
