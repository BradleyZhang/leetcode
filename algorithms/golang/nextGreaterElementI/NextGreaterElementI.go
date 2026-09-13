// Source : https://leetcode.com/problems/next-greater-element-i
// Author : BradleyZhang
// Date   : 2026-09-13

/*****************************************************************************************************
 *
 * The next greater element of some element x in an array is the first greater element that is to the
 * right of x in the same array.
 *
 * You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of
 * nums2.
 *
 * For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the
 * next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for
 * this query is -1.
 *
 * Return an array ans of length nums1.length such that ans[i] is the next greater element as
 * described above.
 *
 * Example 1:
 *
 * Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
 * Output: [-1,3,-1]
 * Explanation: The next greater element for each value of nums1 is as follows:
 * - 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
 * - 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
 * - 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
 *
 * Example 2:
 *
 * Input: nums1 = [2,4], nums2 = [1,2,3,4]
 * Output: [3,-1]
 * Explanation: The next greater element for each value of nums1 is as follows:
 * - 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
 * - 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.
 *
 * Constraints:
 *
 * 	1 <= nums1.length <= nums2.length <= 1000
 * 	0 <= nums1[i], nums2[i] <= 10^4
 * 	All integers in nums1 and nums2 are unique.
 * 	All the integers of nums1 also appear in nums2.
 *
 * Follow up: Could you find an O(nums1.length + nums2.length) solution?
 ******************************************************************************************************/

package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	numToIndex1 := make(map[int]int, len(nums1))
	for i, n := range nums1 {
		numToIndex1[n] = i
	}
	stack := stack[int]{}
	for _, n := range nums2 {
		for !empty(stack) && top(stack) < n {
			var x int
			x, stack = pop(stack)
			if i, ok := numToIndex1[x]; ok {
				result[i] = n
			}
		}
		stack = push(stack, n)
	}
	for !empty(stack) {
		var x int
		x, stack = pop(stack)
		if i, ok := numToIndex1[x]; ok {
			result[i] = -1
		}
	}
	return result
}

// 更优：result初始化为全-1,省去最后清理栈

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	for i := range result {
		result[i] = -1
	}
	numToIndex1 := make(map[int]int, len(nums1))
	for i, n := range nums1 {
		numToIndex1[n] = i
	}
	stack := stack[int]{}
	for _, n := range nums2 {
		for !empty(stack) && top(stack) < n {
			var x int
			x, stack = pop(stack)
			if i, ok := numToIndex1[x]; ok {
				result[i] = n
			}
		}
		stack = push(stack, n)
	}
	return result
}

// 递减栈写法
func nextGreaterElement3(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2))

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]

		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			nextGreater[num] = -1
		} else {
			nextGreater[num] = stack[len(stack)-1]
		}

		stack = append(stack, num)
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}

	return result
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
