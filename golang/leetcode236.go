package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// much simpler idea, using recursion
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return solver(root, p, q)
}

func solver(root, p, q *TreeNode) *TreeNode {
	//if one of which is the root, return immediately
	if root == nil || p == root || q == root {
		return root
	}

	//find the left and right possible LCA
	// if one of which is nil, it means that two of the node are in the same subtree
	// if both are in different subtrees, then l and r will not be nil
	// since we have the p==root || q==root, return root
	// in this case, return root
	l, r := solver(root.Left, p, q), solver(root.Right, p, q)

	if l == nil {
		return r
	} else if r == nil {
		return l
	}
	return root

}

// this uses the property of post and pre traversal
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	pre := preorderTraversal(root)
// 	post := postorderTraversal(root)
// 	qval := q.Val
// 	pval := p.Val
// 	rootMap := make(map[*TreeNode]bool)

// 	for i := range pre {
// 		rootMap[pre[i]] = true
// 		if pre[i].Val == qval || pre[i].Val == pval {
// 			break
// 		}
// 	}

// 	i := 0
// 	for i = len(post) - 1; i >= 0; i-- {
// 		if post[i].Val == qval || post[i].Val == pval {
// 			break
// 		}
// 	}

// 	for i < len(post) {
// 		if rootMap[post[i]] {
// 			return post[i]
// 		}
// 		i++
// 	}
// 	return nil
// }

// func preorderTraversal(root *TreeNode) []*TreeNode {
// 	ans := make([]*TreeNode, 0)
// 	preOrderSolver(root, &ans)
// 	return ans
// }

// func preOrderSolver(root *TreeNode, ans *[]*TreeNode) {
// 	if root != nil {
// 		*ans = append(*ans, root)
// 		preOrderSolver(root.Left, ans)
// 		preOrderSolver(root.Right, ans)
// 	}
// }

// func postorderTraversal(root *TreeNode) []*TreeNode {
// 	ans := make([]*TreeNode, 0)
// 	postOrderSolver(root, &ans)
// 	return ans
// }

// func postOrderSolver(root *TreeNode, ans *[]*TreeNode) {
// 	if root != nil {
// 		postOrderSolver(root.Left, ans)
// 		postOrderSolver(root.Right, ans)
// 		*ans = append(*ans, root)
// 	}
// }
