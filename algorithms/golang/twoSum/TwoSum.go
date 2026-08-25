// Source : https://leetcode.com/problems/two-sum/
// Author : BradleyZ
// Date   : 2026-07-02

/*****************************************************************************************************
 *
 * Given an array of integers nums and an integer target, return indices of the two numbers such that
 * they add up to target.
 *
 * You may assume that each input would have exactly one solution, and you may not use the same
 * element twice.
 *
 * You can return the answer in any order.
 *
 * Example 1:
 *
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 *
 * Example 2:
 *
 * Input: nums = [3,2,4], target = 6
 * Output: [1,2]
 *
 * Example 3:
 *
 * Input: nums = [3,3], target = 6
 * Output: [0,1]
 *
 * Constraints:
 *
 * 	2 <= nums.length <= 10^4
 * 	-10^9 <= nums[i] <= 10^9
 * 	-10^9 <= target <= 10^9
 * 	Only one valid answer exists.
 *
 * Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?
 ******************************************************************************************************/
package twosum

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 其他解法
// 需要快速知道数组中有没有 target-num[x]
// 常用pattern：数组中需要快速查找->用 hash map

// 先遍历一遍构造hash map，再遍历一遍用hash map 找 target-num[x]
func twoSumHashMap(nums []int, target int) []int {
	vToIndex := make(map[int]int)
	for i, v := range nums {
		vToIndex[v] = i
	}
	for i, v := range nums {
		j, ok := vToIndex[target-v]
		if ok {
			if i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 边遍历边填 hash map，并检索已有的hash map
func twoSumHashMapOnce(nums []int, target int) []int {
	vToIndex := make(map[int]int)
	for i, v := range nums {
		if j, ok := vToIndex[target-v]; ok {
			return []int{i, j}
		}
		vToIndex[v] = i
	}
	return []int{}
}
