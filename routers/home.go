package routers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//自定义一个字符串
var jwtkey = []byte("www.100txy.com")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

/*结构体验证*/
type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

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

//结构体验证
func check(c *gin.Context) {
	// 验证url http://localhost:8080/check?age=11&name=ginpom&birthday=2006-01-02
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		return
	}
	c.String(200, fmt.Sprintf("%#v", person))
}

//颁发token
func setting(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	ctx.JSON(200, gin.H{"token": tokenString})
}

//解析token
func getting(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	fmt.Println(111)
	fmt.Println(claims.UserId)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}

func LoadHome(e *gin.Engine) {
	e.GET("/", home)
	e.GET("/index", index)
	e.GET("/about", about)
	e.GET("/website", website)
	e.GET("/login", login)
	e.GET("/info", AuthMiddleWare(), info)
	e.GET("/check", check)
	e.GET("/set", setting)
	e.GET("/get", getting)
}