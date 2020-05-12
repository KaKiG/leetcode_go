package main

func main() {

}

func isSubsequence(s string, t string) bool {
	l1 := 0

	for i := range t {
		if t[i] == s[l1] {
			l1++
			if l1 == len(s) {
				return true
			}
		}
	}

	return false
}
