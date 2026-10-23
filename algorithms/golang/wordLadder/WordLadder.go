// Source : https://leetcode.com/problems/word-ladder
// Author : BradleyZhang
// Date   : 2026-10-22

/*****************************************************************************************************
 *
 * A transformation sequence from word beginWord to word endWord using a dictionary wordList is a
 * sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
 *
 * 	Every adjacent pair of words differs by a single letter.
 * 	Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in
 * wordList.
 * 	sk == endWord
 *
 * Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in
 * the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.
 *
 * Example 1:
 *
 * Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
 * Output: 5
 * Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog",
 * which is 5 words long.
 *
 * Example 2:
 *
 * Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
 * Output: 0
 * Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation
 * sequence.
 *
 * Constraints:
 *
 * 	1 <= beginWord.length <= 10
 * 	endWord.length == beginWord.length
 * 	1 <= wordList.length <= 5000
 * 	wordList[i].length == beginWord.length
 * 	beginWord, endWord, and wordList[i] consist of lowercase English letters.
 * 	beginWord != endWord
 * 	All the words in wordList are unique.
 ******************************************************************************************************/
package wordladder

// 双向BFS+位压缩编码
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[uint64]struct{}, len(wordList))
	for _, w := range wordList {
		dict[encode(w)] = struct{}{}
	}
	bc, ec := encode(beginWord), encode(endWord)
	if _, ok := dict[ec]; !ok {
		return 0
	}
	lset := map[uint64]struct{}{bc: {}} // 准备拓展的一侧
	rset := map[uint64]struct{}{ec: {}} // rset 另一侧的当前边界端点
	delete(dict, bc)
	delete(dict, ec)
	n := len(beginWord)
	steps := 1
	for len(lset) > 0 && len(rset) > 0 {
		if len(lset) > len(rset) { // lset 维持为长度小的,优先拓展更小一侧
			lset, rset = rset, lset
		}
		set := make(map[uint64]struct{})
		for code := range lset { // BFS
			for pos := range n {
				shift := uint(pos * 5)                // 位移量
				base := code & ^(uint64(31) << shift) // 31 即 11111b ，base 为code 在 pos 位置清零后的值
				for x := uint64(1); x <= 26; x++ {
					ncode := base | x<<shift
					if ncode == code {
						continue
					}
					if _, ok := rset[ncode]; ok { // 说明rset 有单词和当前的相连
						return steps + 1
					}
					if _, ok := dict[ncode]; !ok {
						continue
					} else { // dict 找到相邻的
						delete(dict, ncode)
						set[ncode] = struct{}{}
					}
				}
			}
		}
		steps++
		lset = set
	}
	return 0
}

func encode(s string) uint64 { // 位压缩编码
	var code uint64 // 64/5=12 最长编码12个小写字母
	for i := range s {
		code |= uint64(s[i]-'a'+1) << uint(i*5) // 把字母映射到1-26 ，需要5位二进制
	}
	return code
}

// 双向BFS
func ladderLengthBFS2(beginWord string, endWord string, wordList []string) int {
	graph := make(map[string][]string)
	// 构造图
	for _, word := range wordList {
		if isAdjacent(beginWord, word) {
			graph[beginWord] = append(graph[beginWord], word)
			graph[word] = append(graph[word], beginWord)
		}
	}
	for i := 0; i < len(wordList)-1; i++ {
		a := wordList[i]
		for _, b := range wordList[i+1:] {
			if isAdjacent(a, b) {
				graph[a] = append(graph[a], b)
				graph[b] = append(graph[b], a)
			}
		}
	}

	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	visited[endWord] = struct{}{}
	lset := map[string]struct{}{beginWord: {}}
	rset := map[string]struct{}{endWord: {}}
	step := 1
	for len(lset) > 0 && len(rset) > 0 {
		if len(lset) > len(rset) {
			lset, rset = rset, lset
		}
		set := make(map[string]struct{})
		for w := range lset {
			for _, adjW := range graph[w] {
				if _, ok := rset[adjW]; ok {
					return step + 1
				}
				if _, ok := visited[adjW]; !ok {
					visited[adjW] = struct{}{}
					set[adjW] = struct{}{}
				}
			}
		}
		step++
		lset = set
	}
	return 0
}

func ladderLengthBFS(beginWord string, endWord string, wordList []string) int {
	graph := make(map[string][]string)
	for _, word := range wordList {
		if isAdjacent(beginWord, word) {
			graph[beginWord] = append(graph[beginWord], word)
			graph[word] = append(graph[word], beginWord)
		}
	}
	for i := 0; i < len(wordList)-1; i++ {
		a := wordList[i]
		for _, b := range wordList[i+1:] {
			if isAdjacent(a, b) {
				graph[a] = append(graph[a], b)
				graph[b] = append(graph[b], a)
			}
		}
	}
	queue := []string{beginWord}
	visited := make(map[string]int)
	visited[beginWord] = 1
	for len(queue) > 0 {
		vet := queue[0]
		if vet == endWord {
			return visited[vet]
		}
		queue = queue[1:]
		for _, adjVet := range graph[vet] {
			if _, ok := visited[adjVet]; !ok {
				queue = append(queue, adjVet)
				visited[adjVet] = visited[vet] + 1
			}
		}
	}
	return 0
}
func isAdjacent(a, b string) bool {
	n := len(a)
	var diff int
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			if diff > 0 {
				return false
			}
			diff++
		}
	}
	return true
}
