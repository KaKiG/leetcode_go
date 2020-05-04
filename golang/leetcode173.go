package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

// The key concept is how to find the next smallest.
// Here we using a stack of tree nodes.
// We push every node on its left for root.
// Whenever we find the next node, we pop the top one and do push the right node of the top one.
// If right node has some left children. Then new top one on the stack will be smallest node bigger than the one just pop up
// If right node does not has any left child, then itself it the smallest node bigger than the one just pop up
// If it does not have any right node, then the parent node would be the smallest node bigger than the one just pop up

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0)
	var iterator BSTIterator
	iterator.stack = stack
	iterator.push(root)

	return iterator
}

func (iterator *BSTIterator) push(root *TreeNode) {
	for root != nil {
		iterator.stack = append(iterator.stack, root)
		root = root.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	var nextNode *TreeNode
	length := len(this.stack)
	this.stack, nextNode = this.stack[:length-1], this.stack[length-1]
	this.push(nextNode.Right)

	return nextNode.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
