package bind

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

// http://localhost:8080/bind?name=hvn&address=Q7
// curl "http://localhost:8080/bind?name=hvn&address=Q7"
// curl -d "name=hvn&address=Q7" -X POST "http://localhost:8080/bind"
func Index(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		fmt.Println(person)
		fmt.Printf("%T\n", person)
	}

	c.JSON(200, person)

}
