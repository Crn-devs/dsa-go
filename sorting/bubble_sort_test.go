package sorting

import "fmt"

func ExampleBubbleSort() {
	fmt.Println(BubbleSort([]int{43, 1, 5, 2, 8, 9}))
	// Output: [1 2 5 8 9 43]
}
