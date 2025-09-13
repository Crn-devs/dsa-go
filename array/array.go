package array

type Array[T any] struct {
	data []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{data: []T{}}
}

func (arr *Array[T]) Get(index int) T {
	return arr.data[index]
}
