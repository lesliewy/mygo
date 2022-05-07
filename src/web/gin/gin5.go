package main

import "github.com/gin-gonic/gin"

func main() {
	//testGroup1()
	//testGroup2()
	testGroup3()
}

/**
分组路由.
*/
func testGroup1() {
	r := gin.Default()
	// V1版本
	v1Group := r.Group("/v1")
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})

	// V2版本
	v2Group := r.Group("/v2")
	v2Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v2/users")
	})
	v2Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v2/products")
	})
	r.Run(":8080")
}

func testGroup2() {
	r := gin.Default()
	// V1版本  添加公共处理.
	v1Group := r.Group("/v1", func(c *gin.Context) {
		c.String(200, "/v1 public\n")
	})
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})

	// V2版本
	v2Group := r.Group("/v2")
	v2Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v2/users")
	})
	v2Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v2/products")
	})
	r.Run(":8080")
}


/**
	分组路由嵌套.
 */
func testGroup3() {
	r := gin.Default()
	// V1版本
	v1Group := r.Group("/v1")
	// 再增加一个分组路由
	v1Group = v1Group.Group("/admin")
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/admin/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/admin/products")
	})

	r.Run(":8080")
}
