package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB

func OpenConnection() *gorm.DB {
	mySqlHost := "localhost"
	mySqlPort := "3306"
	mySqlUser := "root"
	mySqlPassword := "root"
	mySqlDatabase := "esgo"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mySqlUser,
		mySqlPassword,
		mySqlHost,
		mySqlPort,
		mySqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	instance = db
	return instance
}

func CloseConnection() {
	if instance != nil {
		sqlDb, _ := instance.DB()
		_ = sqlDb.Close()
	}
}
