package array

import (
	"fmt"
	"strings"
)

func ReverseStringArray(s string) string {

	// first check input is correct
	if len(s) <= 0 {
		fmt.Println("incorrect length: length should be greater than 0")
	}

	// create an array to store reversed string
	reversed := NewArray[string]()

	// run reverse loop
	for i := len(s) - 1; i >= 0; i-- {
		reversed.Push(string(s[i]))
	}

	// implement toString method instead
	return strings.Join(reversed.data, "")
}

/**
Problem:
Merge 2 Given arrays of sorted integers
Example:
	Input: nums = [1, 9, 28],[14, 17, 25, 33]
	Output: [1,9,14,17,25,28,33]
**/

func MergeSortedArray(arr1, arr2 []int) []int {

	// check for empty array
	if len(arr1) <= 0 || len(arr2) <= 0 {
		fmt.Println("incorrect length: length should be greater than 0")
	}

	arr := []int{}
	i, j := 0, 0
	for len(arr1) > i && len(arr2) > j {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}

	arr = append(arr, arr1[i:]...)
	arr = append(arr, arr2[j:]...)
	return arr

}

/**
Problem:
Given an array of integers, rotate the array to the right by k steps, where k is non-negative.
Example:
	Input: nums = [1,2,3,4,5,6,7], k = 3
	Output: [5,6,7,1,2,3,4]
**/

func RotateArray(arr1 []int, k int) []int {
	// check for empty
	if len(arr1) <= 0 {
		fmt.Println("incorrect length: length must be greater than 0")
		return nil
	}
	// create new array add elements rotated by k places
	arr := make([]int, len(arr1))

	fmt.Println("length", len(arr1))

	for i := range arr1 {
		pos := (i + k) % len(arr1) // modulus used for 'wrap around'
		arr[pos] = arr1[i]
	}
	return arr
}

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49

area = width * height
width = difference in index
*/
func MaxAreaBrute(height []int) int {

	// check for empty
	if len(height) <= 0 {
		fmt.Println("incorrect length: length must be greater than 0")
		return 0
	}

	var area int
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			// compare the elements,
			s1 := height[i]
			s2 := height[j]
			// select the smallest
			h := min(s1, s2)
			// find the area width * height
			width := j - i
			if area > width*h {
				continue
			}
			area = width * h
			fmt.Printf("for pairs %v %v ", s1, s2)
			fmt.Printf("the area should be %v as width: %v * height: %v\n", area, width, h)
		}
	}
	return area

}

func MaxArea(height []int) int {

	// check for empty
	if len(height) <= 0 {
		fmt.Println("incorrect length: length must be greater than 0")
		return 0
	}
	i := 0
	j := len(height) - 1
	area := 0

	for i < j {
		h := min(height[i], height[j])

		width := j - i

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		if width*h > area {
			area = width * h
		}
	}
	return area
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
**/

func TwoSum(nums []int, target int) []int {

	if len(nums) <= 0 {
		fmt.Println("incorrect length: length must be greater than 0")
		return []int{}
	}

	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		comp := target - v

		if j, ok := m[comp]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
