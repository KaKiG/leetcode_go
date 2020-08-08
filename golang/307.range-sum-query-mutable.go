import (
	"math"
)

/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */

// @lc code=start
type NumArray struct {
	array    []int
	sumArray []int
	window   int
}

func Constructor(nums []int) NumArray {
	win := int(math.Log(float64(len(nums))))
	if win == 0 {
		win = 1
	}
	arr := make([]int, len(nums))
	sumArr := make([]int, 0)
	for i := range nums {
		arr[i] = nums[i]
	}

	count := 0
	sum := 0
	for i := range nums {
		if count < win {
			count += 1
			sum += nums[i]
		}
		if count == win || i == len(nums)-1 {
			if len(sumArr) != 0 {
				sumArr = append(sumArr, sum+sumArr[len(sumArr)-1])
			} else {
				sumArr = append(sumArr, sum)
			}
			sum = 0
			count = 0
		}
	}
	return NumArray{arr, sumArr, win}
}

func (this *NumArray) Update(i int, val int) {
	pos := i / this.window
	for j := pos; j < len(this.sumArray); j++ {
		this.sumArray[j] += (val - this.array[i])
	}
	this.array[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.FindSum(j) - this.FindSum(i-1)
}

func (this *NumArray) FindSum(i int) int {
	if i == -1 {
		return 0
	}
	pos := i / this.window
	sum := 0
	if pos != 0 {
		sum = this.sumArray[pos-1]
	}
	for j := this.window * pos; j <= i; j++ {
		sum += this.array[j]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end
