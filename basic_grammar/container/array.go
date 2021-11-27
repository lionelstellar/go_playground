package container

// 数组

import "fmt"
import "go_playground/basic_grammar/utils"

func ArrayVars() {
	utils.FuncStart("Print array vars")

	// declare and assign
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println("array a is", a)

	// walk array
	fmt.Println("\nwalk array a:")
	for k, v := range a {
		fmt.Printf("k is %d, v is %d\n", k, v)
	}

	b := [...]int{8, 7, 6, 5}
	fmt.Println(len(b))

	// cannot compare arrays of different types
	// fmt.Println(a == b)

	//
	// c := new([2][4]int)
	var c [2][4]int
	c = [2][4]int{{1, 2, 3, 4}, {5, 6, 7}}
	for k, v := range c {
		fmt.Printf("k is %d, v is %d\n", k, v)
	}

	c[1] = b

	// slice is same as Python
	fmt.Println("\nslicing of c0 from index 1 to 3:")
	fmt.Println(c[0][1:3])

	// append
	var d []int
	d = append(d, 1, 2, 3)
	fmt.Println("\ndynamically append:")
	fmt.Println(d)
	fmt.Println("capacity of d is", cap(d))

	// capacity of array is the power of 2
	d = append(d, 5, 4)
	fmt.Println(d)
	fmt.Println("capacity of d is", cap(d))

	// delete from slice
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	e = append(e[:4], e[6:]...)
	fmt.Println("\ndelete 5 and 6 from array e:", e)

	utils.FuncEnd("Print array vars")
}
