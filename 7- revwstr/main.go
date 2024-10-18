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
	ss := strings.Split(arg1, " ")
	for i := len(ss) - 1; i > 0; i-- {
		fmt.Print(ss[i] + " ")
	}
	fmt.Println(ss[0])
}
