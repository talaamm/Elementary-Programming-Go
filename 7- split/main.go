package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", split(s, "HA"))
}

func split(s string, sep string) []string {
	var arr []string
	word := ""

	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) {
			if s[i:i+len(sep)] == sep {
				arr = append(arr, word)
				word = ""
				i += len(sep) - 1
			} else {
				word += string(s[i])
			}
		}
	}
	word += string(s[len(s)-1])
	arr = append(arr, word)
	return arr
}
