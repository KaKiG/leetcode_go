package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(search(nums, 2))
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left := 0
	right := len(nums) - 1
	mid := 0
	if left <= right {
		mid = (right + left) / 2
		if nums[mid] == target || nums[right] == target || nums[left] == target {
			return true
		}
		if nums[mid] < nums[right] {
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[right] {
			right--
		} else {
			if target < nums[mid] && target > nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		mid = (right + left) / 2
		if left <= mid && mid <= right {
			return search(nums[left:mid], target) || search(nums[mid:right], target)
		}
	}
	return false
}

/*func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target || nums[right] == target || nums[left] == target {
			return true
		}
		if nums[mid] < nums[right] {
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[right] {
			right--
		} else {
			if target < nums[mid] && target > nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}*/
