package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < math.MaxInt32; i++ {
		if isPowerOfFour(i) {
			fmt.Println(i)
		}
	}
}

func isPowerOfFour(num int) bool {
	posn := -1
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			if posn != -1 {
				return false
			}
			posn = i
		}
		num = num >> 1
	}

	return posn%2 == 0
}
