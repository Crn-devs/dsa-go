package main

/*
Write a function called isSubsequence which takes in two strings and checks whether
the characters in the first string form a subsequence of the characters in the second string.
In other words, the function should check whether the characters in the first string appear somewhere
in the second string, without their order changing.
*/
func IsSubSequence(s1, s2 string) bool {

	// check empty
	if len(s1) <= 0 || len(s2) <= 0 {
		return false
	}

	// create pointers
	p1 := 0
	// start at the same index
	// if they match increase both
	// if they dont match only increase p2 to check the next element
	// if p1 reaches the end of the string then no match was found

	for p2 := 0; p2 < len(s2); p2++ {

		if s1[p1] == s2[p2] {
			p1++
			if p1 == len(s1) {
				return true
			}
		}
	}
	return false
}
