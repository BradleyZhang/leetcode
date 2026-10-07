// Source : https://leetcode.com/problems/word-search
// Author : BradleyZhang
// Date   : 2026-10-07

/*****************************************************************************************************
 *
 * Given an m x n grid of characters board and a string word, return true if word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are
 * horizontally or vertically neighboring. The same letter cell may not be used more than once.
 *
 * Example 1:
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
 * Output: true
 *
 * Example 2:
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
 * Output: true
 *
 * Example 3:
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
 * Output: false
 *
 * Constraints:
 *
 * 	m == board.length
 * 	n = board[i].length
 * 	1 <= m, n <= 6
 * 	1 <= word.length <= 15
 * 	board and word consists of only lowercase and uppercase English letters.
 *
 * Follow up: Could you use search pruning to make your solution faster with a larger board?
 ******************************************************************************************************/

package wordsearch

func exist(board [][]byte, word string) bool {
	var do func(i int, p, q int) bool
	do = func(i int, p, q int) bool {
		if i == len(word) {
			return true
		}
		if p < 0 || q < 0 || p > len(board)-1 || q > len(board[p])-1 {
			return false
		}
		if board[p][q] == '*' {
			return false
		}

		b := board[p][q]
		if b != word[i] {
			return false
		}
		board[p][q] = '*' // 用特殊字符优化掉 used
		if do(i+1, p, q-1) ||
			do(i+1, p, q+1) ||
			do(i+1, p+1, q) ||
			do(i+1, p-1, q) {
			return true
		}
		board[p][q] = b
		return false
	}
	for i, r := range board {
		for j := range r {
			if find := do(0, i, j); find {
				return true
			}
		}
	}
	return false
}

// // 不记忆 start index
// func exist(board [][]byte, word string) bool {
// 	start := word[0]
// 	// var startIndexes [][]int
// 	// find the first start index
// 	// for i, r := range board {
// 	// 	for j, v := range r {
// 	// 		if v == start {
// 	// 			startIndexes = [][]int{{i, j}}
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }
// 	// if len(startIndexes) == 0 {
// 	// 	return false
// 	// }
// 	used := make([][]bool, len(board))
// 	for i := range used {
// 		used[i] = make([]bool, len(board[0]))
// 	}
// 	var do func(i int, p, q int) bool
// 	do = func(i int, p, q int) bool {
// 		if i == len(word) {
// 			return true
// 		}
// 		if p < 0 || q < 0 || p > len(board)-1 || q > len(board[p])-1 {
// 			return false
// 		}
// 		if used[p][q] {
// 			return false
// 		}

// 		b := board[p][q]
// 		// if b == start {
// 		// startIndexes = append(startIndexes, []int{p, q})
// 		// }
// 		if b != word[i] {
// 			return false
// 		}
// 		used[p][q] = true
// 		if find := do(i+1, p, q-1); find {
// 			return true
// 		}
// 		if find := do(i+1, p, q+1); find {
// 			return true
// 		}
// 		if find := do(i+1, p+1, q); find {
// 			return true
// 		}
// 		if find := do(i+1, p-1, q); find {
// 			return true
// 		}
// 		used[p][q] = false
// 		return false
// 	}
// 	for i, r := range board {
// 		for j, v := range r {
// 			if v == start {
// 				if find := do(0, i, j); find {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	// if find := do(0, startIndexes[0][0], startIndexes[0][1]); find {
// 	// 	return true
// 	// }
// 	// if len(startIndexes) > 1 {
// 	// 	startIndexes = startIndexes[1:]
// 	// 	for _, v := range startIndexes {
// 	// 		// used = make([][]bool, len(board))
// 	// 		// for i := range used {
// 	// 		// 	used[i] = make([]bool, len(board[0]))
// 	// 		// }
// 	// 		if find := do(0, v[0], v[1]); find {
// 	// 			return true
// 	// 		}
// 	// 	}
// 	// }
// 	return false
// }
