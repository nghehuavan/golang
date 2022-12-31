package routes

import (
	"esgo/controllers/about"
	"esgo/controllers/home"

	"github.com/gin-gonic/gin"
)

func Apply(engine *gin.Engine) {
	engine.GET("/", home.Index)
	engine.GET("/about", about.Index)
}