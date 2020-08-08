import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	listArr []int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UTC().UnixNano())
	sol := Solution{[]int{}}
	curr := head
	for curr != nil {
		sol.listArr = append(sol.listArr, curr.Val)
		curr = curr.Next
	}
	return sol
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	randInt := rand.Int() % len(this.listArr)
	return this.listArr[randInt]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
