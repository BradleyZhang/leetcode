// Source : https://leetcode.com/problems/longest-palindromic-substring
// Author : BradleyZhang
// Date   : 2026-11-04

/*****************************************************************************************************
 *
 * Given a string s, return the longest palindromic substring in s.
 *
 * Example 1:
 *
 * Input: s = "babad"
 * Output: "bab"
 * Explanation: "aba" is also a valid answer.
 *
 * Example 2:
 *
 * Input: s = "cbbd"
 * Output: "bb"
 *
 * Constraints:
 *
 * 	1 <= s.length <= 1000
 * 	s consist of only digits and English letters.
 ******************************************************************************************************/
package longestpalindromicsubstring

func longestPalindrome(s string) string {
	var start, end int

	for center, sLen := 0, len(s); center < 2*sLen; center++ {
		left := center / 2
		right := left + center%2

		for left >= 0 && right < sLen && s[left] == s[right] {
			left--
			right++
		}

		if right-left > end-start {
			start = left
			end = right
		}
	}
	return s[start+1 : end]
}
func longestPalindromeDP(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	maxLen := 1
	maxStr := s[:1]

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		// 单个字符一定是回文串
		dp[i][i] = true

		for j := 0; j < i; j++ {
			// 判断首尾字符是否相等
			if s[j] == s[i] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true

				// 更新最长回文子串
				if i-j+1 > maxLen {
					maxLen = i - j + 1
					maxStr = s[j : i+1]
				}
			}
		}
	}

	return maxStr
}
