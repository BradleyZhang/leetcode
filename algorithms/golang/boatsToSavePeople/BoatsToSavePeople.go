// Source : https://leetcode.com/problems/boats-to-save-people
// Author : BradleyZhang
// Date   : 2026-09-08

/*****************************************************************************************************
 *
 * You are given an array people where people[i] is the weight of the i^th person, and an infinite
 * number of boats where each boat can carry a maximum weight of limit. Each boat carries at most two
 * people at the same time, provided the sum of the weight of those people is at most limit.
 *
 * Return the minimum number of boats to carry every given person.
 *
 * Example 1:
 *
 * Input: people = [1,2], limit = 3
 * Output: 1
 * Explanation: 1 boat (1, 2)
 *
 * Example 2:
 *
 * Input: people = [3,2,2,1], limit = 3
 * Output: 3
 * Explanation: 3 boats (1, 2), (2) and (3)
 *
 * Example 3:
 *
 * Input: people = [3,5,3,4], limit = 5
 * Output: 4
 * Explanation: 4 boats (3), (3), (4), (5)
 *
 * Constraints:
 *
 * 	1 <= people.length <= 5 * 10^4
 * 	1 <= people[i] <= limit <= 3 * 10^4
 ******************************************************************************************************/

package boatstosavepeople

import "sort"

// O(n+limit) 但难以维护
func numRescueBoats(people []int, limit int) int {
	result := 0
	weightCount := make([]int, limit)
	for _, w := range people {
		if w == limit {
			result++
			continue
		}
		weightCount[w]++
	}
	leftWeight, rightWeight := 1, len(weightCount)-1
	for leftWeight <= rightWeight {
		for leftWeight < rightWeight && weightCount[leftWeight] == 0 {
			leftWeight++
		}
		for leftWeight < rightWeight && weightCount[rightWeight] == 0 {
			rightWeight--
		}

		if leftWeight == rightWeight {
			if 2*leftWeight <= limit {
				result += (weightCount[leftWeight] + 1) / 2
			} else {
				result += weightCount[leftWeight]
			}
			break
		}
		if leftWeight+rightWeight <= limit {
			if weightCount[leftWeight] < weightCount[rightWeight] {
				result += weightCount[leftWeight]
				weightCount[rightWeight] -= weightCount[leftWeight]
				weightCount[leftWeight] = 0
				leftWeight++
			} else {
				result += weightCount[rightWeight]
				weightCount[leftWeight] -= weightCount[rightWeight]
				weightCount[rightWeight] = 0
				rightWeight--
			}
		} else {
			result += weightCount[rightWeight]
			weightCount[rightWeight] = 0
			rightWeight--
		}
	}
	return result
}

// O(n+limit)
func numRescueBoatsUseSort(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	i, j := 0, len(people)-1

	for i <= j {
		if people[i]+people[j] > limit {
			ans++
			j--
		}

		if people[i]+people[j] <= limit {
			ans++
			i++
			j--
		}

		if i == j {
			ans++
			i++
			j--
		}
	}

	return ans
}
