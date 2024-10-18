package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	arg1, arg2, arg3 := os.Args[1], os.Args[2], os.Args[3]
	var sum int
	if len(arg1) > 10 {
		return
	}
	x, err := strconv.Atoi(arg1)
	checkError(err)
	y, err := strconv.Atoi(arg3)
	checkError(err)

	if arg2 == "+" {
		sum = x + y
	} else if arg2 == "*" {
		sum = x * y
	} else if arg2 == "/" {
		if y == 0 {
			os.Stdout.Write([]byte("Cannot divide by zero"))
			return
		}
		sum = x / y
	} else if arg2 == "-" {
		sum = x - y
	} else if arg2 == "%" {
		if y == 0 {
			os.Stdout.Write([]byte("Cannot module by zero"))
			return
		}
		sum = x % y
	} else { //if it said ex: 1 p 2
		os.Exit(0)
	}

	conv := strconv.Itoa(sum)
	os.Stdout.Write([]byte(conv))
}
func checkError(e error) {
	if e != nil {
		os.Exit(0)
	}
}
