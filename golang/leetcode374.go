package main

func main() {

}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	gus := 1
	left := 1
	right := n

	for guess(gus) != 0 {
		gus := (1 + n) / 2
		tmp := guess(gus)

		if tmp == -1 {
			right = gus
		}

		if tmp == 1 {
			left = gus + 1
		}

		if tmp == 0 {
			return gus
		}
	}
	return gus
}
