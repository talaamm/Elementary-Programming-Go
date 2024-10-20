package main

// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }
func HashCode(dec string) string {
	res := 0
	res1 := ""
	for _, i := range dec {
		res = (int(i) + len(dec)) % 127
		if res >= 0 && res <= 127 {
			if res < 32 && res != 127 {
				res += 33
			}
			res1 += string(res)
		}
	}
	return res1
}
