package main

/**
Frequency Counter - findAllDuplicates

Given an array of positive integers, some elements appear twice and others appear once.
Find all the elements that appear twice in this array.
Note that you can return the elements in any order.

**/

func FindAllDuplicates(nums []int) []int {

	if len(nums) <= 0 {
		return nil
	}

	dups := make(map[int]int)
	dupsSlice := []int{}
	for _, v := range nums {
		dups[v]++
		if dups[v] > 1 {
			dupsSlice = append(dupsSlice, v)
		}
	}

	return dupsSlice
}
