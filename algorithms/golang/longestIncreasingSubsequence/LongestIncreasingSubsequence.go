// Source : https://leetcode.com/problems/longest-increasing-subsequence
// Author : BradleyZhang
// Date   : 2026-11-03

/*****************************************************************************************************
 *
 * Given an integer array nums, return the length of the longest strictly increasing subsequence.
 *
 * Example 1:
 *
 * Input: nums = [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
 *
 * Example 2:
 *
 * Input: nums = [0,1,0,3,2,3]
 * Output: 4
 *
 * Example 3:
 *
 * Input: nums = [7,7,7,7,7,7,7]
 * Output: 1
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 2500
 * 	-10^4 <= nums[i] <= 10^4
 *
 * Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
 ******************************************************************************************************/

package longestincreasingsubsequence

import "sort"

// 二分查找优化
// tails[k] = 所有长度为 k+1 的递增子序列中，结尾元素的最小值。
// 性质 tails[k-1] < tails[k] (因为 若存在长度 k+1 的递增子序列结尾为 x，那它前 k 个元素构成长度 k 的子序列，结尾 < x)
func lengthOfLIS(nums []int) int {
	tails := []int{}

	for i := range nums {
		if len(tails) == 0 || tails[len(tails)-1] < nums[i] {
			tails = append(tails, nums[i])
		} else {
			idx := sort.SearchInts(tails, nums[i])
			tails[idx] = nums[i]
		}
	}
	return len(tails)
}

// dp[i] 0...i 的 答案
// 状态转移 dp[i]=1+max(dp[x]) 需要 dp[x]<dp[i]
// O(n^2)
func lengthOfLISNormal(nums []int) int {
	var ans int
	dp := make([]int, len(nums))
	for i := range dp {
		num := nums[i]
		maxLen := 0
		for j := 0; j < i; j++ {
			if nums[j] < num {
				maxLen = max(maxLen, dp[j])
			}
		}
		dp[i] = maxLen + 1
		ans = max(ans, dp[i])
	}
	return ans
}
