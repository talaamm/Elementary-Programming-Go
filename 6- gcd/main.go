package main

import (
	"fmt"
	"os"
	"strconv"
)
//if it only asks for a function write this down below
func gcd(a, b int) int {
	//notice this is a **for** not if
	for a != b {
		if a > b {
			a = a - b
		}
		if b > a {
			b = b - a
		}
	}
	return a
}

func main() {
	//If the number of arguments is different from 2, the program displays nothing.
	if len(os.Args) != 3 {
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	num1, _ := strconv.Atoi(arg1)
	num2, _ := strconv.Atoi(arg2)

	koko := gcd(num1, num2)

	fmt.Print(koko)
}
