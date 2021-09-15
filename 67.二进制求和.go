/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
package main

import "fmt"

// @lc code=start
func addBinary(a string, b string) string {
	len_a, len_b := len(a), len(b)
	max_len := len_a + 1
	if len_b > len_a {
		max_len = len_b + 1
	}
	res := make([]byte, max_len)
	his := false
	i, j := 0, 0
	for i < len_a && j < len_b {
		if a[len_a-1-i] == '0' && b[len_b-1-j] == '0' {
			if his == false {
				res[max_len-i-1] = '0'
			} else {
				res[max_len-i-1] = '1'
			}
			his = false

		} else if (a[len_a-1-i] == '1' && b[len_b-1-j] == '0') ||
			(a[len_a-1-i] == '0' && b[len_b-1-j] == '1') {
			if his == false {
				res[max_len-i-1] = '1'
			} else {
				res[max_len-i-1] = '0'
				his = true
			}
		} else {
			if his == false {
				res[max_len-i-1] = '0'
			} else {
				res[max_len-i-1] = '1'
			}
			his = true
		}
		i += 1
		j += 1
	}
	for ; i < len_a; i += 1 {
		if a[len_a-i-1] == '0' && his == true {
			res[max_len-i-1] = '1'
			his = false
		} else if a[len_a-i-1] == '1' && his == true {
			res[max_len-i-1] = '0'
			his = true
		} else if a[len_a-i-1] == '0' && his == false {
			res[max_len-i-1] = '0'
		} else {
			res[max_len-i-1] = '1'
		}
	}
	for ; j < len_b; j += 1 {
		if b[len_b-j-1] == '0' && his == true {
			res[max_len-j-1] = '1'
			his = false
		} else if b[len_b-j-1] == '1' && his == true {
			res[max_len-j-1] = '0'
			his = true
		} else if b[len_b-j-1] == '0' && his == false {
			res[max_len-j-1] = '0'
		} else {
			res[max_len-j-1] = '1'
		}
	}
	if his == false {
		return string(res[1:])
	}
	res[0] = '1'
	return string(res)
}

// @lc code=end
func main() {
	fmt.Println(addBinary("11", "1"))
}
