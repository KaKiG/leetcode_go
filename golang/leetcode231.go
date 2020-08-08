package main

import "fmt"

func main() {
	n := 218
	fmt.Println(isPowerOfTwo(n))
}

// using bitwise operation, if an number is the power of two
// then there is only one '1'
func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	count := 0
	for i := 0; i < 32; i++ {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	if count == 1 {
		return true
	}
	return false
}

// a litte bit slower
// func isPowerOfTwo(n int) bool {
// 	if n == 1 {
// 		return true
// 	}
// 	if n == 0 || n%2 == 1 {
// 		return false
// 	}
// 	return isPowerOfTwo(n / 2)
// }
