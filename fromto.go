package main

import "strconv"

func FromTo(from int, to int) string {
	res := ""
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	if to > from {
		for i := from; i <= to; i++ {
			conv1 := strconv.Itoa(i)
			if len(conv1) == 1 {
				conv1 = "0" + conv1
			}
			res += conv1 + ", "
		}
	} else {
		for i1 := from; i1 >= to; i1-- {
			conv2 := strconv.Itoa(i1)
			if len(conv2) == 1 {
				conv2 = "0" + conv2
			}
			res += conv2 + ", "
		}
	}
	return res[:len(res)-2] + "\n"
}
