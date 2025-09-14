package array

import (
	"testing"
)

func TestPush(t *testing.T) {
	arr := NewArray[int]()

	if arr.Len() != 0 {
		t.Errorf("expected: %v got: %v", 0, arr.Len())
	}

	arr.Push(5)
	if arr.Len() != 1 {
		t.Errorf("expected: %v got: %v", 1, arr.Len())
	}
}

func TestLen(t *testing.T) {
	arr := NewArray[int]()

	if arr.Len() != 0 {
		t.Errorf("expected: %v got: %v", 0, arr.Len())
	}

	arr.Push(5)
	arr.Push(5)

	if arr.Len() != 2 {
		t.Errorf("expected: %v got: %v", 2, arr.Len())
	}

	arr.Pop()

	if arr.Len() != 1 {
		t.Errorf("expected: %v got: %v", 1, arr.Len())
	}
}

func TestPop(t *testing.T) {
	arr := NewArray[int]()

	if arr.Len() != 0 {
		t.Errorf("expected: %v got: %v", 0, arr.Len())
	}

	arr.Push(5)
	arr.Push(5)

	if arr.Len() != 2 {
		t.Errorf("expected: %v got: %v", 2, arr.Len())
	}
	arr.Pop()
	arr.Pop()

	if arr.Len() != 0 {
		t.Errorf("expected: %v got: %v", 0, arr.Len())
	}
}

func TestDelete(t *testing.T) {
	arr := NewArray[int]()

	if arr.Len() != 0 {
		t.Errorf("expected: %v got: %v", 0, arr.Len())
	}

	arr.Push(5)
	arr.Push(10)
	arr.Push(15)

	if arr.Len() != 3 {
		t.Errorf("expected: %v got: %v", 3, arr.Len())
	}
	arr.Delete(1)

	if arr.Get(1) == 10 {
		t.Errorf("expected: %v got: %v", 15, arr.Get(1))
	}
}

func TestGet(t *testing.T) {
	arr := NewArray[int]()

	arr.Push(5)
	arr.Push(15)
	arr.Push(25)

	got := arr.Get(1)

	if got != 15 {
		t.Errorf("expected: %v got: %v", 15, got)
	}
}
