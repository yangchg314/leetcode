/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */
package main

import "fmt"

// @lc code=start
var m = map[int]bool{
	1:                   true,
	4:                   true,
	16:                  true,
	64:                  true,
	256:                 true,
	1024:                true,
	4096:                true,
	16384:               true,
	65536:               true,
	262144:              true,
	1048576:             true,
	4194304:             true,
	16777216:            true,
	67108864:            true,
	268435456:           true,
	1073741824:          true,
	4294967296:          true,
	17179869184:         true,
	68719476736:         true,
	274877906944:        true,
	1099511627776:       true,
	4398046511104:       true,
	17592186044416:      true,
	70368744177664:      true,
	281474976710656:     true,
	1125899906842624:    true,
	4503599627370496:    true,
	18014398509481984:   true,
	72057594037927936:   true,
	288230376151711744:  true,
	1152921504606846976: true,
	4611686018427387904: true,
}

func isPowerOfFour(n int) bool {
	_, ok := m[n]
	return ok
}

// @lc code=end
func main() {
	n := 1
	for i := 0; i < 34; i++ {
		fmt.Println(n)
		n <<= 2
	}
}
