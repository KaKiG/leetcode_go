package main

import (
	"fmt"
	"strconv"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeKey struct {
	start, end int
}

func main() {
	ans := generateTrees(5)
	fmt.Println(len(ans))
	/*root := &TreeNode{1, nil, nil}
	PrintTree(root)
	PrintTree(AddLeaf(2, root))
	PrintTree(root)*/
	/*fmt.Println(CountNode(ans[0]))*/
}

func PrintTree(root *TreeNode) {
	fmt.Println(inorderTraversal(root))
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	TMap := make(map[TreeKey][]*TreeNode)
	return Solver(1, n, TMap)
}

func Solver(start, end int, TMap map[TreeKey][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		return append([]*TreeNode{}, &TreeNode{start, nil, nil})
	}
	if _, exist := TMap[TreeKey{start, end}]; exist {
		return TMap[TreeKey{start, end}]
	}

	ans := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		lSubT := Solver(start, i-1, TMap)
		rSubT := Solver(i+1, end, TMap)
		for m := range lSubT {
			for n := range rSubT {
				ans = append(ans, &TreeNode{i, lSubT[m], rSubT[n]})
			}
		}
	}

	TMap[TreeKey{start, end}] = ans
	return ans
}

/*func PartialSolver(n int) []*TreeNode {
	if n == 0 {
    return nil
  }
  if n == 1 {
    return append([]*TreeNode{}, &TreeNode{1,nil,nil})
  }
  ans := make([]*TreeNode, 0)
  for i := range PartialSolver(n-1) {
    ans = append(ans, AddLeaves(n, PartialSolver(n-1)[i]))
  }
  return ans
}

func AddLeaves(n int, root *TreeNode) []*TreeNode {
  ans := make([]*TreeNode, 0)
  cpRoot := CopyTree(root)
  if cpRoot == nil {
    return append([]*TreeNode{}, &TreeNode{n,nil,nil})
  }
  if cpRoot.Val < n {
    ans = append(ans, )
  }

}

func AddLeaf(n int, root *TreeNode) *TreeNode {
	cpRoot := CopyTree(root)
	if cpRoot == nil {
		return &TreeNode{n, nil, nil}
	}

	if n > cpRoot.Val {
		newRoot = &TreeNode{n, nil, cpRoot.Right}
		cpRoot.Right = newRoot
	} else {
		newRoot = &TreeNode{n, cpRoot.Left, nil}
		cpRoot.Left = newRoot
	}
	return cpRoot
}

func CopyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{root.Val, CopyTree(root.Left), CopyTree(root.Right)}
	return newRoot
}

func CountNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return CountNode(root.Left) + CountNode(root.Right) + 1
}*/

func inorderTraversal(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := make([]string, 0)
	PartialSolverT(root, &ans)
	return ans
}

func PartialSolverT(root *TreeNode, ans *[]string) {
	if root.Left != nil {
		PartialSolverT(root.Left, ans)
	} else {
		*ans = append(*ans, "Left")
	}
	*ans = append(*ans, strconv.Itoa(root.Val))
	if root.Right != nil {
		PartialSolverT(root.Right, ans)
	} else {
		*ans = append(*ans, "Right")
	}
}
