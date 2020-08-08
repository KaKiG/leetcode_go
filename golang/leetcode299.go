package main

import (
	"fmt"
)

func main() {
	secret := "1122"
	guess := "1222"

	fmt.Println(getHint(secret, guess))
}

// a faster method using only one loop
func getHint(secret string, guess string) string {
	countA := 0
	countB := 0

	var numS, numG [10]int

	for i := range secret {
		s := int(secret[i] - '0')
		g := int(guess[i] - '0')

		if s == g {
			countA++
			continue
		} else {
			if numG[s] > 0 {
				countB++
				numG[s]--
			} else {
				numS[s]++
			}

			if numS[g] > 0 {
				countB++
				numS[g]--
			} else {
				numG[g]++
			}
		}
	}

	return fmt.Sprintf("%vA%vB", countA, countB)
}

// func getHint(secret string, guess string) string {
// 	countA := 0
// 	countB := 0
// 	countADic := make(map[int]bool)
// 	countBDic := make(map[byte]int)
// 	for i := range secret {
// 		if secret[i] == guess[i] {
// 			countADic[i] = true
// 			countA++
// 		} else {
// 			countBDic[secret[i]]++
// 		}
// 	}

// 	for i := range secret {
// 		if countADic[i] == false && countBDic[guess[i]] > 0 {
// 			countB++
// 			countBDic[guess[i]]--
// 		}
// 	}

// 	return fmt.Sprintf("%vA%vB", countA, countB)
// }
