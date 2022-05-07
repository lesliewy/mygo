package main

import (
	"fmt"
	"time"
)

/**
select 类似于操作系统的select, poll, epoll中的select, 监听多个channel
 */
func main(){
	testSelect1()
	testSelect2()
}

/**
非阻塞的.
select 的作用是同时监听多个 case 是否可以执行，如果多个 Channel 都不能执行，那么运行 default
*/
func testSelect1(){
	fmt.Println("=====testSelect1=====")
	ch := make(chan int)
	select {
	case i := <-ch:
		println(i)

	default:
		println("default")
	}
}

/**
多个channel到来时，随机执行.
 */
func testSelect2(){
	fmt.Println("=====testSelect2=====")
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}
