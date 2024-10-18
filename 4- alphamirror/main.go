package main

import (
	"fmt"
	"os"
)

func main() {
	arg1 := os.Args[1]
	arr := []rune(arg1)
	for i := range arr {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			arr[i] = ('z' - arr[i] + 'a') 
			// if arr[i] = 'c'  --->  122 - 99 +97 = 120 --> 'x'
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = ('Z' - arr[i] + 'A')
		}
	}
	fmt.Println(string(arr))
}
