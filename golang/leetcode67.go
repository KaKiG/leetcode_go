package main

import "fmt"

func main() {
	a := "101111"
	b := "10"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	carrier := 0
	indexA := len(a) - 1
	indexB := len(b) - 1
	tmp := ""
	res := ""
	for indexA >= 0 && indexB >= 0 {
		if string(a[indexA]) == "1" && string(b[indexB]) == "1" && carrier == 1 {
			tmp += "1"
			indexA--
			indexB--
			carrier = 1
			continue
		}
		if (string(a[indexA]) == "1" && string(b[indexB]) == "0" && carrier == 0) ||
			(string(a[indexA]) == "0" && string(b[indexB]) == "1" && carrier == 0) ||
			(string(a[indexA]) == "0" && string(b[indexB]) == "0" && carrier == 1) {
			tmp += "1"
			indexA--
			indexB--
			carrier = 0
			continue
		}
		if (string(a[indexA]) == "1" && string(b[indexB]) == "1" && carrier == 0) ||
			(string(a[indexA]) == "0" && string(b[indexB]) == "1" && carrier == 1) ||
			(string(a[indexA]) == "1" && string(b[indexB]) == "0" && carrier == 1) {
			tmp += "0"
			indexA--
			indexB--
			carrier = 1
			continue
		}
		if string(a[indexA]) == "0" && string(b[indexB]) == "0" && carrier == 0 {
			tmp += "0"
			indexA--
			indexB--
			carrier = 0
			continue
		}
	}
	i := indexB
	if indexB == -1 {
		if carrier == 1 {
			tmp += "1"
		}
	}
	if indexB != -1 {
		if carrier == 1 {
			for i = indexB; i >= 0; i-- {
				if string(b[i]) == "0" {
					for j := indexB; j > i; j-- {
						tmp += "0"
					}
					tmp += "1"
					for j := i - 1; j >= 0; j-- {
						tmp += string(b[j])
					}
					break
				}
				if i == 0 && string(b[0]) == "1" {
					for j := indexB; j >= 0; j-- {
						tmp += "0"
					}
					tmp += "1"
					break
				}
			}
		} else {
			for j := indexB; j >= 0; j-- {
				tmp += string(b[j])
			}
		}
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		res += string(tmp[i])
	}
	return res
}
