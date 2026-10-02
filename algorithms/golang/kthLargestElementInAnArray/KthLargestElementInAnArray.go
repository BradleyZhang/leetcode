// Source : https://leetcode.com/problems/kth-largest-element-in-an-array
// Author : BradleyZhang
// Date   : 2026-10-02

/*****************************************************************************************************
 *
 * Given an integer array nums and an integer k, return the k^th largest element in the array.
 *
 * Note that it is the k^th largest element in the sorted order, not the k^th distinct element.
 *
 * Can you solve it without sorting?
 *
 * Example 1:
 * Input: nums = [3,2,1,5,6,4], k = 2
 * Output: 5
 * Example 2:
 * Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
 * Output: 4
 *
 * Constraints:
 *
 * 	1 <= k <= nums.length <= 10^5
 * 	-10^4 <= nums[i] <= 10^4
 ******************************************************************************************************/
package kthlargestelementinanarray

import (
	"container/heap"
)

// heap 方法的优化
func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums[:k])
	heap.Init(&h)

	for _, num := range nums[k:] {
		if num > h[0] {
			heap.Pop(&h)
			heap.Push(&h, num)
		}
	}

	return h[0]
}

func findKthLargestBackup(nums []int, k int) int {
	maxheap := &IntHeap{}
	heap.Init(maxheap)
	for _, n := range nums {
		heap.Push(maxheap, n)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(maxheap)
	}
	return heap.Pop(maxheap).(int)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
