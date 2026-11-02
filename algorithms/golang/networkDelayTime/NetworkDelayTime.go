// Source : https://leetcode.com/problems/network-delay-time
// Author : BradleyZhang
// Date   : 2026-10-28

/*****************************************************************************************************
 *
 * You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel
 * times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target
 * node, and wi is the time it takes for a signal to travel from source to target.
 *
 * We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to
 * receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.
 *
 * Example 1:
 *
 * Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
 * Output: 2
 *
 * Example 2:
 *
 * Input: times = [[1,2,1]], n = 2, k = 1
 * Output: 1
 *
 * Example 3:
 *
 * Input: times = [[1,2,1]], n = 2, k = 2
 * Output: -1
 *
 * Constraints:
 *
 * 	1 <= k <= n <= 100
 * 	1 <= times.length <= 6000
 * 	times[i].length == 3
 * 	1 <= ui, vi <= n
 * 	ui != vi
 * 	0 <= wi <= 100
 * 	All the pairs (ui, vi) are unique. (i.e., no multiple edges.)
 ******************************************************************************************************/
package networkdelaytime

import (
	"container/heap"
	"math"
	"slices"
)

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][][2]int, n)
	for _, list := range times {
		adj[list[0]-1] = append(adj[list[0]-1], [2]int{list[1] - 1, list[2]})
	}

	dist := dijkstra(n, adj, k-1)
	res := slices.Max(dist)
	if res == math.MaxInt {
		return -1
	}
	return res
}

func dijkstra(n int, adj [][][2]int, start int) []int {
	dist := make([]int, n)
	for i := range dist { // 初始化 distance
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	minHeap := &PQ{{0, start}}
	for minHeap.Len() > 0 {
		it := heap.Pop(minHeap).(Item)
		if it.dist > dist[it.node] {
			continue // 跳过过时 Item
		}
		for _, e := range adj[it.node] {
			v, w := e[0], e[1]
			if it.dist+w < dist[v] {
				dist[v] = it.dist + w
				heap.Push(minHeap, Item{dist[v], v})
			}
		}
	}
	return dist
}

type Item struct{ dist, node int }
type PQ []Item

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *PQ) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
