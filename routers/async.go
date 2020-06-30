package routers
//异步路由
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func asyncHandler(c *gin.Context) {
	// 需要搞一个副本
	copyContext := c.Copy()
	// 异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}


func LoadAsync(e *gin.Engine)  {
	e.GET("/long_async", asyncHandler)

}