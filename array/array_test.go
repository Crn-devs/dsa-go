package array

import (
	"testing"
)

func TestPush(t *testing.T) {
	arr := NewArray[int]()

	if arr.Len() != 0 {
		t.Errorf("expected: %v got: %v", 0, len(arr.data))
	}

	arr.Push(5)
	if arr.Len() != 1 {
		t.Errorf("expected: %v got: %v", 1, len(arr.data))
	}
}
