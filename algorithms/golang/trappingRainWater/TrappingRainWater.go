// Source : https://leetcode.com/problems/trapping-rain-water
// Author : BradleyZhang
// Date   : 2026-09-02

/*****************************************************************************************************
 *
 * Given n non-negative integers representing an elevation map where the width of each bar is 1,
 * compute how much water it can trap after raining.
 *
 * Example 1:
 *
 * Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * Explanation: The above elevation map (black section) is represented by array
 * [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
 *
 * Example 2:
 *
 * Input: height = [4,2,0,3,2,5]
 * Output: 9
 *
 * Constraints:
 *
 * 	n == height.length
 * 	1 <= n <= 2 * 10^4
 * 	0 <= height[i] <= 10^5
 ******************************************************************************************************/
package trappingrainwater

// 先从左边计算每一个块加上水最多可能是多少，不考虑水从右边流掉。再从右边计算从右边水流掉后剩下多少
func trap(height []int) int {
	var result int
	barPlusWaterHeight := make([]int, len(height))
	barPlusWaterHeight[0] = height[0]
	for i := 1; i < len(height); i++ {
		if barPlusWaterHeight[i-1] > height[i] {
			barPlusWaterHeight[i] = barPlusWaterHeight[i-1]
		} else {
			barPlusWaterHeight[i] = height[i]
		}
	}
	barPlusWaterHeight[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i > -1; i-- {
		if barPlusWaterHeight[i] > barPlusWaterHeight[i+1] {
			barPlusWaterHeight[i] = max(height[i], barPlusWaterHeight[i+1])
		}
		result += barPlusWaterHeight[i] - height[i]
	}
	return result
}

// 最优解
// 水位由左右两侧最高柱子中较矮的那根决定
func trapTwoPoint(height []int) int {
	var leftMax, rightMax int
	left, right := 0, len(height)-1
	var result int
	for left < right {

		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}
	return result
}
