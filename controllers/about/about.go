package about

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	paramPairs := c.Request.URL.Query()
	fmt.Printf("paramPairs = %v\n", paramPairs)
	c.JSON(200, gin.H{
		"message": "welcome about",
	})
}
