package crud

import (
	"esgo/databases/entities"
	"esgo/databases/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl "http://localhost:8080/crud?name=hv"
// curl -d "login=test&password=test&name=hvn&address=Q7" -X POST "http://localhost:8080/crud"
// curl -d "id=2&login=test&password=changed&name=hvn_changed&address=Q7_changed" -X PUT "http://localhost:8080/crud"
// curl -X DELETE "http://localhost:8080/crud?id=4"
func Index(c *gin.Context) {
	var user entities.User
	if c.ShouldBind(&user) == nil {
		fmt.Println(user)
	}
	db := mysql.OpenConnection()
	defer func() {
		mysql.CloseConnection()
	}()

	if c.Request.Method == http.MethodGet {
		var results []entities.User
		db.Where("Name LIKE ?", "%"+user.Name+"%").Find(&results)
		c.JSON(http.StatusOK, results)

	} else if c.Request.Method == http.MethodPost {
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	} else if c.Request.Method == http.MethodPut {
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	} else if c.Request.Method == http.MethodDelete {
		db.Delete(&user)
		c.Status(http.StatusOK)
	}
}
