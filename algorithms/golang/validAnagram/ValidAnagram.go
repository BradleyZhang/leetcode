// Source : https://leetcode.com/problems/valid-anagram/
// Author : BradleyZhang
// Date   : 2026-08-25

/*****************************************************************************************************
 *
 * Given two strings s and t, return true if t is an anagram of s, and false otherwise.
 *
 * Example 1:
 *
 * Input: s = "anagram", t = "nagaram"
 *
 * Output: true
 *
 * Example 2:
 *
 * Input: s = "rat", t = "car"
 *
 * Output: false
 *
 * Constraints:
 *
 * 	1 <= s.length, t.length <= 5 * 10^4
 * 	s and t consist of lowercase English letters.
 *
 * Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such
 * a case?
 ******************************************************************************************************/

package validanagram

// first solution
func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterToCountS := make(map[rune]int)
	for _, l := range s {
		letterToCountS[l] += 1
	}
	letterToCountT := make(map[rune]int)
	for _, l := range t {
		letterToCountT[l] += 1
	}
	for letter, count := range letterToCountS {
		if countT, ok := letterToCountT[letter]; !ok || countT != count {
			return false
		}
		delete(letterToCountT, letter) // 删除确认过的字母
	}
	if len(letterToCountT) != 0 { // t中有s没有的字母
		return false
	}
	return true

}

// 更优解法
// 复用同一个map，s在map上加，t在map上减，最后看是否都为0
func validAnagramUseOneMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letterToCount := make(map[rune]int)
	for _, l := range s {
		letterToCount[l]++
	}
	for _, l := range t {
		letterToCount[l]--
	}
	for _, count := range letterToCount {
		if count != 0 {
			return false
		}
	}
	return true
}

// 在仅有小写字母约束下最优解，用长度26的数组/slice
func validAnagramUseArray(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make([]int, 26, 26)
	for i := 0; i < len(s); i++ {
		counts[int(s[i]-'a')]++
		counts[int(t[i]-'a')]--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
