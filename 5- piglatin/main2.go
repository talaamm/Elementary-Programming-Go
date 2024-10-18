package main

import (
	"fmt"
	"os"
)

// function to check if the string contains a vowel
func isVowel(r rune) bool {
	s := "aeiouAEIOU"
	for _, i1 := range s {
		if r == i1 {
			return true
		}
	}
	return false
}

func conatinsV(s string) bool {
	for _, k := range s {
		if isVowel(k) {
			return true
		}
	}
	return false
}

func maiin() {
	if len(os.Args) != 2 {
		return
	}
	arg1 := os.Args[1]

	if !conatinsV(arg1) {
		fmt.Println("no vowels in the string")
		return
	}

	if isVowel(rune(arg1[0])) {
		fmt.Println(arg1 + "ay")
		return
	} else {
		normal_Letters := ""
		resVOWELS := ""
		final := ""
		// else means the first letter is NOT vowel. ex:  count --> there is vowels in the word but the vowel is NOT the 1st letter
		for x, i := range arg1 {
			if isVowel(i) {
				resVOWELS = arg1[x:]
				break
			} else {
				normal_Letters += string(i)
			}
		}
		final = resVOWELS + normal_Letters + "ay"
		fmt.Println(string(final))
	}
}
