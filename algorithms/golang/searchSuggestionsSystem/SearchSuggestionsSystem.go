// Source : https://leetcode.com/problems/search-suggestions-system
// Author : BradleyZhang
// Date   : 2026-10-20

/*****************************************************************************************************
 *
 * You are given an array of strings products and a string searchWord.
 *
 * Design a system that suggests at most three product names from products after each character of
 * searchWord is typed. Suggested products should have common prefix with searchWord. If there are
 * more than three products with a common prefix return the three lexicographically minimums products.
 *
 * Return a list of lists of the suggested products after each character of searchWord is typed.
 *
 * Example 1:
 *
 * Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
 * Output:
 * [["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],["mouse","mous
 * epad"],["mouse","mousepad"]]
 * Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"].
 * After typing m and mo all products match and we show user ["mobile","moneypot","monitor"].
 * After typing mou, mous and mouse the system suggests ["mouse","mousepad"].
 *
 * Example 2:
 *
 * Input: products = ["havana"], searchWord = "havana"
 * Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
 * Explanation: The only word "havana" will be always suggested while typing the search word.
 *
 * Constraints:
 *
 * 	1 <= products.length <= 1000
 * 	1 <= products[i].length <= 3000
 * 	1 <= sum(products[i].length) <= 2 * 10^4
 * 	All the strings of products are unique.
 * 	products[i] consists of lowercase English letters.
 * 	1 <= searchWord.length <= 1000
 * 	searchWord consists of lowercase English letters.
 ******************************************************************************************************/

package searchsuggestionssystem

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		start := sort.SearchStrings(products, prefix)
		cur := []string{}
		for j := start; j < len(products) && len(cur) < 3; j++ {
			if !strings.HasPrefix(products[j], prefix) {
				break
			}
			cur = append(cur, products[j])
		}
		res = append(res, cur)
	}
	return res
}
func suggestedProductsTrie(products []string, searchWord string) [][]string {
	// 先排序
	sort.Strings(products)
	trie := &Trie{}
	for _, product := range products {
		trie.Insert(product)
	}
	ans := make([][]string, len(searchWord))
	for i := range searchWord {
		ans[i] = trie.firstThree(searchWord[:i+1])
	}
	return ans
}

type Trie struct {
	next  [26]*Trie
	last  bool
	words []string
}

func (t *Trie) firstThree(prefix string) []string {
	curr := t
	for _, r := range prefix {
		i := int(r - 'a')
		if curr.next[i] != nil {
			curr = curr.next[i]
		} else {
			return []string{}
		}
	}
	if len(curr.words) >= 3 {
		return curr.words[:3]
	} else {
		return curr.words
	}
}

func (t *Trie) Insert(word string) {
	curr := t
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if curr.next[index] == nil {
			curr.next[index] = &Trie{}
		}
		curr = curr.next[index]
		curr.words = append(curr.words, word)
	}
	curr.last = true
}
