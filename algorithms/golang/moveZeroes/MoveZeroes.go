// Source : https://leetcode.com/problems/move-zeroes
// Author : BradleyZhang
// Date   : 2026-09-08

/*****************************************************************************************************
 *
 * Given an integer array nums, move all 0's to the end of it while maintaining the relative order of
 * the non-zero elements.
 *
 * Note that you must do this in-place without making a copy of the array.
 *
 * Example 1:
 * Input: nums = [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 * Example 2:
 * Input: nums = [0]
 * Output: [0]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10^4
 * 	-2^31 <= nums[i] <= 2^31 - 1
 *
 * Follow up: Could you minimize the total number of operations done?
 ******************************************************************************************************/

package movezeroes

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	left, right := 0, 1
	for right < len(nums) && left < len(nums) {
		for left < len(nums) && nums[left] != 0 {
			left++
		}
		right = left + 1
		for right < len(nums) && nums[right] == 0 {
			right++
		}
		if left >= len(nums) || right >= len(nums) {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right++
	}
}

// 简洁写法
func moveZeroesBest(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			if left != right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
	}
}
