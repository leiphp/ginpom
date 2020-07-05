package main

import (
	"fmt"
	"ginpom/routers"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	//r := routers.SetupRouter()
	//// 3.监听端口，默认在8080
	//// Run("里面不指定端口号默认为8080")
	//if err := r.Run(":8000"); err != nil {
	//	fmt.Println("startup service failed, err:%v\n", err)
	//}
	r := gin.Default()

	r.LoadHTMLGlob("template/**/*")
	//如果你需要引入静态文件需要定义一个静态文件目录
	//r.Static("/assets", "./assets")
	routers.LoadHome(r)
	routers.LoadBlog(r)
	routers.LoadShop(r)
	routers.LoadAsync(r)//异步
	routers.LoadSync(r)//同步
	//注册中间件
	r.Use(MiddleWare())
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

