// Source : https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
// Author : BradleyZhang
// Date   : 2026-09-16

/*****************************************************************************************************
 *
 * Given an array of integers nums sorted in non-decreasing order, find the starting and ending
 * position of a given target value.
 *
 * If target is not found in the array, return [-1, -1].
 *
 * You must write an algorithm with O(log n) runtime complexity.
 *
 * Example 1:
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * Example 2:
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * Example 3:
 * Input: nums = [], target = 0
 * Output: [-1,-1]
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 10^5
 * 	-10^9 <= nums[i] <= 10^9
 * 	nums is a non-decreasing array.
 * 	-10^9 <= target <= 10^9
 ******************************************************************************************************/
package findfirstandlastpositionofelementinsortedarray

// 两次二分
func searchRange(nums []int, target int) []int {
	return []int{
		lowerBound(nums, target),
		upperBound(nums, target) - 1,
	}
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if l == len(nums) || nums[l] != target {
		return -1
	}

	return l
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

// 第一次找完，写法复杂
func searchRangeBackup(nums []int, target int) []int {
	ans := []int{-1, -1}
	leftL, leftR := 0, -1
	rightL, rightR := -1, len(nums)-1
	for leftL <= rightR {
		mid := leftL + (rightR-leftL)/2
		if nums[mid] == target {
			leftR = mid
			rightL = mid
			ans = []int{mid, mid}
			for leftL <= leftR {
				m := leftL + (leftR-leftL)/2
				if nums[m] != target {
					leftL = m + 1
				} else {
					ans[0] = m
					leftR = m - 1
				}
			}
			for rightL <= rightR {
				m := rightL + (rightR-rightL)/2
				if nums[m] != target {
					rightR = m - 1
				} else {
					ans[1] = m
					rightL = m + 1
				}
			}
			return ans
		} else if nums[mid] < target {
			leftL = mid + 1
		} else {
			rightR = mid - 1
		}

	}
	return ans
}
