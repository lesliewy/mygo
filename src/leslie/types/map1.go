package main

import "fmt"

func main() {
	testMap1()
}

func testMap1() {
	fmt.Println("=====testMap1=====")
	// 初始化map.  key:string,  value:int
	ages1 := map[string]int{
		"alice":   14,
		"charlie": 15,
	}
	fmt.Println("ages1 ", ages1)
	// 等价于上面. make() 创建map
	ages2 := make(map[string]int)
	ages2["alice"] = 14
	ages2["charlie"] = 15
	ages2["bob"] = 16
	fmt.Println("ages2 ", ages2)

	// remove item
	delete(ages2, "alice")

	// 遍历一：不关心 key 和 value 的情况
	for range ages2 {
	}

	// 遍历二：只关心 key 的情况
	fmt.Println("keys:")
	for key := range ages2 {
		fmt.Println(key)
	}

	// 遍历二：关心 key 和 value 的情况
	fmt.Println("key and values:")
	for key, value := range ages2 {
		fmt.Println(key, value)
	}

	// 遍历
	//for k, v := range ages2 {
	//	fmt.Printf("%s: %d\n", k, v)
	//}
	// map 元素不能获取地址. rehash
	//a := &ages2["a"]
}
