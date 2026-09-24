// Source : https://leetcode.com/problems/reorder-list
// Author : BradleyZhang
// Date   : 2026-09-24

/*****************************************************************************************************
 *
 * You are given the head of a singly linked-list. The list can be represented as:
 *
 * L0 &rarr; L1 &rarr; &hellip; &rarr; Ln - 1 &rarr; Ln
 *
 * Reorder the list to be on the following form:
 *
 * L0 &rarr; Ln &rarr; L1 &rarr; Ln - 1 &rarr; L2 &rarr; Ln - 2 &rarr; &hellip;
 *
 * You may not modify the values in the list's nodes. Only nodes themselves may be changed.
 *
 * Example 1:
 *
 * Input: head = [1,2,3,4]
 * Output: [1,4,2,3]
 *
 * Example 2:
 *
 * Input: head = [1,2,3,4,5]
 * Output: [1,5,2,4,3]
 *
 * Constraints:
 *
 * 	The number of nodes in the list is in the range [1, 5 * 10^4].
 * 	1 <= Node.val <= 1000
 ******************************************************************************************************/
package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	second := reverseList(slow.Next)
	slow.Next = nil
	for second != nil {
		secNext := second.Next
		second.Next = head.Next
		head.Next = second

		head = second.Next
		second = secNext
	}
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 暴力求解
func reorderListBackup(head *ListNode) {
	curr := head
	for curr != nil && curr.Next != nil {
		right := curr
		prev := right
		for right.Next != nil {
			prev = right
			right = right.Next
		}
		prev.Next = nil
		right.Next = curr.Next
		curr.Next = right
		curr = right.Next
	}
}
