package main

import "fmt"

func main() {
	num := 6
	fmt.Println(isUgly(num))
}

func isUgly(num int) bool {
	if num == 1 {
		return true
	}
	var res2, res3, res5 bool
	if num%2 == 0 {
		res2 = true
	}
	if num%3 == 0 {
		res3 = true
	}
	if num%5 == 0 {
		res5 = true
	}
	return (res2 && isUgly(num/2)) || (res3 && isUgly(num/3)) || (res5 && isUgly(num/5))
}
