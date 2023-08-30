package main

import "fmt"

/*
@author echo huang
@desc 基础类
*/
func main() {
	boolean("hello", true)
	int_arr()
}

func boolean(a string, bool2 bool) {
	var b bool = true
	println("b is of type %t\n", b)
	e := bool(true)
	println("e is of type %t\n", e)
}

func int_arr() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

// 其他基础类型类似于java
