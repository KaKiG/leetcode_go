package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	length := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[length-i] = s[length-i], s[i]
	}
}
