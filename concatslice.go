package main

func ccslice(a, b []int) []int {
	var re []int
	re = append(re, a...)
	re = append(re, b...)
	return re
}
