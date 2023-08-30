package main

import (
	"fmt"
)

/*
*

	@author echo huang
	@desc go 集合类demo
*/
func main() {
	mapCollection()
}

// map case
func mapCollection() {
	a := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(a["a"] + a["b"])

	// 遍历操作
	for key, value := range a {
		fmt.Printf("key:%s value:%d \n", key, value)
	}

	// add
	a["c"] = 3
	fmt.Println(a["c"])

	delete(a, "b")
	for key, value := range a {
		fmt.Printf("key:%s value:%d \n", key, value)
	}
}
