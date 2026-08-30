// Source : https://leetcode.com/problems/majority-element
// Author : BradleyZhang
// Date   : 2026-08-30

/*****************************************************************************************************
 *
 * Given an array nums of size n, return the majority element.
 *
 * The majority element is the element that appears more than &lfloor;n / 2&rfloor; times. You may
 * assume that the majority element always exists in the array.
 *
 * Example 1:
 * Input: nums = [3,2,3]
 * Output: 3
 * Example 2:
 * Input: nums = [2,2,1,1,1,2,2]
 * Output: 2
 *
 * Constraints:
 *
 * 	n == nums.length
 * 	1 <= n <= 5 * 10^4
 * 	-10^9 <= nums[i] <= 10^9
 * 	The input is generated such that a majority element will exist in the array.
 *
 * Follow-up: Could you solve the problem in linear time and in O(1) space?
 ******************************************************************************************************/
package majorityelement

func majorityElement(nums []int) int {
	n := len(nums) / 2
	numToCount := make(map[int]int, len(nums)/2+1)
	for _, num := range nums {
		numToCount[num]++
		if numToCount[num] > n {
			return num
		}
	}
	return 0
}

// Space complexity: O(1) 的最佳算法
// moore voting algorithm
// 用于找到数组中多数元素（> n/2)
// 使用candidate 和 count在count为0是换candidate为当前元素，遍历后candidate为多数元素
func majorityElementMVA(nums []int) int {
	count := 0
	var candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}

	}
	return candidate
}
