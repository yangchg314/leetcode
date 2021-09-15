/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
// @lc code=start
func myAtoi(s string) int {
	for s[0] == ' ' {
		s = s[1:]
	}
	l := 0
	if s[0] == '-' {
		l = 1
	}
	for l < len(s) {
		if s[l] >= '0' && s[l] <= '9' {
			l++
		} else {
			break
		}
	}
	n, _ := strconv.Atoi(s[:l])
	return n
}

// @lc code=end

