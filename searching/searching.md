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

binary search works of the basis that we divide the searching area of the array in half on each iteration we can do this as we know the array is sorted usual implementation takes a low at the first element of the array and a high at the last element and calculates a mid point we then check to see if the value at mid is equal to, greater than or less than our target element and depending on which condition we either reduce the high or increase the low to be that of the mid point + 1 allowing the search window to be cut in half on each iteration

[1,2,3,4,5,6], target = 5
starting iteration values
low = idx: 0 ele: 1
high = idx: 5 ele:6
mid idx: 2 ele: 3

[1,2,3,4,5,6], target = 5
second iteration values
low = idx: 3 ele: 4
high = idx: 4 ele:5
mid idx: 4 ele: 5

target at mid == target specified so value is found

Binary search big O
best case: O(1)
Avg case: O(log n)
Worst case: O(log n)
