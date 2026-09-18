// Source : https://leetcode.com/problems/longest-substring-without-repeating-characters
// Author : BradleyZhang
// Date   : 2026-09-18

/*****************************************************************************************************
 *
 * Given a string s, find the length of the longest substring without duplicate characters.
 *
 * Example 1:
 *
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct
 * answers.
 *
 * Example 2:
 *
 * Input: s = "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 * Example 3:
 *
 * Input: s = "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 *
 * Constraints:
 *
 * 	0 <= s.length <= 10^5
 * 	s consists of English letters, digits, symbols and spaces.
 ******************************************************************************************************/

package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	exist := make(map[rune]int)
	left := 0
	ans := 0

	for i, r := range s {
		if j, ok := exist[r]; ok && j >= left {
			left = j + 1
		}

		exist[r] = i

		if length := i - left + 1; length > ans {
			ans = length
		}
	}

	return ans
}
func lengthOfLongestSubstringBackup2(s string) int {
	exist := make(map[rune]int, 256)
	var ans int
	subLen := 0
	left := 0
	for i, r := range s {
		if j, ok := exist[r]; !ok || left > j {
			exist[r] = i
		} else {
			left = j + 1
			if subLen > ans {
				ans = subLen
			}
			subLen = i - j
			exist[r] = i
		}
	}
	if subLen > ans {
		return subLen
	}
	return ans
}

func lengthOfLongestSubstringBackup(s string) int {
	exist := make(map[rune]int, 256)
	var ans int
	subLen := 0
	for i, r := range s {
		if j, ok := exist[r]; !ok {
			exist[r] = i
			subLen++
		} else {
			if subLen > ans {
				ans = subLen
			}
			subLen = i - j
			exist[r] = i
			for x, k := range exist {
				if k < j {
					delete(exist, x)
				}
			}
		}
	}
	if subLen > ans {
		return subLen
	}
	return ans
}
