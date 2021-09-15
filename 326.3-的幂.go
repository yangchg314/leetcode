/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%3 != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}

// @lc code=end

