package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root, node1, node2, node4, node6 TreeNode
	root = TreeNode{3, &node1, &node2}
	node1 = TreeNode{9, nil, nil}
	node2 = TreeNode{20, &node4, &node6}
	//node3 = TreeNode{3, nil, nil}
	node4 = TreeNode{15, nil, nil}
	//node5 = TreeNode{4, nil, nil}
	node6 = TreeNode{7, nil, nil}
	fmt.Println(LOT(&root))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return solver(root.Left, root.Right)
}

func solver(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}

	if l.Val != r.Val {
		return false
	}

	return solver(l.Left, r.Right) && solver(l.Right, r.Left)
}

/*func isSymmetric(root *TreeNode) bool {
	qNodes := make([]*TreeNode, 0)
	ansInt := make([]int, 0)
	return LOT(root, &qNodes, &ansInt)
}

func isSymmetricArray(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}*/

func LOT(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	preqNodes := make([]*TreeNode, 0)
	currInt := make([]int, 0)
	currNodes := preqNodes

	preqNodes = append(preqNodes, root)
	currInt = append(currInt, root.Val)
	ans = append(ans, currInt)
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
			ans = append(ans, currInt)
			currInt = currInt[len(currInt):]
		}
	}
	return ans
}
