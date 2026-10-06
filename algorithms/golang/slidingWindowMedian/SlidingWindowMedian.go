// Source : https://leetcode.com/problems/sliding-window-median
// Author : BradleyZhang
// Date   : 2026-10-04

/*****************************************************************************************************
 *
 * The median is the middle value in an ordered integer list. If the size of the list is even, there
 * is no middle value. So the median is the mean of the two middle values.
 *
 * 	For examples, if arr = [2,3,4], the median is 3.
 * 	For examples, if arr = [1,2,3,4], the median is (2 + 3) / 2 = 2.5.
 *
 * You are given an integer array nums and an integer k. There is a sliding window of size k which is
 * moving from the very left of the array to the very right. You can only see the k numbers in the
 * window. Each time the sliding window moves right by one position.
 *
 * Return the median array for each window in the original array. Answers within 10^-5 of the actual
 * value will be accepted.
 *
 * Example 1:
 *
 * Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
 * Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
 * Explanation:
 * Window position                Median
 * ---------------                -----
 * [1  3  -1] -3  5  3  6  7        1
 *  1 [3  -1  -3] 5  3  6  7       -1
 *  1  3 [-1  -3  5] 3  6  7       -1
 *  1  3  -1 [-3  5  3] 6  7        3
 *  1  3  -1  -3 [5  3  6] 7        5
 *  1  3  -1  -3  5 [3  6  7]       6
 *
 * Example 2:
 *
 * Input: nums = [1,2,3,4,2,3,1,4,2], k = 3
 * Output: [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]
 *
 * Constraints:
 *
 * 	1 <= k <= nums.length <= 10^5
 * 	-2^31 <= nums[i] <= 2^31 - 1
 ******************************************************************************************************/

// 思路，两个堆，延迟删除

package slidingwindowmedian

import "container/heap"

func medianSlidingWindow(nums []int, k int) []float64 {
	small := &MaxHeap{}
	large := &MinHeap{}

	heap.Init(small)
	heap.Init(large)

	// 需要延迟删除的元素
	delayed := make(map[int]int)

	// 有效元素数量，不计算 delayed 元素
	smallSize := 0
	largeSize := 0

	// 删除堆顶已经失效的元素
	pruneSmall := func() {
		for small.Len() > 0 {
			x := (*small)[0]

			if delayed[x] == 0 {
				break
			}

			delayed[x]--
			if delayed[x] == 0 {
				delete(delayed, x)
			}

			heap.Pop(small)
		}
	}

	pruneLarge := func() {
		for large.Len() > 0 {
			x := (*large)[0]

			if delayed[x] == 0 {
				break
			}

			delayed[x]--
			if delayed[x] == 0 {
				delete(delayed, x)
			}

			heap.Pop(large)
		}
	}

	// 调整两个堆，使：
	//
	// smallSize == largeSize
	// 或
	// smallSize == largeSize + 1
	rebalance := func() {
		if smallSize > largeSize+1 {
			x := heap.Pop(small).(int)
			heap.Push(large, x)

			smallSize--
			largeSize++

			pruneSmall()

		} else if smallSize < largeSize {
			x := heap.Pop(large).(int)
			heap.Push(small, x)

			largeSize--
			smallSize++

			pruneLarge()
		}
	}

	add := func(x int) {
		if small.Len() == 0 || x <= (*small)[0] {
			heap.Push(small, x)
			smallSize++
		} else {
			heap.Push(large, x)
			largeSize++
		}

		rebalance()
	}

	remove := func(x int) {
		// 先标记删除
		delayed[x]++

		// 判断 x 属于哪一边
		if x <= (*small)[0] {
			smallSize--

			// 如果恰好就在堆顶，立即清理
			if x == (*small)[0] {
				pruneSmall()
			}
		} else {
			largeSize--

			if large.Len() > 0 && x == (*large)[0] {
				pruneLarge()
			}
		}

		rebalance()
	}

	getMedian := func() float64 {
		pruneSmall()
		pruneLarge()

		if k%2 == 1 {
			return float64((*small)[0])
		}

		return (float64((*small)[0]) + float64((*large)[0])) / 2
	}

	// 初始化第一个窗口
	for i := 0; i < k; i++ {
		add(nums[i])
	}

	ans := make([]float64, 0, len(nums)-k+1)
	ans = append(ans, getMedian())

	// 滑动窗口
	for i := k; i < len(nums); i++ {
		// 移除最左边
		remove(nums[i-k])

		// 加入新元素
		add(nums[i])

		ans = append(ans, getMedian())
	}

	return ans
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[j] > h[i]
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

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[j] < h[i]
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
