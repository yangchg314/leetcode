/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */
package main

// @lc code=start
func addDigits(num int) int {
	if num/10 == 0 {
		return num
	}
	s := num % 10
	for num/10 != 0 {
		num /= 10
		s += num % 10
	}
	return addDigits(s)
}

// @lc code=end
