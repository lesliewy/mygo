package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	//test1()
	//testHttpRouterParam1()
	//testQueryParam1()
	//testQueryArray1()
	testQueryMap1()
}

func test1() {
	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "王五"}}
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.Run(":8080")
}

/**
路由参数: Param()
*/
func testHttpRouterParam1() {
	r := gin.Default()
	r.GET("/users/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.String(200, "The user id is %s", id)
	})
	r.Run(":8080")
}

/**
查询参数: Query()
*/
func testQueryParam1() {
	r:=gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, c.Query("wechat"))
	})
	r.GET("/id", func(c *gin.Context) {
		c.String(200, c.DefaultQuery("id", "0"))
	})
	r.Run(":8080")
}

/**
	接收数组: QueryArray()   http://localhost:8080/array?media=a&media=b
 */
func testQueryArray1(){
	r := gin.Default()
	r.GET("/array", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})
	r.Run(":8080")
}

/**
 接收map: QueryMap   http://localhost:8080/map?ids[a]=1&ids[b]=2&ids[c]=3
 */
func testQueryMap1(){
	r := gin.Default()
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})
	r.Run(":8080")
}
