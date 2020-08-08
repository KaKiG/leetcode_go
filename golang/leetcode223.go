package main

import "fmt"

func main() {
	A := -3
	B := 0
	C := 3
	D := 4
	E := 0
	F := -1
	G := 9
	H := 2
	fmt.Println(computeArea(A, B, C, D, E, F, G, H))
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sumArea := computeRectangele(A, B, C, D) + computeRectangele(E, F, G, H)
	x, y := 0, 0
	if min2(C, G) > max2(A, E) {
		x = min2(C, G) - max2(A, E)
	}

	if min2(D, H) > max2(B, F) {
		y = min2(D, H) - max2(B, F)
	}
	return sumArea - x*y

}

func computeRectangele(A, B, C, D int) int {
	xDiff := A - C
	yDiff := B - D
	return absolute(xDiff * yDiff)
}

func absolute(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
