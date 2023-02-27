package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"os"
)

func Connect_database() *gorm.DB {
	dsn := os.Getenv("DATABASE_users") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_IP_ADDRESS") + ":3306)/FiberTutorial"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
