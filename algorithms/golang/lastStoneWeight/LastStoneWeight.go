// Source : https://leetcode.com/problems/last-stone-weight
// Author : BradleyZhang
// Date   : 2026-10-03

/*****************************************************************************************************
 *
 * You are given an array of integers stones where stones[i] is the weight of the i^th stone.
 *
 * We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash
 * them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this
 * smash is:
 *
 * 	If x == y, both stones are destroyed, and
 * 	If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y -
 * x.
 *
 * At the end of the game, there is at most one stone left.
 *
 * Return the weight of the last remaining stone. If there are no stones left, return 0.
 *
 * Example 1:
 *
 * Input: stones = [2,7,4,1,8,1]
 * Output: 1
 * Explanation:
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
 *
 * Example 2:
 *
 * Input: stones = [1]
 * Output: 1
 *
 * Constraints:
 *
 * 	1 <= stones.length <= 30
 * 	1 <= stones[i] <= 1000
 ******************************************************************************************************/

package laststoneweight

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	maxHeap := MaxHeap(stones)
	heap.Init(&maxHeap)
	for maxHeap.Len() > 1 {
		a := heap.Pop(&maxHeap).(int)
		b := heap.Pop(&maxHeap).(int)
		if a != b {
			x := max(a, b) - min(a, b)
			heap.Push(&maxHeap, x)
		}
	}
	if maxHeap.Len() > 0 {
		return maxHeap[0]
	}
	return 0
}

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[j] < h[i] // 大顶堆
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
