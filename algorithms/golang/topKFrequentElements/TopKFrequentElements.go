// Source : https://leetcode.com/problems/top-k-frequent-elements
// Author : BradleyZhang
// Date   : 2026-08-27

/*****************************************************************************************************
 *
 * Given an integer array nums and an integer k, return the k most frequent elements. You may return
 * the answer in any order.
 *
 * Example 1:
 *
 * Input: nums = [1,1,1,2,2,3], k = 2
 *
 * Output: [1,2]
 *
 * Example 2:
 *
 * Input: nums = [1], k = 1
 *
 * Output: [1]
 *
 * Example 3:
 *
 * Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2
 *
 * Output: [1,2]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10^5
 * 	-10^4 <= nums[i] <= 10^4
 * 	k is in the range [1, the number of unique elements in the array].
 * 	It is guaranteed that the answer is unique.
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's
 * size.
 ******************************************************************************************************/

package topkfrequentelements

// 使用桶排序，
// Time complexity: O(n)
// Space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	numToFreq := make(map[int]int)
	for _, num := range nums {
		numToFreq[num]++
	}
	freqBucket := make([][]int, len(nums)+1)
	for num, freq := range numToFreq {
		freqBucket[freq] = append(freqBucket[freq], num)
	}
	var result []int
	for i := len(freqBucket) - 1; i > -1; i-- {
		k -= len(freqBucket[i])
		if k < 0 {
			break
		}

		result = append(result, freqBucket[i]...)
	}
	return result
}
