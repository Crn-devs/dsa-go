package main

import (
	"fmt"

	"github.com/Crn-devs/dsa-go/array"
	"github.com/Crn-devs/dsa-go/bigo"
	"github.com/Crn-devs/dsa-go/hashmap"
)

func main() {

	// Big O
	var fish = []string{"nemo"}

	bigo.FindNemo(fish)

	bigo.GetFirstFish(fish)

	bigo.FunChallenge([]int{1, 2, 3, 4, 5})
	// Big O

	// Array

	arr := array.NewArray[int]()

	arr.Push(21)
	arr.Push(15)
	arr.Push(45)
	arr.Push(21)
	arr.Push(15)
	arr.Push(45)

	fmt.Println(arr.Get(0))
	fmt.Println(arr.Get(1))

	fmt.Println(arr)
	arr.Pop()
	arr.Pop()
	fmt.Println(arr)

	arr.Delete(2)
	arr.Delete(2)
	fmt.Println(arr)

	reversed := array.ReverseStringArray("hello")
	fmt.Println(reversed)

	merged := array.MergeSortedArray([]int{1, 9, 28}, []int{14, 17, 25, 33})
	fmt.Println(merged)

	arr1 := []int{1, 2, 3, 4, 5}
	arrRot := array.RotateArray(arr1, 2)

	fmt.Println(arrRot)

	fmt.Println(array.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println(array.TwoSum([]int{5, 8, 3, 5, 9, 8, 0}, 10))

	// Array

	// Hashmaps

	hm := hashmap.NewHashmap(50)

	hm.Set("Apples", "5")
	hm.Set("Oranges", "10")

	val := hm.Get("Oranges")
	fmt.Println(val) // 10

	keys := hm.Keys()

	hm.Set("Oranges", "15")
	val2 := hm.Get("Oranges")
	fmt.Println(val2) // 15

	fmt.Println(keys)

	// Hashmaps

}
