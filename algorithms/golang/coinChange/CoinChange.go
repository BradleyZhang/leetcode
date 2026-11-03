// Source : https://leetcode.com/problems/coin-change
// Author : BradleyZhang
// Date   : 2026-11-02

/*****************************************************************************************************
 *
 * You are given an integer array coins representing coins of different denominations and an integer
 * amount representing a total amount of money.
 *
 * Return the fewest number of coins that you need to make up that amount. If that amount of money
 * cannot be made up by any combination of the coins, return -1.
 *
 * You may assume that you have an infinite number of each kind of coin.
 *
 * Example 1:
 *
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 * Example 3:
 *
 * Input: coins = [1], amount = 0
 * Output: 0
 *
 * Constraints:
 *
 * 	1 <= coins.length <= 12
 * 	1 <= coins[i] <= 2^31 - 1
 * 	0 <= amount <= 10^4
 ******************************************************************************************************/
package coinchange

import "math"

// 更优
// dp[i] amout 为 i 时的最小 coin 数
// 初试状态 dp[0]=0
// 状态转移 dp[cur] = min( dp[cur],  dp[cur - coin] + 1 )
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for cur := coin; cur <= amount; cur++ {
			dp[cur] = min(dp[cur], dp[cur-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// dp[i] amout 为 i 时的最小 coin 数
// 初试状态 dp[0]=0
// 状态转移 dp[i]=min(1+dp[i-coins[0]],...1+dp[i-coins[j]])
func coinChangeBackup(coins []int, amount int) int {
	minCoin := coins[0]
	for _, v := range coins {
		minCoin = min(minCoin, v)
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for _, v := range coins {
		if v > amount {
			continue
		}
		dp[v] = 1
	}
	for i := minCoin + 1; i < amount+1; i++ {
		minNum := dp[i]
		for _, v := range coins {
			if i-v < minCoin || dp[i-v] == math.MaxInt {
				continue
			}
			minNum = min(minNum, 1+dp[i-v])
		}
		dp[i] = minNum
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
