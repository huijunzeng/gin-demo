## gin框架demo

官方文档：https://gin-gonic.com/docs/

下载安装gin依赖包  
工程路径下执行（-u参数表示存在就更新）：  
```
go get -u github.com/gin-gonic/gin  
```    
若报以下相关错：  
```
go get github.com/gin-gonic/gin: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 172.217.160.113:443: 
connectex: A connection attempt failed because the connected party did not properly respond after a period of time, 
or established connection failed because connected host has failed to respond.
```
则是国内无法访问到https://golang.org/

可执行
```
go env  
```  
看到环境参数
```
set GOPROXY=https://proxy.golang.org,direct  
```
可替换成国内七牛云的镜像（前提是环境参数set GO111MODULE=on，不然需先执行go env -w GO111MODULE=on）：   
```
go env -w GOPROXY=https://goproxy.cn,direct    
```    
之后再次下载即可


整合swagger自动生成API接口文档：  
下载go-swagger依赖包：  
```
go get -u github.com/swaggo/swag/cmd/swag
```
可通过输入命令查看swagger是否下载安装成功：  
```
swag -version
```
然后在main.go平级下执行命令生成swagger的docs文件夹：  
```
swag init
```
添加完路由启动项目之前，也需要先执行swag init命令更新swagger文档信息  

安装操作mysql驱动
```
go get -u github.com/go-sql-driver/mysql
```

安装操作gorm
```
go get -u github.com/jinzhu/gorm
```