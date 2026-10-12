// Source : https://leetcode.com/problems/sudoku-solver
// Author : BradleyZhang
// Date   : 2026-10-09

/*****************************************************************************************************
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 * 	Each of the digits 1-9 must occur exactly once in each row.
 * 	Each of the digits 1-9 must occur exactly once in each column.
 * 	Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
 *
 * The '.' character indicates empty cells.
 *
 * Example 1:
 *
 * Input: board =
 * [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".
 * ",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",
 * ".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".",
 * "5"],[".",".",".",".","8",".",".","7","9"]]
 * Output:
 * [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2
 * ","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1",
 * "3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3",
 * "5"],["3","4","5","2","8","6","1","7","9"]]
 * Explanation: The input board is shown above and the only valid solution is shown below:
 *
 * Constraints:
 *
 * 	board.length == 9
 * 	board[i].length == 9
 * 	board[i][j] is a digit or '.'.
 * 	It is guaranteed that the input board has only one solution.
 ******************************************************************************************************/
package sudokusolver

// 优化
func solveSudoku(board [][]byte) {
	// 位运算 替代 map
	rows, cols, boxes := [9]int16{}, [9]int16{}, [9]int16{} // 每个int16 第n位为1代表n存在。例如 0b100 表示2 存在
	var empties [][2]int

	for i, v := range board {
		for j := range v {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				bit := int16(1) << num
				rows[i] |= bit
				cols[j] |= bit
				box := (i/3)*3 + (j / 3)
				boxes[box] |= bit
			} else {
				empties = append(empties, [2]int{i, j})
			}
		}
	}
	var dfs func(k int) bool
	dfs = func(k int) bool {
		if k == len(empties) {
			return true
		}
		best, bestCnt := k, 10
		for i := k; i < len(empties); i++ {
			r, c := empties[i][0], empties[i][1]
			used := rows[r] | cols[c] | boxes[(r/3)*3+c/3]
			cnt := 9 - popcount(int16(used))
			if cnt < bestCnt {
				best, bestCnt = i, cnt
				if cnt <= 1 {
					break
				}
			}
		}
		empties[best], empties[k] = empties[k], empties[best]

		r, c := empties[k][0], empties[k][1]
		used := rows[r] | cols[c] | boxes[(r/3)*3+c/3]
		if used == 0x3FE {
			empties[best], empties[k] = empties[k], empties[best] // 回溯
			return false
		}

		for free := ^used & 0x3FE; free != 0; free &= free - 1 {
			bit := free & -free // 提取 free 中最低位的1
			rows[r] |= bit
			cols[c] |= bit
			boxes[(r/3)*3+c/3] |= bit
			board[r][c] = byte('0' + trailingZeros(bit))
			if dfs(k + 1) {
				return true
			}
			rows[r] &^= bit
			cols[c] &^= bit
			boxes[(r/3)*3+c/3] &^= bit
			board[r][c] = '.'
		}
		empties[best], empties[k] = empties[k], empties[best] // 回溯
		return false

	}
	dfs(0)
}

func popcount(x int16) int { // 统计 1 的个数
	n := 0
	for ; x != 0; x &= x - 1 { // 消除最低位的1
		n++
	}
	return n
}
func trailingZeros(x int16) int {
	n := 0
	for x&1 == 0 {
		x >>= 1
		n++
	}
	return n
}

// 先注册已有的数字
func solveSudokuBackup(board [][]byte) {
	l := 9
	cols := make([]map[int]bool, l)
	for i := range cols {
		cols[i] = make(map[int]bool, l)
	}
	rows := make([]map[int]bool, l)
	for i := range rows {
		rows[i] = make(map[int]bool, l)
	}
	boxes := make([]map[int]bool, l)
	for i := range boxes {
		boxes[i] = make(map[int]bool, l)
	}

	var regis func(i, j, num int)
	regis = func(i, j, num int) {
		rows[i][num] = true
		cols[j][num] = true
		box := (i/3)*3 + (j / 3)
		boxes[box][num] = true
	}

	var unregis func(i, j, num int)
	unregis = func(i, j, num int) {
		rows[i][num] = false
		cols[j][num] = false
		box := (i/3)*3 + (j / 3)
		boxes[box][num] = false
	}

	var canUse func(i, j, num int) bool
	canUse = func(i, j, num int) bool {
		box := (i/3)*3 + (j / 3)
		if rows[i][num] == false &&
			cols[j][num] == false &&
			boxes[box][num] == false {
			return true
		} else {
			return false
		}
	}

	for i, v := range board {
		for j := range v {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				regis(i, j, num)
			}
		}
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == l {
			return true
		}
		if board[i][j] != '.' {
			if j == l-1 {
				return dfs(i+1, 0)
			} else {
				return dfs(i, j+1)
			}
		}

		for num := 1; num < 10; num++ {
			if canUse(i, j, num) {
				board[i][j] = byte('0' + num)
				regis(i, j, num)
				var ok bool
				if j == 8 {
					ok = dfs(i+1, 0)
				} else {
					ok = dfs(i, j+1)
				}
				if ok {
					return ok
				}
				board[i][j] = '.'
				unregis(i, j, num)

			}
		}
		return false
	}
	dfs(0, 0)
}
