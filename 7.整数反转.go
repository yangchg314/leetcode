/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package main

import "fmt"

// @lc code=start
func reverse(x int) int {
	res := 0
	flag := true

	if x < 0 {
		x = -x
		flag = false
	}
	for x != 0 {
		res *= 10
		res += x % 10
		x = x / 10
	}
	if flag == false {
		res = -res
	}
	if (res-2147483647 > 0) || (res+2147483648 < 0) {
		return 0
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(9646324351-2147483647 > 0)
}
