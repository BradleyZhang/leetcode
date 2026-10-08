// Source : https://leetcode.com/problems/combination-sum-ii
// Author : BradleyZhang
// Date   : 2026-10-08

/*****************************************************************************************************
 *
 * Given a collection of candidate numbers (candidates) and a target number (target), find all unique
 * combinations in candidates where the candidate numbers sum to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note: The solution set must not contain duplicate combinations.
 *
 * Example 1:
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8
 * Output:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * Example 2:
 *
 * Input: candidates = [2,5,2,1,2], target = 5
 * Output:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 * Constraints:
 *
 * 	1 <= candidates.length <= 100
 * 	1 <= candidates[i] <= 50
 * 	1 <= target <= 30
 ******************************************************************************************************/

package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := [][]int{}
	path := []int{}
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		if sum == target {
			temp := append([]int{}, path...)
			ans = append(ans, temp)
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
