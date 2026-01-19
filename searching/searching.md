# Searching Algorithms in Go


## Built in searching algorithms in Go 

In Go there are 2 main built in searching algorithms and they are binary search and linear search, they are found heavily inside the sort/slices packages and allow quick searching of collections.

### Built in searching in the sort package
Here are a few examples of the methods that use binary search in the Go sort package.

- sort.Search - Binary Search
- sort.SearchInts (typed helper) - Binary Search
- sort.SearchFloat64s (typed helper) - Binary Search
- sort.SearchStrings (typed helper) - Binary Search

All of these require sorted data. linear search doesnt show up much due to the ease of writing it.

### Built in searching in the slices package
Here are some examples of methods that use binary / linear search in the slices package.

The newer Go slices package (1.21+) does include linear and binary search examples vs the sort package that only really has binary search used.

#### Linear Search
- slices.Contains
- slices.Index
- slices.IndexFunc
- slices.LastIndex
- slices.LastIndexFunc

### Binary Search
- slices.BinarySearch
- slices.BinarySearchFunc

Searching is generally prefered to be done with maps vs slices as we have a much faster average look up when utilising maps

Here are the Big O time complexitys for binary/linear search

### Binary Search
Best case: O(1)
Avg case: O(log n)
Worst case: O(log n)

### Linear Search Big O
best case: O(1)
Avg case: O(n)
Worst case: O(n)

# Algorithms Explained

## Linear Search

Linear Search is quite simple overall we take a collection and a target element and iterate through the entire collection checking on each iteration if the value at that iteration is the target, if it is we return its index and if its not in the collection we return -1. 

At worst this checks all elements in the collection causing its big o to be o(n) linear time complexity this algorithms performs well under constrains when the collection is completely unsorted. 

### Go Implementation
    func LinearSearch(arr []int, target int) int {

        if len(arr) <= 0 {
            return -1
        }

        for i := 0; i < len(arr); i++ {
            if arr[i] == target {
                return i
            }
        }
        return -1
    }

## Binary Search


Binary Search works on sorted collections by checking a given target against a mid point and reducing the collection size depending on the greater than less than conditions.

In basics it works like this, we take a mid point in the collection and check if its the target, if not we check if that mid point is greater than the target if this is true we now search from that mid point +1 to the end of the array by recomputing the mid point from the new section if its false we search from that mid point -1 back to the start of the array this in turn eliminates half of the collection on each iteration giving us O(log n) logarithmic time complexity


### Go Implementation
    func BinarySearch(arr []int, target int) int {

        // check empty
        if len(arr) <= 0 {
            return -1
        }

        low, high := 0, len(arr)-1

        for low <= high {
            // always calc mid high-low not low + high to stop integer overflow 
            mid := low + (high-low)/2

            // if element at mid == target return it
            if arr[mid] == target {
                return mid

                // if element at mid is greater than the target then drop the ceiling
            } else if arr[mid] > target {
                high = mid - 1
                // else if the element is less than the element at mid bring up the floor
            } else {
                low = mid + 1
            }
        }
        // not in the array
        return -1
    }