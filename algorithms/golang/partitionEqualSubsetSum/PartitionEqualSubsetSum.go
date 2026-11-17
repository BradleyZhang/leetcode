// Source : https://leetcode.com/problems/partition-equal-subset-sum
// Author : BradleyZhang
// Date   : 2026-11-17

/*****************************************************************************************************
 *
 * Given an integer array nums, return true if you can partition the array into two subsets such that
 * the sum of the elements in both subsets is equal or false otherwise.
 *
 * Example 1:
 *
 * Input: nums = [1,5,11,5]
 * Output: true
 * Explanation: The array can be partitioned as [1, 5, 5] and [11].
 *
 * Example 2:
 *
 * Input: nums = [1,2,3,5]
 * Output: false
 * Explanation: The array cannot be partitioned into equal sum subsets.
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 200
 * 	1 <= nums[i] <= 100
 ******************************************************************************************************/

package partitionequalsubsetsum

// 前 i 个数能否凑出 sum/2
// 决策： 第 i 个数加入/不加入 能凑出
// dp[j] = 能不能凑出数字 j
// 状态转移：dp[j] = dp[j] || dp[j-num] 对一个数字 num dp[j] 即不加入能否凑出 , dp[j-num] 加入能否凑出
func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 != 0 {
		return false
	}

	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		// 往 num 前面遍历
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
		// 提前剪枝：已经能凑出 target
		if dp[target] {
			return true
		}
	}

	return dp[target]
}
