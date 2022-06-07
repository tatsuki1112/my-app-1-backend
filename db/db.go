package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tatsuki1112/my-app-1-backend/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	godotenv.Load(".env")

	db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})

	fmt.Println(db)
	if err != nil {
		panic(err)
	}

	autoMigration()

}

func GetDB() *gorm.DB {
	return db
}


func Close() {
	db, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func autoMigration() {
	db.AutoMigrate(&entity.TrashUser{})
}