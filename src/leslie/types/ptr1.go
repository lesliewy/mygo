package main

import "fmt"

func main() {
	testPtr1()
	testPtr2()
}

func testPtr1() {
	fmt.Println("=====test1=====")
	x := 1
	p := &x // p是整型指针, 指向x
	fmt.Println(*p)
	*p = 2 // 等于x=2
	fmt.Println(x)

	// 指针比较   当且仅当指向同一个变量时才相等.  值相等指针不一定相等.
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)
}

func testPtr2() {
	fmt.Println("=====test2=====")
	a := 1
	b := 1
	incr(&a)  // a = 2
    incr2(b)  // b = 1
	fmt.Printf("a=%d, b=%d\n", a, b)
}

// 指针参数  直接修改了参数.
func incr(p *int) int {
	*p++ // 递增p指向的的值, p自身不变.
	return *p
}

func incr2(p int) int {
	p++
	return p
}
