package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	counter := make([]int, 26)

	for _, char := range magazine {
		counter[char-'a']++
	}

	for _, char := range ransomNote {
		if counter[char-'a'] == 0 {
			return false
		}
		counter[char-'a']--
	}

	return true
}
