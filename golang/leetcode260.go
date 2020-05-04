package main

func main() {

}

func singleNumber(nums []int) []int {
	repDic := make(map[int]int)
	res := make([]int, 0)
	for i := range nums {
		repDic[nums[i]]++
	}
	for k, v := range repDic {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

//Constant space solution
// wait
func singleNumber(nums []int) []int {

}
