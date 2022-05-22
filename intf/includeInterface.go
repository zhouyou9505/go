package intf

import "fmt"

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {
}

func (c Cat) test3() {
	fmt.Println("test3....")
}

func (c Cat) test2() {
	fmt.Println("test2....")
}

func includeInterface() {

}

