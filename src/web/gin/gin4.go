package main

import "github.com/gin-gonic/gin"

func main() {
	testPostForm1()
}

/**
	form-data 方式.
 */
func testPostForm1() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.String(200, wechat)
	})
	r.Run(":8080")
}
