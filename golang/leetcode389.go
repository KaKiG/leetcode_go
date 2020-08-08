package main

func main() {

}

func findTheDifference(s string, t string) byte {
	counter := make([]int, 26)

	for _, char := range s {
		counter[char-'a']++
	}

	for idx, char := range t {
		if counter[char-'a'] == 0 {
			return t[idx]
		}
		counter[char-'a']--
	}

	return 's'
}
