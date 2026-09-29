// Source : https://leetcode.com/problems/maximum-depth-of-binary-tree
// Author : BradleyZhang
// Date   : 2026-09-29

/*****************************************************************************************************
 *
 * Given the root of a binary tree, return its maximum depth.
 *
 * A binary tree's maximum depth is the number of nodes along the longest path from the root node down
 * to the farthest leaf node.
 *
 * Example 1:
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: 3
 *
 * Example 2:
 *
 * Input: root = [1,null,2]
 * Output: 2
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 10^4].
 * 	-100 <= Node.val <= 100
 ******************************************************************************************************/

package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return do(root)
}
func do(node *TreeNode) (depth int) {
	if node == nil {
		return 0
	}
	depth = 1
	left := do(node.Left)
	right := do(node.Right)
	depth += max(left, right)
	return depth
}
