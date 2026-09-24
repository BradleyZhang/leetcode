// Source : https://leetcode.com/problems/merge-two-sorted-lists
// Author : BradleyZhang
// Date   : 2026-09-24

/*****************************************************************************************************
 *
 * You are given the heads of two sorted linked lists list1 and list2.
 *
 * Merge the two lists into one sorted list. The list should be made by splicing together the nodes of
 * the first two lists.
 *
 * Return the head of the merged linked list.
 *
 * Example 1:
 *
 * Input: list1 = [1,2,4], list2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 *
 * Example 2:
 *
 * Input: list1 = [], list2 = []
 * Output: []
 *
 * Example 3:
 *
 * Input: list1 = [], list2 = [0]
 * Output: [0]
 *
 * Constraints:
 *
 * 	The number of nodes in both lists is in the range [0, 50].
 * 	-100 <= Node.val <= 100
 * 	Both list1 and list2 are sorted in non-decreasing order.
 ******************************************************************************************************/
package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用 dummy 优化
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	curr1, curr2 := list1, list2
	curr := dummy
	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			curr.Next = curr1
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr2 = curr2.Next
		}
		curr = curr.Next
	}
	if curr1 == nil {
		curr.Next = curr2
	} else {
		curr.Next = curr1
	}

	return dummy.Next
}
func mergeTwoListsBackup(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var ans *ListNode
	curr1, curr2 := list1, list2
	if list1.Val <= list2.Val {
		ans = list1
		curr1 = list1.Next
	} else {
		ans = list2
		curr2 = list2.Next
	}
	curr := ans
	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			curr.Next = curr1
			curr = curr.Next
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr = curr.Next
			curr2 = curr2.Next
		}
	}
	if curr1 == nil {
		curr.Next = curr2
	} else {
		curr.Next = curr1
	}

	return ans
}
