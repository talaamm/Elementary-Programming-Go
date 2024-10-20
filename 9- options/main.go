package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
	}
	args := os.Args[1:] //  -abc -ijk

	for _, opt := range args { // -abc
		if len(opt) == 1 {
			fmt.Println("Invalid Option")
			return
		}
		if opt[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
		for _, r := range opt[1:] { // a  ...
			if !(r >= 'a' && r <= 'z') {
				fmt.Println("Invalid Option")
				return
			}
		}
	}
	// Initialize the 4 arrays
	arrays := [4][]int{
		make([]int, 8),
		make([]int, 8),
		make([]int, 8),
		make([]int, 8),
	}
	// Mapping letters to its index and array (position)
	letterMap := map[rune][2]int{
		'a': {3, 7}, 'b': {3, 6}, 'c': {3, 5}, 'd': {3, 4}, 'e': {3, 3}, 'f': {3, 2}, 'g': {3, 1}, 'h': {3, 0},
		'i': {2, 7}, 'j': {2, 6}, 'k': {2, 5}, 'l': {2, 4}, 'm': {2, 3}, 'n': {2, 2}, 'o': {2, 1}, 'p': {2, 0},
		'q': {1, 7}, 'r': {1, 6}, 's': {1, 5}, 't': {1, 4}, 'u': {1, 3}, 'v': {1, 2}, 'w': {1, 1}, 'x': {1, 0},
		'y': {0, 7}, 'z': {0, 6},
	}
	for _, opt1 := range args {
		for _, harf := range opt1[1:] {
			arrNum, letterPos := letterMap[harf][0], letterMap[harf][1]
			arrays[arrNum][letterPos] = 1
		}
	}
	fmt.Println(arrays[0], arrays[1], arrays[2], arrays[3])

	/*

		the first ever time how i solved it (the second part) LOL i had limited time
		and this was the only way i could think of

			arr1 := []int{0, 0, 0, 0, 0, 0, 0, 0}
			arr2 := []int{0, 0, 0, 0, 0, 0, 0, 0}
			arr3 := []int{0, 0, 0, 0, 0, 0, 0, 0}
			arr4 := []int{0, 0, 0, 0, 0, 0, 0, 0}

			for _, opt1 := range args {
				for _, harf := range opt1[1:] {
					if harf == 'a' {
						arr4[7] = 1
					} else if harf == 'b' {
						arr4[6] = 1
					} else if harf == 'c' {
						arr4[5] = 1
					} else if harf == 'd' {
						arr4[4] = 1
					} else if harf == 'e' {
						arr4[3] = 1
					} else if harf == 'f' {
						arr4[2] = 1
					} else if harf == 'g' {
						arr4[1] = 1
					} else if harf == 'h' {
						arr4[0] = 1
					} else if harf == 'i' {
						arr3[7] = 1
					} else if harf == 'j' {
						arr3[6] = 1
					} else if harf == 'k' {
						arr3[5] = 1
					} else if harf == 'l' {
						arr3[4] = 1
					} else if harf == 'm' {
						arr3[3] = 1
					} else if harf == 'n' {
						arr3[2] = 1
					} else if harf == 'o' {
						arr3[1] = 1
					} else if harf == 'p' {
						arr3[0] = 1
					} else if harf == 'q' {
						arr2[7] = 1
					} else if harf == 'r' {
						arr2[6] = 1
					} else if harf == 's' {
						arr2[5] = 1
					} else if harf == 't' {
						arr2[4] = 1
					} else if harf == 'u' {
						arr2[3] = 1
					} else if harf == 'v' {
						arr2[2] = 1
					} else if harf == 'w' {
						arr2[1] = 1
					} else if harf == 'x' {
						arr2[0] = 1
					} else if harf == 'y' {
						arr1[7] = 1
					} else if harf == 'z' {
						arr1[6] = 1
					}
				}
			}
			fmt.Println(arr1, arr2, arr3, arr4)*/
}
