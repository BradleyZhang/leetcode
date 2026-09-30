// Source : https://leetcode.com/problems/serialize-and-deserialize-binary-tree
// Author : BradleyZhang
// Date   : 2026-09-29

/*****************************************************************************************************
 *
 * Serialization is the process of converting a data structure or object into a sequence of bits so
 * that it can be stored in a file or memory buffer, or transmitted across a network connection link
 * to be reconstructed later in the same or another computer environment.
 *
 * Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your
 * serialization/deserialization algorithm should work. You just need to ensure that a binary tree can
 * be serialized to a string and this string can be deserialized to the original tree structure.
 *
 * Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do
 * not necessarily need to follow this format, so please be creative and come up with different
 * approaches yourself.
 *
 * Example 1:
 *
 * Input: root = [1,2,3,null,null,4,5]
 * Output: [1,2,3,null,null,4,5]
 *
 * Example 2:
 *
 * Input: root = []
 * Output: []
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 10^4].
 * 	-1000 <= Node.val <= 1000
 ******************************************************************************************************/

package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

// dfs
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("#,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val) + ",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if idx >= len(vals) || vals[idx] == "#" || vals[idx] == "" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(vals[idx])
		idx++
		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}

// backup

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

var (
	sep = "|"
)

// Serializes a tree to a single string.
func (this *Codec) serializeBackup(root *TreeNode) string {
	if root == nil {
		return ""
	}
	builder := strings.Builder{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr != nil {
			builder.WriteString(strconv.Itoa(curr.Val))
			builder.WriteString(sep)
		} else {
			builder.WriteString(sep)
		}

		if curr != nil {
			queue = append(queue, curr.Left, curr.Right)
		}
	}
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeBackup(data string) *TreeNode {
	nodes := strings.Split(data, sep)
	if len(nodes) < 1 {
		return nil
	}
	root := newNode(nodes[0])
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]

		l := newNode(nodes[i])
		r := newNode(nodes[i+1])
		i += 2
		curr.Left = l
		curr.Right = r

		if l != nil {
			queue = append(queue, l)
		}
		if r != nil {
			queue = append(queue, r)
		}
	}
	return root

}
func newNode(data string) *TreeNode {
	var node *TreeNode
	if data != "" {
		val, _ := strconv.Atoi(data)
		node = &TreeNode{
			Val: val,
		}
	}
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
