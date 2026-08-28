// Source : https://leetcode.com/problems/product-of-array-except-self
// Author : BradleyZhang
// Date   : 2026-08-28

/*****************************************************************************************************
 *
 * Given an integer array nums, return an array answer such that answer[i] is equal to the product of
 * all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the division operation.
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 * Constraints:
 *
 * 	2 <= nums.length <= 10^5
 * 	-30 <= nums[i] <= 30
 * 	The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not
 * count as extra space for space complexity analysis.)
 ******************************************************************************************************/

package productofarrayexceptself

// 计算所有数乘积后再除，处理特殊情况0
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	productOfAllExceptZero := 1
	zeroIndex := -1
	for i, num := range nums {
		if num == 0 {
			if zeroIndex > -1 { //nums有两个0
				return result //全是0
			}
			zeroIndex = i
		} else {
			productOfAllExceptZero *= num
		}
	}

	if zeroIndex > -1 { // nums有一个0
		result[zeroIndex] = productOfAllExceptZero
		return result
	}

	for i, _ := range result {
		result[i] = productOfAllExceptZero
	}
	for i, num := range nums {
		result[i] /= num
	}
	return result
}

// 更优雅解法：nums 从前和从后各遍历。核心是 result[i]=Prefix Product * Suffix Product
func productExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	curr := 1
	for i, num := range nums {
		result[i] *= curr
		curr *= num
	}

	curr = 1
	for i := len(nums) - 1; i > -1; i-- {
		result[i] *= curr
		curr *= nums[i]
	}
	return result
}
