package main

import (
	"fmt"
	"os"
)

func mmain() {
	if len(os.Args) != 2 {
		return
	}
	rs1 := os.Args[1]
	for rs1[len(rs1)-1] == ' ' {
		rs1 = rs1[:len(rs1)-1]
	}
	for rs1[0] == ' ' {
		rs1 = rs1[1:]
	}
	for x, i := range rs1 {
		if i == ' ' {
			continue
		}
		if x > 0 && (i >= 33 && i <= 126) && (rs1[x-1] == ' ') {
			fmt.Print("   " + string(i))
		} else {
			fmt.Print(string(i))
		}
	}
}
