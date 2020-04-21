package main

func main() {

}

// Binary serach

func hIndex(citations []int) int {
	left := 0
	right := len(citations)
	middle := 0
	for left < right {
		middle = (left + right) / 2
		if citations[middle] >= len(citations)-middle {
			if middle == 0 || citations[middle-1] <= len(citations)-middle {
				return len(citations) - middle
			}
			right = middle
		} else {
			left = middle + 1
		}
	}
	return 0
}

// Naive method

// func hIndex(citations []int) int {
// 	if len(citations) == 0 {
// 		return 0
// 	}
//     for i := len(citations); i >= 0; i-- {
// 		if i == len(citations) && citations[len(citations)-i] >= i {
// 			return i
// 		} else if i == 0 {
// 			return 0
// 		} else if citations[len(citations)-i] >= i && citations[len(citations)-1-i] <= i {
// 			return i
// 		}
// 	}
// 	return 0
// }
