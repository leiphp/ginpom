package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func LoadHome(e *gin.Engine) {
	e.GET("/", home)
	e.GET("/index", index)
	e.GET("/about", about)
	e.GET("/website", website)
}