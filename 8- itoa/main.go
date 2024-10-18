package main

func Itoa(num int) string {
	if num == 0 { //if zero just return "0" duh return "0"
		return "0"
	}

	b := num // to save its REAL value (if its negative)
	res := ""
	if num < 0 { //if num is negative --> make it positive
		num = num * -1
	}
	for num != 0 { // as long as the num is not

		k := rune(num%10 + '0') //take the LAST digit (num%10) AND convert it to rune (num + '0') num = num / 10 //REMOVE the last digit
		res = string(k) + res   //ADD the digit from the back of res
	}
	if b < 0 { //if the ORIGINAL num was negative
		res = "_" + res // add a negative sign in the beggining
	}
	return res
}
