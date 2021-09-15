/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
package main

// @lc code=start
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s); {
		c := s[i]
		if i == len(s)-1 {
			res += m[c]
			i++
		} else if c == 'I' && s[i+1] == 'V' {
			res += 4
			i += 2
		} else if c == 'I' && s[i+1] == 'X' {
			res += 9
			i += 2
		} else if c == 'X' && s[i+1] == 'L' {
			res += 40
			i += 2
		} else if c == 'X' && s[i+1] == 'C' {
			res += 90
			i += 2
		} else if c == 'C' && s[i+1] == 'D' {
			res += 400
			i += 2
		} else if c == 'C' && s[i+1] == 'M' {
			res += 900
			i += 2
		} else {
			res += m[c]
			i++
		}
	}
	return res
}

// @lc code=end
