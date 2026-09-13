// Source : https://leetcode.com/problems/remove-k-digits
// Author : BradleyZhang
// Date   : 2026-09-13

/*****************************************************************************************************
 *
 * Given string num representing a non-negative integer num, and an integer k, return the smallest
 * possible integer after removing k digits from num.
 *
 * Example 1:
 *
 * Input: num = "1432219", k = 3
 * Output: "1219"
 * Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
 *
 * Example 2:
 *
 * Input: num = "10200", k = 1
 * Output: "200"
 * Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain
 * leading zeroes.
 *
 * Example 3:
 *
 * Input: num = "10", k = 2
 * Output: "0"
 * Explanation: Remove all the digits from the number and it is left with nothing which is 0.
 *
 * Constraints:
 *
 * 	1 <= k <= num.length <= 10^5
 * 	num consists of only digits.
 * 	num does not have any leading zeros except for the zero itself.
 ******************************************************************************************************/

package removekdigits

func removeKdigits(num string, k int) string {
	limit := k
	stack := stack[rune]{}
	for _, s := range num {
		if empty(stack) && s != '0' {
			stack = push(stack, s)
		} else {
			for !empty(stack) && limit > 0 && top(stack) > s {
				_, stack = pop(stack)
				limit--
			}
			if !empty(stack) || s != '0' {
				stack = push(stack, s)
			}
		}
	}
	for !empty(stack) && limit > 0 {
		_, stack = pop(stack)
		limit--
	}
	if empty(stack) {
		return "0"
	}
	return string(stack)
}

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
func top[T any](s stack[T]) T {
	return s[len(s)-1]
}
func empty[T any](s stack[T]) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// 其他写法
func Removekdigits(num string, k int) string {
	stack := make([]rune, 0)
	for _, ch := range num {
		for len(stack) > 0 && stack[len(stack)-1] > ch && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, ch)
	}

	// If all k digits haven't been removed, remove them.
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// Remove leading zeroes
	for i := 0; i < len(stack); i++ {
		// Return fast if you find a non-zero digit.
		if stack[i] != '0' {
			return string(stack[i:])
		}
	}

	// If you reach here, result must be zero.
	return "0"
}
