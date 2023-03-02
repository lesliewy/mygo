package main

import (
	"fmt"
	"log"
)

func main() {
	//testPanic1()
	testPanic2()
}

/*
*

	内置的宕机函数: panic
*/
func testPanic1() {
	log.Println("=====test1=====")
	panicInTest1()
}

func panicInTest1() {
	panic("error...")
}

/*
*

	defer 倒序输出, panic 输出堆栈信息.
*/
func testPanic2() {
	log.Println("=====testPanic2=====")
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
