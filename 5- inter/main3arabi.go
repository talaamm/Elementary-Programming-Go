package main

import (
	"fmt"
	"os"
)

//  funcنحتاج إلى وظيفة .
// stringللتحقق من وجود حرف في

// resنحتاج إلى هذه الوظيفة لأننا لا نريد أن تتكرر الحروف في النتيجة
// example:  res = noor --> WRONG ... it should be --> res = nor
// لا تكرار للحروف

func findLetterll(r rune, s string) bool {
	for _, letter := range s {
		if letter == r {
			// string (s)إذا كان إذا كان هناك حرف من
			//  (r) هو نفس الحرف
			//  أعد صحيح. الحرف موجود
			return true
		}
	}

	// إذا لم يكن هناك في السلسلة الحرف فارجع خطأ
	return false
}

func mainn() {

	//If the number of arguments is different from 2, the program displays nothing.
	if len(os.Args) != 3 {
		return
	}

	//take argument 1
	arg1 := os.Args[1]

	//take argument 2
	arg2 := os.Args[2]

	//define a variable to save in it the result
	res := ""

	for _, letter1 := range arg1 { // range over arg1 ... because it said: (in the order they appear in the first one(first argument))

		// نحن بحاجة إلى الأحرف التي تظهر في كلا السلسلتين.
		// str1لذلك، نتحقق مما إذا كان الحرف من
		// str2موجودًا في
		if findLetter(letter1, arg2) == true {

			// resإذا كان الحرف موجودًا في السلسلة الثانية، فتحقق مما إذا كان قد تمت إضافته من قبل إلى النتيجة
			//	لأننا لا نريد التكرار
			if findLetter(letter1, res) == false {

				//  resأضف الحرف إلى
				// falseعندما تكون النتيجة
				// لأننا لا نريد تكرار الأحرف
				res += string(letter1)

			}
		}

	}
	fmt.Println(res)
}
