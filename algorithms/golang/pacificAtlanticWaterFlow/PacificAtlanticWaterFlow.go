// Source : https://leetcode.com/problems/pacific-atlantic-water-flow/
// Author : BradleyZhang
// Date   : 2026-10-21

/*****************************************************************************************************
 *
 * There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The
 * Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's
 * right and bottom edges.
 *
 * The island is partitioned into a grid of square cells. You are given an m x n integer matrix
 * heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
 *
 * The island receives a lot of rain, and the rain water can flow to neighboring cells directly north,
 * south, east, and west if the neighboring cell's height is less than or equal to the current cell's
 * height. Water can flow from any cell adjacent to an ocean into the ocean.
 *
 * Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can
 * flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
 *
 * Example 1:
 *
 * Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
 * Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
 * Explanation: The following cells can flow to the Pacific and Atlantic oceans, as shown below:
 * [0,4]: [0,4] -> Pacific Ocean
 *        [0,4] -> Atlantic Ocean
 * [1,3]: [1,3] -> [0,3] -> Pacific Ocean
 *        [1,3] -> [1,4] -> Atlantic Ocean
 * [1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean
 *        [1,4] -> Atlantic Ocean
 * [2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean
 *        [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
 * [3,0]: [3,0] -> Pacific Ocean
 *        [3,0] -> [4,0] -> Atlantic Ocean
 * [3,1]: [3,1] -> [3,0] -> Pacific Ocean
 *        [3,1] -> [4,1] -> Atlantic Ocean
 * [4,0]: [4,0] -> Pacific Ocean
 *        [4,0] -> Atlantic Ocean
 * Note that there are other possible paths for these cells to flow to the Pacific and Atlantic oceans.
 *
 * Example 2:
 *
 * Input: heights = [[1]]
 * Output: [[0,0]]
 * Explanation: The water can flow from the only cell to the Pacific and Atlantic oceans.
 *
 * Constraints:
 *
 * 	m == heights.length
 * 	n == heights[r].length
 * 	1 <= m, n <= 200
 * 	0 <= heights[r][c] <= 10^5
 ******************************************************************************************************/
package pacificatlanticwaterflow

func pacificAtlantic(heights [][]int) [][]int {
	var ans [][]int
	pacific := make(map[[2]int]struct{})
	atlantic := make(map[[2]int]struct{})
	var dfs func(i, j, old int, visited map[[2]int]struct{})
	dfs = func(i, j, old int, visited map[[2]int]struct{}) {
		if i > len(heights)-1 || i < 0 || j > len(heights[0])-1 || j < 0 {
			return
		}
		if old > heights[i][j] {
			return
		}
		coordinate := [2]int{i, j}
		if _, ok := visited[coordinate]; ok {
			return
		}
		visited[coordinate] = struct{}{}
		myHeight := heights[i][j]
		dfs(i+1, j, myHeight, visited)
		dfs(i-1, j, myHeight, visited)
		dfs(i, j+1, myHeight, visited)
		dfs(i, j-1, myHeight, visited)
	}
	for j := 0; j < len(heights[0]); j++ {
		dfs(0, j, heights[0][j], pacific)
	}
	for i := 0; i < len(heights); i++ {
		dfs(i, 0, heights[i][0], pacific)
	}
	for j := 0; j < len(heights[0]); j++ {
		i := len(heights) - 1
		dfs(i, j, heights[i][j], atlantic)
	}
	for i := 0; i < len(heights); i++ {
		j := len(heights[0]) - 1
		dfs(i, j, heights[i][j], atlantic)
	}
	for coordinate := range pacific {
		if _, ok := atlantic[coordinate]; ok {
			ans = append(ans, coordinate[:])
		}
	}
	return ans

}

type Queue struct {
	q [][]int
}

func (q *Queue) Push(point []int) {
	q.q = append(q.q, point)
}
func (q *Queue) Pop() ([]int, bool) {
	if len(q.q) == 0 {
		return []int{}, false
	}
	ans := q.q[0]
	q.q = q.q[1:]
	return ans, true
}
