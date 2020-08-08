// package main

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
// func main() {
// 	obj := Constructor()
// 	param_1 := obj.Insert(0)
// 	fmt.Println(obj, param_1)
// 	param_2 := obj.Insert(1)
// 	fmt.Println(obj, param_2)
// 	param_3 := obj.Remove(0)
// 	fmt.Println(obj, param_3)
// 	param_4 := obj.Insert(2)
// 	fmt.Println(obj, param_4)
// 	// param_5 := obj.Remove(1)
// 	// fmt.Println(obj, param_5)
// 	param_6 := obj.GetRandom()
// 	fmt.Println(obj, param_6)
// 	param_7 := obj.GetRandom()
// 	fmt.Println(obj, param_7)
// }

type RandomizedSet struct {
	numInd map[int]int
	numArr []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rand.Seed(time.Now().UTC().UnixNano())
	newSet := RandomizedSet{make(map[int]int), []int{}}
	return newSet
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.numInd[val]; !exist {
		this.numInd[val] = len(this.numArr)
		this.numArr = append(this.numArr, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, exist := this.numInd[val]; exist {
		ind := this.numInd[val]
		end := len(this.numArr) - 1
		if ind != end {
			this.numArr[ind] = this.numArr[end]
			this.numInd[this.numArr[ind]] = ind
		}
		this.numArr = this.numArr[:end]
		delete(this.numInd, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	randInt := rand.Int() % len(this.numArr)
	return this.numArr[randInt]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
