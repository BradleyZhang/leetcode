// Source : https://leetcode.com/problems/sum-of-prefix-scores-of-strings
// Author : BradleyZhang
// Date   : 2026-10-17

/*****************************************************************************************************
 *
 * You are given an array words of size n consisting of non-empty strings.
 *
 * We define the score of a string term as the number of strings words[i] such that term is a prefix
 * of words[i].
 *
 * 	For example, if words = ["a", "ab", "abc", "cab"], then the score of "ab" is 2, since "ab"
 * is a prefix of both "ab" and "abc".
 *
 * Return an array answer of size n where answer[i] is the sum of scores of every non-empty prefix of
 * words[i].
 *
 * Note that a string is considered as a prefix of itself.
 *
 * Example 1:
 *
 * Input: words = ["abc","ab","bc","b"]
 * Output: [5,4,3,2]
 * Explanation: The answer for each string is the following:
 * - "abc" has 3 prefixes: "a", "ab", and "abc".
 * - There are 2 strings with the prefix "a", 2 strings with the prefix "ab", and 1 string with the
 * prefix "abc".
 * The total is answer[0] = 2 + 2 + 1 = 5.
 * - "ab" has 2 prefixes: "a" and "ab".
 * - There are 2 strings with the prefix "a", and 2 strings with the prefix "ab".
 * The total is answer[1] = 2 + 2 = 4.
 * - "bc" has 2 prefixes: "b" and "bc".
 * - There are 2 strings with the prefix "b", and 1 string with the prefix "bc".
 * The total is answer[2] = 2 + 1 = 3.
 * - "b" has 1 prefix: "b".
 * - There are 2 strings with the prefix "b".
 * The total is answer[3] = 2.
 *
 * Example 2:
 *
 * Input: words = ["abcd"]
 * Output: [4]
 * Explanation:
 * "abcd" has 4 prefixes: "a", "ab", "abc", and "abcd".
 * Each prefix has a score of one, so the total is answer[0] = 1 + 1 + 1 + 1 = 4.
 *
 * Constraints:
 *
 * 	1 <= words.length <= 1000
 * 	1 <= words[i].length <= 1000
 * 	words[i] consists of lowercase English letters.
 ******************************************************************************************************/

package sumofprefixscoresofstrings

import (
	"slices"
	"strings"
)

var idxs [1000]uint16
var prefix [1000]uint16
var result [1000]int

// 最优解
func sumPrefixScores(words []string) []int {
	n := uint16(len(words))
	for i := range n {
		idxs[i] = i
	}
	// 按字典序排序
	slices.SortFunc(idxs[:n], func(a, b uint16) int {
		return strings.Compare(words[a], words[b])
	})
	// LCP(Longest Common Prefix)
	for i := range n - 1 {
		prev, curr := words[idxs[i]], words[idxs[i+1]]
		size, m := uint16(0), uint16(min(len(prev), len(curr)))
		for size < m && prev[size] == curr[size] {
			size++
		}
		prefix[i+1] = size
	}
	// 任意两个字符串的最长公共前缀长度，等于它们之间所有相邻字符串的 LCP 长度的最小值。
	// LCP(i, j)
	clear(result[:n])
	for i := range n {
		wi := idxs[i]
		size := len(words[wi])
		total := size // 自己的所有前缀对自己的score
		for j := i + 1; j < n; j++ {
			size = min(size, int(prefix[j])) // 与word[idx[j]] 的公共前缀word[:i+1] 对word[idx[j]]的score
			total += size
			result[idxs[j]] += size // 在word[idx[j]] 该公共前缀对自己也有score
		}
		result[wi] += total
	}
	return result[:n]
}

// appe  apple applicatoin

// 用Trie
func sumPrefixScoresTrie(words []string) []int {
	var trie Trie
	for _, w := range words {
		trie.Insert(w)
	}
	ans := make([]int, len(words))
	// 优化1: 每个word只走一次 Trie
	for idx, w := range words {
		var i int
		curr := &trie
		for ; i < len(w); i++ {
			j := int(w[i] - 'a')
			curr = curr.next[j]
			ans[idx] += curr.remaining
			if curr.remaining == 1 {
				i++
				break
			}
		}
		if i < len(w) {
			for ; i < len(w); i++ {
				ans[idx] += 1
			}
		}
	}
	return ans
}

func sumPrefixScoresBackup(words []string) []int {
	var trie Trie
	for _, w := range words {
		trie.Insert(w)
	}
	memo := make(map[string]int)
	ans := make([]int, len(words))
	for idx, w := range words {
		var i int
		for ; i < len(w); i++ {
			prefix := w[:i+1]
			if score, ok := memo[prefix]; ok {
				ans[idx] += score
			} else {
				score := trie.ScoreOf(prefix)
				ans[idx] += score
				memo[prefix] = score
				if score == 1 { // 剪枝，更长的score都为0
					i++
					break
				}
			}
		}
		if i < len(w) {
			for ; i < len(w); i++ {
				ans[idx] += 1
				prefix := w[:i+1]
				memo[prefix] = 1
			}
		}
	}
	return ans
}

type Trie struct {
	next      [26]*Trie
	last      bool
	remaining int
}

func (t *Trie) ScoreOf(word string) int { // 优化1： 用 remaining 保存节点的score
	curr := t
	for _, r := range word {
		i := int(r - 'a')
		if curr.next[i] != nil {
			curr = curr.next[i]
		} else {
			return 0
		}
	}
	return curr.remaining
}
func (t *Trie) ScoreOfBackup(word string) int {
	var ans int
	curr := t
	for _, r := range word {
		i := int(r - 'a')
		if curr.next[i] != nil {
			curr = curr.next[i]
		} else {
			return 0
		}
	}
	var wordNum func(node *Trie) int
	wordNum = func(node *Trie) int {
		var num int
		if node.last {
			num++
		}
		for _, n := range node.next {
			if n != nil {
				num += wordNum(n)
			}
		}
		return num
	}
	ans += wordNum(curr)
	return ans
}

func (t *Trie) Insert(word string) {
	curr := t
	for i := 0; i < len(word); i++ {
		curr.remaining++
		index := int(word[i] - 'a')
		if curr.next[index] == nil {
			curr.next[index] = &Trie{}
		}
		curr = curr.next[index]
	}
	curr.remaining++
	curr.last = true
}
