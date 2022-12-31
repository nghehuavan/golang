package about

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	paramPairs := context.Request.URL.Query()
	fmt.Printf("paramPairs = %v\n", paramPairs)
	context.JSON(200, gin.H{
		"message": "welcome about",
	})
}
