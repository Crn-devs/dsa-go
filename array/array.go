package array

type Array[T any] struct {
	data []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{data: []T{}}
}

func (arr Array[T]) Len() int {
	return len(arr.data)
}

func (arr *Array[T]) Push(item T) int {
	arr.data = append(arr.data, item)
	return len(arr.data)
}

func (arr *Array[T]) Get(index int) T {
	return arr.data[index]
}

func (arr *Array[T]) Pop() int {

	arr.data = arr.data[0 : len(arr.data)-1]
	return len(arr.data)
}

func (arr *Array[T]) Delete(index int) {
	arr.data = append(arr.data[:index], arr.data[index+1:]...)
}
