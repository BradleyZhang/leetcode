// Source : https://leetcode.com/problems/evaluate-reverse-polish-notation
// Author : BradleyZhang
// Date   : 2026-09-10

/*****************************************************************************************************
 *
 * You are given an array of strings tokens that represents an arithmetic expression in a Reverse
 * Polish Notation.
 *
 * Evaluate the expression. Return an integer that represents the value of the expression.
 *
 * Note that:
 *
 * 	The valid operators are '+', '-', '*', and '/'.
 * 	Each operand may be an integer or another expression.
 * 	The division between two integers always truncates toward zero.
 * 	There will not be any division by zero.
 * 	The input represents a valid arithmetic expression in a reverse polish notation.
 * 	The answer and all the intermediate calculations can be represented in a 32-bit integer.
 *
 * Example 1:
 *
 * Input: tokens = ["2","1","+","3","*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 *
 * Example 2:
 *
 * Input: tokens = ["4","13","5","/","+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 *
 * Example 3:
 *
 * Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
 * Output: 22
 * Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 * Constraints:
 *
 * 	1 <= tokens.length <= 10^4
 * 	tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200,
 * 200].
 ******************************************************************************************************/
package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := stackWrapper{}
	for _, t := range tokens {
		switch t {
		case "+", "-", "*", "/":
			x := stack.top()
			right, _ := strconv.Atoi(x)
			stack.pop()
			x = stack.top()
			left, _ := strconv.Atoi(x)
			stack.pop()
			switch t {
			case "+":
				stack.push(strconv.Itoa(left + right))
			case "-":
				stack.push(strconv.Itoa(left - right))
			case "*":
				stack.push(strconv.Itoa(left * right))
			case "/":
				stack.push(strconv.Itoa(left / right))
			}
		default:
			stack.push(t)
		}
	}
	result, _ := strconv.Atoi(stack.top())
	return result
}

type stackWrapper struct {
	stack []string
}

func (s *stackWrapper) push(c string) {
	s.stack = append(s.stack, c)
}
func (s *stackWrapper) top() string {
	return s.stack[len(s.stack)-1]
}
func (s *stackWrapper) pop() (ok bool) {
	if len(s.stack) == 0 {
		return false
	}
	s.stack = s.stack[:len(s.stack)-1]
	return true
}
func (s *stackWrapper) empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

// 更优雅写法
type stack[T any] []T

func push[T any](s stack[T], ts ...T) stack[T] {
	for _, t := range ts {
		s = append(s, t)
	}
	return s
}

func pop[T any](s stack[T]) (T, stack[T]) {
	var t T
	if len(s) == 0 {
		return t, s
	}
	i := len(s) - 1
	t = s[i]
	s = s[:i]
	return t, s
}

func evalRPNBest(tokens []string) int {
	s := make(stack[int], len(tokens))
	for _, tok := range tokens {
		if i, err := strconv.Atoi(tok); err == nil {
			s = push(s, i)
		} else {
			var a, b, c int
			b, s = pop(s)
			a, s = pop(s)
			switch tok {
			case "+":
				c = a + b
			case "-":
				c = a - b
			case "*":
				c = a * b
			case "/":
				c = a / b
			default:
				panic("not implemented")
			}
			s = push(s, c)
		}
	}
	i, _ := pop(s)
	return i
}
