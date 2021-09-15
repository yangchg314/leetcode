/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
**/
type Stack struct {
	values []*TreeNode
	size   int
	top    int
}

func NewStack() *Stack {
	return &Stack{
		make([]*TreeNode, 10), 10, -1,
	}
}
func (s *Stack) Push(node *TreeNode) {
	s.top++
	if s.top >= s.size {
		s.values = append(s.values, node)
		s.size++
		return
	}
	s.values[s.top] = node
}
func (s *Stack) Empty() bool {
	if s.top == -1 {
		return true
	}
	return false
}
func (s *Stack) Pop() *TreeNode {
	if s.Empty() {
		return nil
	}
	s.top--
	return s.values[s.top+1]
}
func (s *Stack) Top() *TreeNode {
	if s.Empty() {
		return nil
	}
	return s.values[s.top]
}
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := NewStack()
	node := root
	var pre *TreeNode
	for node != nil || !stack.Empty() {
		if node != nil {
			stack.Push(node)
			node = node.Left
		} else {
			node = stack.Top()
			if node.Right == nil || node.Right == pre {
				node = stack.Pop()
				pre = node
				res = append(res, node.Val)
				node = nil
			} else {
				node = node.Right
			}
		}
	}
	return res
}

// @lc code=end
