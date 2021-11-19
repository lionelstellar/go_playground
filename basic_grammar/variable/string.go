package variable

// 字符串

import "fmt"

import "go_playground/basic_grammar/utils"

func StringVar() {
	utils.FuncStart("Print StringVar")
	var s1 string
	s1 = "this is a"
	s1 += " string"
	fmt.Printf("s1 is \"%s\"\n", s1)
	fmt.Printf("initial char of s1 is '%c'\n", s1[0])

	// ` ... ` in go 类似于 """  ...   """ in Python
	s2 := `this
		is 
		also
		a
		string
		`
	fmt.Printf("s2 is \"%s\"\n", s2)
	utils.FuncEnd("Print StringVar")

}
