/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	table := make([]bool, 811)
	tmp := func(val int) int {
		res := 0
		for val != 0 {
			res += (val % 10) * (val % 10)
			val = val / 10
		}
		return res
	}
	n = tmp(n)
	for n != 1 && table[n] == false {
		table[n] = true
		n = tmp(n)
	}
	if n == 1 {
		return true
	}
	return false
}

// @lc code=end

