package main

import "fmt"

func main() {
	for i := 'z'; i >= 97; i -= 2 {
		fmt.Print(string(i) + string(i-32-1))
	}
	fmt.Println()
}

/*
z01 solution

func main() {
	for i := 'z'; i >= 97; i -= 2 {
		z01.PrintRune(i)
		z01.PrintRune(i-32-1)
	}
	z01.PrintRune('\n')

}

*/
