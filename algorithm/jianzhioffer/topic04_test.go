package jianzhioffer

import (
	"fmt"
	"testing"
)

func Test_FindANumIn2DArray(t *testing.T) {
	array := [][]int{
		{1, 2, 3},
		{2, 3, 4},
	}
	result := FindANumIn2DArray(array, 4)
	fmt.Println("FindANumIn2DArray->result：", result)
}

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/?favorite=xb9nqhhg
func FindANumIn2DArray(array [][]int, num int) bool {
	if len(array) == 0 {
		return false
	}
	row := len(array) - 1
	col := len(array[0]) - 1
	for i, j := 0, col; i <= row && j >= 0; {
		if num == array[i][j] {
			return true
		} else if num > array[i][j] {
			i++
		} else {
			j--
		}
	}
	return false
}
