// Source : https://leetcode.com/problems/k-closest-points-to-origin
// author : bradleyzhang
// date   : 2026-10-02

/*****************************************************************************************************
 *
 * given an array of points where points[i] = [xi, yi] represents a point on the x-y plane and an
 * integer k, return the k closest points to the origin (0, 0).
 *
 * the distance between two points on the x-y plane is the euclidean distance (i.e., &radic;(x1 -
 * x2)^2 + (y1 - y2)^2).
 *
 * you may return the answer in any order. the answer is guaranteed to be unique (except for the order
 * that it is in).
 *
 * example 1:
 *
 * input: points = [[1,3],[-2,2]], k = 1
 * output: [[-2,2]]
 * explanation:
 * the distance between (1, 3) and the origin is sqrt(10).
 * the distance between (-2, 2) and the origin is sqrt(8).
 * since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
 * we only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
 *
 * example 2:
 *
 * input: points = [[3,3],[5,-1],[-2,4]], k = 2
 * output: [[3,3],[-2,4]]
 * explanation: the answer [[-2,4],[3,3]] would also be accepted.
 *
 * constraints:
 *
 * 	1 <= k <= points.length <= 10^4
 * 	-10^4 <= xi, yi <= 10^4
 ******************************************************************************************************/

package kclosestpointstoorigin

import (
	"container/heap"
	"math"
)

func kClosest(points [][]int, k int) [][]int {
	maxHeap := MaxPointHeap(points[:k]) // push k 个元素
	heap.Init(&maxHeap)
	for _, p := range points[k:] {
		if less(p, maxHeap[0]) { // 堆中只维护k个最小（最近）元素
			heap.Pop(&maxHeap)
			heap.Push(&maxHeap, p)
		}
	}
	return maxHeap
}

type MaxPointHeap [][]int

func (h MaxPointHeap) Len() int { return len(h) }
func less(p, q []int) bool {
	d1 := math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))
	d2 := math.Sqrt(float64(q[0]*q[0] + q[1]*q[1]))
	return d1 < d2
}
func (h MaxPointHeap) Less(i, j int) bool {
	return less(h[j], h[i]) // 大顶堆
}
func (h MaxPointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxPointHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MaxPointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
