// Source : https://leetcode.com/problems/house-robber
// Author : BradleyZhang
// Date   : 2026-11-02

/*****************************************************************************************************
 *
 * You are a professional robber planning to rob houses along a street. Each house has a certain
 * amount of money stashed, the only constraint stopping you from robbing each of them is that
 * adjacent houses have security systems connected and it will automatically contact the police if two
 * adjacent houses were broken into on the same night.
 *
 * Given an integer array nums representing the amount of money of each house, return the maximum
 * amount of money you can rob tonight without alerting the police.
 *
 * Example 1:
 *
 * Input: nums = [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 2:
 *
 * Input: nums = [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 100
 * 	0 <= nums[i] <= 400
 ******************************************************************************************************/

package houserobber

// 决策：rob 或 not rob
// dp[i] 表示 0...i 能 rob 的最大金额
// 初始状态 dp[0]=nums[0] dp[1]=max(dp1,dp2)
// 状态转移: dp[i] = max(dp[i-1], dp[i-2] + nums[i]) 不抢 i dp[i]==dp[i+1] , dp
func rob(nums []int) int {
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
func robDFS(nums []int) int {
	n := len(nums)
	var dfs func(i int, total int) int
	robed := make([]int, len(nums))
	dfs = func(i int, total int) int {
		if i > n-1 {
			return total
		}
		var a, b int
		if robed[i-1] == 0 {
			robed[i] = 1
			a = dfs(i+1, total+nums[i])
			robed[i] = 0
		}
		b = dfs(i+1, total)
		return max(a, b)
	}
	robed[0] = 1
	rob0 := dfs(1, nums[0])
	robed[0] = 0
	nrob0 := dfs(1, 0)
	return max(rob0, nrob0)
}
