// Source : https://leetcode.com/problems/median-of-two-sorted-arrays
// Author : BradleyZhang
// Date   : 2026-09-15

/*****************************************************************************************************
 *
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two
 * sorted arrays.
 *
 * The overall run time complexity should be O(log (m+n)).
 *
 * Example 1:
 *
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 *
 * Example 2:
 *
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 *
 * Constraints:
 *
 * 	nums1.length == m
 * 	nums2.length == n
 * 	0 <= m <= 1000
 * 	0 <= n <= 1000
 * 	1 <= m + n <= 2000
 * 	-10^6 <= nums1[i], nums2[i] <= 10^6
 ******************************************************************************************************/

package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxX := math.MinInt64
		if partitionX > 0 {
			maxX = nums1[partitionX-1]
		}

		minX := math.MaxInt64
		if partitionX < m {
			minX = nums1[partitionX]
		}

		maxY := math.MinInt64
		if partitionY > 0 {
			maxY = nums2[partitionY-1]
		}

		minY := math.MaxInt64
		if partitionY < n {
			minY = nums2[partitionY]
		}

		if maxX <= minY && maxY <= minX {
			if (m+n)%2 == 0 {
				return (float64(max(maxX, maxY)) + float64(min(minX, minY))) / 2.0
			}
			return float64(max(maxX, maxY))
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}
