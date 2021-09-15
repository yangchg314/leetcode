/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
package main

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	base := 10
	count := 1
	for x/base != 0 {
		count += 1
		base *= 10
	}
	base /= 10
	for base > 1 {
		h := x / base
		t := x % 10
		if h != t {
			return false
		}
		x = (x - h*base) / 10
		base /= 100
	}
	return true
}

// @lc code=end
