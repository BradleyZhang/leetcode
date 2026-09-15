// Source : https://leetcode.com/problems/search-a-2d-matrix
// Author : BradleyZhang
// Date   : 2026-09-15

/*****************************************************************************************************
 *
 * You are given an m x n integer matrix matrix with the following two properties:
 *
 * 	Each row is sorted in non-decreasing order.
 * 	The first integer of each row is greater than the last integer of the previous row.
 *
 * Given an integer target, return true if target is in matrix or false otherwise.
 *
 * You must write a solution in O(log(m * n)) time complexity.
 *
 * Example 1:
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 * Example 2:
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 * Constraints:
 *
 * 	m == matrix.length
 * 	n == matrix[i].length
 * 	1 <= m, n <= 100
 * 	-10^4 <= matrix[i][j], target <= 10^4
 ******************************************************************************************************/
package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	var floor = -1
	{
		up, down := 0, len(matrix)-1
		for up <= down {
			mid := up + (down-up)/2
			if matrix[mid][0] > target {
				down = mid - 1
			} else if matrix[mid][len(matrix[0])-1] < target {
				up = mid + 1
			} else {
				floor = mid
				break
			}
		}
	}
	if floor == -1 {
		return false
	}
	left, right := 0, len(matrix[floor])-1
	for left <= right {
		mid := left + (right-left)/2
		if x := matrix[floor][mid]; x == target {
			return true
		} else if x < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		x := matrix[mid/n][mid%n]
		if x == target {
			return true
		} else if x < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
