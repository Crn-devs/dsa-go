package searching

import "fmt"

func ExampleBinarySearch() {

	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	// Output: 9
}
