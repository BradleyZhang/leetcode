// Source : https://leetcode.com/problems/n-queens
// Author : BradleyZhang
// Date   : 2026-10-07

/*****************************************************************************************************
 *
 * The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two
 * queens attack each other.
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer
 * in any order.
 *
 * Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.'
 * both indicate a queen and an empty space, respectively.
 *
 * Example 1:
 *
 * Input: n = 4
 * Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
 *
 * Example 2:
 *
 * Input: n = 1
 * Output: [["Q"]]
 *
 * Constraints:
 *
 * 	1 <= n <= 9
 ******************************************************************************************************/

package nqueens

import (
	"bytes"
)

func solveNQueens(n int) [][]string {
	var ans [][]string

	board := make([][]byte, n)
	for i := range board {
		board[i] = bytes.Repeat([]byte{'.'}, n)
	}

	cols := make([]bool, n)
	// 每个方向对角线有 2n-1 条。向右下的对角线上格子 row-col 一样，再因数组下标非负，加 n-1 偏移量；向左下对角线上格子 row + col 一样
	diag1 := make([]bool, 2*n-1) // row - col + n - 1
	diag2 := make([]bool, 2*n-1) // row + col

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := range board {
				temp[i] = string(board[i])
			}
			ans = append(ans, temp)
			return
		}

		for col := 0; col < n; col++ {
			d1 := row - col + n - 1
			d2 := row + col

			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			// 选择
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true
			board[row][col] = 'Q'

			dfs(row + 1)

			// 撤销
			board[row][col] = '.'
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	dfs(0)
	return ans
}

func solveNQueensBackup(n int) [][]string {
	var board []string
	var ans [][]string
	exist := [][]int{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			temp := append([]string{}, board...)
			ans = append(ans, temp)
			return
		}
		for j := 0; j < n; j++ {
			valid := true
			for _, q := range exist {
				if canAttack(i, j, q[0], q[1]) {
					valid = false
					break
				}
			}
			if valid {
				exist = append(exist, []int{i, j})
				board = append(board, line(j, n))
				dfs(i + 1)
				exist = exist[:len(exist)-1]
				board = board[:len(board)-1]
			}
		}
	}
	dfs(0)
	return ans
}
func canAttack(i, j, p, q int) bool {
	if i == p || j == q {
		return true
	}
	if (i - p) == (j - q) {
		return true
	}
	return abs(i-p) == abs(j-q)
}
func line(i, n int) string {
	row := bytes.Repeat([]byte{'.'}, n)
	row[i] = 'Q'
	return string(row)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
