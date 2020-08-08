package main

func main() {

}

// func addDigits(num int) int {
// 	if num <= 9 {
// 		return num
// 	}
// 	return addDigits(num/10 + num%10)
// }

// Alternative O(1) method
func addDigits(num int) int {
	mod := num % 9
	if num == 0 {
		return 0
	}
	if mod == 0 {
		return 9
	}
	return mod
}
