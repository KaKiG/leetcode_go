package main

func main() {

}

// loop method
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}

	for {
		if n%3 == 0 {
			n = n / 3
		} else if n == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}

// recursion method
// func isPowerOfThree(n int) bool {
// 	if n == 0 {
// 		return false
// 	}
// 	if n == 1 {
// 		return true
// 	}

// 	return (n%3 == 0) && isPowerOfThree(n/3)
// }
