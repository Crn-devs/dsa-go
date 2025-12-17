package recursionproblems

import "strings"

/**
Write a recursive function called capitalizeFirst.
Given an array of strings, capitalize the first letter of each string in the array.

**/

func CapitaliseFirst(arr []string) []string {

	if len(arr) == 0 {
		return []string{}
	}

	first := strings.ToUpper(arr[0][:1]) + arr[0][1:]
	rest := CapitaliseFirst(arr[1:])

	return append([]string{first}, rest...)
}
