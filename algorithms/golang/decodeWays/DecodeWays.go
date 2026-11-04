// Source : https://leetcode.com/problems/decode-ways
// Author : BradleyZhang
// Date   : 2026-11-04

/*****************************************************************************************************
 *
 * You have intercepted a secret message encoded as a string of numbers. The message is decoded via
 * the following mapping:
 *
 * "1" -> 'A'
 * "2" -> 'B'
 * ...
 * "25" -> 'Y'
 * "26" -> 'Z'
 *
 * However, while decoding the message, you realize that there are many different ways you can decode
 * the message because some codes are contained in other codes ("2" and "5" vs "25").
 *
 * For example, "11106" can be decoded into:
 *
 * 	"AAJF" with the grouping (1, 1, 10, 6)
 * 	"KJF" with the grouping (11, 10, 6)
 * 	The grouping (1, 11, 06) is invalid because "06" is not a valid code (only "6" is valid).
 *
 * Note: there may be strings that are impossible to decode.
 *
 * Given a string s containing only digits, return the number of ways to decode it. If the entire
 * string cannot be decoded in any valid way, return 0.
 *
 * The test cases are generated so that the answer fits in a 32-bit integer.
 *
 * Example 1:
 *
 * Input: s = "12"
 *
 * Output: 2
 *
 * Explanation:
 *
 * "12" could be decoded as "AB" (1 2) or "L" (12).
 *
 * Example 2:
 *
 * Input: s = "226"
 *
 * Output: 3
 *
 * Explanation:
 *
 * "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
 *
 * Example 3:
 *
 * Input: s = "06"
 *
 * Output: 0
 *
 * Explanation:
 *
 * "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06"). In this
 * case, the string is not a valid encoding, so return 0.
 *
 * Constraints:
 *
 * 	1 <= s.length <= 100
 * 	s contains only digits and may contain leading zero(s).
 ******************************************************************************************************/

package decodeways

import "strconv"

// dp[i] = s[0:i] 的解码方式数 // 这样避免了对 s[1]=='0' 的判断
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	// prev2 相当于 dp[i-2]，prev1 相当于 dp[i-1]
	prev2, prev1 := 1, 1 // dp[0]=1, dp[1]=1

	for i := 2; i <= n; i++ {
		cur := 0
		if s[i-1] != '0' {
			cur += prev1
		}
		twoDigit := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
		if twoDigit >= 10 && twoDigit <= 26 {
			cur += prev2
		}

		prev2, prev1 = prev1, cur
	}

	return prev1
}

// dp[i] = s[0:i+1] 的解码方式数
// dp[i]= dp[i-1]+dp[i-2]
func numDecodingsBakup(s string) int {
	n := len(s)
	dp := make([]int, n)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	if n == 1 {
		return 1
	}
	if s[1] == '0' {
		if ok(s[:2]) {
			dp[1] = 1
		} else {
			return 0
		}
	} else {
		dp[1] = 1
		if ok(s[:2]) {
			dp[1]++
		}
	}
	for i := 2; i < n; i++ {
		if s[i] == '0' {
			if ok(s[i-1 : i+1]) {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else {
			dp[i] = dp[i-1]
			if ok(s[i-1 : i+1]) {
				dp[i] += dp[i-2]
			}

		}

	}
	return dp[n-1]
}
func ok(s string) bool {
	if s[0] == '0' {
		return false
	}
	sx, _ := strconv.Atoi(s[:2])
	if sx < 27 {
		return true
	}
	return false
}
