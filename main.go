package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

const (
	DefaultReadTimeOut  = 1
	DefaultWriteTimeOut = 1
	Host                = "localhost"
	Port                = 9876
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	engine := gin.New()

	v1Group := engine.Group("/api/v1")
	// API PATH
	v1Group.POST("/topic/:topic/data", addDeviceData)

	endless.DefaultReadTimeOut = DefaultReadTimeOut
	endless.DefaultWriteTimeOut = DefaultWriteTimeOut
	addr := fmt.Sprintf("%s:%d", Host, Port)

	// Start api server
	go func() {
		err := endless.ListenAndServe(addr, engine)
		if err != nil {
			fmt.Printf("start server %s failed: %s", addr, err)
			os.Exit(2)
		}
	}()

	select {}
}