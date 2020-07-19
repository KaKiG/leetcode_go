/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	num := 0
	distDic := make(map[int]int)
	for i := range points {
		for j := range points {
			if j == i {
				continue
			}
			distDic[dist(points[i], points[j])]++
		}
		for _, v := range distDic {
			num += v * (v - 1)
		}
		for k := range distDic {
			delete(distDic, k)
		}
	}
	return num
}

func dist(pt1 []int, pt2 []int) int {
	return (pt1[0]-pt2[0])*(pt1[0]-pt2[0]) + (pt1[1]-pt2[1])*(pt1[1]-pt2[1])
}

// @lc code=end
