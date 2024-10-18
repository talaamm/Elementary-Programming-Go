package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	arg1, arg2, arg3 := os.Args[1], os.Args[2], os.Args[3]
	arr := []rune(arg1)
	for i, r := range arr {
		if r == rune(arg2[0]) {
			arr[i] = rune(arg3[0])
		}
		fmt.Print(string(arr[i]))
	}
	fmt.Println()
}
