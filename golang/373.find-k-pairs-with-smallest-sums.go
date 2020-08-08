// package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start

/* Straight forward idea with somehow bruteforce, but without using the numbers in ascending order */
// type Pair struct {
// 	num1, num2 int
// }

// type Pairs []Pair

// func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
// 	if k < len(nums1) {
// 		nums1 = nums1[:k]
// 	}
// 	if k < len(nums2) {
// 		nums2 = nums2[:k]
// 	}
// 	var pairs Pairs
// 	pairs = make([]Pair, len(nums1)*len(nums2))
// 	for i := range nums1 {
// 		for j := range nums2 {
// 			pairs[i*len(nums2)+j].num1 = nums1[i]
// 			pairs[i*len(nums2)+j].num2 = nums2[j]
// 		}
// 	}

// 	sort.Sort(Pairs(pairs))
// 	if len(pairs) < k {
// 		return Convert(pairs)
// 	}
// 	return Convert(pairs)[:k]
// }

// func Convert(pairs Pairs) [][]int {
// 	res := make([][]int, len(pairs))
// 	for i := range res {
// 		res[i] = make([]int, 2)
// 		res[i][0] = pairs[i].num1
// 		res[i][1] = pairs[i].num2
// 	}
// 	return res
// }

// func (pairs Pairs) Len() int {
// 	return len(pairs)
// }

// func (pairs Pairs) Swap(i, j int) {
// 	pairs[i], pairs[j] = pairs[j], pairs[i]
// }

// func (pairs Pairs) Less(i, j int) bool {
// 	return (pairs[i].num1 + pairs[i].num2) < (pairs[j].num1 + pairs[j].num2)
// }

/* another idea is to use minHeap to keep track of k elements */
// An IntHeap is a min-heap of ints.

// func main() {
// 	nums1 := []int{1, 1, 2}
// 	nums2 := []int{1, 2, 3}
// 	k := 10
// 	fmt.Println(kSmallestPairs(nums1, nums2, k))
// }

type Pair struct {
	num1, ind1, num2, ind2 int
}

type minHeap []*Pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return (h[i].num1 + h[i].num2) < (h[j].num1 + h[j].num2) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Pair))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}
	res := []Pair{}

	h := make(minHeap, 0)
	heap.Init(&h)
	for i := range nums1 {
		heap.Push(&h, &Pair{nums1[i], i, nums2[0], 0})
	}
	for k > 0 && h.Len() > 0 {
		tmp := heap.Pop(&h).(*Pair)
		k--
		res = append(res, *tmp)
		if tmp.ind2 < len(nums2)-1 {
			heap.Push(&h, &Pair{nums1[tmp.ind1], tmp.ind1, nums2[tmp.ind2+1], tmp.ind2 + 1})
		}

	}
	return Convert(res)
}

func Convert(pairs []Pair) [][]int {
	res := make([][]int, len(pairs))
	for i := range res {
		res[i] = make([]int, 2)
		res[i][0] = pairs[i].num1
		res[i][1] = pairs[i].num2
	}
	return res
}

// @lc code=end
