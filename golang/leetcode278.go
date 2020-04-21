package main

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left := 1
	right := n
	middle := 0
	for left < right {
		middle = (left + right) / 2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	if isBadVersion(middle) {
		return middle
	}
	return right
}
