package main

import (
	"fmt"
	"os"
)

// we need a func to check if a letter exists in a string
// we need this func because we dont want letters to be doubled
// we dont want letters to appear twice in the result
// ex: noor X
// it should be --> nor
func findLetter(r rune, s string) bool {
	for _, letter := range s {
		if letter == r {
			// if a letter from the string s is the same as the letter r
			return true
		}
	}
	//if there is no (r) in the string return false
	return false
}

func main() {

	//If the number of arguments is different from 2, the program displays nothing.
	if len(os.Args) != 3 {
		return
	}

	//take argument 1
	arg1 := os.Args[1]

	//take argument 2
	arg2 := os.Args[2]

	//define result to save in it the result
	res := ""

	for _, letter1 := range arg1 { // range over arg1 ... because it said in the order they appear in the first one

		// we need characters that appear in both strings so we check if the letter from str1 exists in str2
		if findLetter(letter1, arg2) == true {

			//if the letter exists in the second string, check if it is already appended to the result
			if findLetter(letter1, res) == false {

				//add the letter when false (the letter is not in the result) because we dont want doubles
				res += string(letter1)

			}
		}

	}
	fmt.Println(res)
}
