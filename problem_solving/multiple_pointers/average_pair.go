package main

/**
Multiple Pointers - averagePair

Write a function called averagePair. Given a sorted array of integers and a target average,
determine if there is a pair of values in the array where the average of the pair equals the target average.
There may be more than one pair that matches the average target.

**/

func AveragePair(nums []int, avg float64) bool {

	if len(nums) <= 0 || avg <= 0 {
		return false
	}

	end := len(nums) - 1
	start := 0
	for start < end {

		if float64((nums[start]+nums[end]))/2 == avg {
			return true
		}

		if float64((nums[start]+nums[end]))/2 < avg {
			start++
		} else {
			end--
		}
	}
	return false
}
