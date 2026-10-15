// Source : https://leetcode.com/problems/word-search-ii
// Author : BradleyZhang
// Date   : 2026-10-14

/*****************************************************************************************************
 *
 * Given an m x n board of characters and a list of strings words, return all words on the board.
 *
 * Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are
 * horizontally or vertically neighboring. The same letter cell may not be used more than once in a
 * word.
 *
 * Example 1:
 *
 * Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words =
 * ["oath","pea","eat","rain"]
 * Output: ["eat","oath"]
 *
 * Example 2:
 *
 * Input: board = [["a","b"],["c","d"]], words = ["abcb"]
 * Output: []
 *
 * Constraints:
 *
 * 	m == board.length
 * 	n == board[i].length
 * 	1 <= m, n <= 12
 * 	board[i][j] is a lowercase English letter.
 * 	1 <= words.length <= 3 * 10^4
 * 	1 <= words[i].length <= 10
 * 	words[i] consists of lowercase English letters.
 * 	All the strings of words are unique.
 ******************************************************************************************************/

package wordsearchii

func findWords(board [][]byte, words []string) []string {
	root := Trie{}

	w, h := len(board), len(board[0])
	visited := make(map[string]bool)
	// build trie
	// 优化: 提前过滤不可能的单词
	for i := range w {
		for j := range h - 1 {
			p := string(board[i][j]) + string(board[i][j+1])
			visited[p] = true
		}
	}
	for i := range w - 1 {
		for j := range h {
			p := string(board[i][j]) + string(board[i+1][j])
			visited[p] = true
		}
	}
	for _, w := range words {
		valid := true
		for i := range len(w) - 1 {
			a, b := string(w[i]), string(w[i+1])
			if !visited[a+b] && !visited[b+a] {
				valid = false
				break
			}
		}
		if valid {
			root.Insert(w)
		}
	}

	var ans []string
	var path []byte
	var dfs func(i, j int, trie *Trie)
	dfs = func(i, j int, trie *Trie) {

		origin := board[i][j]
		trie = trie.next[int(origin-'a')]
		if trie == nil {
			return
		}
		path = append(path, board[i][j])
		board[i][j] = '.'
		if trie.last {
			ans = append(ans, string(path))
			trie.last = false         // 防止重复
			root.Remove(string(path)) // 优化：剪枝，删除已经没有单词的子树
		}

		if i < len(board)-1 {
			r, c := i+1, j
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		if i > 0 {
			r, c := i-1, j
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		if j < len(board[0])-1 {
			r, c := i, j+1
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		if j > 0 {
			r, c := i, j-1
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		board[i][j] = origin
		path = path[:len(path)-1]
	}
	for r, row := range board {
		for c, ch := range row {
			i := int(ch - 'a')
			if root.next[i] != nil {
				dfs(r, c, &root)
			}
		}
	}
	return ans
}

func findWordsBackup(board [][]byte, words []string) []string {
	trie := Trie{}
	for _, w := range words {
		trie.Insert(w)
	}

	var ans []string
	var path []byte
	var dfs func(i, j int, trie *Trie)
	dfs = func(i, j int, trie *Trie) {

		origin := board[i][j]
		trie = trie.next[int(origin-'a')]
		if trie == nil {
			return
		}
		path = append(path, board[i][j])
		board[i][j] = '.'
		if trie.last {
			ans = append(ans, string(path))
			trie.last = false // 防止重复
		}

		if i < len(board)-1 {
			r, c := i+1, j
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		if i > 0 {
			r, c := i-1, j
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		if j < len(board[0])-1 {
			r, c := i, j+1
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		if j > 0 {
			r, c := i, j-1
			if board[r][c] != '.' {
				dfs(r, c, trie)
			}
		}
		board[i][j] = origin
		path = path[:len(path)-1]
	}
	for r, row := range board {
		for c, ch := range row {
			i := int(ch - 'a')
			if trie.next[i] != nil {
				dfs(r, c, &trie)
			}
		}
	}
	return ans
}

type Trie struct {
	last      bool
	next      [26]*Trie
	remaining int // 子树还有多少单词
}

func (this *Trie) Insert(word string) {
	curr := this
	curr.remaining++
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if curr.next[index] == nil {
			curr.next[index] = &Trie{}
		}
		curr = curr.next[index]
		curr.remaining++
	}
	curr.last = true
}
func (this *Trie) IsEmpty() bool {
	for _, n := range this.next {
		if n != nil {
			return false
		}
	}
	return true
}
func (this *Trie) Remove(existsWord string) bool {
	this.remaining--
	if len(existsWord) == 0 {
		this.last = false
		return this.remaining <= 0
	}
	c := this.next[existsWord[0]-'a']
	if c == nil {
		return false
	}
	if c.Remove(existsWord[1:]) {
		this.next[existsWord[0]-'a'] = nil
		return this.remaining <= 0
	}
	return false
}
