package main

import (
	"log"

	r "github.com/dosReady/dlog/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	r.SettingRouters(router)
	log.Fatal(router.Run(":80"))
}
