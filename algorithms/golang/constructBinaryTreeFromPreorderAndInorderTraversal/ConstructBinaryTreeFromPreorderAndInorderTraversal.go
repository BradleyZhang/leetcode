// Source : https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// Author : BradleyZhang
// Date   : 2026-10-01

/*****************************************************************************************************
 *
 * Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary
 * tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
 *
 * Example 1:
 *
 * Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * Output: [3,9,20,null,null,15,7]
 *
 * Example 2:
 *
 * Input: preorder = [-1], inorder = [-1]
 * Output: [-1]
 *
 * Constraints:
 *
 * 	1 <= preorder.length <= 3000
 * 	inorder.length == preorder.length
 * 	-3000 <= preorder[i], inorder[i] <= 3000
 * 	preorder and inorder consist of unique values.
 * 	Each value of inorder also appears in preorder.
 * 	preorder is guaranteed to be the preorder traversal of the tree.
 * 	inorder is guaranteed to be the inorder traversal of the tree.
 ******************************************************************************************************/

package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// inorder 中每个值对应的下标
	indexMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		indexMap[v] = i
	}

	preIndex := 0

	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// preorder 的当前元素就是当前子树的根
		rootVal := preorder[preIndex]
		preIndex++

		root := &TreeNode{Val: rootVal}

		// 根节点在 inorder 中的位置
		mid := indexMap[rootVal]

		root.Left = build(left, mid-1)
		root.Right = build(mid+1, right)

		return root
	}

	return build(0, len(inorder)-1)
}
