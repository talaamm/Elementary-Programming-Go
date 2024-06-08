// printhex same as printbits but easier
package main

import (
	"fmt"
	"os"
	"strconv"
)

// allowed pkgs:  strconv / z01 / os

func main5656() {
	// it gives you an argument in os. ex: 30
	// and u have to print it in hex. ex: 1e
	//ex:  (255)  in hex is (ff)

	if len(os.Args) != 2 {
		return
	}

	arg1 := os.Args[1] // it should be a number

	//convert from string to int
	// and you should check for errors as the question wants ex:
	// $ go run . "123 132 1" | cat -e
	// ERROR$
	myNum, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	//use a ready func to convert from decimal to hex
	// func FormatInt takes int64 and returns a string
	binary := strconv.FormatInt(int64(myNum), 16) // take myNum convert it to hex (16)
	// binary is a string  EX:
	// binary = strconv.FormatInt(int64(10), 16)
	// binary = "a"  --> a string default value is lowercase

	// print the string
	fmt.Println(binary)

}
