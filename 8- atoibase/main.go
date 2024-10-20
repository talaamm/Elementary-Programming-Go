package main

import "strings"

func atoiBase(s, base string) int {
	if !validBase(base) {
		return 0
	}
	/*
	   String number must contain only elements that are in base.
	*/
	for _, l := range s {
		if !strings.Contains(base, string(l)) {
			return 0
		}
	}
	var (
		reversed string
		res      int
	)
	for k := len(s) - 1; k >= 0; k-- {
		reversed += string(s[k])
	}
	for i, h := range reversed {
		pp := powerOf(len(base), i)
		pos := where(h, base)
		res += pp * pos
	}
	return res
}
func where(r rune, s string) int {
	for i, j := range s {
		if r == j {
			return i
		}
	}
	return -1
}
func powerOf(i1, power int) int {
	if power == 0 {
		return 1
	} else if power == 1 {
		return i1
	}
	sum := 1
	for i := 1; i <= power; i++ {
		sum *= sum
	}
	return sum
}
func validBase(b string) bool {
	/*
			A base must contain at least 2 characters.
		if len(b) < 2 {
			return false
		}*/
	if len(b) < 2 || checkIfRepeated(b) || strings.Contains(b, "+") || strings.Contains(b, "-") {
		return false
	}
	/*
			Each character of a base must be unique.
		if checkIfRepeated(b) {
			return false
		}*/

	/*
		   A base should not contain + or - characters.
		if strings.Contains(b, "+") || strings.Contains(b, "-") {
			return false
		}*/

	return true
}
func checkIfRepeated(s string) bool {
	res := ""
	for _, t := range s {
		if strings.Contains(res, string(t)) {
			return true //it is repeated
		} else {
			res += string(t)
		}
	}
	return false
}
