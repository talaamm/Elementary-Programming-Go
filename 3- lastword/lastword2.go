package main

import (
	"fmt"
	"os"
)

func main2() {
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

	var word string
	//it saves the words until the ast one
	//but when it sees a space it knows that there is more words
	//so we empty the variable word and re-save the runes
	// since we got rid of the spaces from the end of the sentence
	//the last value of the var word IS the last word
	for _, w := range arg {
		if w == ' ' {
			word = ""
			continue
		}
		word += string(w)
	}
	fmt.Print(word)
}
