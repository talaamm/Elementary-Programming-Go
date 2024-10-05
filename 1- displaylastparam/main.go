package main

import (
	"fmt"
	"os"
)

func main() {
fmt.Println(os.Args[len(os.Args)-1])
}

/*
z01 Solution

func main(){

	for _, k:= range os.Args[len(os.Args)-1]{
	z01.PrintRune(k)
	}
z01.PrintRune('\n')
}
*/