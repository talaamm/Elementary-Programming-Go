package main

import "strings"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	res := ""
	for _, i := range str2 {
		if strings.Contains(str1, string(i)) {
			continue
		}
		res = res + string(i)
	}
	for _, i := range str1 {
		if strings.Contains(str2, string(i)) {
			continue
		}
		if strings.Contains(res, string(i)) {
			continue
		}
		res += string(i)
	}
	return len(res)
}
