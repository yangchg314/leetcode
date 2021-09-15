/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */
package main

//异或运算满足交换律和结合律

// @lc code=start
func missingNumber(nums []int) int {
	s := 0
	for i := 1; i <= len(nums); i++ {
		s ^= i ^ nums[i-1]
	}
	return s
}

// @lc code=end
