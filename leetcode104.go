package main

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	preqNodes := make([]*TreeNode, 0)
	currNodes := preqNodes
	level := 1

	preqNodes = append(preqNodes, root)
	for !(len(preqNodes) == 0 && len(currNodes) == 0) {

		if preqNodes[0].Left != nil {
			currNodes = append(currNodes, preqNodes[0].Left)
		}
		if preqNodes[0].Right != nil {
			currNodes = append(currNodes, preqNodes[0].Right)
		}
		preqNodes = preqNodes[1:]
		if len(preqNodes) == 0 && len(currNodes) != 0 {
			preqNodes = currNodes
			currNodes = currNodes[len(currNodes):]
			level++
		}
	}
	return level
}
