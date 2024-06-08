package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	res1 := ""
	res2 := ""

	arg1 := os.Args[1]

	//The program should have a limit of 4000.
	// In case of an invalid number,
	//for example "hello" or 0
	//the program should print
	//ERROR: cannot convert to roman digit.
	myNum, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	if myNum <= 0 || myNum > 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	v1 := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	v2 := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	for indx := range nums {
		for myNum >= nums[indx] {
			myNum -= nums[indx]
			res1 += v1[indx]
			res2 += v2[indx]
			if myNum > 0 {
				res2 += "+"
			}
		}
	}
	fmt.Println(res2)
	fmt.Println(res1)
}
