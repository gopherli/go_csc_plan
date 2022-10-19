package jianzhioffer

import (
	"testing"
)

func Test_ReplaceSpaces(t *testing.T) {
	str := "0"
	rpStr := "%20"
	ways := 2
	for way := 1; way <= ways; way++ {
		result := ReplaceSpacesHandle(str, rpStr, way)
		t.Logf("ReplaceSpaces->result：%v", result)
	}
}

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/?favorite=xb9nqhhg
func ReplaceSpacesHandle(str, rpStr string, way int) string {
	switch way {
	case 1:
		return ReplaceSpacesWay01(str, rpStr)
	case 2:
		return ReplaceSpacesWay02(str, rpStr)
	}
	return ""
}

func ReplaceSpacesWay01(str, rpStr string) string {
	replaceStr := ""
	for i := range str {
		switch string(str[i]) == " " {
		case true:
			replaceStr += rpStr
		case false:
			replaceStr += string(str[i])
		}
	}
	return replaceStr
}

// NO
func ReplaceSpacesWay02(str, rpStr string) string {
	emptyCount := 0
	for x := range str {
		if string(str[x]) == " " {
			emptyCount++
		}
	}
	if emptyCount == len(str) || emptyCount == 0 {
		return ReplaceSpacesWay01(str, rpStr)
	}
	emptyStr := make([]byte, 2)
	for i := range str {
		if string(str[i]) == " " {
			str = str[:i] + string(str[i]) + string(emptyStr) + str[i+1:]
		}
	}
	strRune := []rune(str)
	for j := 0; j < len(strRune); j++ {
		if string(str[j]) == " " {
			strRune[j] = '%'
			strRune[j+1] = '2'
			strRune[j+2] = '0'
			j += 2
		}
	}
	return string(strRune)
}
