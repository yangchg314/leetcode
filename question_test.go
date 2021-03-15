package leetcode

import (
	"fmt"
	"testing"
)

func Test_PermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	res := PermuteUnique(nums)
	fmt.Println(res)
	t.Log("ok")
}
func Test_GroupAnagrams(t *testing.T) {
	strs := []string{"", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(GroupAnagrams(strs))
	t.Log("ok")
}
