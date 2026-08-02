package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const keyRequestId = "request_id"

func main() {
	router := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	router.Use(func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	}, func(c *gin.Context) {
		s := time.Now()
		c.Next()
		var requestId = 0
		if id, exists := c.Get(keyRequestId); exists {
			//  id.(int) 是单返回值形式的类型断言，语义是：
			//
			//  - 如果 id 的动态类型确实是 int，断言"通过"，表达式的结果就是那个 int 值，然后正常赋值给 requestId——和普通表达式赋值没有区别。
			//  - 如果 id 的动态类型不是 int，断言"不通过"，会直接 panic（interface conversion: interface {} is T, not int 之类的错误），程序会在这一行崩溃，不会进入赋值语句。
			//
			//  也就是说，断言只有两种结果：成功并拿到值完成赋值，或者直接 panic 而不会执行赋值——不存在"断言失败但仍然赋值（比如赋零值）"这种情况。如果想要"失败时不赋值、也不 panic 的安全写法，要用双返回值形式：
			//
			//  if v, ok := id.(int); ok {
			//      requestId = v
			//  }
			if v, ok := id.(int); ok {
				requestId = v
			}
		}
		// path, status, elapsed
		logger.Info("incoming request",
			zap.Int(keyRequestId, requestId),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	})

	router.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if id, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = id
		}
		c.JSON(200, h)
	})
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
