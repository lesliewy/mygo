package main

import "github.com/gin-gonic/gin"

func main(){
	// 实例化一个默认的gin示例
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		// c.JSON 设置返回.
		c.JSON(200, gin.H{
			"Blog":"aaa",
		})
	})
	// 启动服务.
	r.Run(":8080")
}
