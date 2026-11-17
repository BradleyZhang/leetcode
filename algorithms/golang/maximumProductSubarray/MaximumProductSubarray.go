// Source : https://leetcode.com/problems/maximum-product-subarray
// Author : BradleyZhang
// Date   : 2026-11-17

/*****************************************************************************************************
 *
 * Given an integer array nums, find a subarray that has the largest product, and return the product.
 *
 * The test cases are generated so that the answer will fit in a 32-bit integer.
 *
 * Note that the product of an array with a single element is the value of that element.
 *
 * ​​​​​​​Example 1:
 *
 * Input: nums = [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 * Example 2:
 *
 * Input: nums = [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 2 * 10^4
 * 	-10 <= nums[i] <= 10
 * 	The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
 ******************************************************************************************************/

package maximumproductsubarray

func maxProduct(nums []int) int {
	prefix, suffix := 1, 1
	maxi := 0
	if len(nums) == 1 {
		return nums[0]
	}
	for i := range nums {
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
		prefix *= nums[i]

		suffix *= nums[len(nums)-i-1]

		maxi = max(maxi, max(prefix, suffix))
	}
	return maxi
}

// maxDp minDp
func maxProductDp(nums []int) int {
	maxDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	minDp := make([]int, len(nums))
	minDp[0] = nums[0]
	ans := maxDp[0]

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		maxDp[i] = max(x, x*maxDp[i-1], x*minDp[i-1])
		minDp[i] = min(x, x*maxDp[i-1], x*minDp[i-1])
		ans = max(ans, maxDp[i])
	}
	return ans
}
