package main

import (
	"fmt"
	routers "routers"

	"modules/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	middleware.InitTables(true)
	r.Use(middleware.InitException())
	r.Use(middleware.InitConnection())
	routers.InitRouter(r)
	err := r.Run(":80")
	if err != nil {
		fmt.Println(err)
	}
}
