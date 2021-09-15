/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func readWord(s string, start int) string {
	ss := []byte{}
	for i := start; i < len(s); i++ {
		if s[i] == ' ' {
			return string(ss)
		}
		ss = append(ss, s[i])
	}
	return string(ss)
}
func wordPattern(pattern string, s string) bool {
	m := make(map[string]byte)
	dic := make([]int, 28)
	start := 0
	for i := 0; i < len(pattern); i++ {
		if start >= len(s) {
			return false
		}
		t := readWord(s, start)
		start += len(t) + 1
		c, ok := m[t]
		if ok {
			if c != pattern[i] {
				return false
			}
		} else {
			m[t] = pattern[i]
		}
	}
	if start <= len(s) {
		return false
	}
	for _, values := range m {
		if dic[int(values-'a')] != 0 {
			return false
		}
		dic[int(values-'a')] += 1
	}
	return true
}

// @lc code=end

