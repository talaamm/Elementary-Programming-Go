package main

import "fmt"

func Chunk(slice []int, size int) {
	if len(slice) < size {
		fmt.Println(slice)
		return
	}
	if size == 0 {
		fmt.Println()
		return
	}
	arr := [][]int{}
	for len(slice) > size {
		arr = append(arr, slice[:size])
		slice = slice[size:]
	}
	arr = append(arr, slice)
	fmt.Println(arr)
}
