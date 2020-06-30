package main

import (
	"fmt"
	"ginpom/routers"
	"github.com/gin-gonic/gin"
)

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
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

