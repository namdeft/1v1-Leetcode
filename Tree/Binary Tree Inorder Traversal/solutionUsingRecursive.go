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
	stack := []*TreeNode{}
	currentNode := root
	for currentNode != nil || len(stack) > 0 {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}

		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		results = append(results, currentNode.Val)

		currentNode = currentNode.Right
	}

	return results
}