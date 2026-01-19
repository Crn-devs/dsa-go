package sorting

import "fmt"

func ExampleSelectionSort() {
	fmt.Println(SelectionSort([]int{4, 3, 2, 1}))
	//Output: [1 2 3 4]
}
