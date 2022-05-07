package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	testLog1()
	testLog2()
}

func testLog1(){
	fmt.Println("=====testLog1=====")
	log.Print("abc")
	log.Println("def")
	log.Printf("ghi: %s", "aaa")
	// 输出到标准错误后，exit 整个程序.
	//log.Fatal("fatal..")
}

// 自定义logger, 输出到文件中.
func testLog2(){
	fmt.Println("=====testLog2=====")
	// 在src/debug1.log
	fileName := "debug1.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error.")
	}
	debugLog := log.New(logFile, "[Debug]", log.Llongfile)
	debugLog.Println("A message here.")

	debugLog.SetPrefix("[Info]")
	debugLog.Println("A message here.")

	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A differenct prefix.")
}
