// Source : https://leetcode.com/problems/minimum-window-substring
// Author : BradleyZhang
// Date   : 2026-09-18

/*****************************************************************************************************
 *
 * Given two strings s and t of lengths m and n respectively, return the minimum window substring of s
 * such that every character in t (including duplicates) is included in the window. If there is no
 * such substring, return the empty string "".
 *
 * The testcases will be generated such that the answer is unique.
 *
 * Example 1:
 *
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
 *
 * Example 2:
 *
 * Input: s = "a", t = "a"
 * Output: "a"
 * Explanation: The entire string s is the minimum window.
 *
 * Example 3:
 *
 * Input: s = "a", t = "aa"
 * Output: ""
 * Explanation: Both 'a's from t must be included in the window.
 * Since the largest window of s only has one 'a', return empty string.
 *
 * Constraints:
 *
 * 	m == s.length
 * 	n == t.length
 * 	1 <= m, n <= 10^5
 * 	s and t consist of uppercase and lowercase English letters.
 *
 * Follow up: Could you find an algorithm that runs in O(m + n) time?
 ******************************************************************************************************/

package minimumwindowsubstring

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := [128]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	count := len(t)

	ansL, ansR := 0, len(s)+1

	for right := 0; right < len(s); right++ {
		c := s[right]

		// c 是 t 中需要的字符
		if need[c] > 0 {
			count--
		}
		need[c]--

		// 当前窗口已经包含 t 的全部字符
		for count == 0 {
			// 更新答案
			if right-left+1 < ansR-ansL {
				ansL = left
				ansR = right + 1
			}

			// 移除左边字符
			c = s[left]
			need[c]++
			if need[c] > 0 {
				count++
			}
			left++
		}
	}

	if ansR == len(s)+1 {
		return ""
	}

	return s[ansL:ansR]
}
