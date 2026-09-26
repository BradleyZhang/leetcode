// Source : https://leetcode.com/problems/merge-k-sorted-lists
// Author : BradleyZhang
// Date   : 2026-09-24

/*****************************************************************************************************
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
 *
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 * Example 1:
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 *   1->4->5,
 *   1->3->4,
 *   2->6
 * ]
 * merging them into one sorted linked list:
 * 1->1->2->3->4->4->5->6
 *
 * Example 2:
 *
 * Input: lists = []
 * Output: []
 *
 * Example 3:
 *
 * Input: lists = [[]]
 * Output: []
 *
 * Constraints:
 *
 * 	k == lists.length
 * 	0 <= k <= 10^4
 * 	0 <= lists[i].length <= 500
 * 	-10^4 <= lists[i][j] <= 10^4
 * 	lists[i] is sorted in ascending order.
 * 	The sum of lists[i].length will not exceed 10^4.
 ******************************************************************************************************/
package mergeksortedlists

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	minHeap := &lHeap{}
	heap.Init(minHeap)
	for _, node := range lists {
		if node != nil {
			heap.Push(minHeap, node)
		}
	}
	for minHeap.Len() > 0 {
		curr.Next = heap.Pop(minHeap).(*ListNode)
		curr = curr.Next
		if curr.Next != nil {
			heap.Push(minHeap, curr.Next)
		}
	}
	return dummy.Next
}

type lHeap []any

func (h *lHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作为参数
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(*ListNode))
}

func (h *lHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *lHeap) Len() int {
	return len(*h)
}

func (h *lHeap) Less(i, j int) bool {
	return (*h)[i].(*ListNode).Val < (*h)[j].(*ListNode).Val
}

func (h *lHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *lHeap) Top() any {
	return (*h)[0]
}

func mergeKListsBackup(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	i2node := make(map[int]*ListNode)
	for i, l := range lists {
		if l != nil {
			i2node[i] = l
		}
	}
	curr := dummy
	for len(i2node) > 0 {
		var minI int
		for i, _ := range i2node {
			minI = i
			break
		}
		for i, node := range i2node {
			if node.Val < i2node[minI].Val {
				minI = i
			}
		}
		curr.Next = i2node[minI]
		curr = curr.Next

		i2node[minI] = i2node[minI].Next
		if i2node[minI] == nil {
			delete(i2node, minI)
		}
	}
	return dummy.Next
}
