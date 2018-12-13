package main

import (
	"log"

	. "github.com/dosReady/dlog/backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	SettingRouters(router)
	log.Fatal(router.Run(":80"))
}
