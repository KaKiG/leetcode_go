import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

// @lc code=start
type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UTC().UnixNano())
	sol := Solution{nums}
	return sol
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.arr)
	res := append([]int{}, this.arr...)
	for currPos := range this.arr {
		ind := rand.Int() % length
		res[currPos], res[ind+currPos] = res[ind+currPos], res[currPos]
		length--
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
