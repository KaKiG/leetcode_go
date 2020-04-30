package main

func main() {

}

// this problem is equivalent to find integer solution to z = m*x + n*y
// from Bézout's lemma, we know
// z%gcd(x, y) == 0 is the condition

// be careful to the gcd implementation (fast!)
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if z > x+y {
		return false
	}

	return (z%gcd(x, y) == 0)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
