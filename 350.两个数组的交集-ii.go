/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */
package main

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m2 := make(map[int]int, 0)
	for i := 0; i < len(nums2); i++ {
		_, ok := m2[nums2[i]]
		if !ok {
			m2[nums2[i]] = 1
		} else {
			m2[nums2[i]] += 1
		}
	}
	for i := 0; i < len(nums1); i++ {
		_, ok := m2[nums1[i]]
		if ok && m2[nums1[i]] > 0 {
			res = append(res, nums1[i])
			m2[nums1[i]]--
		}
	}
	return res
}

// @lc code=end
