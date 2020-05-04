package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 3, 2, 1, 2, 6, 5, 4, 3, 3, 2, 1, 1, 3, 3, 3, 4, 2}
	fmt.Println(candy(n))
}

func candy(ratings []int) int {
	candy := make([]int, len(ratings))
	lowIndex := make([]int, 0)

	if len(ratings) == 1 {
		return 1
	}

	for i := range ratings {
		if i == 0 && ratings[1] >= ratings[0] {
			lowIndex = append(lowIndex, 0)
		}
		if i == len(ratings)-1 && ratings[len(ratings)-2] >= ratings[len(ratings)-1] {
			lowIndex = append(lowIndex, i)
		}

		if i > 0 && i < len(ratings)-1 {
			if ratings[i] <= ratings[i-1] && ratings[i] <= ratings[i+1] {
				lowIndex = append(lowIndex, i)
			}
		}
	}

	for _, i := range lowIndex {
		candy[i] = 1
		left := i - 1
		right := i + 1
		for left >= 0 || right < len(ratings) {
			check := false
			if left > 0 && ratings[left-1] >= ratings[left] && ratings[left] > ratings[left+1] {
				candy[left] = candy[left+1] + 1
				left--
				check = true
			}
			if left == 0 && ratings[left] > ratings[left+1] {
				candy[left] = candy[left+1] + 1
				left--
				check = true
			}
			if right < len(ratings)-1 && ratings[right+1] >= ratings[right] && ratings[right] > ratings[right-1] {
				candy[right] = candy[right-1] + 1
				right++
				check = true
			}
			if right == len(ratings)-1 && ratings[right] > ratings[right-1] {
				candy[right] = candy[right-1] + 1
				right++
				check = true
			}
			if !check {
				break
			}
		}
	}
	sum := 0

	for i := range candy {
		if candy[i] == 0 {
			sum += max2(candy[i-1], candy[i+1]) + 1
		} else {
			sum += candy[i]
		}
	}
	return sum
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
