package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	tmpMap := make(map[string][]string)

	for i := range strs {
		tmpStr := strs[i]
		strSplit := strings.Split(strs[i], "")
		sort.Strings(strSplit)
		tmpStr = strings.Join(strSplit, "")
		tmpMap[tmpStr] = append(tmpMap[tmpStr], strs[i])
	}
	for _, v := range tmpMap {
		res = append(res, v)
	}
	return res
}

/*func groupAnagrams(strs []string) [][]string {
	strFreqMap := make(map[int](map[string]int))
	letterFreqMap := make(map[string]int)
	res := make([][]string, 0)

	for i := range strs {
		letterFreqMap = CreateLetterFreqMap(strs[i])
		strFreqMap[i] = letterFreqMap
	}

	for i := range strs {
		_, existi := strFreqMap[i]
		if existi {
			tmp := make([]string, 0)
			tmp = append(tmp, strs[i])
			for j := i + 1; j < len(strs); j++ {
				_, existj := strFreqMap[j]
				if existj {
					if MapCompare(strFreqMap[i], strFreqMap[j]) {
						tmp = append(tmp, strs[j])
						delete(strFreqMap, j)
					}
				}
			}
			delete(strFreqMap, i)
			res = append(res, tmp)
		}
	}
	return res
}

func CreateLetterFreqMap(str string) map[string]int {
	freqMap := make(map[string]int)
	for i := range str {
		_, exist := freqMap[str[i:i+1]]
		if exist {
			freqMap[str[i:i+1]]++
		} else {
			freqMap[str[i:i+1]] = 1
		}
	}
	return freqMap
}

func MapCompare(strMap1, strMap2 map[string]int) bool {
	if len(strMap1) != len(strMap2) {
		return false
	}

	for k, v := range strMap1 {
		if v != strMap2[k] {
			return false
		}
	}

	return true
}*/
