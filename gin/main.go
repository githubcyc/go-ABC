package main

import "github.com/gin-gonic/gin"
import (
    "log"
	"time"
	"reflect"
	"fmt"
)
func typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	/*
	  // 获取路由参数，假设有路由为"/user/:name"
		c.Params.ByName("name")

		// 获取query参数
		c.Query("name")
		c.DefaultQuery("name", "Guest")

		// 获取表单参数
		c.PostForm("name")
		c.DefaultPostForm("name")
		// response
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
		c.String(200, "pong")
	*/
	r.GET("/ping/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		n1 := c.Query("name")
		fmt.Println(typeof(n1))
		c.JSON(200, gin.H{
			"message": name + ", "+ n1,
		})
	})
	// gin可以借助协程来实现异步任务，但是这时候得手动copy上下文，
	// 并且只能是可读取的。

	r.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5)
			log.Println("Done! in path" + cCp.Request.URL.Path)
			
		}()
		c.String(200, "pong")
	})
	// Group
	// v1 := r.Group("/api/v1/userinfo")
	// {
	// 	v1.POST("/", CreateUser)
	// 	v1.GET("/", FetchAllUsers)
	// 	v1.GET("/:id", FetchSingleUser)
	// 	v1.PUT("/:id", UpdateUser)
	// 	v1.DELETE("/:id", DeleteUser)
	// }
	r.Run()
}
