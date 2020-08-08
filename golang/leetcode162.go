package main

func main() {

}

//binary search
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	low, high := 0, len(nums)-1
	mid := (low + high) / 2

	for {
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] < nums[mid+1] {
			low = mid
		} else if nums[mid] > nums[mid+1] {
			high = mid
		}

		mid = (low + high) / 2
	}
	return -1
}

//linear search
// func findPeakElement(nums []int) int {
// 	if len(nums) == 1 {
// 		return 0
// 	}
//
// 	for i := range nums {
// 		if i == 0 {
// 			if nums[0] > nums[1] {
// 				return 0
// 			}
// 		}
// 		if i == len(nums)-1 {
// 			return i
// 		}
//
// 		if i != 0 && nums[i] > nums[i-1] && nums[i] > nums[i+1] {
// 			return i
// 		}
// 	}
// 	return -1
// }
