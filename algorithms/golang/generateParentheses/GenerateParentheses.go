// Source : https://leetcode.com/problems/generate-parentheses
// Author : BradleyZhang
// Date   : 2026-09-09

/*****************************************************************************************************
 *
 * Given n pairs of parentheses, write a function to generate all combinations of well-formed
 * parentheses.
 *
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 *
 * Constraints:
 *
 * 	1 <= n <= 8
 ******************************************************************************************************/
package generateparentheses

// 根据有效括号的约束，探索搜索空间 （剪枝）
// 约束：left即左括号数量小于n、右括号数量不大于左括号
// 有效结果(搜索空间中的有效搜索结果）：len(s) 为 n*2
func generateParenthesis(n int) []string {
	result := []string{}
	var dfs func(left, right int, s string) //Depth-First Search 深度优先搜索
	dfs = func(left, right int, s string) {
		if len(s) == n*2 {
			result = append(result, s)
			return
		}
		if left < n {
			dfs(left+1, right, s+"(")
		}
		if right < left {
			dfs(left, right+1, s+")")
		}
	}
	dfs(0, 0, "")
	return result
}
