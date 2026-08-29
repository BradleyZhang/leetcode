// Source : https://leetcode.com/problems/longest-consecutive-sequence
// Author : BradleyZhang
// Date   : 2026-08-29

/*****************************************************************************************************
 *
 * Given an unsorted array of integers nums, return the length of the longest consecutive elements
 * sequence.
 *
 * You must write an algorithm that runs in O(n) time.
 *
 * Example 1:
 *
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
 *
 * Example 2:
 *
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 *
 * Example 3:
 *
 * Input: nums = [1,0,1,2]
 * Output: 3
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 10^5
 * 	-10^9 <= nums[i] <= 10^9
 ******************************************************************************************************/

package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return 1
	}

	numToNeighbour := make(map[int][]int, len(nums)) // num 对 pre长度，succ长度
	for _, num := range nums {
		if _, ok := numToNeighbour[num]; ok {
			continue
		}
		numToNeighbour[num] = make([]int, 2)

		pre := num - 1
		successor := num + 1
		if _, ok := numToNeighbour[pre]; ok {
			numToNeighbour[num][0]++
			numToNeighbour[pre][1]++
		}
		if _, ok := numToNeighbour[successor]; ok {
			numToNeighbour[num][1]++
			numToNeighbour[successor][0]++
		}
	}

	var result int
	for k, v := range numToNeighbour {
		length := 1
		if v[0] != 0 {
			length += getPreLen(k, numToNeighbour)
		}
		if v[1] != 0 {
			length += getSucLen(k, numToNeighbour)
		}

		if length > result {
			result = length
		}
	}
	return result
}
func getPreLen(k int, m map[int][]int) int {
	preLen := 0
	if _, ok := m[k-1]; ok {
		preLen++
		preLen += getPreLen(k-1, m)
		delete(m, k-1)
	}
	return preLen
}
func getSucLen(k int, m map[int][]int) int {
	sucLen := 0
	if _, ok := m[k+1]; ok {
		sucLen++
		sucLen += getSucLen(k+1, m)
		delete(m, k+1)
	}
	return sucLen
}

// 更优雅解法：化为set后找序列的开头，再遍历
func longestConsecutiveBest(nums []int) int {
	numsSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numsSet[num] = struct{}{}
	}
	result := 0
	for num := range numsSet {
		if _, ok := numsSet[num-1]; !ok { // num是序列开头
			y := num + 1
			for {
				if _, ok := numsSet[y]; ok {
					y++
				} else {
					break
				}
			}
			if result < y-num {
				result = y - num
			}
		}
	}
	return result
}
