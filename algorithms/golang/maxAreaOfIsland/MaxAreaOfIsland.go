// Source : https://leetcode.com/problems/max-area-of-island
// Author : BradleyZhang
// Date   : 2026-10-23

/*****************************************************************************************************
 *
 * You are given an m x n binary matrix grid. An island is a group of 1's (representing land)
 * connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are
 * surrounded by water.
 *
 * The area of an island is the number of cells with a value 1 in the island.
 *
 * Return the maximum area of an island in grid. If there is no island, return 0.
 *
 * Example 1:
 *
 * Input: grid =
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,
 * 0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,
 * 0,0,0,0,0,0,1,1,0,0,0,0]]
 * Output: 6
 * Explanation: The answer is not 11, because the island must be connected 4-directionally.
 *
 * Example 2:
 *
 * Input: grid = [[0,0,0,0,0,0,0,0]]
 * Output: 0
 *
 * Constraints:
 *
 * 	m == grid.length
 * 	n == grid[i].length
 * 	1 <= m, n <= 50
 * 	grid[i][j] is either 0 or 1.
 ******************************************************************************************************/
package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == 1 {
				area := dfs_stack(grid, cell{i, j})
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

type cell struct {
	x int
	y int
}

// 栈代替递归
func dfs_stack(grid [][]int, source cell) (area int) {
	stack := []cell{source}
	grid[source.x][source.y] = 0
	area = 1
	m := len(grid)
	n := len(grid[0])
	for len(stack) != 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		neighbours := []cell{
			{node.x - 1, node.y},
			{node.x + 1, node.y},
			{node.x, node.y - 1},
			{node.x, node.y + 1},
		}

		for _, l := range neighbours {
			if l.x < 0 || l.x >= m || l.y < 0 || l.y >= n {
				continue
			}
			if grid[l.x][l.y] == 1 {
				area++
				grid[l.x][l.y] = 0
				stack = append(stack, l)
			}
		}
	}
	return
}

func maxAreaOfIslandDFS(grid [][]int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		ans := 1
		dir := [4][2]int{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}}
		grid[i][j] = 0
		for _, d := range dir {
			ans += dfs(d[0], d[1])
		}
		return ans
	}

	var ans int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
