package main

import (
	"esgo/databases/entities"
	"esgo/databases/mysql"

	"gorm.io/gorm"
)

func main() {
	mysql.OpenConnection()
	migrate(mysql.Instance)
	mysql.CloseConnection()
}
func migrate(db *gorm.DB) {

	db.AutoMigrate(&entities.User{})
	// if !db.Migrator().HasTable(&entities.User{}) {
	// 	err := db.Migrator().CreateTable(&entities.User{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
