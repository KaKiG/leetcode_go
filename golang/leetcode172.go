package main

import "fmt"

func main() {
	n := 200
	fmt.Println(trailingZeroes(n))
}

//this problem is to count 5....
//to compute trailingZeroes, we can see zero is produced from 5 and 2.
//Through the numbers, 2 is much more than 5, therefore we only need to compute the number of 5
//be cautious that something like 25 could produce more than 1 five
func trailingZeroes(n int) int {
	count := 0
	for n != 0 {
		count += n / 5
		n /= 5
	}
	return count
}
