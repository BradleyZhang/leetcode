// Source : https://leetcode.com/problems/subsets-ii
// Author : BradleyZhang
// Date   : 2026-10-08

/*****************************************************************************************************
 *
 * Given an integer array nums that may contain duplicates, return all possible subsets (the power
 * set).
 *
 * The solution set must not contain duplicate subsets. Return the solution in any order.
 *
 * Example 1:
 * Input: nums = [1,2,2]
 * Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
 * Example 2:
 * Input: nums = [0]
 * Output: [[],[0]]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10
 * 	-10 <= nums[i] <= 10
 ******************************************************************************************************/
package subsetsii

import "sort"

// 通常解法
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	path := []int{}

	var dfs func(start int)
	dfs = func(start int) {
		result = append(result, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			// 同一层，跳过重复元素
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}

// path 中 n 选择不出现(不选的分支），则后面的都必须不出现
func subsetsWithDupBackup(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	sign := make(map[int]struct{})

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			copyPath := append([]int{}, path...)
			result = append(result, copyPath)
			return
		}
		if _, ok := sign[nums[i]]; ok {
			dfs(i + 1)
		} else {
			sign[nums[i]] = struct{}{}
			dfs(i + 1)
			delete(sign, nums[i])
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}
