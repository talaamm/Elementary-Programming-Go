package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	arg1 := os.Args[1]
	conv1, e := strconv.Atoi(arg1)
	if e != nil {
		return
	}
	sum := 1
	var res string
	for i := 1; i <= 9; i++ {
		sum = i * conv1
		con2 := strconv.Itoa(i)
		con3 := strconv.Itoa(sum)
		res += con2 + " x " + arg1 + " = " + con3 + "\n"
	}
	fmt.Print(res)
}
