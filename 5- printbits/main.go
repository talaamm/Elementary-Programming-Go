package main

import (
	"fmt"
	"os"
	"strconv"
)

// allowed pkgs:  strconv / z01 / os

func main() {
	// it gives you an argument in os. ex: 192
	// and u have to print it in binay. ex: 01101010
	// it should be ALWAYS from 8 digits
	//example:  2  in binary is 10
	// but u should print 00000010

	if len(os.Args) != 2 {
		return
	}

	arg1 := os.Args[1] // it should be a number

	//convert from string to int
	// and you should check for errors as the question wants ex:
	// 	$ go run . a
	//  00000000$
	myNum, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println("00000000")
		return
	}

	//use a ready func to convert from decimal to binary
	// func FormatInt takes int64 and returns a string
	binary := strconv.FormatInt(int64(myNum), 2) // take myNum convert it to binary (2)
	// returns a string  EX:
	// binary = strconv.FormatInt(int64(192), 2)
	// binary = "11000000"  --> a string

	// it always should print 8 digits so
	//while the length of the binary string is less than 8
	//keep adding zeros at the beginning of the string
	for len(binary) < 8 {
		binary = "0" + binary
	}

	// print the string
	fmt.Println(binary)

}
