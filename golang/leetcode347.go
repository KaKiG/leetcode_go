package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)
	freqMap := make(map[int]int)

	for _, v := range nums {
		freqMap[v]++
	}

	for k, v := range freqMap {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, 0)
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) != 0 && k != 0 {
			// guarantee unique solution, so we can do that directly
			res = append(res, bucket[i]...)
			k -= len(bucket[i])
		}
	}
	return res
}
