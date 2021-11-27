package container

// 数组

import (
	"fmt"
	"sync"
)
import "go_playground/basic_grammar/utils"

func MapVars() {
	utils.FuncStart("Print map vars")

	// declare and assign
	// var mapname map[keytype]valuetype
	var mapA map[string]string
	fmt.Println("map a is", mapA)
	fmt.Println("length of map a is", len(mapA), "\n")

	// 用make来构造map
	// make(map[keytype]valuetype, cap), cap设定初始容量
	mapB := make(map[string]int, 10)
	mapB["ten"] = 10
	mapB["handred"] = 100
	mapB["thousand"] = 1000
	for k, v := range mapB {
		fmt.Printf("k is %s, v is %d\n", k, v)
	}

	// delete key from map
	delete(mapB, "ten")
	fmt.Println("\ndelete key 'ten'")
	for k, v := range mapB {
		fmt.Printf("k is %s, v is %d\n", k, v)
	}

	// lock the shared map
	var shared_map sync.Map
	shared_map.Store("apple", 10)
	shared_map.Store("banana", 5)
	fmt.Println("\nsync.Map")
	shared_map.Range(func(key, value interface{}) bool {
		fmt.Printf("k is %s, v is %d\n", key, value)
		return true
	})

	utils.FuncEnd("Print map vars")
}
