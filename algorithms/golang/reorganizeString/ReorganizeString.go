// Source : https://leetcode.com/problems/reorganize-string
// Author : BradleyZhang
// Date   : 2026-10-04

/*****************************************************************************************************
 *
 * Given a string s, rearrange the characters of s so that any two adjacent characters are not the
 * same.
 *
 * Return any possible rearrangement of s or return "" if not possible.
 *
 * Example 1:
 * Input: s = "aab"
 * Output: "aba"
 * Example 2:
 * Input: s = "aaab"
 * Output: ""
 *
 * Constraints:
 *
 * 	1 <= s.length <= 500
 * 	s consists of lowercase English letters.
 ******************************************************************************************************/
package reorganizestring

import (
	"container/heap"
	"strings"
)

func reorganizeString(s string) string {
	letters := [26]int{}
	maxCount := 0

	for _, r := range s {
		letters[int(r-'a')]++
		maxCount = max(maxCount, letters[int(r-'a')])
	}
	if maxCount > (len(s)+1)/2 {
		return ""
	}
	maxHeap := MaxHeap{}
	for i, v := range letters {
		if v > 0 {
			heap.Push(&maxHeap, Letter{
				l:   rune('a' + i),
				num: v,
			})
		}
	}

	builder := strings.Builder{}
	for maxHeap.Len() > 1 {
		l := heap.Pop(&maxHeap).(Letter)
		r := heap.Pop(&maxHeap).(Letter)
		l.num--
		r.num--
		builder.WriteRune(l.l)
		builder.WriteRune(r.l)
		if l.num > 0 {
			heap.Push(&maxHeap, l)
		}
		if r.num > 0 {
			heap.Push(&maxHeap, r)
		}
	}
	if maxHeap.Len() > 0 {
		l := heap.Pop(&maxHeap).(Letter)
		l.num--
		builder.WriteRune(l.l)
	}
	return builder.String()
}

type Letter struct {
	l   rune
	num int
}
type MaxHeap []Letter

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[j].num < h[i].num
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Letter))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
