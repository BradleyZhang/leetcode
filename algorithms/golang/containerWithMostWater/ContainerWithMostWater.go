// Source : https://leetcode.com/problems/container-with-most-water
// Author : BradleyZhang
// Date   : 2026-09-02

/*****************************************************************************************************
 *
 * You are given an integer array height of length n. There are n vertical lines drawn such that the
 * two endpoints of the i^th line are (i, 0) and (i, height[i]).
 *
 * Find two lines that together with the x-axis form a container, such that the container contains the
 * most water.
 *
 * Return the maximum amount of water a container can store.
 *
 * Notice that you may not slant the container.
 *
 * Example 1:
 *
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case,
 * the max area of water (blue section) the container can contain is 49.
 *
 * Example 2:
 *
 * Input: height = [1,1]
 * Output: 1
 *
 * Constraints:
 *
 * 	n == height.length
 * 	2 <= n <= 10^5
 * 	0 <= height[i] <= 10^4
 ******************************************************************************************************/

package containerwithmostwater

// You start with the widest possible container and move the pointer on the shorter side inward,
// since that is the only move that could increase the area.
// This problem teaches the key insight that moving the longer side never helps,
// which is a non-obvious greedy choice that the two pointer pattern makes visible.
func maxArea(height []int) int {
	var result int
	left, right := 0, len(height)-1
	for left < right {
		if a := area(left, right, height); a > result {
			result = a
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return result
}
func area(i int, j int, height []int) int {
	if i == j {
		return 0
	}
	return min(height[i], height[j]) * (max(i, j) - min(i, j))
}
