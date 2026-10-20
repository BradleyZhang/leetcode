// Source : https://leetcode.com/problems/number-of-islands
// Author : BradleyZhang
// Date   : 2026-10-20

/*****************************************************************************************************
 *
 * Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return
 * the number of islands.
 *
 * An island is surrounded by water and is formed by connecting adjacent lands horizontally or
 * vertically. You may assume all four edges of the grid are all surrounded by water.
 *
 * Example 1:
 *
 * Input: grid = [
 *   ["1","1","1","1","0"],
 *   ["1","1","0","1","0"],
 *   ["1","1","0","0","0"],
 *   ["0","0","0","0","0"]
 * ]
 * Output: 1
 *
 * Example 2:
 *
 * Input: grid = [
 *   ["1","1","0","0","0"],
 *   ["1","1","0","0","0"],
 *   ["0","0","1","0","0"],
 *   ["0","0","0","1","1"]
 * ]
 * Output: 3
 *
 * Constraints:
 *
 * 	m == grid.length
 * 	n == grid[i].length
 * 	1 <= m, n <= 300
 * 	grid[i][j] is '0' or '1'.
 ******************************************************************************************************/
package numberofislands

var visitedFlag = byte('*')

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = visitedFlag
		if i > 0 {
			dfs(i-1, j)
		}
		if j < n-1 {
			dfs(i, j+1)
		}
		if i < m-1 {
			dfs(i+1, j)
		}
		if j > 0 {
			dfs(i, j-1)
		}
	}
	var ans int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
