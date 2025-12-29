## Searching Algorithms


# Linear Search

Linear search is a method of finding an element within a list, it checks each element of the list one by one until either a match is found or it reaches the end. 

linear search is used to power various algorithms such as indexOf and Contains both of which are in the linear_search.go file 

Linear search big o
best case: O(1)
Avg case: O(n)
Worst case: O(n)

# Binary Search

Binary search is a method of finding an element in a list similar to linear search but rather than eleminating 1 element at a time it can eliminate half the remaining elements on each step

a rule of binary search is the input MUST be sorted

in go the slices package has a binary search implementation which works on all types by using generics

Linear search big O
best case: O(1)
Avg case: O(log n)
Worst case: O(log n)
