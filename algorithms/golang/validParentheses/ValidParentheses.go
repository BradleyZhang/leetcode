// Source : https://leetcode.com/problems/valid-parentheses
// Author : BradleyZhang
// Date   : 2026-09-09

/*****************************************************************************************************
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the
 * input string is valid.
 *
 * An input string is valid if:
 *
 * 	Open brackets must be closed by the same type of brackets.
 * 	Open brackets must be closed in the correct order.
 * 	Every close bracket has a corresponding open bracket of the same type.
 *
 * Example 1:
 * Input: s = "()"
 * Output: true
 *
 * Example 2:
 * Input: s = "()[]{}"
 * Output: true
 *
 * Example 3:
 * Input: s = "(]"
 * Output: false
 *
 * Example 4:
 * Input: s = "([])"
 * Output: true
 *
 * Example 5:
 * Input: s = "([)]"
 * Output: false
 *
 * Constraints:
 * 	1 <= s.length <= 10^4
 * 	s consists of parentheses only '()[]{}'.
 ******************************************************************************************************/
package validparentheses

func isValid(s string) bool {
	stack := stack{}
	for _, c := range s {
		if stack.empty() {
			stack.push(c)
		} else {
			if pair(stack.top(), c) {
				stack.pop()
			} else {
				stack.push(c)
			}
		}
	}
	return stack.empty()
}
func pair(left, right rune) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	return false
}

type stack struct {
	stack []rune
}

func (s *stack) push(c rune) {
	s.stack = append(s.stack, c)
}
func (s *stack) top() rune {
	return s.stack[len(s.stack)-1]
}
func (s *stack) pop() rune {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}
func (s *stack) empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}
