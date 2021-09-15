/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
package main

// @lc code=start
func longestCommonPrefix(strs []string) string {
	res := 0
	for len(strs) != 0 {
		if res >= len(strs[0]) {
			return strs[0][:res]
		}
		c := strs[0][res]
		for _, ss := range strs[1:] {
			if res >= len(ss) || ss[res] != c {
				return ss[:res]
			}
		}
		res++
	}
	return ""
}

// @lc code=end
