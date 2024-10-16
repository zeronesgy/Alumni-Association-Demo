package main

import (
	"Alumni-Association-Demo/common"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run(":8080"))
}
