package main

import (
	"fmt"
)

/*
*
person struct
*/
type Mem struct {
	id   int
	name string
	age  int
}

/*
*
anonymous object 匿名对象
*/
type Person struct {
	Mem
	string
}

/*
*

	@author echo huang
	@desc go 集合类demo
*/
func main() {
	mapCollection()
	sliceCase()
	// test struct
	mem := Mem{
		1, "hsm", 10,
	}
	fmt.Println(mem)
	// test anonymous obj
	person := Person{mem, "这是一个男人"}
	fmt.Println(person)

	// test if condition
	conditionIf(false)
}

// 切片（slice）是对数组一个连续片段的引用。它是指向底层数组的指针。不需要定义长度。和数组在写法上的区别就是不需要指定长度。
func sliceCase() {
	var mySlice []string
	mySlice = append(mySlice, "a", "b", "c")
	fmt.Println(mySlice)
	// make build slice
	makeSlice := make([]string, 10)
	makeSlice = append(makeSlice, "flink", "spark")
	fmt.Println(makeSlice)
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

	// make build
	makeMap := make(map[string]string)
	makeMap["a"] = "1"
	makeMap["a"] = "2"
	fmt.Println(makeMap)
}

func conditionIf(isPass bool) {
	if isPass {
		println("恭喜你通过了!")
	} else {
		println("请继续努力...")
	}
}
