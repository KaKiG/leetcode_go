package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	num := []int{}

	for head != nil {
		num = append(num, head.Val)
		head = head.Next
	}

	return isPa(num)
}

func isPa(num []int) bool {
	left := 0
	right := len(num) - 1

	for left < right {
		if num[left] != num[right] {
			return false
		}
		left++
		right--
	}
	return true
}
