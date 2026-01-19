## Sorting Algorithms

Sorting algorithms are everywhere in programming, alot of languages include various built in sorting algorithms directly in the language or on certain types this is why its crucial to understand how they work under the hood

Some sorting algorithms have advanteges and disadvanteges when used and are comprised on different techiques they also have some odd outcomes when the data passed is partly sorted vs completely unsorted
A website which shows the basics can be found [Here](https://www.toptal.com/developers/sorting-algorithms)


# Objectives
Implement the following sorting algorithms in Go either via type or generically doing so inline with the standard library implementations, 

- Bubble Sort
- Selection Sort
- Insertion Sort

# Built in Sorting in Go (production sorting patterns)
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

A Sorting algorithm where larger collection elements bubble to the top

bubble sort in basics steps through a collection of elements and checks if the first element is larger than the second
if so the elements are swapped and the the algorithm continues passing until the array is sorted

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

After each outer iteration, the largest element is in its final position at the end, so the inner loop can safely ignore those elements

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

elements are essentially selected when they are the current minimum and moved to the start

we keep track of a current minimum as an index and use that as a gauge for our comparison if the newer element is less than the current minimum, minimum is updated to reflect that

here is an example

inital array [4,6,2,0]

4 becomes the minimum and is checked to see if its the minimum value of a pair 
4 < 6 

4 is kept as its still the minimum

4 < 2
minimum is set to 2 since its smaller than 4

2<0
minimum is set to the index of the element 0 since its smaller than 2
since we have reached the end of the array and there are no more numbers to compare the element at minimum is swapped to be in the first place of the array
[0,6,2,4]
outer loop begins again at the next index: min = 6 skipping the last sorted element, the inner loop also needs setting to be i+1 so that it also doesnt check against the sorted elements
6<2 would be the next comparison

Selection Sort big O
best case: $$O(n)$$ // list is already sorted
Avg case: $$O(n^2)$$
Worst case: $$O(n^2)$$

Real World Usage Rating: Not often used, there are faster and more optimal sorting algos but usually if we want to prioritise swapping less (less memory usage) then we can use selection sort over Bubble sort


full algorithm

	func SelectionSort[S ~[]E, E cmp.Ordered](s S) S {

		var min int

		for i := 0; i < len(s)-1; i++ {
			// set min to be the first elements index
			min = i
			for j := i + 1; j < len(s); j++ {
				// check the first element the element at index(min) againt the rest of the slice
				// if min > next element swap min
				if s[j] < s[min] {
					min = j
				}
			}

			s[i], s[min] = s[min], s[i]
		}
		return s
	}


# Insertion Sort

Insertion sort works by inserting values into there correct position in the collection by checking them against a known value it works due to each value being re-checked and again moved upwards this creates a sort of build up of a fully sorted collection

to implement insertion sort we need to keep track of the element currently infront of our first loops element, this way we have a reference allowing us to overwrite the actual value at the position with its adjacent element


[4,1,5,7,18,11,100]

for a collection like this we would start our first loop from index 1

on our nested loop we check to see if the element before our current element is smaller
if the element is smaller then we move the element at i upwards, we do this as we have a reference to the element before it which we can use to write the value

to correctly write the value in place though we now need to check those elements that have been sorted to find the correct placement this is where our loop clause comes in we loop while j >= 0 and if the element at position j > than the current we continue looping as it means we have not yet found its placement when all conditions are exhausted we break for the loop writing the element at its correct position

the flow looks like this 

[1,9,6,19]

i index = 1
i element = 9

j index = 0
j element = 1

while ever j is greater than or equal to -1 and the element at j is greater than the element at i
move the element at j upwards one place

in our current iteration j is not greater than the element at i so we do nothing

j gets decremented after this operation which leaves us with this 

i index = 1
i element = 9

j index = -1
j element = nil not set as loop isnt run

Second iteration

[1,9,6,19]

i index = 2
i element = 6

j index = 1
j element = 9

while ever j is greater than or equal to -1 and the element at j is greater than the element at i
move the element at j upwards one place

in our current iteration j is greater than the element at i so we move j to be at position j+1

[1,9,9,19] here we lose our array element in place but we have a reference to it in the variable

j gets decremented after this operation which leaves us with this 

i index = 2
i element = 6

j index = 0
j element = 1

again our conditional is satisfied so we again iterate on j

i index = 2
i element = 6

j index = -1
j element = nil not set loop isnt run

if we now look at our array we have [1,6,9,19] which leaves us with a sorted array 6 is written to the index of j+1 when the loop breaks causing the element to be correctly placed in the array

Selection Sort big O
best case: $$O(n)$$ // list is already sorted
Avg case: $$O(n^2)$$
Worst case: $$O(n^2)$$


Full Algorithm

	func InsertionSort[S ~[]E, E cmp.Ordered](s S) S {

		for i := 1; i < len(s); i++ {

			curr := s[i]
			var j int

			for j = i - 1; j >= 0 && s[j] > curr; j-- {
				s[j+1] = s[j]
			}
			s[j+1] = curr
		}
		return s
	}



# Sorting Algorithm comparisons
## Quadratic Sorting Functions

| Algorithm    | Time Complexity (avg) | Time Complexity (worst) | Time Complexity (best) |
| ------------ | --------------------- | ----------------------- | ---------------------- |
| BubbleSort   | $$O(n^2)$$            | $$O(n^2)$$				 | $$O(n)$$				  |
| SelectionSort| $$O(n^2)$$            | $$O(n^2)$$				 | $$O(n^2)$$			  |
| InsertionSort| $$O(n^2)$$            | $$O(n^2)$$				 | $$O(n)$$				  |




# Faster Sorting Algorithms


## Merge Sort

A Divide and conquer style algorithm

divides up work 
conquers it 
merges the output and returns it 

works off the model that arrays of 0 or 1 items are always sorted

works by chopping a single array of multiple elements into arrays of 1 element and then building up a new array

an 8 element array split would look like this 
```
[1, 2, 3, 4, 5, 6, 7, 8]
            |		 
           / \
[1, 2, 3, 4]  [5, 6, 7, 8]
      |            |
     / \          / \
[1, 2] [3,4] [5,6] [7,8]
   |     |     |     |
  / \   / \   / \   / \
[1][2] [3][4] [5][6] [7][8]
```

once the elements have been split down to single digits then techically they are each a sorted array of a single element

from here we just merge the sorted arrays by comparing the elements from the array we are merging together and placing them at the correct position in the new array respective to which element is larger

arr1[0] < arr2[0] {
	[]{arr1[0], arr2[0]}
}else {
	[]{arr2[0], arr1[0]}
}

```
[4][1] [16][11] [6][6] [1][25]
 \ /	\ /    \/     \/
[1,4]   [11,16]  [6,6]  [1,25]
    \   /            \   /  
     [1,4,11,16]     [1,6,6,25]
               \     /
            [1,4,6,6,11,16,25]
```

the algorithm uses recursion to process the arrays first splitting the array passed into 2 and passing those back to the same function, this happens until there is a single element - the base case and when the base case is satisfied the stack rolls back up merging each of the processed arrays until it returns a sorted single array from the origonal caller.

Merge Sort big O (Time)
best case: $$O(n log n)$$ // list is already sorted
Avg case: $$O(n log n)$$
Worst case: $$O(n log n)$$

## Quick Sort

recursive algorithm

similar to merge it exploits the fact that arrays of 0 or 1 element are always sorted

works off of a pivot principle where an element is chosen as a pivot and all elements greater are put to its right and all elements less than are put to its left

lets look at an example flow of quick sort

lets say we have the following array

[1,5,7,32,77,18,12,22,5]

quick sort has a few methods of implementation for this we will use the lomuto partition algorithm to implement the partitioning

for quick sort to work we need the following 

a pivot/partition
a base case to stop the recursion
a starting index in which to compare the pivot to

these can be defined as the following

high = pivot or the last element of the array
low = 0 beginning index
base case is while low < high each call is with a differing input array

[1,5,7,32,4,5]

first iteration
j = 0
smaller window = -1
pivot = last index value 5

1<5 true so window is increased and element is swapped to be in place of smaller window index
smaller window = 0
element 1 is swapped to position 0

j=1
5<5 false so nothing happens
smaller window still 0 

j=2
7<5 false so nothing happens
smaller window still 0 

j=3
32<5 false so nothing happens
smaller window still 0 

j=4
4<5 true so window is increased and the elements are swapped
smaller window increases to 1
element 4 is swapped to be in place of smaller window index
array goes from  [1,5,7,32,4,5] to [1,4,7,32,5,5]

j=5 so the loop finishes as its hit its condition j < high

once the loop finishes we have the following 

smaller window counter value at 1
our pivot value which was 5
our less than elements contained inside the smaller window counter up to its value as an index
we now take our pivot value and swap it to be at the end of the smaller than window

smaller than window = 1 pivot value goes at index 2 leaving us with the following array

[1,4,5,7,32,5] all elements to the left are less than / all to the right are greater than

the element is in its correct position
now the recursion begins 

we re-call quick sort passing in the array segments we have created by using pivot as our differing value

quickSort(arr, low, pivot-1) // this is the left side of the array (Less than)
[1,4,5]

quickSort(arr, pivot+1, high)// this is the right side of the array (greater)
[7,32,5]

these 2 partitions of the array are again processed with the above steps here is a brief example

[7,32,5]

5 is the pivot
j = 0
smaller window = -1

7<5 false nothing happens

j=1
32<5 false nothing happens

j=2
5<5 false nothing happens

loop finishes
pivot is moved to smaller window index + 1 so index 0
array goes from [7,32,5] to [5,7,32] leaving us with a sorted array half

Quick Sort big O (Time)
best case: $$O(n log n)$$
Avg case: $$O(n log n)$$
Worst case: $$O(n log n)$$


