/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */
package main

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	start, end := 1, n
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		x := guess(mid)
		if x == 0 {
			return mid
		} else if x > 0 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// @lc code=end
