/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}

	leftNodes := postorderTraversal(root.Left)
	rightNodes := postorderTraversal(root.Right)
	results = append(results, leftNodes...)
	results = append(results, rightNodes...)
	results = append(results, root.Val)

	return results
}