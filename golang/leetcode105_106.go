package main

import (
	"fmt"
	"strconv"
)

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(preorder, inorder)
	root2 := buildTree2(inorder, postorder)
	fmt.Println(inorderTraversal(root2))
	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}

	var i int
	for i = range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Left = buildTree2(inorder[:i], postorder[:i])
	root.Right = buildTree2(inorder[i+1:], postorder[i:len(postorder)-1])
	return root

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	var i int
	for i = range inorder {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:1+i], inorder[:i])
	root.Right = buildTree(preorder[1+i:], inorder[i+1:])
	return root
}

func inorderTraversal(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := make([]string, 0)
	PartialSolver1(root, &ans)
	return ans
}

func PartialSolver1(root *TreeNode, ans *[]string) {
	if root.Left != nil {
		PartialSolver1(root.Left, ans)
	} else {
		*ans = append(*ans, "L")
	}
	*ans = append(*ans, strconv.Itoa(root.Val))
	if root.Right != nil {
		PartialSolver1(root.Right, ans)
	} else {
		*ans = append(*ans, "R")
	}
}
