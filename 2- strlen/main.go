package main

func StrLen(s string) int {
	return len(s)
}

//////for exam where they want only to count the runes(letters in ascii tbl)

func StrLen2(s string) int {
	return len([]rune(s))
}
