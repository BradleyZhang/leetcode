// Source : https://leetcode.com/problems/replace-words
// Author : BradleyZhang
// Date   : 2026-10-15

/*****************************************************************************************************
 *
 * In English, we have a concept called root, which can be followed by some other word to form another
 * longer word - let's call this word derivative. For example, when the root "help" is followed by the
 * word "ful", we can form a derivative "helpful".
 *
 * Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces,
 * replace all the derivatives in the sentence with the root forming it. If a derivative can be
 * replaced by more than one root, replace it with the root that has the shortest length.
 *
 * Return the sentence after the replacement.
 *
 * Example 1:
 *
 * Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
 * Output: "the cat was rat by the bat"
 *
 * Example 2:
 *
 * Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
 * Output: "a a b c"
 *
 * Constraints:
 *
 * 	1 <= dictionary.length <= 1000
 * 	1 <= dictionary[i].length <= 100
 * 	dictionary[i] consists of only lower-case letters.
 * 	1 <= sentence.length <= 10^6
 * 	sentence consists of only lower-case letters and spaces.
 * 	The number of words in sentence is in the range [1, 1000]
 * 	The length of each word in sentence is in the range [1, 1000]
 * 	Every two consecutive words in sentence will be separated by exactly one space.
 * 	sentence does not have leading or trailing spaces.
 ******************************************************************************************************/

package replacewords

import "strings"

// 使用 map
func replaceWords(dictionary []string, sentence string) string {
	// Track the minimum and maximum lengths of root words in the dictionary
	min, max := len(dictionary[0]), len(dictionary[0])

	// Key optimization 1: Use a map to store root words for O(1) lookup
	dict := make(map[string]bool, len(dictionary))
	for _, root := range dictionary {
		dict[root] = true

		// Update min and max lengths
		n := len(root)
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}

	// Iterate over each word in the sentence
	words := strings.Fields(sentence)
	for idx, word := range words {
		// Key optimization 2: Skip words shorter than the minimum root length
		if len(word) < min {
			continue
		}
		// Key optimization 3: Check prefixes from min to max length
		ceil := len(word)
		if ceil > max {
			ceil = max
		}
		for i := min; i <= ceil; i++ {
			prefix := word[:i]
			if dict[prefix] {
				words[idx] = prefix
				break
			}
		}
	}
	// Return the reconstructed sentence
	return strings.Join(words, " ")
}

// 使用 Trie
func replaceWordsTrie(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	trie := Trie{}
	for _, v := range dictionary {
		trie.Insert(v)
	}
	for i, w := range words {
		if root, ok := trie.FindRoot(w); ok {
			words[i] = root
		}
	}
	return strings.Join(words, " ")
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
func (t *Trie) FindRoot(word string) (root string, ok bool) {
	curr := t
	for i, r := range word {
		if curr.last {
			return word[:i], true
		}
		i := int(r - 'a')
		if curr.next[i] != nil {
			curr = curr.next[i]
		} else {
			return "", false
		}
	}
	if curr.last {
		return word, true
	}
	// unreachable
	return "", false
}
