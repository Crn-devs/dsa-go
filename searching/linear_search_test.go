package searching

import (
	"fmt"
	"testing"
)

func TestIndexOfString(t *testing.T) {

	arr := []string{"1", "2", "3", "4"}
	val := "1"

	got := IndexOfString(arr, val)

	if got != 0 {
		t.Errorf(`IndexOf({"1", "2", "3", "4"}, "1") = %d; want 0`, got)
	}
}

func ExampleLinearSearch() {

	fmt.Println(LinearSearch([]int{1, 2, 3}, 3))
	// Output: 2

}
