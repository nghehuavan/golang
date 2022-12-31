package main

import (
	"esgo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	controllers.ApplyRoutes(engine)
	engine.Run() // listen and serve on 0.0.0.0:8080
}
