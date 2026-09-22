// Source : https://leetcode.com/problems/sliding-window-maximum
// Author : BradleyZhang
// Date   : 2026-09-22

/*****************************************************************************************************
 *
 * You are given an array of integers nums, there is a sliding window of size k which is moving from
 * the very left of the array to the very right. You can only see the k numbers in the window. Each
 * time the sliding window moves right by one position.
 *
 * Return the max sliding window.
 *
 * Example 1:
 *
 * Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
 * Output: [3,3,5,5,6,7]
 * Explanation:
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 *  1 [3  -1  -3] 5  3  6  7       3
 *  1  3 [-1  -3  5] 3  6  7       5
 *  1  3  -1 [-3  5  3] 6  7       5
 *  1  3  -1  -3 [5  3  6] 7       6
 *  1  3  -1  -3  5 [3  6  7]      7
 *
 * Example 2:
 *
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10^5
 * 	-10^4 <= nums[i] <= 10^4
 * 	1 <= k <= nums.length
 ******************************************************************************************************/

package slidingwindowmaximum

import "math"

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{} // 存下标
	ans := []int{}

	for i := 0; i < len(nums); i++ {

		// 1. 移除已经离开窗口的下标
		if len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}

		// 2. 从队尾删除比 nums[i] 小的元素
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		// 3. 当前下标入队
		queue = append(queue, i)

		// 4. 窗口形成后，队首就是最大值
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}

	return ans
}

var INFINITY = math.MaxInt

func maxSlidingWindowBackup(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i, _ := range ans {
		ans[i] = -INFINITY
	}
	for right, num := range nums {
		for i := 0; i < k; i++ {
			if right-i > len(nums)-k {
				continue
			}
			if right-i > -1 {
				ans[right-i] = max(ans[right-i], num)
			} else {
				break
			}
		}
	}
	return ans
}
