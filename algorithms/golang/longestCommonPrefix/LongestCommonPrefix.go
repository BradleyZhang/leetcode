// Source : https://leetcode.com/problems/longest-common-prefix
// Author : BradleyZhang
// Date   : 2026-10-15

/*****************************************************************************************************
 *
 * Write a function to find the longest common prefix string amongst an array of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 * Example 2:
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 * Constraints:
 *
 * 	1 <= strs.length <= 200
 * 	0 <= strs[i].length <= 200
 * 	strs[i] consists of only lowercase English letters if it is non-empty.
 ******************************************************************************************************/

package longestcommonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	root := &Trie{}
	curr := root
	for _, b := range strs[0] {
		curr.next[int(b-'a')] = &Trie{}
		curr = curr.next[int(b-'a')]
	}
	var ans strings.Builder
	for i := 1; i < len(strs); i++ {
		curr := root
		ans.Reset()
		for _, b := range strs[i] {
			if curr.next[int(b-'a')] == nil || curr.last {
				curr.last = true
				break
			}
			ans.WriteRune(b)
			curr = curr.next[int(b-'a')]
		}
		curr.last = true
	}
	return ans.String()
}

type Trie struct {
	last bool
	next [26]*Trie
}
