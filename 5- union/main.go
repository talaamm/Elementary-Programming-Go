package main

import (
	"fmt"
	"os"
)

// func to check if a letter exist in a string
func isLetter(r rune, s string) bool {
	for _, k := range s {
		if r == k {
			return true
		}
	}
	return false
}

func main() {
	//If the number of arguments is different from 2, then the program displays a newline ('\n').
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// you should take 2 strings
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	//we have to print the letters as they appear in order (arg1 then arg2)
	//we have to print all letters in both strings without repeating letters

	str := arg1 + arg2
	result := ""
	for _, l := range str {

		if isLetter(l, result) == false { //if the letter is NOT in the result, add it!
			result += string(l)
		}
	}
	fmt.Println(result)
}
