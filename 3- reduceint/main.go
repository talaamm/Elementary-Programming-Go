package main

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	sum := f(a[0], a[1])
	a = a[2:]
	for i := range a {
		sum = f(sum, a[i])
	}
	fmt.Println(sum)
}
