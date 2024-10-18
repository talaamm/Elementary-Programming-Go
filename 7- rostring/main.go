package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg1 := os.Args[1]
	str := strings.Split(arg1, " ")
	modified := []string{}
	for _, p := range str {
		if p != "" {
			modified = append(modified, p)
		}
	}
	for i := 1; i < len(modified); i++ {
		fmt.Print(modified[i] + " ")
	}
	fmt.Println(modified[0])
}
