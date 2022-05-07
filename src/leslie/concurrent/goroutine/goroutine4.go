package main

import "fmt"

func main() {
	//testPipeline1()
	//testPipeline2()
	testPipeline3()
}

func testPipeline1() {
	naturals := make(chan int) // 类型为int的通道. 输出自然数.
	squares := make(chan int)  // 输出平方值.

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x // 发送，往通道里传值.
		}
	}()

	// squarer
	go func() {
		for {
			x := <-naturals // 接收通道里的值.
			squares <- x * x
		}
	}()

	// printer 在主goroutine中.
	for {
		fmt.Println(<-squares)
	}
}

func testPipeline2() {
	naturals := make(chan int) // 类型为int的通道. 输出自然数.
	squares := make(chan int)  // 输出平方值.

	// counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x // 发送，往通道里传值.
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer 在主goroutine中.
	for x := range squares {
		fmt.Println(x)
	}
}

/**
  只发送通道.
*/
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

/**
  只接收通道.
*/
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func testPipeline3() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
