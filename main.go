package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var port int
	var addr string
	var root string

	flag.IntVar(&port, "p", 8075, "as")
	flag.StringVar(&addr, "a", "0.0.0.0", "asd")
	flag.StringVar(&root, "r", ".", "")

	router := gin.Default()

	// 设置静态文件路由
	router.StaticFS("/", http.Dir(root))

	// 启动服务
	router.Run(fmt.Sprint(addr, ":", port))
}
