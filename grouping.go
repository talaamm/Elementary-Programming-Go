package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1, arg2 := os.Args[1], os.Args[2]

	if !(arg1[0] == '(' || arg1[len(arg1)-1] == ')') {
		return
	}
	var result []string
	search := strings.Split(arg1[1:len(arg1)-1], "|")
	allWords := strings.Split(arg2, " ")
	for _, word := range allWords {
		for _, sub := range search {
			if strings.Contains(word, sub) {
				if unicode.IsPunct(rune(word[len(word)-1])) {
					result = append(result, word[:len(word)-1])
				} else {
					result = append(result, word)
				}
			}
		}
	}
	for i, k := range result {
		fmt.Println((strconv.Itoa(i + 1)) + ": " + k)
	}
}
