package main

import (
	"fmt"
	"math"
	"slices"
)

/**

Write a function that accepts a sorted array of integers
and returns the first pair where the sum is 0
return an array that includes both elements or undefined if non exist

**/

// O(N^2)
func SumZeroQuadratic(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == 0 {
				return []int{arr[i], arr[j]}
			}
		}
	}
	return nil
}

// O(n) / O(log(n)) with sorting function
// converging double pointer pattern
func SumZero(arr []int) []int {

	if !slices.IsSorted(arr) {
		slices.Sort(arr)
	}

	// start and end didnt really feel like pointers so i added explicit p1,p2 pointers for readability
	start := 0
	end := len(arr) - 1

	p1 := &arr[start]
	p2 := &arr[end]

	for i := 0; i < len(arr); i++ {

		if *p1+*p2 == 0 {
			return []int{arr[start], arr[end]}
		}

		if *p1+*p2 > 0 {
			end = end - 1 // end--
			p2 = &arr[end]
			continue
		}
		if *p1+*p2 < 0 {
			start = start + 1 // start++
			p1 = &arr[start]
			continue
		}
	}
	return nil
}

// implement a function that takes in a sorted array and counts the unique values in the array
// the array may have negatives but will always be sorted

// slow fast double pointer pattern
func CountUniqueVals(arr []int) int {

	if !slices.IsSorted(arr) {
		slices.Sort(arr)
	}

	// set a pointer at the first element
	i := 0
	// set second pointer at second element
	for j := 1; j < len(arr); j++ {
		// check to see if the elements are not the same
		if arr[i] != arr[j] {
			// if the elements are not the same then increment the first pointer as the value is unique and we want to keep it
			i++
			// set the element at the new i position to be that of index j
			arr[i] = arr[j]
			continue
		}
	}
	// return i plus 1 giving us the unique elements swapped to the start and i sitting at the most recent unique values position
	return i + 1
}

/**
Problem: Count Pairs With Sum ≤ Target

You are given a sorted array of integers (ascending) and a target integer T.
Return the number of unique index pairs (i, j) with i < j such that
arr[i] + arr[j] <= T.

Use an O(n) solution with a converging two-pointer approach
(one at the start, one at the end, move inward).

arr := []int{1, 2, 3, 4, 6}
T := 7
**/

// Given a sorted array of integers, return all unique pairs that sum to a given target.

// Example: arr = [1,2,3,4,5,6,7]
// target = 8

// result → [[1,7], [2,6], [3,5]]

// converging double pointer pattern
func TargetPairs(arr []int, target int) [][]int {

	var pairs [][]int

	i := 0
	j := len(arr) - 1
	for i < j {

		if arr[i]+arr[j] > target {
			j--
		}

		if arr[i]+arr[j] < target {
			i++
		}

		if arr[i]+arr[j] == target {
			pairs = append(pairs, []int{arr[i], arr[j]})
			j--
			i++
		}
	}
	return pairs
}

/**
Find the Closest Pair to Target (Sorted Array)

Given a sorted array of integers and a target value,
return the pair whose sum is closest to the target.

If two pairs are equally close, return either one.
**/

// converging double pointer pattern
func ClosestPair(arr []int, target int) []int {

	var best []int

	// place pointers
	start := 0
	end := len(arr) - 1
	closestDiff := math.MaxInt64

	for i := 0; i < len(arr); i++ {
		sum := arr[start] + arr[end]

		diff := target - sum
		if diff < 0 {
			diff = -diff
		}

		if diff < closestDiff {
			closestDiff = diff
			best = []int{arr[start], arr[end]}
		}
		if sum > target {
			end--
		}
		if sum < target {
			start++
		}
	}
	return best
}

/**
Given an array of integers and a number k,
find the maximum possible sum of any subarray of length k.

input: [2, 6, 9, 2, 1, 8, 5, 6, 3], k=3

// sliding window
**/

func maxSubarraySum(arr []int, target int) int {

	// check for empty input
	if len(arr) <= 0 || target <= 0 || len(arr) < target {
		fmt.Println("please provide a valid array and target")
		return 0
	}

	// create the Max and Temp sums
	var maxSum, tempSum int

	// loop over the sub array and create the sum
	for i := 0; i < target; i++ {
		maxSum += arr[i]
	}

	tempSum = maxSum
	// loop from the target IE 3 to the end of the array
	// on each iteration assign tempSum to be equal to tempSum - the sub arrays first element + the element after the sub array
	for i := target; i < len(arr); i++ {
		// this is the magic
		// i = 3 in our example, i - target(also 3) - 0 giving us the first element in the array
		// next slide becomes 4 - target(3) = 1
		// next slide becomes 5 - target(3) = 2
		// so the window essentially slides one place forward and does a single calculation on each step
		// MaxSum - firstElement + NextElement
		tempSum = tempSum - arr[i-target] + arr[i]
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}
