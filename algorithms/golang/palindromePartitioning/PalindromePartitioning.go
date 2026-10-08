// Source : https://leetcode.com/problems/palindrome-partitioning
// Author : BradleyZhang
// Date   : 2026-10-08

/*****************************************************************************************************
 *
 * Given a string s, partition s such that every substring of the partition is a palindrome. Return
 * all possible palindrome partitioning of s.
 *
 * Example 1:
 * Input: s = "aab"
 * Output: [["a","a","b"],["aa","b"]]
 * Example 2:
 * Input: s = "a"
 * Output: [["a"]]
 *
 * Constraints:
 *
 * 	1 <= s.length <= 16
 * 	s contains only lowercase English letters.
 ******************************************************************************************************/

package palindromepartitioning

// 通常解法
// 0-start 已经切好，判断start+i能不能切
func partition(s string) [][]string {
	var ans [][]string
	path := []string{}

	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			ans = append(ans, append([]string{}, path...))
			return
		}

		for i := start; i < len(s); i++ {
			chip := s[start : i+1]

			if !isPalindrome(chip) {
				continue
			}

			path = append(path, chip)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}

// 判断当前切不切
func partitionBackup(s string) [][]string {
	ans := [][]string{}
	path := []string{}
	last := 0
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			chip := s[last:i]
			if isPalindrome(chip) {
				path = append(path, chip)
				ans = append(ans, append([]string{}, path...))
				path = path[:len(path)-1]
			}
			return
		}
		dfs(i + 1)
		origin := last
		chip := s[last:i]
		if isPalindrome(chip) {
			path = append(path, chip)
			last = i
			dfs(i + 1)
			path = path[:len(path)-1]
			last = origin
		}
	}
	dfs(1)
	return ans
}
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
