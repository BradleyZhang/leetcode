// Source : https://leetcode.com/problems/rotting-oranges
// Author : BradleyZhang
// Date   : 2026-10-23

/*****************************************************************************************************
 *
 * You are given an m x n grid where each cell can have one of three values:
 *
 * 	0 representing an empty cell,
 * 	1 representing a fresh orange, or
 * 	2 representing a rotten orange.
 *
 * Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
 *
 * Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is
 * impossible, return -1.
 *
 * Example 1:
 *
 * Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
 * Output: 4
 *
 * Example 2:
 *
 * Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
 * Output: -1
 * Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because
 * rotting only happens 4-directionally.
 *
 * Example 3:
 *
 * Input: grid = [[0,2]]
 * Output: 0
 * Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
 *
 * Constraints:
 *
 * 	m == grid.length
 * 	n == grid[i].length
 * 	1 <= m, n <= 10
 * 	grid[i][j] is 0, 1, or 2.
 ******************************************************************************************************/

package rottingoranges

func orangesRotting(grid [][]int) int {
	result := 0
	visited := make([][]bool, len(grid))
	for i, g := range grid {
		visited[i] = make([]bool, len(g))
	}

	q := [][3]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				q = append(q, [3]int{i, j, 1})
				visited[i][j] = true
			}
		}
	}
	valid := func(i, j int) bool {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || visited[i][j] || grid[i][j] == 0 {
			return false
		}
		return true
	}

	for len(q) > 0 {
		i, j, c := q[0][0], q[0][1], q[0][2]
		q = q[1:]

		if valid(i+1, j) {
			result = max(result, c)
			visited[i+1][j] = true
			q = append(q, [3]int{i + 1, j, c + 1})
		}
		if valid(i-1, j) {
			result = max(result, c)
			visited[i-1][j] = true
			q = append(q, [3]int{i - 1, j, c + 1})
		}
		if valid(i, j+1) {
			result = max(result, c)
			visited[i][j+1] = true
			q = append(q, [3]int{i, j + 1, c + 1})
		}
		if valid(i, j-1) {
			result = max(result, c)
			visited[i][j-1] = true
			q = append(q, [3]int{i, j - 1, c + 1})
		}

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				return -1
			}
		}
	}
	return result
}

const (
	empty = iota
	fresh
	rotten
)

// 基于map，时间复杂度一致，但更慢
func orangesRottingMap(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := -1
	minGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		minGrid[i] = make([]int, n)
	}
	set := make(map[[2]int]struct{})
	var freshNum int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case rotten:
				set[[2]int{i, j}] = struct{}{}
			case fresh:
				freshNum++
			}
		}
	}
	if freshNum == 0 {
		return 0
	}

	for len(set) > 0 {
		ans++
		newSet := make(map[[2]int]struct{})
		for c := range set {
			i, j := c[0], c[1]
			if grid[i][j] == fresh {
				freshNum--
				grid[i][j] = rotten
			}
			pos := [][2]int{{i + 1, j}, {i - 1, j}, {i, j - 1}, {i, j + 1}}
			for _, p := range pos {
				i, j := p[0], p[1]
				if i > -1 && i < m && j > -1 && j < n {
					if _, ok := set[[2]int{i, j}]; !ok {
						if grid[i][j] == fresh {
							newSet[[2]int{i, j}] = struct{}{}
						}
					}
				}
			}

		}
		set = newSet
	}
	if freshNum > 0 {
		return -1
	}
	return ans
}
