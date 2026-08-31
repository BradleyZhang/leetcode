// Source : https://leetcode.com/problems/valid-palindrome
// Author : BradleyZhang
// Date   : 2026-08-31

/*****************************************************************************************************
 *
 * A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
 * removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric
 * characters include letters and numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 *
 * Example 1:
 *
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 * Example 2:
 *
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 *
 * Example 3:
 *
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric characters.
 * Since an empty string reads the same forward and backward, it is a palindrome.
 *
 * Constraints:
 *
 * 	1 <= s.length <= 2 * 10^5
 * 	s consists only of printable ASCII characters.
 ******************************************************************************************************/

package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	s = normalize(s)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

// 瓶颈在规范化了
func normalize(s string) string {
	s = strings.ToLower(s)
	result := ""
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			result += string(c)
		}
	}
	return result
}

// 最优写法
func isPalindromeBest(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && skip(s[left]) {
			left++
		}
		for left < right && skip(s[right]) {
			right--
		}
		if left >= right {
			return true
		}
		if sanitize(s[left]) == sanitize(s[right]) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func skip(char byte) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
		return false
	}
	return true
}

func sanitize(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return (char - 'A') + 'a'
	}
	return char
}
