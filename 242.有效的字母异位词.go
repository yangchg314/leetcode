/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
package main

import "fmt"

// @lc code=start
func isAnagram(s string, t string) bool {
	m := make(map[byte]int, 0)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		_, ok := m[t[i]]
		if !ok || m[t[i]] <= 0 {
			return false
		}
		m[t[i]]--
	}
	return true
}

// @lc code=end
func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
