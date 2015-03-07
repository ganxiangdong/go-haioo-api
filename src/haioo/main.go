package main

import (
	"github.com/gin-gonic/gin"
	"haioo/routers"
)

func main() {
	r := gin.Default()
	routers.Init(r)
	r.Run(":8080")
}
