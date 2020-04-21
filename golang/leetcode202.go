package main

import "fmt"

func main() {
	num := 2
	fmt.Println(isHappy(num))
}

func isHappy(n int) bool {
	numMap := make(map[int]bool)

	for {
		fmt.Println(n)
		if SumSquare(n) == 1 {
			return true
		}
		numMap[n] = true
		n = SumSquare(n)
		if numMap[n] {
			return false
		}
	}
	return false
}

func SumSquare(num int) int {
	sum := 0

	for num/10 != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	sum += num * num
	return sum
}
