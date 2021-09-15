/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */
package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 */
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	array := []int{}
	var PreSearch func(*TreeNode)
	PreSearch = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			s := ""
			for _, n := range array {
				s += strconv.Itoa(n) + "->"
			}
			s += strconv.Itoa(node.Val)
			res = append(res, s)
			return
		}
		array = append(array, node.Val)
		if node.Left != nil {
			PreSearch(node.Left)
		}
		if node.Right != nil {
			PreSearch(node.Right)
		}
		array = array[0 : len(array)-1]
	}
	if root == nil {
		return res
	}
	PreSearch(root)
	return res
}

// @lc code=end
