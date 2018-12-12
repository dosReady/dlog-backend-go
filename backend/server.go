package main

import (
	"fmt"
	routers "routers"

	"modules/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.InitConnection())
	r.Use(middleware.HanddleException())
	routers.InitRouter(r)
	err := r.Run(":80")
	if err != nil {
		fmt.Println(err)
	}
}
