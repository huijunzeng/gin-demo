package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	router := gin.Default() // 函数返回的是一个Engine指针，Engine代表的是整个框架的一个实例，查看源码可发现实际就是调用New()方法创建实例,并且为实例添加了Logger和Recovery中间件.

	// 调用Engine的GET方法（其他请求方式POST PUT DELETE等）  第一个参数为相对路径，第二个参数为多个handle
	router.GET("/hello", func(c *gin.Context)  {
		// 当响应码为200时，返回JSON格式数据
		c.JSON(http.StatusOK, gin.H  {
			"message": "success",
		})
	})
	// 不指定ip地址和端口时，默认是监听并在 0.0.0.0:8080 上启动服务，另外的写法还有Run(":8080")、Run("0.0.0.0:8080")、Run("localhost:8080")都是指定http://localhost:8080或者http://127.0.0.1:8080/
	router.Run()
}