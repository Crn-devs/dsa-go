package recursionproblems

/**

reverse

Write a recursive function called reverse which accepts a string and returns a new string in reverse.

**/

func Reverse(s string) string {

	rev := []byte{}
	var inner func(s string)

	if len(s) == 0 {
		return string(rev)
	}

	inner = func(s string) {
		if len(s) == 0 {
			return
		}
		rev = append(rev, s[len(s)-1])
		inner(s[:len(s)-1])
	}
	inner(s)
	return string(rev)
}

// pass each last element into the func then cut the end off
