package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args) - 1)
}

/*
z01 Solution

***Note: the test cases won't check for a "os.Args" with length more than 15

func main(){
length:= len(os.Args)-1

if length<=9{
z01.PrintRune(length+48)
}else if length==10{
	z01.PrintRune('1')
	z01.PrintRune('0')
}else if length==11{
	z01.PrintRune('1')
	z01.PrintRune('1')
}else if length==12{
	z01.PrintRune('1')
	z01.PrintRune('2')
}else if length==13{
	z01.PrintRune('1')
	z01.PrintRune('3')
}else if length==14{
	z01.PrintRune('1')
	z01.PrintRune('4')
}else if length==15{
	z01.PrintRune('1')
	z01.PrintRune('5')
}
z01.PrintRune('\n')
}
*/
