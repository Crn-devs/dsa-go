package recursionproblems

/*
*
isPalindrome

Write a recursive function called isPalindrome which returns true
if the string passed to it is a palindrome (reads the same forward and backward).
Otherwise it returns false.

*
*/

func IsPalindrome(s string) bool {

	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return IsPalindrome(s[1 : len(s)-1])
}

// is palindrome wants to check if the last
// element and first element are the same and then remove them both and recheck
