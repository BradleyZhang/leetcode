// Source : https://leetcode.com/problems/permutations
// Author : BradleyZhang
// Date   : 2026-10-07

/*****************************************************************************************************
 *
 * Given an array nums of distinct integers, return all the possible permutations. You can return the
 * answer in any order.
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * Example 2:
 * Input: nums = [0,1]
 * Output: [[0,1],[1,0]]
 * Example 3:
 * Input: nums = [1]
 * Output: [[1]]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 6
 * 	-10 <= nums[i] <= 10
 * 	All the integers of nums are unique.
 ******************************************************************************************************/

package permutations

func permute(nums []int) [][]int {
	var ans [][]int
	temp := make([]int, len(nums))
	used := make(map[int]struct{}, len(nums))
	var chose func(index int)
	chose = func(index int) {
		if index == len(nums) {
			t := append([]int{}, temp...)
			ans = append(ans, t)
			return
		}
		for i := 0; i < len(nums); i++ {
			n := nums[i]
			if _, ok := used[n]; ok {
				continue
			}
			temp[index] = n
			used[n] = struct{}{}
			chose(index + 1)
			delete(used, n)
		}
	}
	chose(0)
	return ans
}
