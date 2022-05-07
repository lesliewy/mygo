package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
}

func test1() {
	// 新起一个goroutine
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-|/` {
			fmt.Printf("\r%c", r)   // \r  return, 回到行头.
			time.Sleep(delay)
		}
	}
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
