/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}

	leftNodes := inorderTraversal(root.Left)
	rightNodes := inorderTraversal(root.Right)
	results = append(results, leftNodes...)
	results = append(results, root.Val)
	results = append(results, rightNodes...)

	return results
}