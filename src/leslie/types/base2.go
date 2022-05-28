package main

func main(){

}

/**
	new(int) 是语法糖，下面两个是等价的。
 */
func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var x int
	return &x
}
