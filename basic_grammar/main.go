package main

import (
	"go_playground/basic_grammar/container"
	"go_playground/basic_grammar/utils"
	"go_playground/basic_grammar/variable"
)

func test_variable() {
	utils.PackageStart("test variable")
	variable.PrintVars()
	variable.AnonymousVar()
}

func test_string() {
	utils.PackageStart("test string")
	variable.StringVar()
}

func test_char() {
	utils.PackageStart("test char")
	variable.CharVar()
}

func test_pointer() {
	utils.PackageStart("test pointer")
	variable.PointerVar()
}

func test_array() {
	utils.PackageStart("test array")
	container.ArrayVars()
}

func test_map() {
	utils.PackageStart("test map")
	container.MapVars()
}

func test_list() {
	utils.PackageStart("test list")
	container.ListVars()
}

func main() {
	// test_variable()
	// test_string()
	// test_char()
	// test_pointer()
	// test_array()
	// test_map()
	test_list()
}
