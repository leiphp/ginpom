package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*中间件*/
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

/*方法区*/
func home(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是测试", "address": "www.100txy.com"})
}

func about(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello this page is about!",
	})
}

func website(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.100txy.com")
}
//设置cookie
func login(c *gin.Context) {
	// 设置cookie
	c.SetCookie("abc", "123", 60, "/",
		"localhost", false, true)
	// 返回信息
	c.String(200, "Login success!")
}

//设置cookie
func info(c *gin.Context) {
	// 验证cookie
	c.JSON(200, gin.H{"data": "home"})
}


func LoadHome(e *gin.Engine) {
	e.GET("/", home)
	e.GET("/index", index)
	e.GET("/about", about)
	e.GET("/website", website)
	e.GET("/login", login)
	e.GET("/info", AuthMiddleWare(), info)
}