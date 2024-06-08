package main

import (
	"fmt"
	"os"
)

func main() {
	//it gives u a string
	//you should print the string with
	//repeating each letter according to its place in the alphabet
	// u should print (a) one time because its the first letter
	// (b) should be printed twice
	// (c) should be printed 3 times
	// (z) should be printed 26 times

	//ex: abc  ---> result: abbccc

	//If the number of arguments is different from 1, the program displays nothing.
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	res := ""

	for _, i := range str {
		if i >= 'a' && i <= 'z' {

			for ll := 0; ll <= int(i-'a'); ll++ {
				// add the letter (i-'a') times
				// ex:  'b' - 'a'  -->  98 - 97 = 1
				// so the loop goes 2 times bc it starts from 0
				res += string(i)
			}

		} else if i >= 'A' && i <= 'Z' {
			for ll := 0; ll <= int(i-'A'); ll++ {
				res += string(i) //same as for uppercase
			}
		} else { // for other letter as: 0 1 ! . , ] ...
			res += string(i)
		}
	}
	fmt.Println(res)
}
