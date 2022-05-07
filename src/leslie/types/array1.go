package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	test1()
	test2()
}
func test1() {
	fmt.Println("=======test1========")
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// 索引、元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	fmt.Println("aaaaa")
	// 数组初始化, 最后的是0
	b := [5]int{77, 88, 99}
	for _, v := range b {
		fmt.Printf("%d\n", v)
	}

	fmt.Println("bbbbb")
	// ... 表示个数由初始化决定.
	c := [...]int{2, 3, 4, 5, 6}
	fmt.Printf("%T\n", c)   // type
	for _, v := range c {
		fmt.Printf("%d\n", v)
	}

	// 直接 --, !=   必须是同样类型的数组. [3]int 和 [4]int 属于不同类型.
	fmt.Printf("a == b: %t", b == c)   // %t 输出 boolean

}

func test2(){
	fmt.Println("======test2======")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
