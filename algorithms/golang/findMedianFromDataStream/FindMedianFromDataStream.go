// Source : https://leetcode.com/problems/find-median-from-data-stream
// Author : BradleyZhang
// Date   : 2026-10-03

/*****************************************************************************************************
 *
 * The median is the middle value in an ordered integer list. If the size of the list is even, there
 * is no middle value, and the median is the mean of the two middle values.
 *
 * 	For example, for arr = [2,3,4], the median is 3.
 * 	For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
 *
 * Implement the MedianFinder class:
 *
 * 	MedianFinder() initializes the MedianFinder object.
 * 	void addNum(int num) adds the integer num from the data stream to the data structure.
 * 	double findMedian() returns the median of all elements so far. Answers within 10^-5 of the
 * actual answer will be accepted.
 *
 * Example 1:
 *
 * Input
 * ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
 * [[], [1], [2], [], [3], []]
 * Output
 * [null, null, null, 1.5, null, 2.0]
 *
 * Explanation
 * MedianFinder medianFinder = new MedianFinder();
 * medianFinder.addNum(1);    // arr = [1]
 * medianFinder.addNum(2);    // arr = [1, 2]
 * medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
 * medianFinder.addNum(3);    // arr[1, 2, 3]
 * medianFinder.findMedian(); // return 2.0
 *
 * Constraints:
 *
 * 	-10^5 <= num <= 10^5
 * 	There will be at least one element in the data structure before calling findMedian.
 * 	At most 5 * 10^4 calls will be made to addNum and findMedian.
 *
 * Follow up:
 *
 * 	If all integer numbers from the stream are in the range [0, 100], how would you optimize
 * your solution?
 * 	If 99% of all integer numbers from the stream are in the range [0, 100], how would you
 * optimize your solution?
 ******************************************************************************************************/

package findmedianfromdatastream

import "container/heap"

type MedianFinder struct {
	smallHalf MaxHeap
	bigHalf   MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		smallHalf: MaxHeap{},
		bigHalf:   MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.smallHalf.Len() == 0 || num <= (this.smallHalf)[0] {
		heap.Push(&this.smallHalf, num)
	} else {
		heap.Push(&this.bigHalf, num)
	}
	// 平衡两个堆
	if this.smallHalf.Len() > this.bigHalf.Len()+1 {
		heap.Push(&this.bigHalf, heap.Pop(&this.smallHalf))
	} else if this.bigHalf.Len() > this.smallHalf.Len() {
		heap.Push(&this.smallHalf, heap.Pop(&this.bigHalf))
	}
}
func (this *MedianFinder) AddNumBackup(num int) {
	if this.smallHalf.Len() == 0 {
		heap.Push(&this.smallHalf, num)
		return
	}
	if this.bigHalf.Len() == 0 {
		if num > this.smallHalf[0] {
			heap.Push(&this.bigHalf, num)
		} else {
			temp := heap.Pop(&this.smallHalf)
			heap.Push(&this.bigHalf, temp)
			heap.Push(&this.smallHalf, num)
		}
		return
	}
	left := this.smallHalf[0]
	if num < left {
		heap.Push(&this.smallHalf, num)
	} else {
		heap.Push(&this.bigHalf, num)
	}
	for this.bigHalf.Len() > this.smallHalf.Len() {
		temp := heap.Pop(&this.bigHalf)
		heap.Push(&this.smallHalf, temp)
	}
	for this.smallHalf.Len()-this.bigHalf.Len() > 1 {
		temp := heap.Pop(&this.smallHalf)
		heap.Push(&this.bigHalf, temp)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smallHalf.Len() > this.bigHalf.Len() {
		return float64(this.smallHalf[0])
	} else {
		return float64(this.smallHalf[0]+this.bigHalf[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
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
