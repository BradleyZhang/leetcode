// Source : https://leetcode.com/problems/reverse-nodes-in-k-group
// Author : BradleyZhang
// Date   : 2026-09-27

/*****************************************************************************************************
 *
 * Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified
 * list.
 *
 * k is a positive integer and is less than or equal to the length of the linked list. If the number
 * of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
 *
 * You may not alter the values in the list's nodes, only nodes themselves may be changed.
 *
 * Example 1:
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [2,1,4,3,5]
 *
 * Example 2:
 *
 * Input: head = [1,2,3,4,5], k = 3
 * Output: [3,2,1,4,5]
 *
 * Constraints:
 *
 * 	The number of nodes in the list is n.
 * 	1 <= k <= n <= 5000
 * 	0 <= Node.val <= 1000
 *
 * Follow-up: Can you solve the problem in O(1) extra memory space?
 ******************************************************************************************************/

package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy

	for {
		// 找到这一组的第 k 个节点
		kth := groupPrev
		for i := 0; i < k; i++ {
			kth = kth.Next
			if kth == nil {
				return dummy.Next
			}
		}

		groupNext := kth.Next

		// 反转 [groupPrev.Next, kth]
		prev := groupNext
		curr := groupPrev.Next

		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// 接回上一组
		oldGroupStart := groupPrev.Next
		groupPrev.Next = kth

		// oldGroupStart 已经变成这一组最后一个节点
		groupPrev = oldGroupStart
	}
}
