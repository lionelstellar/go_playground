package variable

// 字符: uint8(byte)和uint32(rune)

import (
	"fmt"
	"unicode"
)

import "go_playground/basic_grammar/utils"

func CharVar() {
	utils.FuncStart("Print CharVar")

	var c1 byte = '\x61'
	fmt.Printf("c1 is '%c'\n", c1)

	c2 := 'A'
	fmt.Printf("c2 is '%c'\n", c2)

	// unicode, \u四字节， \U八字节
	var c3 int = '\u0041'
	fmt.Printf("c3 is '%c'\n", c3)

	if unicode.IsLetter(rune(c3)) {
		fmt.Println("c3 is a letter")
	} else {
		fmt.Println("c3 is not a letter")
	}

	utils.FuncEnd("Print CharVar")

}
