## Sorting Algorithms

Sorting algorithms are everywhere in programming, alot of languages include various built in sorting algorithms directly in the language or on certain types this is why its crucial to understand how they work under the hood

Some sorting algorithms have advanteges and disadvanteges when used and are comprised on different techiques they also have some odd outcomes when the data passed is partly sorted vs completely unsorted
A website which shows the basics can be found [Here](https://www.toptal.com/developers/sorting-algorithms)


# Objectives
Implement the following sorting algorithms in Go either via type or generically inline with the standard library implementations

- Bubble Sort
- Selection Sort
- Insertion Sort

# Built in Sorting in Go
In go the sorting algorithms live inside the slices package. 

Sorting functions are generic and work for any ordered built in type orderable types are inside the cmp package on go.dev [Ordered_types](https://pkg.go.dev/cmp#Ordered)

Here is the interface that the sorting functions depend on these are defined by there ability to support the operators < <= >= >

    type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

really they support the < operator but thats another story.

if we want custom sorting vs a types natural order we can use the slices.SortFunc method and specify our own sorting algorithm this can be done either by using cmp.compare method or by writing a method that satisfys the following

returns -1 if x < y
returns 0 if x == y
returns +1 if x > y

here is an example using cmp.Compare with the lengths of words instead of alphabetically

	func main() {

		words := []string{"hi", "Bye", "hello", "Good Morning", "Good Day", "Ey-Up", "Whatsup", "👋"}

		fmt.Println(words)

		// cmp.Compare returns 1 if a>b 0 if a==b and -1 if a<b
		byLength := func(a, b string) int {
			return cmp.Compare(len(a), len(b))
		}
		// we can then use that function as the argument to SortFunc
		slices.SortFunc(words, byLength)

		fmt.Println(words)

	}


# Bubble Sort

A Sorting algorithm where larger numbers bubble to the top

bubble sort in basics steps through a list of elements and checks if the first element is larger than the second
if so the elements are swapped and the the algorithm continues passing one loop for each element in the array

so for the array [4,3,2,1]

step 1
comparing 4-3
3/4 are swapped
[3,4,2,1]

step 2
comparing 4-2
4/2 are swapped
[3,2,4,1]

step 3
comparing 4-1
4/1 are swapped
[3,2,1,4]

step 4
nothing else to compare - loop again; it would take a total of 4 outer loops to sort the array each outer loop has 4 elements to compare however we dont need to compare the last item each time so we can reduce the inner loops comparison by 1 each time


inital array [4,3,2,1]

after first pass [3,2,1,4]

only 3 elements to compare now

after second pass [2,1,3,4]

only 2 elements to compare now

after third pass [1,2,3,4] array is sorted 

We dont need to pass through the full array each time, we can begin at the last element in the outer loop
and only loop up to the outerloop-1 (i - 1) on the inner loop giving us only the left over comparisons 

i=4, j goes up to 4-1
i=3, j goes up to 3-1
i=2, j goes up to 2-1
i=1, j goes up to 1-1
this makes sure we dont compare elements that have already been sorted


Bubble Sort big O
best case: $$O(n)$$ // list is already sorted
Avg case: $$O(n^2)$$
Worst case: $$O(n^2)$$

Real World Usage Rating: Not often used, there are faster and more optimal sorting algos

Full Algorithm

	func BubbleSort[S ~[]E, E cmp.Ordered](s S) S {

		// optimise
		// using a bool no swap
		var noSwap bool
		// optimise

		if len(s) < 2 {
			return s
		}

		for i := len(s); i > 1; i-- {
			noSwap = true
			for j := 0; j < i-1; j++ {
				if s[j] > s[j+1] {
					s[j], s[j+1] = s[j+1], s[j]
					noSwap = false
				}
			}
			if noSwap {
				return s
			}
		}
		return s
	}


# Selection Sort

Selection sort is similar to bubble sort but instead of elements being moved to the end with adjacent comparisons elements are moved towards the beginning based on a minimum condition


inital array [4,6,2,0]

4 becomes the minimum and is checked to see if its the minimum value of a pair 
4 < 6 

4 is kept as its still the minimum

4 < 2
minimum is set to 2 since its smaller than 4

2<0
minimum is set to 0 since its smaller than 2
begginning element is swapped with zero
[0,6,2,4]

outer loop begins again at the next index: min = 6








