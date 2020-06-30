package routers
//同步路由
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func syncHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}


func LoadSync(e *gin.Engine)  {
	e.GET("/long_sync", syncHandler)

}