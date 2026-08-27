// Source : https://leetcode.com/problems/group-anagrams
// Author : BradleyZhang
// Date   : 2026-08-27

/*****************************************************************************************************
 *
 * Given an array of strings strs, group the anagrams together. You can return the answer in any order.
 *
 * Example 1:
 *
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 *
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * Explanation:
 *
 * 	There is no string in strs that can be rearranged to form "bat".
 * 	The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
 * 	The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each
 * other.
 *
 * Example 2:
 *
 * Input: strs = [""]
 *
 * Output: [[""]]
 *
 * Example 3:
 *
 * Input: strs = ["a"]
 *
 * Output: [["a"]]
 *
 * Constraints:
 *
 * 	1 <= strs.length <= 10^4
 * 	0 <= strs[i].length <= 100
 * 	strs[i] consists of lowercase English letters.
 ******************************************************************************************************/
package groupanagrams

import "slices"

// first solution
func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	notGrouped := strs
	var result [][]string
	for len(notGrouped) > 0 {
		var newNotGroupd []string
		group := []string{notGrouped[0]}
		for i := 1; i < len(notGrouped); i++ {
			if validAnagram(notGrouped[0], notGrouped[i]) {
				group = append(group, notGrouped[i])
			} else {
				newNotGroupd = append(newNotGroupd, notGrouped[i])
			}
		}
		result = append(result, group)
		notGrouped = newNotGroupd
	}
	return result
}

func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make([]int, 26, 26)
	for i := 0; i < len(s); i++ {
		counts[int(s[i]-'a')]++
		counts[int(t[i]-'a')]--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}

// 用sort的字符串做key
// Time complexity: O(m*nlogn)
// Space complexity: O(mn)
func groupAnagramsUseSort(strs []string) [][]string {
	sortedStrToGroup := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		sorted := string(b)
		sortedStrToGroup[sorted] = append(sortedStrToGroup[sorted], str)
	}
	var result [][]string
	for _, group := range sortedStrToGroup {
		result = append(result, group)
	}
	return result
}

// 用频率做key
// Time complexity: O(mn)
// Space complexity: O(mn)
func groupAnagrams2(strs []string) [][]string {
	charFrequenciesToGroup := make(map[[26]int][]string)
	for _, str := range strs {
		var counts [26]int
		for i := 0; i < len(str); i++ {
			counts[int(str[i]-'a')]++
		}
		charFrequenciesToGroup[counts] = append(charFrequenciesToGroup[counts], str)
	}
	var result [][]string
	for _, group := range charFrequenciesToGroup {
		result = append(result, group)
	}
	return result
}
