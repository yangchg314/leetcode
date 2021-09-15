/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
func moveZeroes(nums []int) {
	start, end := 0, len(nums)
	for start < end {
		if nums[start] == 0 {
			for i := start + 1; i < end; i++ {
				nums[i-1] = nums[i]
			}
			end--
			nums[end] = 0
		} else {
			start += 1
		}
	}
}

// @lc code=end
