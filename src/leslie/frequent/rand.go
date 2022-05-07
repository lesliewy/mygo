package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testRand1()
}

/**
相同种子，每次运行的结果都是一样的
 */
func testRand1() {
	fmt.Println("=====testRand1=====")
	for i := 0; i < 10; i++ {
		fmt.Printf("current:%d\n", time.Now().Unix())
		rand.Seed(time.Now().Unix())
		fmt.Println(rand.Intn(100))
	}
}
