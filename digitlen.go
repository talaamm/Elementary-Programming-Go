package main

func DigitLen(n, base int) int {
	con := 1
	if !(base >= 2 && base <= 36) {
		return -1
	}
	if n < 0 {
		n = n * -1
	}
	for n/base != 0 {
		con++
		n = n / base
	}
	return con
}
