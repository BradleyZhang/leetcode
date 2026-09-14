// Source : https://leetcode.com/problems/search-in-rotated-sorted-array
// Author : BradleyZhang
// Date   : 2026-09-14

/*****************************************************************************************************
 *
 * There is an integer array nums sorted in ascending order (with distinct values).
 *
 * Prior to being passed to your function, nums is possibly left rotated at an unknown index k (1 <= k
 * < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0],
 * nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be left rotated by 3
 * indices and become [4,5,6,7,0,1,2].
 *
 * Given the array nums after the possible rotation and an integer target, return the index of target
 * if it is in nums, or -1 if it is not in nums.
 *
 * You must write an algorithm with O(log n) runtime complexity.
 *
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2: 4+0
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 5000
 * 	-10^4 <= nums[i] <= 10^4
 * 	All values of nums are unique.
 * 	nums is an ascending array that is possibly rotated.
 * 	-10^4 <= target <= 10^4
 ******************************************************************************************************/
package searchinrotatedsortedarray

// 1 3 0
// l 0 r 1

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	iOfMax := 0
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	pivot := left
	if pivot == 0 {
		iOfMax = len(nums) - 1
	} else {
		iOfMax = pivot - 1
	}

	if target > nums[0] {
		left = 1
		right = iOfMax
	} else if target < nums[0] {
		left = iOfMax + 1
		right = len(nums) - 1
	} else {
		return 0
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func searchBest(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] { // 左半有序
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // 右半有序
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
