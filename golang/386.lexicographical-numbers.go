/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
// type lexiArr []int

// func lexicalOrder(n int) []int {
// 	res := make([]int, n)
// 	for i := range res {
// 		res[i] = i + 1
// 	}
// 	sort.Sort(lexiArr(res))
// 	return res
// }

// func (nums lexiArr) Len() int {
// 	return len(nums)
// }

// func (nums lexiArr) Swap(i, j int) {
// 	nums[i], nums[j] = nums[j], nums[i]
// }

// func (nums lexiArr) Less(i, j int) bool {
// 	s1 := strconv.Itoa(nums[i])
// 	s2 := strconv.Itoa(nums[j])
// 	ind1 := 0
// 	ind2 := 0
// 	for ind1 < len(s1) && ind2 < len(s2) {
// 		if s1[ind1] < s2[ind2] {
// 			return true
// 		}
// 		if s1[ind1] > s2[ind2] {
// 			return false
// 		}
// 		ind1++
// 		ind2++
// 	}

// 	if ind1 != len(s1) && ind2 == len(s2) {
// 		return false
// 	}
// 	return true
// }

// dfs solution
func lexicalOrder(n int) []int {
	ans := make([]int, n)
	pos := 0
	solver(n, 1, &pos, &ans)
	return ans
}

func solver(max int, num int, pos *int, ans *[]int) {
	if num > max {
		return
	}

	for i := 0; i < 10; i++ {
		if i != 0 && (num+i)%10 == 0 {
			continue
		}
		if num+i <= max {
			(*ans)[*pos] = num + i
			*pos += 1
			solver(max, (num+i)*10, pos, ans)
		}
	}
}

// @lc code=end
