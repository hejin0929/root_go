package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"modTest/component/login"
	_ "modTest/docs"
	chum2 "modTest/paths/chum"
	"modTest/paths/getCode"
	home2 "modTest/paths/home"
	loginPaths "modTest/paths/login"
	"modTest/paths/uploadFiles"
	"modTest/web_socket"
	"net/http"
	"strings"
)

var router *gin.Engine

func init() {
	router = gin.New()
}

func main() {
	router.Use(Cors())
	router.Use(UserVerify())

	// socket 通信
	Api := router.Group("/api")

	//Api.GET("/ws", func(g *gin.Context) {
	//	ws := web_socket.NewWsServer()
	//	err := ws.Start()
	//	if err != nil {
	//		g.JSON(http.StatusBadRequest, err)
	//		return
	//	}
	//
	//})

	Api.GET("/ws", func(g *gin.Context) {
		web_socket.WebSocketNews(g)
	})

	//ws := web_socket.NewWsServer()
	//_ = ws.Start()

	loginGroup := Api.Group("/login")

	loginGroup.GET("/user/:name", login.GetSingCode) // 用户获取验证码登录接口

	//Login.POST("/user/login", login.LoginsUser) // 用户登录接口

	loginGroup.POST("/user/login", loginPaths.LoginInPasswordPaths) // 用户登录接口

	loginGroup.POST("/user/login_code", loginPaths.LoginInCodePaths) // 用户验证码登陆

	loginGroup.POST("/user/sign", loginPaths.LoginNewsUserPaths) // 用户注册接口
	//
	//Login.POST("/user/set_password", login.SetPasswordUser) // 修改密码

	// 验证码模块
	code := Api.Group("/phone_code")

	code.GET("/user/:phone", getCode.GetPathsCode) // 获取验证码

	// 文件上传模块
	file := Api.Group("/upload")

	file.POST("/images", uploadFiles.UploadImages)          // 上传图片
	file.POST("/videos", uploadFiles.UploadVideos)          // 上传视频
	file.GET("/deleteImg", uploadFiles.UploadImagesDelete)  // 删除图片
	file.GET("/deleteVideo", uploadFiles.UploadVideoDelete) // 删除视频
	file.POST("/test", uploadFiles.UploadTestPaths)         // 测试

	// 首页
	home := Api.Group("/home")

	home.GET("/message", home2.AppHomePaths) // 首页信息模块

	// 个人信息获取
	user := Api.Group("/user")

	user.GET("/user_message/:id", home2.GetUserMessagePaths)

	// 好友功能
	chum := Api.Group("/chum")

	chum.GET("/search", chum2.SearchUserPaths) // 搜素好友

	router.StaticFS("/oss/images", http.Dir("./static/images/"))
	router.StaticFS("/oss/videos", http.Dir("./static/videos/"))

	ginSwagger.URL("http://localhost:8080/docs/swagger.json")

	Api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8081").Error()

}

// Cors //// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")

			c.Set("Access-Control-Allow-Origin", "*")    //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")    //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "multipart/form-data") // 设置返回格式是json
			c.Set("content-type", "text/plain; charset=utf-8")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
