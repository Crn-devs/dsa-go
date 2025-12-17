package recursionproblems

/**
power

Write a function called power which accepts a base and an exponent.
The function should return the power of the base to the exponent.
This function should mimic the functionality of Math.pow()
- do not worry about negative bases and exponents.

example:
	base = 2
	exponent = 4

	2 x 2 x 2 x 2
**/

func Power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * Power(base, exponent-1)
}

/**

base = 2
exponent = 2

first pass ->
	2 * Power(2, 2-1)
Second pass ->
	2 * Power(2, 1-1)
Third Pass ->
	returns 1

	results in ->>
	2 * 2 * 1 = 4

**/
