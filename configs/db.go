package configs

import (
	"fmt"
	"my-gram-1/helpers"
	"my-gram-1/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	userDB := helpers.GoDotEnvVariable("DB_USER")
	passDB := helpers.GoDotEnvVariable("DB_PASS")
	dbName := helpers.GoDotEnvVariable("DB_NAME")
	host := helpers.GoDotEnvVariable("DB_PORT")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userDB, passDB, host, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(entity.User{})

}

func GetDB() *gorm.DB {
	return DB
}
