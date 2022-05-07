package main

import (
	"github.com/gin-gonic/gin"
	aaa "hello/leslie/frequent/protobuf/message"
	"net/http"
)

func main() {
	r := gin.Default()
	// http://localhost:8080/protobuf 下载序列化的数据.
	r.GET("/protobuf", func(c *gin.Context) {
		data := &aaa.User{
			Name: "张三",
			Age:  20,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}
