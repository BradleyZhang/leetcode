// Source : https://leetcode.com/problems/word-break
// Author : BradleyZhang
// Date   : 2026-11-04

/*****************************************************************************************************
 *
 * Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a
 * space-separated sequence of one or more dictionary words.
 *
 * Note that the same word in the dictionary may be reused multiple times in the segmentation.
 *
 * Example 1:
 *
 * Input: s = "leetcode", wordDict = ["leet","code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet code".
 *
 * Example 2:
 *
 * Input: s = "applepenapple", wordDict = ["apple","pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 * Example 3:
 *
 * Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 * Output: false
 *
 * Constraints:
 *
 * 	1 <= s.length <= 300
 * 	1 <= wordDict.length <= 1000
 * 	1 <= wordDict[i].length <= 20
 * 	s and wordDict[i] consist of only lowercase English letters.
 * 	All the strings of wordDict are unique.
 ******************************************************************************************************/

package wordbreak

// dp[i] = s[0:i+1] 是否可由 wordDict 组成
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		dict[w] = struct{}{}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0 && i-j <= maxLen; j-- { // 用 maxLen 剪枝
			if dp[j] {
				if _, ok := dict[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}

// 使用布尔型动态规划数组来追踪可达性
// dp[i] = s[0:i+1] 的答案
func wordBreakBackup(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	for _, v := range wordDict {
		dict[v] = struct{}{}
	}

	n := len(s)
	dp := make([]bool, n)
	todoI := []int{-1}
	for i := 0; i < n; i++ {
		for j := len(todoI) - 1; j > -1; j-- {
			v := todoI[j]
			if _, ok := dict[s[v+1:i+1]]; ok {
				dp[i] = true
				todoI = append(todoI, i)
				break
			}
		}
	}
	return dp[n-1]
}

// Time Limit Exceeded
func wordBreakDFS(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	for _, v := range wordDict {
		dict[v] = struct{}{}
	}
	var dfs func(start int) bool
	dfs = func(start int) bool {
		for i := start; i < len(s)-1; i++ {
			if _, ok := dict[s[start:i+1]]; ok {
				if dfs(i + 1) {
					return true
				}
			}
		}
		if _, ok := dict[s[start:]]; ok {
			return true
		}
		return false
	}
	return dfs(0)
}
