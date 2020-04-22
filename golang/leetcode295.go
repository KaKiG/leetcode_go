package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddNum(41)
	fmt.Println(obj.FindMedian())
	// fmt.Println(obj.maxHeap)
	// fmt.Println(obj.minHeap)

	obj.AddNum(35)
	fmt.Println(obj.FindMedian())
	// fmt.Println(obj.maxHeap)
	// fmt.Println(obj.minHeap)

	obj.AddNum(62)
	fmt.Println(obj.FindMedian())
	// fmt.Println(obj.maxHeap)
	// fmt.Println(obj.minHeap)

	obj.AddNum(4)
	fmt.Println(obj.FindMedian())
	// fmt.Println(obj.maxHeap)
	// fmt.Println(obj.minHeap)

	obj.AddNum(97)
	fmt.Println(obj.FindMedian())
	// fmt.Println(obj.maxHeap)
	// fmt.Println(obj.minHeap)

	obj.AddNum(108)
	fmt.Println(obj.FindMedian())
	// fmt.Println(obj.maxHeap)
	// fmt.Println(obj.minHeap)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.

	// since we are defining intheap, we need to specify the type of x as int.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	maxHeap *IntHeap
	minHeap *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var mf MedianFinder
	mf.maxHeap = &IntHeap{}
	mf.minHeap = &IntHeap{}
	heap.Init(mf.maxHeap)
	heap.Init(mf.minHeap)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, -num)
	// fmt.Println(this.maxHeap, "step 1")
	// here we need .(int) because actually heap.pop is a interface
	// to apply -, we need to give it a specific type, such as int
	tmp := -heap.Pop(this.maxHeap).(int)
	// fmt.Println(this.maxHeap, tmp, "step 2")
	heap.Push(this.minHeap, tmp)
	if this.minHeap.Len() > this.maxHeap.Len() {
		tmp = -heap.Pop(this.minHeap).(int)
		heap.Push(this.maxHeap, tmp)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(-(*this.maxHeap)[0])
	} else {
		return (float64(-(*this.maxHeap)[0]) + float64((*this.minHeap)[0])) / 2
	}
}

// /** initialize your data structure here. */
// func Constructor() MedianFinder {
// 	var mf MedianFinder
// 	mf.nums = make([]int, 0)
// 	return mf
// }

// func (this *MedianFinder) AddNum(num int) {
// 	if len(this.nums) == 0 {
// 		this.nums = append(this.nums, num)
// 	} else if num > this.nums[len(this.nums)-1] {
// 		this.nums = append(this.nums, num)
// 	} else if num < this.nums[0] {
// 		this.nums = append([]int{num}, this.nums...)
// 	} else {
// 		ind := BinarySearch(this.nums, num)
// 		arr := append([]int{}, this.nums[:ind]...)
// 		arr = append(arr, num)
// 		arr = append(arr, this.nums[ind:]...)
// 		this.nums = arr
// 	}
// }

// func (this *MedianFinder) FindMedian() float64 {
// 	if len(this.nums)%2 == 1 {
// 		return float64(this.nums[len(this.nums)/2])
// 	} else {
// 		return (float64(this.nums[len(this.nums)/2]) + float64(this.nums[len(this.nums)/2-1])) / 2.0
// 	}
// }

// func BinarySearch(nums []int, target int) int {
// 	left := 0
// 	right := len(nums) - 1
// 	mid := (left + right) / 2

// 	for left < right {
// 		mid = (left + right) / 2

// 		if nums[mid] == target {
// 			return mid
// 		}

// 		if nums[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
// 	return right
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
