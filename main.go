package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var port int
var addr string
var root string

func main() {

	flag.IntVar(&port, "p", 8075, "set port")
	flag.StringVar(&addr, "a", "0.0.0.0", "set ip addr")
	flag.StringVar(&root, "r", ".", "set root path")

	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-store")
		c.Next()
	})

	// 设置静态文件路由
	router.StaticFS("/", http.Dir(root))

	// 启动服务
	router.Run(fmt.Sprint(addr, ":", port))
}
