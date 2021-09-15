/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	m := make([]int, n)
	if n == 1 || n == 2 {
		return n
	}
	m[0], m[1] = 1, 2
	for i := 2; i < n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n-1]
}

// @lc code=end

