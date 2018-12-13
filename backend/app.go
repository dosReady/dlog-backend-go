package main

import (
	"fmt"

	. "github.com/dosReady/dlog/backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	SettingRouters(router)
	if err := router.Run(":80"); err != nil {
		fmt.Println(err)
	}
}
