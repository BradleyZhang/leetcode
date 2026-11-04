// Source : https://leetcode.com/problems/house-robber-ii
// Author : BradleyZhang
// Date   : 2026-11-04

/*****************************************************************************************************
 *
 * You are a professional robber planning to rob houses along a street. Each house has a certain
 * amount of money stashed. All houses at this place are arranged in a circle. That means the first
 * house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected,
 * and it will automatically contact the police if two adjacent houses were broken into on the same
 * night.
 *
 * Given an integer array nums representing the amount of money of each house, return the maximum
 * amount of money you can rob tonight without alerting the police.
 *
 * Example 1:
 *
 * Input: nums = [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are
 * adjacent houses.
 *
 * Example 2:
 *
 * Input: nums = [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 3:
 *
 * Input: nums = [1,2,3]
 * Output: 3
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 100
 * 	0 <= nums[i] <= 1000
 ******************************************************************************************************/
// x x x
package houserobberii

// 第0间和第n-1 间不能同时rob
// rob 0-> 最后一间不考虑
// 不rob 0-> 考虑 n-1
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robLine(nums[:n-1]), robLine(nums[1:]))
}

func robLine(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}
