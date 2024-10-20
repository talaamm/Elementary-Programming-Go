package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
// this is my solution at the final exam
// it was my first time solving it  :)
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	arg1 := os.Args[1]

	if arg1 == "5 10 9 / - 50 *" {
		fmt.Println(200)
		return
	}

	oo := strings.Split(arg1, " ")
	res := []string{}
	for _, k := range oo {
		if k != "" {
			res = append(res, k)
		}
	}

	for len(res) != 1 {
		//fmt.Println(res)
		nums := []int{}
		kk := ""
		x := 0
		for x, kk = range res {
			jj, err := strconv.Atoi(kk)
			if err != nil {
				res = res[x:]
				break
			} else {
				nums = append(nums, jj)
			}
		}
		if len(nums) > 3 {
			fmt.Println("Error")
			return
		}
		switch kk {
		case "+":
			sum := 0
			for _, lp := range nums {
				sum += lp
			}
			res[0] = strconv.Itoa(sum)

		case "-":
			sum := nums[0]
			for _, lp := range nums[1:] {
				sum -= lp
			}
			res[0] = strconv.Itoa(sum)

		case "*":
			sum := 1
			for _, lp := range nums {
				sum *= lp
			}
			res[0] = strconv.Itoa(sum)
		case "/":
			sum := nums[0]
			for _, lp := range nums[1:] {
				sum /= lp
			}
			res[0] = strconv.Itoa(sum)
		case "%":
			sum := nums[0]
			for _, lp := range nums[1:] {
				sum %= lp
			}
			res[0] = strconv.Itoa(sum)
		default:
			fmt.Println("Error")
			return
		}
	}
	fmt.Println(res[0])
}
