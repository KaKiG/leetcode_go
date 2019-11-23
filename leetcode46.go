package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ans := make([][]int, 0)
	BacktrackPer(nums, nil, &ans)
	return ans
}

func BacktrackPer(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
	}

	for i := range nums {
		BacktrackPer(append(append([]int{}, nums[0:i]...), nums[i+1:]...), append(prev, nums[i]), ans)
	}
}

/*func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums[0:1]}
	}
	return InsertPers(permute(nums[:len(nums)-1]), nums[len(nums)-1])
}

func InsertPers(numsPers [][]int, num int) [][]int {
	res := make([][]int, 0)
	for i := range numsPers {
		for j := 0; j <= len(numsPers[i]); j++ {
			res = append(res, InsertOnePer(numsPers[i], num, j))
		}
	}
	//fmt.Println(res)
	return res
}

func InsertOnePer(numsPer []int, num, index int) []int {
	res := make([]int, 0)
	tmp := make([]int, len(numsPer))
	copy(tmp, numsPer)
	//fmt.Println(numsPer, num, index, -1)
	if index == 0 {
		res = append(res, num)
		res = append(res, tmp...)
		//fmt.Println(res, num, index, -2)
		return res
	}
	if index == len(numsPer) {
		res = append(res, tmp...)
		res = append(res, num)
		//fmt.Println(res, num, index, -3)
		return res
	}

	res = append(tmp[0:index], num)
	//fmt.Println(res, num, index, tmp, numsPer, -4.1)
	res = append(res, numsPer[index:]...)
	//fmt.Println(res, num, numsPer, -4)
	return res
}*/
