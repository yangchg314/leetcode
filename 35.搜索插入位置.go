/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		// if target<nums[start] {
		// 	return start
		// }else if target>nums[end]{
		// 	return end+1
		// }
		mid := (start + end) / 2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

// @lc code=end

