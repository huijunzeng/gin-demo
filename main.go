package main

import (
	"gin-demo/common"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 初始化数据库连接池
	err := common.InitDB()
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer common.Close()

	// 初初始化路由
	router := router.InitRouter() // 函数返回的是一个Engine指针，Engine代表的是整个框架的一个实例，查看源码可发现实际就是调用New()方法创建实例,并且为实例添加了Logger和Recovery中间件.

	//v1组路由  把Print和Hello方法添加到同一组路由，即访问时需要在前面加上/api/v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/print", Print)
		v1.GET("/hello", Hello)
	}

	// 不指定ip地址和端口时，默认是监听并在 0.0.0.0:8080 上启动服务，另外的写法还有Run(":8080")、Run("0.0.0.0:8080")、Run("localhost:8080")都是指定http://localhost:8080或者http://127.0.0.1:8080/
	router.Run()
}

// @Summary 打印测试功能
// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
// @Host 127.0.0.1:8080
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
// @Router /print [get]
func Print(c *gin.Context) {
	var (
		name string
	)
	name = c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": name,
	})
}

// @Summary Hello接口
// @Description Hello接口
// @Tags 用户信息
// @Success 200 {string} json "{"message":"success"}"
// @Router /hello [get]
func Hello(c *gin.Context) {
	// 当响应码为200时，返回JSON格式数据
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// @Summary 测试接口
// @Description 描述信息
// @Success 200 {string} string  "ok"
// @Router /test [get]
func Test(c *gin.Context) {
	c.JSON(200, "ok")
}
