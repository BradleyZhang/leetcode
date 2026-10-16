// Source : https://leetcode.com/problems/extra-characters-in-a-string
// Author : BradleyZhang
// Date   : 2026-10-16

/*****************************************************************************************************
 *
 * You are given a 0-indexed string s and a dictionary of words dictionary. You have to break s into
 * one or more non-overlapping substrings such that each substring is present in dictionary. There may
 * be some extra characters in s which are not present in any of the substrings.
 *
 * Return the minimum number of extra characters left over if you break up s optimally.
 *
 * Example 1:
 *
 * Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
 * Output: 1
 * Explanation: We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to
 * 8. There is only 1 unused character (at index 4), so we return 1.
 *
 * Example 2:
 *
 * Input: s = "sayhelloworld", dictionary = ["hello","world"]
 * Output: 3
 * Explanation: We can break s in two substrings: "hello" from index 3 to 7 and "world" from index 8
 * to 12. The characters at indices 0, 1, 2 are not used in any substring and thus are considered as
 * extra characters. Hence, we return 3.
 *
 * Constraints:
 *
 * 	1 <= s.length <= 50
 * 	1 <= dictionary.length <= 50
 * 	1 <= dictionary[i].length <= 50
 * 	dictionary[i] and s consists of only lowercase English letters
 * 	dictionary contains distinct words
 ******************************************************************************************************/

package extracharactersinastring

// Trie 解法。
// 动态规划:
// 思路：s从右往左，
// dp[i] 表示 s[i:] 的 minExtraChar
// 初始状态：dp[len(s)]=0
// 状态转移方程：dp[i] = 1 + dp[i+1] (假设从当前字符后面断开，自己追加仅多余数) ; 当找到 i 开始的子串匹配时，dp[i] = min(dp[i], dp[i+size])
// leetcode 18ms
func minExtraCharTrie(s string, dictionary []string) int {
	trie := &Trie{}
	for _, w := range dictionary {
		trie.Insert(w)
	}
	dp := make([]int, len(s)+1)
	dp[len(s)] = 0
	for i := len(s) - 1; i > -1; i-- {
		dp[i] = 1 + dp[i+1]
		curr := trie
		var size int
		for j := i; j < len(s); j++ {
			r := s[j]
			idx := int(r - 'a')
			if curr.next[idx] != nil {
				curr = curr.next[idx]
				size++
			} else {
				break
			}
			if curr.last {
				dp[i] = min(dp[i], dp[i+size])
			}
		}
	}
	return dp[0]
}

// 26ms
func minExtraCharMemo(s string, dictionary []string) int {
	wordSet := make(map[string]struct{}, len(dictionary))
	for _, w := range dictionary {
		wordSet[w] = struct{}{}
	}
	n := len(s)
	memo := make(map[int]int)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}
		if v, ok := memo[i]; ok {
			return v
		}
		ans := dfs(i+1) + 1
		for j := i; j < n; j++ {
			if _, ok := wordSet[s[i:j+1]]; ok {
				size := j + 1 - i
				ans = min(ans, dfs(i+size))
			}
		}
		memo[i] = ans
		return ans
	}
	return dfs(0)
}

// leetcode 19ms
func minExtraCharDP(s string, dictionary []string) int {
	words := make(map[string]struct{}, len(dictionary))
	for _, w := range dictionary {
		words[w] = struct{}{}
	}
	dp := make([]int, len(s)+1)
	dp[len(s)] = 0
	for i := len(s) - 1; i > -1; i-- {
		dp[i] = 1 + dp[i+1]
		for j := i; j < len(s); j++ {
			if _, ok := words[s[i:j+1]]; ok {
				size := j + 1 - i
				dp[i] = min(dp[i], dp[i+size])
			}
		}
	}
	return dp[0]
}

type Trie struct {
	next [26]*Trie
	last bool
}

func (t *Trie) Insert(word string) {
	curr := t
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if curr.next[index] == nil {
			curr.next[index] = &Trie{}
		}
		curr = curr.next[index]
	}
	curr.last = true
}
