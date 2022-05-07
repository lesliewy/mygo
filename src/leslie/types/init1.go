package main

import "fmt"

/**
	init 函数先于main执行， 同一个程序里可以有多个init
 */
func init(){
	fmt.Println("init...")
}

func init(){
	fmt.Println("another init...")
}

func main(){
	test1()
}

func testInit1(){
	fmt.Println("=====test1=====")
}
