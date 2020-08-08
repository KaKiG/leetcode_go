package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// borrow the idea from level order traversal
// we extract the rightmost value of each level
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	preqNodes := make([]*TreeNode, 0)
	currInt := make([]int, 0)
	currNodes := preqNodes

	preqNodes = append(preqNodes, root)
	currInt = append(currInt, root.Val)
	ans = append(ans, currInt...)
	currInt = currInt[len(currInt):]
	for !(len(preqNodes) == 0 && len(currNodes) == 0) {

		if preqNodes[0].Left != nil {
			currNodes = append(currNodes, preqNodes[0].Left)
			currInt = append(currInt, preqNodes[0].Left.Val)
		}
		if preqNodes[0].Right != nil {
			currNodes = append(currNodes, preqNodes[0].Right)
			currInt = append(currInt, preqNodes[0].Right.Val)
		}
		preqNodes = preqNodes[1:]
		if len(preqNodes) == 0 && len(currNodes) != 0 {
			preqNodes = currNodes
			currNodes = currNodes[len(currNodes):]
			ans = append(ans, currInt[len(currInt)-1])
			currInt = currInt[len(currInt):]
		}
	}
	return ans
}
