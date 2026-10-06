// Source : https://leetcode.com/problems/subsets
// Author : BradleyZhang
// Date   : 2026-10-06

/*****************************************************************************************************
 *
 * Given an integer array nums of unique elements, return all possible subsets (the power set).
 *
 * The solution set must not contain duplicate subsets. Return the solution in any order.
 *
 * Example 1:
 *
 * Input: nums = [1,2,3]
 * Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 * Example 2:
 *
 * Input: nums = [0]
 * Output: [[],[0]]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10
 * 	-10 <= nums[i] <= 10
 * 	All the numbers of nums are unique.
 ******************************************************************************************************/

package subsets

func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			copyPath := append([]int{}, path...)
			result = append(result, copyPath)
			return
		}

		dfs(i + 1)
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}

	dfs(0)
	return result
}
