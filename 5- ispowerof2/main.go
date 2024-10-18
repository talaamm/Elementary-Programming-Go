package main

import (
	"fmt"
	"os"
	"strconv"
)

// allowed pkgs:  z01.PrintRune / Atoi / Itoa / os

func main() {
	arg1 := os.Args[1]

	sum := 1 // equals 1 because we are multiplying.   (number) * 0 = 0   so  sum = 1

	mynum, e := strconv.Atoi(arg1) //convert the number from string to int
	if e != nil {
		return
	}

	for sum < mynum {
		// while the number we are calculating is less than the given number, keep multiplying 2 by itself
		// at some point the number will get equal to or higher than the given number
		// at this point the loop will stops multiplying 2 by itslef

		sum *= 2

		// assume   mynum = 16
		// sum = 1
		// enter the for loop and asks   sum < mynum ?    1 < 16 ?  --> yes
		// sum = 2 * sum       sum = 2*1  --> sum =2
		// because its a for loop so ask again:  sum < mynum ?    2 < 16 ?  --> yes
		// sum = 2*2  --> sum =4
		// because its a for loop so ask again:  sum < mynum ?    4 < 16 ?  --> yes
		// sum = 2*4  --> sum =8
		// because its a for loop so ask again:  sum < mynum ?    8 < 16 ?  --> yes
		// sum = 2*8  --> sum =16

		// because its a for loop so ask again:  sum < mynum ?    16 < 16 ?  --> NO, 16 equal to 16

	}
	// so we get out of the loop and mynum = 16 and sum =16

	if sum == mynum { //if the number given (16) equal to sum print true
		fmt.Println(true)
	} else {
		fmt.Print(false)
	}
}
