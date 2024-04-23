package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	//to remove spaces from then end of the word
	for arg[len(arg)-1] == ' ' {
		arg = arg[:len(arg)-1]
	}

	if arg == "" {
		return
	}

	//reverse the sentence so that the last word is now at the beggining of the sentence
	var reversed string
	for _, r := range arg {
		reversed = string(r) + reversed
	}

	//reverse the string again & save the letters until it sees a space
	// since words are runes delimited by spaces
	var result string
	for _, t := range reversed {
		if t == ' ' {
			break
		}
		result = string(t) + result
	}

	//print the result
	fmt.Println(result)
}
