package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.MaxInt32
	fmt.Println(countPrimes(n))
}

// count the number of primes less than n
func countPrimes(n int) int {
	count := 0
	if n < 3 {
		return count
	}
	numList := make([]bool, n)
	numList[0] = true
	count++
	for i := 2; i < n; i++ {
		if numList[i] == false {
			j := 0
			for j = i * i; j < n; j++ {
				if j%i == 0 {
					if numList[j] == false {
						count++
					}
					numList[j] = true
					break
				}
			}

			for j < n {
				if numList[j] == false {
					count++
				}
				numList[j] = true
				j += i
			}
		}
	}
	return n - count - 1
}
