// Source : https://leetcode.com/problems/task-scheduler
// Author : BradleyZhang
// Date   : 2026-10-02

/*****************************************************************************************************
 *
 * You are given an array of CPU tasks, each labeled with a letter from A to Z, and a number n. Each
 * CPU interval can be idle or allow the completion of one task. Tasks can be completed in any order,
 * but there's a constraint: there has to be a gap of at least n intervals between two tasks with the
 * same label.
 *
 * Return the minimum number of CPU intervals required to complete all tasks.
 *
 * Example 1:
 * Input: tasks = ["A","A","A","B","B","B"], n = 2
 *
 * Output: 8
 *
 * Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
 *
 * After completing task A, you must wait two intervals before doing A again. The same applies to task
 * B. In the 3^rd interval, neither A nor B can be done, so you idle. By the 4^th interval, you can do
 * A again as 2 intervals have passed.
 *
 * Example 2:
 *
 * Input: tasks = ["A","C","A","B","D","B"], n = 1
 *
 * Output: 6
 *
 * Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
 *
 * With a cooling interval of 1, you can repeat a task after just one other task.
 *
 * Example 3:
 *
 * Input: tasks = ["A","A","A", "B","B","B"], n = 3
 *
 * Output: 10
 *
 * Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
 *
 * There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads
 * to idling twice between repetitions of these tasks.
 *
 * Constraints:
 *
 * 	1 <= tasks.length <= 10^4
 * 	tasks[i] is an uppercase English letter.
 * 	0 <= n <= 100
 ******************************************************************************************************/

package taskscheduler

import "container/heap"

func leastInterval(tasks []byte, n int) int {
	var count [26]int

	for _, task := range tasks {
		count[task-'A']++
	}

	maxFreq := 0
	maxKinds := 0

	for _, freq := range count {
		if freq > maxFreq {
			maxFreq = freq
			maxKinds = 1
		} else if freq == maxFreq && freq > 0 {
			maxKinds++
		}
	}

	// 最高频 task 之间必须隔 n 个位置
	result := (maxFreq-1)*(n+1) + maxKinds

	if result < len(tasks) {
		return len(tasks)
	}

	return result
}

// 用最大堆
func leastIntervalHeap(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}
	h := &MaxHeap{}
	for _, cnt := range count {
		if cnt > 0 {
			heap.Push(h, cnt)
		}
	}

	time := 0
	for h.Len() > 0 {
		used := 0
		var temp []int

		// 一个 cycle 最多 n+1 个任务
		for used <= n && h.Len() > 0 {
			cnt := heap.Pop(h).(int)
			cnt--

			if cnt > 0 {
				temp = append(temp, cnt)
			}

			used++
			time++
		}

		for _, cnt := range temp {
			heap.Push(h, cnt)
		}

		// 还有任务，说明剩余位置必须 idle
		if h.Len() > 0 {
			time += n + 1 - used
		}
	}

	return time
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
