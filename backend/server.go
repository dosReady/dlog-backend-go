package main

import (
	routers "routers"

	"modules/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.InitException())
	routers.InitRouter(r)
	r.Run(":80")
}
