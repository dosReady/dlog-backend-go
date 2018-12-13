package main

import (
	"log"

	r "github.com/dosReady/dlog/backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	r.SettingRouters(router)
	log.Fatal(router.Run(":80"))
}
