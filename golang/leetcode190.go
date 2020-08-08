package main

func main() {

}

// bit operation
// each time, we left shift res by one and extract the last bit of numbers
// then numbers right shift by one
func reverseBits(num uint32) uint32 {
	var res uint32

	for i := 0; i < 32; i++ {
		res = res<<1 | (num & 1)
		num = num >> 1
	}
	return res
}
