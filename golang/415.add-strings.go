// package main

// import "fmt"

/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
// func main() {
// 	num1 := "9999999901"
// 	num2 := "99"
// 	fmt.Println(addStrings(num1, num2))
// }
func addStrings(num1 string, num2 string) string {
	length := max2(len(num1), len(num2))
	res := ""
	carrier := 0
	for i := 0; i < length; i++ {
		dec1 := 0
		dec2 := 0
		if len(num1)-1-i >= 0 {
			dec1 = int(num1[len(num1)-1-i] - '0')
		}
		if len(num2)-1-i >= 0 {
			dec2 = int(num2[len(num2)-1-i] - '0')
		}
		res = string((dec1+dec2+carrier)%10+'0') + res
		carrier = (dec1 + dec2 + carrier) / 10
	}
	if carrier != 0 {
		res = string(carrier+'0') + res
	}
	return res
}

func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
