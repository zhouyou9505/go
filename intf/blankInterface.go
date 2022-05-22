package intf

import "fmt"

type Blank interface {
}

type BlankImpl1 struct {
	name string
}

type BlankImpl2 struct {
	age int
}

func testBlank(b Blank) {
	fmt.Println(b)
}

func testBlank2(b interface{}) {
	fmt.Println(b)
}

/**
interface{}  类似于Object
*/
func mapIntf() {

	map1 := make(map[string]interface{})
	map1["a"] = 1
	map1["b"] = 1.1

	fmt.Println(map1)
}

