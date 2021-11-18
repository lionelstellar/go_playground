package utils

import "fmt"

func FuncStart(s string) {
	fmt.Printf("[*] %s:\n", s)
}

func FuncEnd(s string) {
	fmt.Printf("[*] %s OK!\n", s)
	fmt.Println()
}

func PackageStart(s string) {
	fmt.Printf("**********************  %s  **********************\n", s)
}
