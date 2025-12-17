

## Recursion 

a function that calls itself

must have a base case to avoid stack over flow
must have differing input passed on each subsequent call 

examples in main include
- factorial
- the sum range of all numbers up to a target
- countdown function


## Helper method recursion
A Design pattern for recursion allowing access to an object we need to manipulate without resetting it

usually implemented with a inner function or a helper method
in go if we use an inner function it must be defined before calling it as go doesnt have named inner functions

-
	var Inner func([]int)

	Inner = func(helperInput []int) {//body}
-
