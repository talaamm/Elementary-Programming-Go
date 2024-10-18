package main

import "math"

// ////////////solution with builtin funcs
func Mmax(a []int) (res int) {
	for _, i := range a[1:] {
		res = max(i, res)
	}
	return
}

// ////////////solution with math library
func max1(a []int) int {
	m := float64(a[0])
	for _, l := range a {
		m = math.Max(float64(m), float64(l))
	}
	return int(m)
}

// ////////////solution without math library for exam
func Max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
