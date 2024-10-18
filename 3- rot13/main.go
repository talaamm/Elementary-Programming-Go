package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg1 := os.Args[1]
	arr := []rune(arg1)
	for i := range arg1 {
		if (arr[i] >= 'a' && arr[i] <= 'm') || (arr[i] >= 'A' && arr[i] <= 'M') {
			arr[i] += 13
		} else if (arr[i] >= 'n' && arr[i] <= 'z') || (arr[i] >= 'N' && arr[i] <= 'Z') {
			arr[i] -= 13
		}
	}
	fmt.Println(string(arr))
}
