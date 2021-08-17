package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type LoginForm struct {
	User string `form:"user"`
	Password string `form:"password"`
}

type formA struct {
	Foo string `json:"foo" xml:"foo"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar"`
}

type checkForm struct {
  Colors []string `form:"colors[]"`
}
func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		message := c.PostForm("name")
		c.JSON(200, gin.H{"message" : message})
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.GET("/reader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil && response.StatusCode != http.StatusOK{
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.POST("/bind", func(c *gin.Context) {
		var checkForm checkForm
		c.ShouldBind(&checkForm)
		c.JSON(200, gin.H{"color": checkForm})

		objA := formA{}
		objB := formB{}
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA != nil {
			c.JSON(http.StatusOK,gin.H{"data" : objA})
		}
		if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB != nil {
			c.JSON(http.StatusOK,gin.H{"data": objB})
		}
		fmt.Println(objA, objB)
	})

	srv := &http.Server{
		Addr: ":8081",
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}