package main

import (
	"fmt"
	"log"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/mattli002/apiproxy/lib/multiproxy"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

var proxy = multiproxy.NewMultiProtocolSingleHostReverseProxy(fmt.Sprintf("127.0.0.1:5603"))

func proxyHandler(c *gin.Context) {
	//spew.Dump(c)
	proxy.ServeHTTP(c.Writer, c.Request)
	return
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Any("/api/v1/*param", proxyHandler)
	router.Any("/socket.io/*param", proxyHandler)
	port := "8531"
	log.Printf("proxy service on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
