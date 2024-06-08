package main

func SortWordArr(a []string) {
	/* sort the array according to the ascii
	of each string */

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

/*
func main() {
	result := []string{"alaa", "Abu alamm", "123", "bobo", "Barhoom", "231", "coco", "Cooo", "312"}
	SortWordArr(result)

	fmt.Println(result)
}
*/
