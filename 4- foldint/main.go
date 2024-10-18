package main

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, i := range a {
		n = f(n, i)
	}
	fmt.Println(n)
}

//func main to test the func

/*func main() {
	table := []int{1, 2, 3}
	ac := 93
	Mul := func(acc int, cur int) int {
		return acc * cur
	}
	Add := func(acc int, cur int) int {
		return acc + cur
	}
	Sub := func(acc int, cur int) int {
		return acc - cur
	}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()
	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
*/
