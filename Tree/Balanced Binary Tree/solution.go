/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if depthTree(root) == -1 {
		return false
	}

	return true
}

func depthTree(root *TreeNode) int {
	var result int

	if root == nil {
		return 0
	}

	leftNodes := depthTree(root.Left)
	if leftNodes == -1 {
		return -1
	}
	rightNodes := depthTree(root.Right)
	if rightNodes == -1 {
		return -1
	}
	result = max(leftNodes, rightNodes) + 1

	if leftNodes-rightNodes > 1 || rightNodes-leftNodes > 1 {
		return -1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}