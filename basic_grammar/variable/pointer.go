package variable

import (
	"fmt"
)

import "go_playground/basic_grammar/utils"

func PointerVar() {
	utils.FuncStart("Print PointerVar")

	var c1 byte = '\x61'
	fmt.Printf("ptr of c1 is %p\n", &c1)

	s2 := "yinmin"
	fmt.Printf("value of s2 addr is %s\n", *(&s2))

	s3 := new(string)
	fmt.Printf("ptr of s3 is %p\n", &s3)

	utils.FuncEnd("Print PointerVar")

}
