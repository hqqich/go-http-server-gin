package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goHttpServerGin/common"
	"net/http"
	"path"
	"time"
)

func main() {

	// 初始化 redis
	err := common.InitRedisClient()
	if err != nil {
		common.Info("redis连接失败! err : %v\n", err)
		return
	}
	common.Info("redis连接成功! ")

	r := gin.Default()

	// 中间件注册
	r.Use(costTimeMiddleware())

	// 静态资源映射
	r.StaticFile("/", "./dist/index.html")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/css", http.Dir("./dist/css"))
	r.StaticFS("/js", http.Dir("./dist/js"))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// redis get
	r.GET("/redis/get/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		value := common.Get(key)

		ctx.JSON(200, gin.H{
			"key":   key,
			"value": value,
		})
	})
	// redis set
	r.GET("/redis/set/:key/:value", func(ctx *gin.Context) {
		key := ctx.Param("key")
		value := ctx.Param("value")
		common.Set(key, value)

		ctx.JSON(200, gin.H{
			"status": "success",
		})
	})

	r.PUT("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := ctx.DefaultQuery("user", "user")
		password := ctx.Query("password")

		ctx.JSON(200, gin.H{
			"message":  "success",
			"id":       id,
			"user":     user,
			"password": password,
		})
	})

	// 上传文件
	r.POST("/fileupload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		fmt.Println("原始文件名：" + file.Filename)

		fileName := c.PostForm("text")

		if err == nil {
			dst := path.Join("./static", fileName)
			saveErr := c.SaveUploadedFile(file, dst)
			if saveErr == nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "success",
					"data": dst,
				})
			}
		}
	})

	err = r.Run(":8888")
	if err != nil {
		return
	}

}

func costTimeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()
		//亲求处理
		ctx.Next()
		//请求处理完获取花费的时间
		costTime := time.Since(nowTime)
		//获取当前访问的URL
		requestUrl := ctx.Request.URL.String()
		//打印输出
		common.Info(requestUrl + " : " + costTime.String())
	}
}
