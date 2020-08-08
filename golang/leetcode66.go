package main

import "fmt"

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}
	if len(digits) == 1 {
		if digits[0] == 9 {
			return []int{1, 0}
		}
		digits[0]++
		return digits
	}
	carrier := 0
	tmp := digits[len(digits)-1] + 1
	fmt.Println(tmp)
	if tmp == 10 {
		carrier = 1
		digits[len(digits)-1] = 0
	} else {
		digits[len(digits)-1] = tmp
		return digits
	}

	for i := len(digits) - 2; i > 0; i-- {
		if digits[i]+carrier == 10 {
			digits[i] = 0
		} else {
			digits[i] += carrier
			carrier = 0
			return digits
		}
	}

	if digits[0]+carrier == 10 {
		digits[0] = 0
	} else {
		digits[0] += carrier
		return digits
	}
	tmpSlice := append([]int{1}, digits...)
	return tmpSlice
}
