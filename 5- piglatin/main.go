package main

import (
	"fmt"
	"os"
)

// function to check if the string contains a vowel
func containV(s string) bool {
	for _, i1 := range s {
		if i1 == 'a' || i1 == 'e' || i1 == 'i' || i1 == 'o' || i1 == 'u' || i1 == 'A' || i1 == 'E' || i1 == 'I' || i1 == 'O' || i1 == 'U' {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) != 2 {
		return
	}

	arg1 := os.Args[1]

	//this is error handling
	//if the string does not contain any vowels print the error
	if containV(arg1) != true {
		fmt.Println("no vowel in the string")
		return
	}

	// check if the first letter ( index 0 ) is a vowel
	if arg1[0] == 'a' || arg1[0] == 'e' || arg1[0] == 'i' || arg1[0] == 'o' || arg1[0] == 'u' || arg1[0] == 'A' || arg1[0] == 'E' || arg1[0] == 'I' || arg1[0] == 'O' || arg1[0] == 'U' {

		//if the first letter is a vowel. add the word with "ay"
		//ex:  agkl --> first letter (a) is vowel so
		// res =   agkl + ay = agklay
		arg1 = arg1 + "ay"
		fmt.Println(string(arg1))
		return

	} else {
		normal_Letters := ""
		resVOWELS := ""
		final := ""
		// else means the first letter is NOT vowel. ex:  count --> there is vowels in the word but the vowel is NOT the 1st letter
		for x, i := range arg1 {

			//check if the letter is a vowel
			if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' || i == 'A' || i == 'E' || i == 'I' || i == 'O' || i == 'U' {
				//	if the letter is a vowel add ALL letters to resVOWELS
				//starting from where it find the vowel to the end
				resVOWELS = arg1[x:]
				//and stop the for loop
				break

			} else {
				// if the letter is NOT  a vowel add it to normal letters
				normal_Letters += string(i)
			}

		}
		// if arg1 = client
		//the result we want is:
		// c l i e n t --->  it finds the first vowel at 'i'. so it splits the string to 2 pieces
		// cl  +  ient --->  ient + cl + ay

		//in the loop: normal_Letters = "cl"    resVowels = "ient"

		//result = ientclay

		final = resVOWELS + normal_Letters + "ay"
		fmt.Println(string(final))
	}
}
