package main

func main() {

}

func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}
	return b
}
