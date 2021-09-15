/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// @lc code=start
func expansion(s string) []byte {
	res := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		res[2*i] = '#'
		res[2*i+1] = s[i]
	}
	res[2*len(s)] = '#'
	return res
}
func longestPalindrome(s string) string {
	t := make([]int, 2*len(s)+1)
	bs := expansion(s)
	m, r := 0, 0
	res := 0
	for i := 1; i < 2*len(s)+1; i++ {
		if i >= m+t[m] {
			r = 0
		} else {
			if i+t[2*m-i] >= m+t[m] {
				r = m + t[m] - i
			} else {
				r = t[2*m-i]
			}
		}
		for i-r-1 >= 0 && i+r+1 < 2*len(s)+1 {
			if bs[i-r-1] == bs[i+r+1] {
				r++
			} else {
				break
			}
		}
		t[i] = r
		if r > t[res] {
			res = i
		}
		if i+r > m+t[m] {
			m = i
		}
	}
	ss := make([]byte, 0, t[res])
	for i := res - t[res]; i <= res+t[res]; i++ {
		if bs[i] != '#' {
			ss = append(ss, bs[i])
		}
	}
	return string(ss)
}

// @lc code=end
func main() {
	longestPalindrome("babadada")
}
