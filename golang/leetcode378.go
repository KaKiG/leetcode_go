package main

import "fmt"

func main() {
	// matrix := [][]int{
	// 	{1, 5, 9},
	// 	{10, 11, 13},
	// 	{12, 13, 15},
	// }
	// k := 8
	matrix := [][]int{
		{1, 1},
		{1, 7},
	}
	k := 3
	fmt.Println(kthSmallest(matrix, k))
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]
	right := matrix[n-1][n-1]
	mid := 0

	for left < right {
		count := 0
		mid = (left + right) / 2
		fmt.Println(left, mid, right)
		j := n - 1
		i := 0
		for ; i < n; i++ {
			for j >= 0 && mid < matrix[i][j] {
				j--
			}
			count = count + j + 1
		}

		// if count == k {
		// 	return matrix[i][j]
		// }

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println(left, mid, right)
	return left
}
