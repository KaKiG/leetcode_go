package main

func main() {

}

func findRepeatedDnaSequences(s string) []string {
	wordMap := make(map[string]int)
	res := make([]string, 0)
	for i := 0; i < len(s)-9; i++ {
		str := s[i : i+10]
		if wordMap[str] != -1 {
			wordMap[str]++
		}
		if wordMap[str] > 1 {
			res = append(res, str)
			wordMap[str] = -1
		}

	}
	return res
}
