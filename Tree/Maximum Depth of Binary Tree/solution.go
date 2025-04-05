/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}

	leftNodes := maxDepth(root.Left)
	rightNodes := maxDepth(root.Right)
	result = max(leftNodes, rightNodes) + 1

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}