package router

import (
	. "gin-demo/controllers"
	_ "gin-demo/docs" // 执行swag init生成的docs文件夹路径 _引包表示只执行init函数
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() (r *gin.Engine) {

	router := gin.Default() // 函数返回的是一个Engine指针，Engine代表的是整个框架的一个实例，查看源码可发现实际就是调用New()方法创建实例,并且为实例添加了Logger和Recovery中间件.
	// swagger访问地址 http://localhost:8080/swagger/index.html
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	// 添加swagger的路由  不然报错404 page not found
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/users", GetUserListByPage)
	return router
}
