package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1, arg2 := os.Args[1], os.Args[2]
	index := 0
	var res string
	for _, i := range arg2 {
		if i == rune(arg1[index]) {
			res += string(i)
			if len(res) < len(arg1) {
				index++
			} else {
				break
			}
		}
	}
	if res == arg1 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
