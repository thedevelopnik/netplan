package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/app.js", "./dist/app.js")
	r.StaticFile("/about.js", "./dist/about.js")
	r.StaticFile("/", "./dist/index.html")
	// v1 := router.Group("/v1")
	r.Run() // listen and serve on 0.0.0.0:8080
}
