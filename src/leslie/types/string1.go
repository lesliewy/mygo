package main

import "fmt"

func main() {
	test1()
}

func testString1() {
	fmt.Println("=====test1=====")
	// backquote ``:  里可以包含任何字符，除了back quote
	a := `aa
bb
cc`
	fmt.Println(a)
}
