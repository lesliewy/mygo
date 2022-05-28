package main

import (
	"fmt"
	"math/rand"
)

func main() {
	testVar1()
	testDeclare1()
	testIf1()
}

// 在任何函数外，语句都应该以关键字开头，例如 type、var这样的关键字

/**
  golang 中不存在未初始化的变量. 默认是零值.
*/
func testVar1() {
	fmt.Println("=====testVar1=====")
	// 标准变量声明方式
	var m int = 3
	fmt.Println(m)

	// 不显式赋初始值声明方式.
	var s1, s2 string // ""
	var i int         // 0
	var j float32     // 0
	var a, b bool     // false
	fmt.Println(s1, i, j, s2, a, b)

	var u, v, w = 2, true, "abc"
	fmt.Println(u, v, w)

	// 省略类型变量声明方式.
	// 类型和显式赋初始值同时省略
	var n = 3
	fmt.Println(n)
}

/**
短变量声明:=  局部变量声明
多赋值区别
*/
func testDeclare1() {
	fmt.Println("=====testDeclare1=====")
	var i = 3
	// := 左侧的变量必须是未声明过的.
	//i := 3
	fmt.Println(i)

	// var声明: 初始值类型不同时使用;
	var j float32 = 4
	fmt.Println(j)

	// := 声明左侧至少要有一个new variable,  此时其他的相当于赋值.
	m, i := false, 33 // i 重新赋值
	fmt.Print(m, i)

	// := 声明不可以重复, 多赋值可以重复.
	i, j = 3, 4.5
	fmt.Println(i, j)
	i, j = 4, 4.6
	fmt.Println(i, j)
}

/**
	类型转换
 */
func testDelare2() {
	fmt.Println("=====testDeclare2=====")
	var a = int8(100)
	b := int8(60)
	fmt.Println(a, b)
}

func testIf1() {
	fmt.Println("=====testIf1=====")
	a := 2
	// if 中可以带有simple statement, 这里的赋值语句就是simple statement
	if b := 3; a < b {
		fmt.Println("a < b")
	}

	if b := rand.Int(); a < b {
		fmt.Println("a < b", b)
	} else {
		fmt.Println("a >= b", b)
	}
}
