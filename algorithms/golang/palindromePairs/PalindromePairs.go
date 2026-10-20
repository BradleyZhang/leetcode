// Source : https://leetcode.com/problems/palindrome-pairs
// Author : BradleyZhang
// Date   : 2026-10-19

/*****************************************************************************************************
 *
 * You are given a 0-indexed array of unique strings words.
 *
 * A palindrome pair is a pair of integers (i, j) such that:
 *
 * 	0 <= i, j < words.length,
 * 	i != j, and
 * 	words[i] + words[j] (the concatenation of the two strings) is a palindrome.
 *
 * Return an array of all the palindrome pairs of words.
 *
 * You must write an algorithm with O(sum of words[i].length) runtime complexity.
 *
 * Example 1:
 *
 * Input: words = ["abcd","dcba","lls","s","sssll"]
 * Output: [[0,1],[1,0],[3,2],[2,4]]
 * Explanation: The palindromes are ["abcddcba","dcbaabcd","slls","llssssll"]
 *
 * Example 2:
 *
 * Input: words = ["bat","tab","cat"]
 * Output: [[0,1],[1,0]]
 * Explanation: The palindromes are ["battab","tabbat"]
 *
 * Example 3:
 *
 * Input: words = ["a",""]
 * Output: [[0,1],[1,0]]
 * Explanation: The palindromes are ["a","a"]
 *
 * Constraints:
 *
 * 	1 <= words.length <= 5000
 * 	0 <= words[i].length <= 300
 * 	words[i] consists of lowercase English letters.
 ******************************************************************************************************/

package palindromepairs

// *****

type Trie struct {
	next [26]*Trie
	idx  int   // 完整单词在此节点结束时的下标
	pal  []int // 从此节点继续到单词末尾的部分是回文的单词下标
}

func palindromePairs(words []string) [][]int {
	root := &Trie{idx: -1}

	// 插入每个单词的反向字符。
	for idx, word := range words {
		odd, even := manacher(word)
		isPal := palindromeChecker(word, odd, even)

		node := root
		for i := len(word) - 1; i >= 0; i-- {
			if isPal(0, i+1) {
				node.pal = append(node.pal, idx)
			}

			c := word[i] - 'a'
			if node.next[c] == nil {
				node.next[c] = &Trie{idx: -1}
			}
			node = node.next[c]
		}

		node.idx = idx
		node.pal = append(node.pal, idx)
	}

	var ans [][]int

	for idx, word := range words {
		odd, even := manacher(word)
		isPal := palindromeChecker(word, odd, even)

		node := root
		j := 0

		for ; j < len(word); j++ {
			if node.idx != -1 && node.idx != idx && isPal(j, len(word)) {
				ans = append(ans, []int{idx, node.idx})
			}

			c := word[j] - 'a'
			if node.next[c] == nil {
				break
			}
			node = node.next[c]
		}

		if j == len(word) {
			for _, other := range node.pal {
				if other != idx {
					ans = append(ans, []int{idx, other})
				}
			}
		}
	}

	return ans
}

// manacher 返回奇数长度和偶数长度回文的半径数组。
func manacher(s string) (odd, even []int) {
	n := len(s)
	odd = make([]int, n)

	for i, left, right := 0, 0, -1; i < n; i++ {
		radius := 1
		if i <= right {
			mirror := left + right - i
			radius = odd[mirror]
			if right-i+1 < radius {
				radius = right - i + 1
			}
		}

		for i-radius >= 0 && i+radius < n &&
			s[i-radius] == s[i+radius] {
			radius++
		}

		odd[i] = radius
		if i+radius-1 > right {
			left, right = i-radius+1, i+radius-1
		}
	}

	even = make([]int, n)
	for i, left, right := 0, 0, -1; i < n; i++ {
		radius := 0
		if i <= right {
			mirror := left + right - i + 1
			radius = even[mirror]
			if right-i+1 < radius {
				radius = right - i + 1
			}
		}

		for i-radius-1 >= 0 && i+radius < n &&
			s[i-radius-1] == s[i+radius] {
			radius++
		}

		even[i] = radius
		if i+radius-1 > right {
			left, right = i-radius, i+radius-1
		}
	}

	return odd, even
}

// palindromeChecker 返回一个检查半开区间 [left, right) 是否回文的函数。
func palindromeChecker(s string, odd, even []int) func(left, right int) bool {
	return func(left, right int) bool {
		length := right - left
		if length <= 1 {
			return true
		}

		if length%2 == 1 {
			center := (left + right - 1) / 2
			return odd[center] >= length/2+1
		}

		center := (left + right) / 2
		return even[center] >= length/2
	}
}
