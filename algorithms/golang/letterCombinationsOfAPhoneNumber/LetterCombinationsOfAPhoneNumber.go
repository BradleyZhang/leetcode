// Source : https://leetcode.com/problems/letter-combinations-of-a-phone-number
// Author : BradleyZhang
// Date   : 2026-10-08

/*****************************************************************************************************
 *
 * Given a string containing digits from 2-9 inclusive, return all possible letter combinations that
 * the number could represent. Return the answer in any order.
 *
 * A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1
 * does not map to any letters.
 *
 * Example 1:
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 * Example 2:
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 * Constraints:
 *
 * 	1 <= digits.length <= 4
 * 	digits[i] is a digit in the range ['2', '9'].
 ******************************************************************************************************/
package lettercombinationsofaphonenumber

var i2l = map[int][]byte{
	2: {'a', 'b', 'c'},
	3: {'d', 'e', 'f'},
	4: {'g', 'h', 'i'},
	5: {'j', 'k', 'l'},
	6: {'m', 'n', 'o'},
	7: {'p', 'q', 'r', 's'},
	8: {'t', 'u', 'v'},
	9: {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	s := make([]byte, len(digits))
	var ans []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(s))
			return
		}
		n := int(digits[i] - '0')
		for _, v := range i2l[n] {
			s[i] = v
			dfs(i + 1)
		}
	}
	dfs(0)
	return ans
}
