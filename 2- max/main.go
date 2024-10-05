package main

import "math"

func max(a []int) int {
	m := float64(a[0])
	for _, l := range a {
		m = math.Max(float64(m), float64(l))
	}
	return int(m)
}
//////////////solution without math library
func Max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
