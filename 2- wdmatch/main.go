package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 { // 3 because we need to have 2 args + the name of file = 3
		return
	}
	arg1, arg2, res := os.Args[1], os.Args[2], ""
	var index int
	for _, k := range arg2 {
		if index < len(arg1) && k == rune(arg1[index]) {
			res += string(k)
			index++
		}
		if index == len(arg1) {
			//if we already found all the letters..
			//no need to continue ranging over all other letters :)
			break
		}
	}
	if res != arg1 {
		return
	}
	fmt.Println(res) //if in exam no fmt; just range over the string and print the runes
}
