package main

import (
	"go_playground/basic_grammar/utils"
	"go_playground/basic_grammar/variable"
)

func test_variable() {
	utils.PackageStart("test variable")
	variable.PrintVars()
	variable.AnonymousVar()
}

func main() {
	test_variable()
}
