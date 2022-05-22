package intf

import "fmt"

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("mouse" + m.name + " 启动成功....")
}

func (f FlashDisk) start() {
	fmt.Println("flashDisk" + f.name + "  启动成功....")
}

func (m Mouse) end() {
	fmt.Println("mouse" + m.name + "  关闭成功....")
}

func (f FlashDisk) end() {
	fmt.Println("flashDisk" + f.name + "  关闭成功....")
}

func (f FlashDisk) delete() {
	fmt.Println("flashDisk" + f.name + "  删除成功....")
}
