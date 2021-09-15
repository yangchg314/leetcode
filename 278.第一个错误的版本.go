/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */
package main

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start, end := 1, n
	mid := 0
	if isBadVersion(1) == true {
		return 1
	}
	for start < end {
		mid = (start + end) / 2
		if isBadVersion(mid) == false && isBadVersion(mid+1) == true {
			return mid + 1
		} else if isBadVersion(mid) == false {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}

// @lc code=end
