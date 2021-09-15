/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */
package main

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := 1; i <= num/2; i++ {
		if i*i == num {
			return true
		}
		if i*i > num {
			return false
		}
	}
	return false
}

// @lc code=end
