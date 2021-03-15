package leetcode

//47.给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列
func PermuteUnique(nums []int) [][]int {
	res := [][]int{}
	len_nums := len(nums)
	var backTrack func(int, []int)
	table := make([]int, 21)
	for _, num := range nums {
		table[num+10]++
	}
	backTrack = func(n int, permute []int) {
		if len_nums == n {
			t := make([]int, len_nums)
			copy(t, permute)
			res = append(res, t)
		}
		for i := 0; i < 21; i++ {
			if table[i] > 0 {
				table[i] -= 1
				permute = append(permute, i-10)
				backTrack(n+1, permute)
				table[i] += 1
				permute = permute[:len(permute)-1]
			}
		}
	}
	permute := make([]int, 0, len(nums))
	backTrack(0, permute)
	return res
}

/*
49.给定一个字符串数组，将字母异位词组合在一起。
字母异位词指字母相同，但排列不同的字符串。
*/
func GroupAnagrams(strs []string) [][]string {
	var byteSetEqual func([]byte, []byte) bool
	byteSetEqual = func(A []byte, B []byte) bool {
		table := make([]int, 26)
		if len(A) != len(B) {
			return false
		}
		for _, a := range A {
			table[int(a-'a')]++
		}
		for _, b := range B {
			if table[int(b-'a')] == 0 {
				return false
			}
			table[int(b-'a')] -= 1
		}
		return true
	}
	res := [][]string{}
	for _, str := range strs {
		i := 0
		for ; i < len(res); i++ {
			if byteSetEqual([]byte(str), []byte(res[i][0])) {
				res[i] = append(res[i], str)
				break
			}
		}
		if i == len(res) {
			res = append(res, []string{str})
		}
	}
	return res
}
