// Source : https://leetcode.com/problems/largest-rectangle-in-histogram
// Author : BradleyZhang
// Date   : 2026-09-09

/*****************************************************************************************************
 *
 * Given an array of integers heights representing the histogram's bar height where the width of each
 * bar is 1, return the area of the largest rectangle in the histogram.
 *
 * Example 1:
 *
 * Input: heights = [2,1,5,6,2,3]
 * Output: 10
 * Explanation: The above is a histogram where width of each bar is 1.
 * The largest rectangle is shown in the red area, which has an area = 10 units.
 *
 * Example 2:
 *
 * Input: heights = [2,4]
 * Output: 4
 *
 * Constraints:
 *
 * 	1 <= heights.length <= 10^5
 * 	0 <= heights[i] <= 10^4
 ******************************************************************************************************/

package largestrectangleinhistogram

// 求，对每个柱子，找到左右第一个比他矮的，来计算 heights[i] 可延伸到哪
// 用栈，跟踪某个稍后才被解决的状态，即跟踪right
func largestRectangleArea(heights []int) int {
	result := 0
	indexStack := stack{}
	heights = append(heights, 0) // 增加0，处理最后一个height
	for x, curH := range heights {
		for top, ok := indexStack.top(); ok && heights[top] > curH; top, ok = indexStack.top() {
			indexStack.pop()
			height := heights[top]
			var width int
			if left, ok := indexStack.top(); ok {
				width = x - left - 1
			} else { // pop 后栈空，说明能延伸到最左边
				width = x
			}
			area := width * height
			if result < area {
				result = area
			}
		}
		indexStack.push(x)
	}
	return result
}

type stack struct {
	stack []int
}

func (s *stack) push(c int) {
	s.stack = append(s.stack, c)
}
func (s *stack) top() (top int, ok bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	return s.stack[len(s.stack)-1], true
}
func (s *stack) pop() (ok bool) {
	if len(s.stack) == 0 {
		return false
	}
	s.stack = s.stack[:len(s.stack)-1]
	return true
}
func (s *stack) empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}
