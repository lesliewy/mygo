package main

import "fmt"

func main() {
	testSlice1()
	testSlice2()
}

var months = [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "Augest", 9: "September", 10: "October", 11: "November", 12: "December"}
func testSlice1() {
	fmt.Println("=====testSlice1=====")
	// 初始化数组，同时指定 index

	// April, May, June
	Q2 := months[4:7]      // len:3, caps 12-4+1=9
	summer := months[6:9]  // len:3, caps:7
	fmt.Println(Q2)
	fmt.Println(summer)
	// 公共.
	fmt.Print("common: ")
	for _, s := range Q2 {
		for _, q := range summer {
			if s == q {
				fmt.Printf(" %s ", s)
			}
		}
	}
	fmt.Println()

	// 前6个，第0是空
	a := months[:6]
	// 后5个.
	b := months[len(months)-5:]
	fmt.Println(a)
	fmt.Println(b)
}

func testSlice2() {
	fmt.Println("=====testSlice2=====")
	// 超过caps 报错  caps: 7
	summer := months[6:9]
	fmt.Println("summer: ", summer)
	a := summer[2:20]
	//fmt.Println(a)
	// 超过lens: 扩展
	b := summer[1:6]
	fmt.Println(b)
}
