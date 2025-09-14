package array

import (
	"fmt"
	"strings"
)

func ReverseStringArray(s string) string {

	// first check input is correct
	if len(s) <= 0 {
		fmt.Println("incorrect length: length should be greater than 0")
	}

	// create an array to store reversed string
	reversed := NewArray[string]()

	// run reverse loop
	for i := len(s) - 1; i >= 0; i-- {
		reversed.Push(string(s[i]))
	}

	return strings.Join(reversed.data, "")
}
