// Source : https://leetcode.com/problems/diameter-of-binary-tree
// Author : BradleyZhang
// Date   : 2026-09-30

/*****************************************************************************************************
 *
 * Given the root of a binary tree, return the length of the diameter of the tree.
 *
 * The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
 * This path may or may not pass through the root.
 *
 * The length of a path between two nodes is represented by the number of edges between them.
 *
 * Example 1:
 *
 * Input: root = [1,2,3,4,5]
 * Output: 3
 * Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 * Example 2:
 *
 * Input: root = [1,2]
 * Output: 1
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [1, 10^4].
 * 	-100 <= Node.val <= 100
 ******************************************************************************************************/
package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var maxD func(node *TreeNode) int
	maxD = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		var l, r int
		if node.Left != nil {
			l++
		}
		if node.Right != nil {
			r++
		}
		l += maxD(node.Left)
		r += maxD(node.Right)
		lrMax := max(l, r)
		maxAll := l + r
		ans = max(ans, maxAll)
		return lrMax
	}
	maxD(root)
	return ans
}
