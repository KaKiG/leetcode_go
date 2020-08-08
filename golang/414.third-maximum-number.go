/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {
	max1 := 0
	max2 := 0
	max3 := 0
	max1Assign := false
	max2Assign := false
	max3Assign := false

	for i := range nums {
		if !max1Assign {
			max1 = nums[i]
			max1Assign = true
		} else if !max2Assign && nums[i] != max1 {
			if nums[i] > max1 {
				max2 = max1
				max1 = nums[i]
			} else {
				max2 = nums[i]
			}
			max2Assign = true
		} else if !max3Assign && nums[i] != max1 && nums[i] != max2 {
			if nums[i] > max1 {
				max3 = max2
				max2 = max1
				max1 = nums[i]
			} else if nums[i] > max2 {
				max3 = max2
				max2 = nums[i]
			} else {
				max3 = nums[i]
			}
			max3Assign = true
			continue
		}

		if max1Assign && max2Assign && max3Assign {
			if nums[i] > max1 {
				max3 = max2
				max2 = max1
				max1 = nums[i]
			} else if nums[i] < max1 && nums[i] > max2 {
				max3 = max2
				max2 = nums[i]
			} else if nums[i] < max2 && nums[i] > max3 {
				max3 = nums[i]
			}
		}
	}
	if !max3Assign {
		return max1
	}
	return max3
}

// @lc code=end
