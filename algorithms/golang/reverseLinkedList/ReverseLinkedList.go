// Source : https://leetcode.com/problems/reverse-linked-list
// Author : BradleyZhang
// Date   : 2026-09-23

/*****************************************************************************************************
 *
 * Given the head of a singly linked list, reverse the list, and return the reversed list.
 *
 * Example 1:
 *
 * Input: head = [1,2,3,4,5]
 * Output: [5,4,3,2,1]
 *
 * Example 2:
 *
 * Input: head = [1,2]
 * Output: [2,1]
 *
 * Example 3:
 *
 * Input: head = []
 * Output: []
 *
 * Constraints:
 *
 * 	The number of nodes in the list is the range [0, 5000].
 * 	-5000 <= Node.val <= 5000
 *
 * Follow up: A linked list can be reversed either iteratively or recursively. Could you implement
 * both?
 ******************************************************************************************************/

package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
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
