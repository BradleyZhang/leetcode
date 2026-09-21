// Source : https://leetcode.com/problems/longest-repeating-character-replacement
// Author : BradleyZhang
// Date   : 2026-09-21

/*****************************************************************************************************
 *
 * You are given a string s and an integer k. You can choose any character of the string and change it
 * to any other uppercase English character. You can perform this operation at most k times.
 *
 * Return the length of the longest substring containing the same letter you can get after performing
 * the above operations.
 *
 * Example 1:
 *
 * Input: s = "ABAB", k = 2
 * Output: 4
 * Explanation: Replace the two 'A's with two 'B's or vice versa.
 *
 * Example 2:
 *
 * Input: s = "AABABBA", k = 1
 * Output: 4
 * Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
 * The substring "BBBB" has the longest repeating letters, which is 4.
 * There may exists other ways to achieve this answer too.
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^5
 * 	s consists of only uppercase English letters.
 * 	0 <= k <= s.length
 ******************************************************************************************************/
package longestrepeatingcharacterreplacement

// 优化对numMap的遍历
func characterReplacement(s string, k int) int {
	var ans int
	left := 0
	numMap := make(map[byte]int)
	maxNum := 0

	for right, r := range s {
		numMap[byte(r)]++
		maxNum = max(maxNum, numMap[byte(r)])
		if othersNum := right - left + 1 - maxNum; othersNum > k {
			numMap[s[left]]--
			left++
		}

		ans = max(ans, right-left+1)
	}
	return ans
}

func characterReplacementBackup(s string, k int) int {
	var ans int
	left := 0
	numMap := make(map[byte]int)
	maxNum := 0

	for right, r := range s {
		numMap[byte(r)]++
		maxNum = max(maxNum, numMap[byte(r)])
		for othersNum := right - left + 1 - maxNum; othersNum > k; {
			numMap[s[left]]--
			for _, num := range numMap {
				maxNum = max(maxNum, num)
			}

			left++
			othersNum = right - left + 1 - maxNum
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
