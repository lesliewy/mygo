package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	test1()
	test2()
	test3()
	testMultiDefer1()
}

/**
函数执行完毕后, 宕机后都可以执行.
*/
func test1() string {
	fmt.Println("=====test1=====")
	defer defer1("a")
	fmt.Println("this is test1....")
	return "a"
}

func defer1(msg string) {
	fmt.Println("defer1: ", msg)
}

/**
  函数调用前添加defer关键字. 则该函数调用会被defer执行.
*/
func test2() {
	fmt.Println("=====test2=====")
	defer trace("bigSlowOperation")() // 最后要有(), 否则只会在进入时执行trace(), 退出后不回执行trace return 的func
	time.Sleep(2 * time.Second)
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

/**
  defer 能修改函数返回结果.
  defer 能获取到定义的func result 和 实参.
  defer 不能带有return
*/
func test3() {
	log.Println("=====test3=====")
	log.Printf("triple(3): %d", triple(3))
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}() // 定义defer func, 最后要有(), 因为defer 的是后面的函数调用，而不是函数.
	return x + x
}

/*
  多个defer，按照定义顺序倒序执行.
*/
func testMultiDefer1() {
	log.Println("=====testMultiDefer1=====")
	defer defer1("a")
	defer defer1("b")
	defer defer1("c")
}

/**
  循环里的defer, 同样是等整个函数执行完毕后才会执行.
*/
