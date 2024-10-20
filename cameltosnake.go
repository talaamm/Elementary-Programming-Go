package main

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	if checkcamel(s) == false {
		return s
	} else {
		return conv2Snack(s)
	}
}
func checkcamel(s1 string) bool {
	if s1[len(s1)-1] >= 'A' && s1[len(s1)-1] <= 'Z' {
		return false // Check if it ends with an uppercase letter
	}
	for x := 0; x < len(s1); x++ {
		char := s1[x]
		if char >= 'A' && char <= 'Z' {
			if x+1 < len(s1) && s1[x+1] >= 'A' && s1[x+1] <= 'Z' {
				return false // Check if two uppercase letters follow each other
			}
		}
		if char >= '0' && char <= '9' {
			return false // Check if there are any numbers
		}
		specialChars := []rune{'.', ',', '/', '-', '_', '&', '*', '$', '@', '#', '!', '?', '+', '='}
		for _, special := range specialChars {
			if rune(char) == special {
				return false // Check if any special characters are present
			}
		}
	}
	return true
}
func conv2Snack(s string) string {
	arr := []rune(s)
	res1 := ""
	res, index := checkupper(s)
	for res == true {
		// Convert uppercase to lowercase with underscore
		res1 += string(arr[:index]) + "_" + string(arr[index])
		s = s[index+1:]
		arr = []rune(s) // Update the rune slice for further slicing
		res, index = checkupper(s)
	}
	res1 += s // Append the remaining part of the string
	return res1
}
func checkupper(s string) (bool, int) {
	for x, i := range s {
		if i >= 'A' && i <= 'Z' {
			return true, x // Return true with the index of the first uppercase letter
		}
	}
	return false, -1
}
