// Source : https://leetcode.com/problems/binary-tree-level-order-traversal
// Author : BradleyZhang
// Date   : 2026-09-29

/*****************************************************************************************************
 *
 * Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from
 * left to right, level by level).
 *
 * Example 1:
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: [[3],[9,20],[15,7]]
 *
 * Example 2:
 *
 * Input: root = [1]
 * Output: [[1]]
 *
 * Example 3:
 *
 * Input: root = []
 * Output: []
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 2000].
 * 	-1000 <= Node.val <= 1000
 ******************************************************************************************************/
package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, 0, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, level)
	}

	return ans
}

func levelOrderBackup(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		level := []int{}
		for len(queue) > 0 {
			var node *TreeNode
			node, queue = pop(queue)
			level = append(level, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		ans = append(ans, level)
		queue = nextLevel
	}
	return ans
}

func pop(q []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(q) <= 0 {
		return &TreeNode{}, q
	}
	ans := q[0]
	q = q[1:]
	return ans, q
}
