package main

import "math"

func main() {

}

// we can store the currently smallest and second smallest value
// if we find another value greater than this two return true
// otherwise update the two values

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	small1, small2 := math.MaxInt32, math.MaxInt32

	for i := range nums {
		if nums[i] <= small1 {
			small1 = nums[i]
		} else if nums[i] <= small2 {
			small2 = nums[i]
		} else {
			return true
		}
	}
	return false
}
