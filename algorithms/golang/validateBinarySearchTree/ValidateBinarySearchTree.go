// Source : https://leetcode.com/problems/validate-binary-search-tree
// Author : BradleyZhang
// Date   : 2026-09-30

/*****************************************************************************************************
 *
 * Given the root of a binary tree, determine if it is a valid binary search tree (BST).
 *
 * A valid BST is defined as follows:
 *
 * 	The left subtree of a node contains only nodes with keys strictly less than the node's key.
 * 	The right subtree of a node contains only nodes with keys strictly greater than the node's
 * key.
 * 	Both the left and right subtrees must also be binary search trees.
 *
 * Example 1:
 *
 * Input: root = [2,1,3]
 * Output: true
 *
 * Example 2:
 *
 * Input: root = [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [1, 10^4].
 * 	-2^31 <= Node.val <= 2^31 - 1
 ******************************************************************************************************/
package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用min max传递范围
func isValidBST(root *TreeNode) bool {
	return valid(root, nil, nil)
}

func valid(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= *min {
		return false
	}

	if max != nil && node.Val >= *max {
		return false
	}

	return valid(node.Left, min, &node.Val) &&
		valid(node.Right, &node.Val, max)
}

// backup (unuseful)
func isValidBSTBackup(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}
	if !leftOK(root.Val, root.Left) {
		return false
	}
	if !rightOk(root.Val, root.Right) {
		return false
	}
	return true
}
func leftOK(parVal int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil {
		if l := node.Left.Val; l >= node.Val || l >= parVal {
			return false
		}
		if !leftOK(node.Val, node.Left) {
			return false
		}
	}
	if node.Right != nil {
		if r := node.Right.Val; r <= node.Val || r >= parVal {
			return false
		}
		if !rightOk(node.Val, node.Right) {
			return false
		}
	}
	return true
}
func rightOk(parVal int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil {
		if l := node.Left.Val; l >= node.Val || l <= parVal {
			return false
		}
		if !leftOK(node.Val, node.Left) {
			return false
		}
	}
	if node.Right != nil {
		if r := node.Right.Val; r <= node.Val || r <= parVal {
			return false
		}
		if !rightOk(node.Val, node.Right) {
			return false
		}
	}
	return true
}
