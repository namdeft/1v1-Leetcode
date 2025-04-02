/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}

	leftNodes := preorderTraversal(root.Left)
	rightNodes := preorderTraversal(root.Right)

	results = append(results, root.Val)
	results = append(results, leftNodes...)
	results = append(results, rightNodes...)

	return results
}