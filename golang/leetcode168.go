package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(convertToTitle(i))
	}
}

func convertToTitle(n int) string {
	rest := n
	res := ""
	for rest != 0 {
		if rest%26 == 0 {
			res = string('Z') + res
			rest--
		} else {
			res = string('@'+rest%26) + res
		}
		rest /= 26
	}
	return res
}
