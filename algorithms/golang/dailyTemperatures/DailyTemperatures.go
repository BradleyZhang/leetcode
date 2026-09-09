// Source : https://leetcode.com/problems/daily-temperatures
// Author : BradleyZhang
// Date   : 2026-09-09

/*****************************************************************************************************
 *
 * Given an array of integers temperatures represents the daily temperatures, return an array answer
 * such that answer[i] is the number of days you have to wait after the i^th day to get a warmer
 * temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.
 *
 * Example 1:
 * Input: temperatures = [73,74,75,71,69,72,76,73]
 * Output: [1,1,4,2,1,1,0,0]
 * Example 2:
 * Input: temperatures = [30,40,50,60]
 * Output: [1,1,1,0]
 * Example 3:
 * Input: temperatures = [30,60,90]
 * Output: [1,1,0]
 *
 * Constraints:
 *
 * 	1 <= temperatures.length <= 10^5
 * 	30 <= temperatures[i] <= 100
 ******************************************************************************************************/

package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	indexStack := stack{}
	result := make([]int, len(temperatures))
	for i, t := range temperatures {
		for {
			top, ok := indexStack.top()
			if !ok || temperatures[top] >= t {
				break
			} else if temperatures[top] < t {
				result[top] = i - top
				indexStack.pop()
			}
		}
		indexStack.push(i)
	}
	for {
		top, ok := indexStack.top()
		if !ok {
			break
		}
		indexStack.pop()
		result[top] = 0
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
