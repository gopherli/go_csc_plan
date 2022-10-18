package jianzhioffer

import "testing"

func Test_LookForARepeatingNumber(t *testing.T) {
	nums := []int{1, 22, 3, 7, 31, 4, 22, 1, 7}
	aRepeatingNum := LookForARepeatingNumber(nums)
	t.Logf("LookForARepeatingNumber->result：%d", aRepeatingNum)
}

// https://leetcode.cn/problems/shu-zu-zahong-zhong-fu-de-shu-zi-lcof/?favorite=xb9nqhhg
func LookForARepeatingNumber(nums []int) int {
	record := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := record[num]; ok {
			return num
		}
		record[num] = true
	}
	return 0
}
