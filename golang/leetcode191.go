package main

func main() {

}

// just a naive solution, very simple
func hammingWeight(num uint32) int {
	count := 0

	for i := 0; i < 32; i++ {
		if (num & 1) == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
