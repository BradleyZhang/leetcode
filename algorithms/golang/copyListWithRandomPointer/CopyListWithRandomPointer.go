// Source : https://leetcode.com/problems/copy-list-with-random-pointer
// Author : BradleyZhang
// Date   : 2026-09-26

/*****************************************************************************************************
 *
 * A linked list of length n is given such that each node contains an additional random pointer, which
 * could point to any node in the list, or null.
 *
 * Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where
 * each new node has its value set to the value of its corresponding original node. Both the next and
 * random pointer of the new nodes should point to new nodes in the copied list such that the pointers
 * in the original list and copied list represent the same list state. None of the pointers in the new
 * list should point to nodes in the original list.
 *
 * For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for
 * the corresponding two nodes x and y in the copied list, x.random --> y.
 *
 * Return the head of the copied linked list.
 *
 * The linked list is represented in the input/output as a list of n nodes. Each node is represented
 * as a pair of [val, random_index] where:
 *
 * 	val: an integer representing Node.val
 * 	random_index: the index of the node (range from 0 to n-1) that the random pointer points
 * to, or null if it does not point to any node.
 *
 * Your code will only be given the head of the original linked list.
 *
 * Example 1:
 *
 * Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
 * Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
 *
 * Example 2:
 *
 * Input: head = [[1,1],[2,1]]
 * Output: [[1,1],[2,1]]
 *
 * Example 3:
 *
 * Input: head = [[3,null],[3,0],[3,null]]
 * Output: [[3,null],[3,0],[3,null]]
 *
 * Constraints:
 *
 * 	0 <= n <= 1000
 * 	-10^4 <= Node.val <= 10^4
 * 	Node.random is null or is pointing to some node in the linked list.
 ******************************************************************************************************/

package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 更简洁
func copyRandomListB(head *Node) *Node {
	alreadyExist := make(map[*Node]*Node)

	var cpNode func(*Node) *Node
	cpNode = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if node, ok := alreadyExist[n]; ok {
			return node
		}

		newNode := &Node{
			Val: n.Val,
		}
		alreadyExist[n] = newNode

		newNode.Next = cpNode(n.Next)
		newNode.Random = cpNode(n.Random)

		return newNode
	}

	return cpNode(head)
}

func copyRandomListBackup(head *Node) *Node {

	dummy := &Node{}
	alreadyExist := make(map[*Node]*Node, 10)
	pre := dummy
	var cpNode func(n *Node) *Node
	cpNode = func(n *Node) *Node {
		if node, ok := alreadyExist[n]; ok {
			return node
		}
		new := &Node{
			Val: n.Val,
		}
		alreadyExist[n] = new
		if n.Random != nil {
			if node, ok := alreadyExist[n.Random]; ok {
				new.Random = node
			} else {
				new.Random = cpNode(n.Random)
			}
		}
		return new
	}

	for head != nil {
		new := cpNode(head)
		pre.Next = new
		pre = new
		head = head.Next
	}
	return dummy.Next
}
