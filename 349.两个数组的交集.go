/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
package main

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m1, m2 := make(map[int]bool, 0), make(map[int]bool, 0)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		_, ok := m2[nums2[i]]
		if !ok {
			m2[nums2[i]] = true
			_, ok2 := m1[nums2[i]]
			if ok2 {
				res = append(res, nums2[i])
			}
		}
	}
	return res

}

// @lc code=end
