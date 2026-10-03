// Source : https://leetcode.com/problems/kth-largest-element-in-a-stream
// Author : BradleyZhang
// Date   : 2026-10-03

/*****************************************************************************************************
 *
 * You are part of a university admissions office and need to keep track of the kth highest test score
 * from applicants in real-time. This helps to determine cut-off marks for interviews and admissions
 * dynamically as new applicants submit their scores.
 *
 * You are tasked to implement a class which, for a given integer k, maintains a stream of test scores
 * and continuously returns the kth highest test score after a new score has been submitted. More
 * specifically, we are looking for the kth highest score in the sorted list of all scores.
 *
 * Implement the KthLargest class:
 *
 * 	KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of
 * test scores nums.
 * 	int add(int val) Adds a new test score val to the stream and returns the element
 * representing the k^th largest element in the pool of test scores so far.
 *
 * Example 1:
 *
 * Input:
 * ["KthLargest", "add", "add", "add", "add", "add"]
 * [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
 *
 * Output: [null, 4, 5, 5, 8, 8]
 *
 * Explanation:
 *
 * KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
 * kthLargest.add(3); // return 4
 * kthLargest.add(5); // return 5
 * kthLargest.add(10); // return 5
 * kthLargest.add(9); // return 8
 * kthLargest.add(4); // return 8
 *
 * Example 2:
 *
 * Input:
 * ["KthLargest", "add", "add", "add", "add"]
 * [[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]
 *
 * Output: [null, 7, 7, 7, 8]
 *
 * Explanation:
 * KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
 * kthLargest.add(2); // return 7
 * kthLargest.add(10); // return 7
 * kthLargest.add(9); // return 7
 * kthLargest.add(9); // return 8
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 10^4
 * 	1 <= k <= nums.length + 1
 * 	-10^4 <= nums[i] <= 10^4
 * 	-10^4 <= val <= 10^4
 * 	At most 10^4 calls will be made to add.
 ******************************************************************************************************/
package kthlargestelementinastream

import "container/heap"

type KthLargest struct {
	heap MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := MinHeap{}
	heap.Init(&minHeap)
	for _, n := range nums {
		if minHeap.Len() < k {
			heap.Push(&minHeap, n)
		} else if n > minHeap[0] {
			heap.Pop(&minHeap)
			heap.Push(&minHeap, n)
		}
	}
	return KthLargest{
		heap: minHeap,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.heap.Len() < this.k {
		heap.Push(&this.heap, val)
		return this.heap[0]
	}
	if val > this.heap[0] {
		heap.Pop(&this.heap)
		heap.Push(&this.heap, val)
	}
	return this.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[j] > h[i] // 小顶堆
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
