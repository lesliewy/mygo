package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// logrus 完全兼容标准库的Log

func main(){
	testLogrus1()
	testLogrus2()
}

// Field机制：logrus鼓励通过Field机制进行精细化的、结构化的日志记录，而不是通过冗长的消息来记录日志。
func testLogrus1(){
	fmt.Println("=====testLogrus1=====")
	logrus.WithFields(logrus.Fields{"animal":"walrus"}).Info("a walrus appears.")

	// 不建议使用
	//logrus.Info("Failed to send event %s to topic %s with the key %d", event, topic, key)
	// 建议改成
	/*
	logrus.WithFields(logrus.Fields{
		"event":event,
		"topic":topic,
		"key":key
	}).Info("Failed to send event.")
	*/

}

func testLogrus2(){
	fmt.Println("=====testLogrus2=====")
	// 设置log
	// 日志格式为json
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置输出为标准输出，默认是标准错误.
	logrus.SetOutput(os.Stdout)
	// 指定级别以上才会输出日志.
	logrus.SetLevel(logrus.WarnLevel)

	logrus.WithFields(logrus.Fields{
		"animal":"walrus",
		"size":10,
	}).Warn("A group of walrus emerges from ocean")

}
