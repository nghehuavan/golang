package routes

import (
	"esgo/controllers/crud"
	"esgo/controllers/home"

	"github.com/gin-gonic/gin"
)

func Apply(engine *gin.Engine) {
	engine.GET("/", home.Index)

	engine.GET("/crud", crud.Index)
	engine.POST("/crud", crud.Index)
	engine.PUT("/crud", crud.Index)
	engine.DELETE("/crud", crud.Index)
}
